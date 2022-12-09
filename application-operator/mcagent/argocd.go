// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package mcagent

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

const (
	KubeSystemNamespace    = "kube-system"
	CaCrtKey               = "ca.crt"
	ServiceAccountName     = "argocd-manager"
	SecName                = "argocd-manager-token"
	ClusterRoleName        = "argocd-manager-role"
	ClusterRoleBindingName = "argocd-manager-role-binding"
)

func (s *Syncer) createArgoCDServiceAccount() error {
	var serviceAccount corev1.ServiceAccount
	serviceAccount.Name = ServiceAccountName
	serviceAccount.Namespace = KubeSystemNamespace

	_, err := controllerruntime.CreateOrUpdate(s.Context, s.LocalClient, &serviceAccount, func() error {
		mutateServiceAccount(serviceAccount)
		s.Log.Infof("createArgoCDServiceAccount: ArgoCD ServiceAccount created successfully")
		return nil
	})
	return err
}

func (s *Syncer) createArgoCDSecret(secretData []byte) error {
	var secret corev1.Secret
	secret.Name = SecName
	secret.Namespace = KubeSystemNamespace

	// Create or update on the local cluster
	_, err := controllerruntime.CreateOrUpdate(s.Context, s.LocalClient, &secret, func() error {
		mutateArgoCDSecret(secret, secretData)
		s.Log.Infof("createArgoCDSecret: ArgoCD secret created successfully")
		return nil
	})
	return err
}

func (s *Syncer) createArgoCDRole() error {
	role := rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: ClusterRoleName}}

	_, err := controllerruntime.CreateOrUpdate(s.Context, s.LocalClient, &role, func() error {
		mutateClusterRole(role)
		s.Log.Infof("createArgoCDRole: ArgoCD Role created successfully")
		return nil
	})
	return err
}

func (s *Syncer) createArgoCDRoleBinding() error {
	binding := rbacv1.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: ClusterRoleBindingName}}

	_, err := controllerruntime.CreateOrUpdate(s.Context, s.LocalClient, &binding, func() error {
		mutateRoleBinding(binding)
		s.Log.Infof("createArgoCDRoleBinding: ArgoCD Rolebinding created successfully")
		return nil
	})
	return err
}

func mutateServiceAccount(sa corev1.ServiceAccount) {
	sa.Secrets = []corev1.ObjectReference{
		{
			Name: SecName,
		},
	}
}

func mutateArgoCDSecret(secret corev1.Secret, secretData []byte) {
	if secret.Data == nil {
		secret.Data = make(map[string][]byte)
	}
	secret.Type = corev1.SecretTypeServiceAccountToken
	secret.Data[CaCrtKey] = secretData
	secret.Annotations = map[string]string{
		corev1.ServiceAccountNameKey: ServiceAccountName,
	}
}

func mutateClusterRole(role rbacv1.ClusterRole) {
	role.Rules = []rbacv1.PolicyRule{
		{
			APIGroups: []string{"*"},
			Resources: []string{"*"},
			Verbs:     []string{"*"},
		},
	}
}

func mutateRoleBinding(binding rbacv1.ClusterRoleBinding) {
	binding.RoleRef = rbacv1.RoleRef{
		APIGroup: rbacv1.GroupName,
		Kind:     "ClusterRole",
		Name:     ClusterRoleName,
	}
	binding.Subjects = []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      ServiceAccountName,
			Namespace: KubeSystemNamespace,
		},
	}
}

func (s *Syncer) createArgocdResources(secretData []byte) error {
	if err := s.createArgoCDServiceAccount(); err != nil {
		return err
	}
	if err := s.createArgoCDSecret(secretData); err != nil {
		return err
	}
	if err := s.createArgoCDRole(); err != nil {
		return err
	}
	if err := s.createArgoCDRoleBinding(); err != nil {
		return err
	}
	return nil
}
