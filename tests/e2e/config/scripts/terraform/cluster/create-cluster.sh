#!/bin/bash
#
# Copyright (c) 2020, 2023, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.
#


. ./init.sh

$SCRIPT_DIR/terraform init -no-color -reconfigure

set -o pipefail

# retry 3 times, 30 seconds apart
tries=0
MAX_TRIES=3
while true; do
   tries=$((tries+1))
   echo "terraform plan iteration ${tries}"
   $SCRIPT_DIR/terraform plan -var-file=$TF_VAR_nodepool_config.tfvars -var-file=$TF_VAR_region.tfvars -no-color && break
   if [ "$tries" -ge "$MAX_TRIES" ];
   then
      echo "Terraform plan tries exceeded.  Cluster creation has failed!"
      exit 1
   fi
   sleep 30
done

# retry 3 times, 30 seconds apart
tries=0
MAX_TRIES=3
while true; do
   tries=$((tries+1))
   echo "terraform apply iteration ${tries}"
   $SCRIPT_DIR/terraform apply -var-file=$TF_VAR_nodepool_config.tfvars -var-file=$TF_VAR_region.tfvars -auto-approve -no-color && break
   if [ "$tries" -ge "$MAX_TRIES" ];
   then
      echo "Terraform apply tries exceeded.  Cluster creation has failed!"
      break
   fi
   echo "Deleting Cluster Terraform and applying again"
   $SCRIPT_DIR/delete-cluster.sh
   sleep 30
done

if [ "$tries" -ge "$MAX_TRIES" ];
then
  exit 1
fi

echo "updating OKE private_workers_seclist to allow pub_lb_subnet access to workers"

# the script would return 0 even if it fails to update OKE private_workers_seclist
# because the OKE still could work if it didn't hit the rate limiting

# find vcn id "${var.label_prefix}-${var.vcn_name}"
VCN_ID=$(oci network vcn list \
  --compartment-id "${TF_VAR_compartment_id}" \
  --display-name "${TF_VAR_label_prefix}-oke-vcn" \
  | jq -r '.data[0].id')

if [ -z "$VCN_ID" ]; then
    echo "Failed to get the id for OKE cluster vcn ${TF_VAR_label_prefix}-oke-vcn"
    exit 0
fi

# find private_workers_seclist id
SEC_LIST_ID=$(oci network security-list list \
  --compartment-id "${TF_VAR_compartment_id}" \
  --display-name "${TF_VAR_label_prefix}-workers" \
  --vcn-id "${VCN_ID}" \
  | jq -r '.data[0].id')

if [ -z "$SEC_LIST_ID" ]; then
    echo "Failed to get the id for security-list ${TF_VAR_label_prefix}-workers"
    exit 0
fi

# find pub_lb_subnet CIDR
LB_SUBNET_CIDR=$(oci network subnet list \
  --compartment-id "${TF_VAR_compartment_id}" \
  --display-name "${TF_VAR_label_prefix}-pub_lb" \
  --vcn-id "${VCN_ID}" \
  | jq -r '.data[0]."cidr-block"')

if [ -z "$LB_SUBNET_CIDR" ]; then
    echo "Failed to get the cidr-block for subnet ${TF_VAR_label_prefix}-pub_lb"
    exit 0
fi

# get current ingress-security-rules
oci network security-list get --security-list-id "${SEC_LIST_ID}" | jq '.data."ingress-security-rules"' > ingress-security-rules.json
if [ $? -eq 0 ]; then
  echo "ingress-security-rules for security-list ${TF_VAR_label_prefix}-private-workers:"
  cat ingress-security-rules.json
else
  echo "Failed to retrieve the ingress-security-rules for security-list ${TF_VAR_label_prefix}-private-workers"
  exit 0
fi

# add pub_lb_subnet ingress-security-rule
cat ingress-security-rules.json | jq --arg LB_SUBNET_CIDR "${LB_SUBNET_CIDR}" '. += [{"description": "allow pub_lb_subnet access to workers","is-stateless": false,"protocol": "6","source": $LB_SUBNET_CIDR,"tcp-options": {"destination-port-range": {"max": 32767,"min": 30000}}},{"description": "allow pub_lb_subnet health check access to workers","is-stateless": false,"protocol": "6","source": $LB_SUBNET_CIDR,"tcp-options": {"destination-port-range": {"max": 10256,"min": 10256}}}]' > new.ingress-security-rules.json

# update private_workers_seclist
oci network security-list update --force --security-list-id "${SEC_LIST_ID}" --ingress-security-rules "file://${PWD}/new.ingress-security-rules.json"
if [ $? -eq 0 ]; then
  echo "Updated the OKE private_workers_seclist"
else
  echo "Failed to update the OKE private_workers_seclist"
fi

# Block docker.io and docker.com at the VCN level
# Get the VCN resolver
VCN_RESOLVER_ID=$(/usr/local/bin/oci dns resolver list \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --display-name "${TF_VAR_label_prefix}-oke-vcn" --all \
  | jq -r '.data[0].id')

echo "Resolver $VCN_RESOLVER_ID"

# Create a "blocked" DNS private view (not protected)
/usr/local/bin/oci dns view create \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --display-name "${TF_VAR_label_prefix}-oke-docker-blocker-view" \
  --wait-for-state ACTIVE
if [ $? -eq 0 ]; then
  echo "Created docker blocker view"
else
  echo "Failed to create docker blocker view"
fi

BLOCKER_VIEW_ID=$(/usr/local/bin/oci dns view list \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --display-name "${TF_VAR_label_prefix}-oke-docker-blocker-view" --all \
  | jq -r '.data[0].id')

echo "Blocker view ID: $BLOCKER_VIEW_ID"


/usr/local/bin/oci dns resolver update \
  --region ${TF_VAR_region}  \
  --force \
  --resolver-id $VCN_RESOLVER_ID \
  --attached-views "[ { \"viewId\": \"$BLOCKER_VIEW_ID\" } ]"
echo "Resolver updated to attach the view"

# Create docker.io blocker private zone
/usr/local/bin/oci dns zone create \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --name "docker.io" \
  --view-id $BLOCKER_VIEW_ID \
  --zone-type PRIMARY \
  --scope PRIVATE \
  --wait-for-state ACTIVE
if [ $? -eq 0 ]; then
  echo "Created docker.io blocker private zone"
else
  echo "Failed to create docker.io blocker private zone"
fi
DOCKER_IO_ZONE_ID=$(/usr/local/bin/oci dns zone list \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --name "docker.io" --all \
  | jq -r '.data[0].id')

echo "docker.io zone ID: $DOCKER_IO_ZONE_ID"

# Add docker.io A record to 127.0.0.1
/usr/local/bin/oci dns record zone patch \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --zone-name-or-id "docker.io" \
  --view-id $BLOCKER_VIEW_ID \
  --scope PRIVATE \
  --items '[ {"domain": "docker.io", "isProtected": false, "operation": "ADD", "rdata": "127.0.0.1", "rtype": "A", "ttl": "86400"  } ]'
if [ $? -eq 0 ]; then
  echo "Added record for docker.io blocker private zone"
else
  echo "Failed to add record for docker.io blocker private zone"
fi


# Create docker.io blocker private zone
/usr/local/bin/oci dns zone create \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --name "docker.com" \
  --view-id $BLOCKER_VIEW_ID \
  --zone-type PRIMARY \
  --scope PRIVATE \
  --wait-for-state ACTIVE
if [ $? -eq 0 ]; then
  echo "Created docker.com blocker private zone"
else
  echo "Failed to create docker.com blocker private zone"
fi

DOCKER_COM_ZONE_ID=$(/usr/local/bin/oci dns zone list \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --name "docker.com" --all \
  | jq -r '.data[0].id')

echo "docker.com zone ID: $DOCKER_COM_ZONE_ID"

# Add docker.io A record to 127.0.0.1
/usr/local/bin/oci dns record zone patch \
  --region ${TF_VAR_region}  \
  --compartment-id "${TF_VAR_compartment_id}" \
  --zone-name-or-id "docker.com" \
  --view-id $BLOCKER_VIEW_ID \
  --scope PRIVATE \
  --items '[ {"domain": "docker.com", "isProtected": false, "operation": "ADD", "rdata": "127.0.0.1", "rtype": "A", "ttl": "86400" } ]'
if [ $? -eq 0 ]; then
  echo "Added record for docker.io blocker private zone"
else
  echo "Failed to add record for docker.io blocker private zone"
fi
