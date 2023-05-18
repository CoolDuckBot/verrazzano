// Copyright (c) 2021, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package transform

import (
	"strings"

	"github.com/verrazzano/verrazzano/pkg/constants"
	vzprofiles "github.com/verrazzano/verrazzano/pkg/profiles"
	"github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1alpha1"
	"github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1beta1"
	"github.com/verrazzano/verrazzano/platform-operator/internal/config"
)

const (
	// implicit base profile (defaults)
	baseProfile = "base"
)

// GetEffectiveCR Creates an "effective" v1alpha1.Verrazzano CR based on the user defined resource merged with the profile definitions
// - Effective CR == base profile + declared profiles + ActualCR (in order)
// - last definition wins
func GetEffectiveCR(actualCR *v1alpha1.Verrazzano) (*v1alpha1.Verrazzano, error) {
	if actualCR == nil {
		return nil, nil
	}
	// Identify the set of profiles, base + declared
	profiles := []string{baseProfile, string(v1alpha1.Prod)}
	if len(actualCR.Spec.Profile) > 0 {
		profiles = append([]string{baseProfile}, strings.Split(string(actualCR.Spec.Profile), ",")...)
	}
	var profileFiles []string
	for _, profile := range profiles {
		profileFiles = append(profileFiles, config.GetProfile(v1alpha1.SchemeGroupVersion, profile))
	}
	// Merge the profile files into an effective profile YAML string
	effectiveCR, err := vzprofiles.MergeProfiles(actualCR, profileFiles...)
	if err != nil {
		return nil, err
	}
	effectiveCR.Status = v1alpha1.VerrazzanoStatus{} // Don't replicate the CR status in the effective config

	// Align the ClusterIssuer configurations between CertManager and ClusterIssuer components
	alignClusterIssuerConfigV1Alpha1(effectiveCR)

	return effectiveCR, nil
}

// GetEffectiveV1beta1CR Creates an "effective" v1beta1.Verrazzano CR based on the user defined resource merged with the profile definitions
// - Effective CR == base profile + declared profiles + ActualCR (in order)
// - last definition wins
func GetEffectiveV1beta1CR(actualCR *v1beta1.Verrazzano) (*v1beta1.Verrazzano, error) {
	if actualCR == nil {
		return nil, nil
	}
	// Identify the set of profiles, base + declared
	profiles := []string{baseProfile, string(v1beta1.Prod)}
	if len(actualCR.Spec.Profile) > 0 {
		profiles = append([]string{baseProfile}, strings.Split(string(actualCR.Spec.Profile), ",")...)
	}
	var profileFiles []string
	for _, profile := range profiles {
		profileFiles = append(profileFiles, config.GetProfile(v1beta1.SchemeGroupVersion, profile))
	}
	// Merge the profile files into an effective profile YAML string
	effectiveCR, err := vzprofiles.MergeProfilesForV1beta1(actualCR, profileFiles...)
	if err != nil {
		return nil, err
	}
	effectiveCR.Status = v1beta1.VerrazzanoStatus{} // Don't replicate the CR status in the effective config

	// Align the ClusterIssuer configurations between CertManager and ClusterIssuer components
	alignClusterIssuerConfigV1Beta1(effectiveCR)

	return effectiveCR, nil
}

// alignClusterIssuerConfigV1Alpha1 aligns the ClusterIssuer configurations between CertManager and the newer ClusterIssuer
// configurations.  The webhook validators will ensure only one is set to a non-defaulted value, so if one is
// configured align it with the other.
func alignClusterIssuerConfigV1Beta1(effectiveCR *v1beta1.Verrazzano) {
	// if Certificate in CertManager is empty, set it to default CA
	var emptyCertConfig = v1beta1.Certificate{}
	defaultCertConfig := v1beta1.Certificate{
		CA: v1beta1.CA{
			SecretName:               constants.DefaultVerrazzanoCASecretName,
			ClusterResourceNamespace: constants.CertManagerNamespace,
		},
	}

	certManagerConfig := effectiveCR.Spec.Components.CertManager
	if certManagerConfig.Certificate == emptyCertConfig {
		certManagerConfig.Certificate = defaultCertConfig
	}
	clusterIssuerConfig := effectiveCR.Spec.Components.ClusterIssuer
	if clusterIssuerConfig == nil {
		trueVal := true
		clusterIssuerConfig = &v1beta1.ClusterIssuerComponent{Enabled: &trueVal}
	}
	// if Certificate in CertManager is empty/defaulted, align it with the ClusterIssuer config
	if clusterIssuerConfig.Certificate == emptyCertConfig || clusterIssuerConfig.Certificate == defaultCertConfig {
		clusterIssuerConfig.Certificate = certManagerConfig.Certificate
	}
	// if Certificate in ClusterIssuer is empty/defaulted, align it with the CertManager config
	if certManagerConfig.Certificate == emptyCertConfig || certManagerConfig.Certificate == defaultCertConfig {
		certManagerConfig.Certificate = clusterIssuerConfig.Certificate
	}
}

// alignClusterIssuerConfigV1Alpha1 aligns the ClusterIssuer configurations between CertManager and the newer ClusterIssuer
// configurations.  The webhook validators will ensure only one is set to a non-defaulted value, so if one is
// configured align it with the other.
func alignClusterIssuerConfigV1Alpha1(effectiveCR *v1alpha1.Verrazzano) {
	// if Certificate in CertManager is empty, set it to default CA
	var emptyCertConfig = v1alpha1.Certificate{}
	defaultCertConfig := v1alpha1.Certificate{
		CA: v1alpha1.CA{
			SecretName:               constants.DefaultVerrazzanoCASecretName,
			ClusterResourceNamespace: constants.CertManagerNamespace,
		},
	}

	certManagerConfig := effectiveCR.Spec.Components.CertManager
	if certManagerConfig.Certificate == emptyCertConfig {
		certManagerConfig.Certificate = defaultCertConfig
	}
	clusterIssuerConfig := effectiveCR.Spec.Components.ClusterIssuer
	if clusterIssuerConfig == nil {
		trueVal := true
		clusterIssuerConfig = &v1alpha1.ClusterIssuerComponent{Enabled: &trueVal}
	}
	// if Certificate in CertManager is empty/defaulted, align it with the ClusterIssuer config
	if clusterIssuerConfig.Certificate == emptyCertConfig || clusterIssuerConfig.Certificate == defaultCertConfig {
		clusterIssuerConfig.Certificate = certManagerConfig.Certificate
	}
	// if Certificate in ClusterIssuer is empty/defaulted, align it with the CertManager config
	if certManagerConfig.Certificate == emptyCertConfig || certManagerConfig.Certificate == defaultCertConfig {
		certManagerConfig.Certificate = clusterIssuerConfig.Certificate
	}
}
