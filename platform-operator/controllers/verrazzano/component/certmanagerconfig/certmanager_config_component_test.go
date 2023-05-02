// Copyright (c) 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package certmanagerconfig

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1beta1"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/common"
	"math/big"
	"net"
	clipkg "sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
	"time"

	cmutil "github.com/cert-manager/cert-manager/pkg/api/util"
	certv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	certv1fake "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned/fake"
	certv1client "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1"
	"github.com/stretchr/testify/assert"
	constants2 "github.com/verrazzano/verrazzano/pkg/constants"
	"github.com/verrazzano/verrazzano/pkg/k8sutil"
	"github.com/verrazzano/verrazzano/pkg/log/vzlog"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/spi"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	vzapi "github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1alpha1"
)

const (
	testDNSDomain  = "example.dns.io"
	testOCIDNSName = "ociDNS"
)

// TestIsInstalledCMNotPresent tests the IsInstalled function
// GIVEN a call to IsInstalled
//
//	WHEN the CM CRDs are not present
//	THEN an error is returned if anything is misconfigured
func TestIsInstalledCMNotPresent(t *testing.T) {
	runIsInstalledTest(t, false, false)
}

// TestIsInstalledCMNotPresent tests the IsInstalled function
// GIVEN a call to IsInstalled
//
//	WHEN the CM CRDs are present and the VZ cluster issuer exists
//	THEN no error is returned and the result is true
func TestIsInstalledCMAndIssuerArePresent(t *testing.T) {
	runIsInstalledTest(t, true, false, createCertManagerCRDs()...)
}

func runIsInstalledTest(t *testing.T, expectInstalled bool, expectErr bool, objs ...clipkg.Object) {
	var clientObjs []clipkg.Object
	if expectInstalled {
		clientObjs = append(clientObjs, &certv1.ClusterIssuer{ObjectMeta: metav1.ObjectMeta{Name: constants2.VerrazzanoClusterIssuerName}})
	}
	if len(objs) > 0 {
		clientObjs = append(clientObjs, objs...)
	}
	client := fake.NewClientBuilder().WithScheme(testScheme).WithObjects(clientObjs...).Build()
	defer func() { common.ResetNewClientFunc() }()
	common.SetNewClientFunc(func(opts clipkg.Options) (clipkg.Client, error) {
		return client, nil
	})

	fakeContext := spi.NewFakeContext(client, &vzapi.Verrazzano{}, nil, false)
	installed, err := fakeComponent.IsInstalled(fakeContext)
	assert.Equal(t, expectInstalled, installed, "Did not get expected result")
	if expectErr {
		assert.Error(t, err, "Did not get expected error")
	} else {
		assert.NoError(t, err, "Got unexpected error")
	}
}

// TestIsNotReadyNoCertManagerResourcesPresent tests the IsReady function
// GIVEN a call to IsReady
//
//	WHEN the CM CRDs are NOT present
//	THEN false is returned
func TestIsNotReadyNoCertManagerResourcesPresent(t *testing.T) {
	runIsReadyTest(t, false)
}

// TestIsNotReady tests the IsReady function
// GIVEN a call to IsReady
//
//	WHEN the CM CRDs are present but the issuer is not
//	THEN false is returned
func TestIsNotReady(t *testing.T) {
	runIsReadyTest(t, false, createCertManagerCRDs()...)
}

// TestIsReady tests the IsReady function
// GIVEN a call to IsReady
//
//	WHEN the CM CRDs and the ClusterIssuer are present
//	THEN true is returned
func TestIsReady(t *testing.T) {
	runIsReadyTest(t, true, createCertManagerCRDs()...)
}

func runIsReadyTest(t *testing.T, expectedReady bool, objs ...clipkg.Object) {
	var clientObjs []clipkg.Object
	if expectedReady {
		clientObjs = append(clientObjs, &certv1.ClusterIssuer{ObjectMeta: metav1.ObjectMeta{Name: constants2.VerrazzanoClusterIssuerName}})
	}
	if len(objs) > 0 {
		clientObjs = append(clientObjs, objs...)
	}
	client := fake.NewClientBuilder().WithScheme(testScheme).WithObjects(clientObjs...).Build()
	defer func() { common.ResetNewClientFunc() }()
	common.SetNewClientFunc(func(opts clipkg.Options) (clipkg.Client, error) {
		return client, nil
	})

	fakeContext := spi.NewFakeContext(client, &vzapi.Verrazzano{}, nil, false)
	ready := fakeComponent.IsReady(fakeContext)
	assert.Equal(t, expectedReady, ready, "Did not get expected result")
}

// TestValidateInstall tests the ValidateInstall function
// GIVEN a call to ValidateInstall
//
//	WHEN for various CM configurations
//	THEN an error is returned if anything is misconfigured
func TestValidateInstall(t *testing.T) {
	validationTests(t, false)
}

// TestPostInstallCA tests the PostInstall function
// GIVEN a call to PostInstall
//
//	WHEN the cert type is CA
//	THEN no error is returned
func TestPostInstallCA(t *testing.T) {
	localvz := defaultVZConfig.DeepCopy()
	localvz.Spec.Components.CertManager.Certificate.CA = ca

	defer func() { getClientFunc = k8sutil.GetCoreV1Client }()
	getClientFunc = createClientFunc(localvz.Spec.Components.CertManager.Certificate.CA, "defaultVZConfig-cn")

	defer func() { getCMClientFunc = GetCertManagerClientset }()
	getCMClientFunc = func() (certv1client.CertmanagerV1Interface, error) {
		cmClient := certv1fake.NewSimpleClientset()
		return cmClient.CertmanagerV1(), nil
	}

	client := fake.NewClientBuilder().WithScheme(testScheme).Build()
	err := fakeComponent.PostInstall(spi.NewFakeContext(client, localvz, nil, false))
	assert.NoError(t, err)
}

// TestPostUpgradeUpdateCA tests the PostUpgrade function
// GIVEN a call to PostUpgrade
//
//	WHEN the type is CA and the CA is updated
//	THEN the ClusterIssuer is updated correctly and no error is returned
func TestPostUpgradeUpdateCA(t *testing.T) {
	runCAUpdateTest(t, true)
}

// TestPostInstallUpdateCA tests the PostInstall function
// GIVEN a call to PostInstall
//
//	WHEN the type is CA and the CA is updated
//	THEN the ClusterIssuer is updated correctly and no error is returned
func TestPostInstallUpdateCA(t *testing.T) {
	runCAUpdateTest(t, false)
}

func runCAUpdateTest(t *testing.T, upgrade bool) {
	localvz := defaultVZConfig.DeepCopy()
	localvz.Spec.Components.CertManager.Certificate.CA = ca

	updatedVZ := defaultVZConfig.DeepCopy()
	newCA := vzapi.CA{
		SecretName:               "newsecret",
		ClusterResourceNamespace: "newnamespace",
	}
	updatedVZ.Spec.Components.CertManager.Certificate.CA = newCA

	defer func() { getClientFunc = k8sutil.GetCoreV1Client }()
	getClientFunc = createClientFunc(updatedVZ.Spec.Components.CertManager.Certificate.CA, "defaultVZConfig-cn")

	defer func() { getCMClientFunc = GetCertManagerClientset }()
	cmClient := certv1fake.NewSimpleClientset()
	getCMClientFunc = func() (certv1client.CertmanagerV1Interface, error) {
		return cmClient.CertmanagerV1(), nil
	}

	expectedIssuer := &certv1.ClusterIssuer{
		Spec: certv1.IssuerSpec{
			IssuerConfig: certv1.IssuerConfig{
				CA: &certv1.CAIssuer{
					SecretName: newCA.SecretName,
				},
			},
		},
	}

	client := fake.NewClientBuilder().WithScheme(testScheme).WithObjects(localvz).Build()
	ctx := spi.NewFakeContext(client, updatedVZ, nil, false)

	var err error
	if upgrade {
		err = fakeComponent.Upgrade(ctx)
	} else {
		err = fakeComponent.Install(ctx)
	}
	assert.NoError(t, err)

	actualIssuer := &certv1.ClusterIssuer{}
	assert.NoError(t, client.Get(context.TODO(), types.NamespacedName{Name: constants2.VerrazzanoClusterIssuerName}, actualIssuer))
	assert.Equal(t, expectedIssuer.Spec.CA, actualIssuer.Spec.CA)
}

// TestPostInstallAcme tests the PostInstall function
// GIVEN a call to PostInstall
//
//	WHEN the cert type is Acme
//	THEN no error is returned
func TestPostInstallAcme(t *testing.T) {
	localvz := defaultVZConfig.DeepCopy()
	localvz.Spec.Components.CertManager.Certificate.Acme = acme
	client := fake.NewClientBuilder().WithScheme(testScheme).Build()
	// set OCI DNS secret value and create secret
	localvz.Spec.Components.DNS = &vzapi.DNSComponent{
		OCI: &vzapi.OCI{
			OCIConfigSecret: testOCIDNSName,
			DNSZoneName:     testDNSDomain,
		},
	}
	_ = client.Create(context.TODO(), &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testOCIDNSName,
			Namespace: ComponentNamespace,
		},
	})
	err := fakeComponent.PostInstall(spi.NewFakeContext(client, localvz, nil, false))
	assert.NoError(t, err)
}

// TestPostUpgradeAcmeUpdate tests the PostUpgrade function
// GIVEN a call to PostUpgrade
//
//	WHEN the cert type is Acme and the config has been updated
//	THEN the ClusterIssuer is updated as expected no error is returned
func TestPostUpgradeAcmeUpdate(t *testing.T) {
	runAcmeUpdateTest(t, true, false)
}

// TestPostInstallAcme tests the PostInstall function
// GIVEN a call to PostInstall
//
//	WHEN the cert type is Acme and the config has been updated
//	THEN the ClusterIssuer is updated as expected no error is returned
func TestPostInstallAcmeUpdate(t *testing.T) {
	runAcmeUpdateTest(t, false, false)
}

// TestPostInstallIPAuthAcmeUpdate tests the PostInstall function
// GIVEN a call to PostInstall
//
//	WHEN the cert type is Acme with IP auth and the config has been updated
//	THEN the ClusterIssuer is updated as expected no error is returned
func TestPostInstallIPAuthAcmeUpdate(t *testing.T) {
	runAcmeUpdateTest(t, false, true)
}

func runAcmeUpdateTest(t *testing.T, upgrade bool, useIPInSecret bool) {
	localvz := defaultVZConfig.DeepCopy()
	localvz.Spec.Components.CertManager.Certificate.Acme = acme
	// set OCI DNS secret value and create secret
	compartmentOCID := "compartmentID"
	oci := &vzapi.OCI{
		OCIConfigSecret:        testOCIDNSName,
		DNSZoneName:            testDNSDomain,
		DNSZoneCompartmentOCID: compartmentOCID,
	}
	localvz.Spec.Components.DNS = &vzapi.DNSComponent{
		OCI: oci,
	}

	oldSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testOCIDNSName,
			Namespace: ComponentNamespace,
		},
		Data: map[string][]byte{},
	}

	newSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "newociDNSSecret",
			Namespace: ComponentNamespace,
		},
		Data: map[string][]byte{},
	}

	if useIPInSecret {
		ipAuthSnippet := `
auth: 
  authtype: instance_principal 
`
		oldSecret.Data["oci.yaml"] = []byte(ipAuthSnippet)
		newSecret.Data["oci.yaml"] = []byte(ipAuthSnippet)
	}

	existingIssuerTemplateData := templateData{
		Email:          acme.EmailAddress,
		AcmeSecretName: caAcmeSecretName,
		Server:         acme.Environment,
		SecretName:     oci.OCIConfigSecret,
		OCIZoneName:    oci.DNSZoneName,
	}
	existingIssuer, _ := createAcmeClusterIssuer(vzlog.DefaultLogger(), existingIssuerTemplateData)

	updatedVz := defaultVZConfig.DeepCopy()
	newAcme := vzapi.Acme{
		Provider:     "letsEncrypt",
		EmailAddress: "slbronkowitz@gmail.com",
		Environment:  "production",
	}
	const newCompartmentOCID = "somenewocid"
	newOCI := &vzapi.OCI{
		DNSZoneCompartmentOCID: newCompartmentOCID,
		OCIConfigSecret:        newSecret.Name,
		DNSZoneName:            "newzone.dns.io",
	}
	updatedVz.Spec.Components.CertManager.Certificate.Acme = newAcme
	updatedVz.Spec.Components.DNS = &vzapi.DNSComponent{OCI: newOCI}

	expectedIssuerTemplateData := templateData{
		Email:           newAcme.EmailAddress,
		AcmeSecretName:  caAcmeSecretName,
		Server:          letsEncryptProdEndpoint,
		SecretName:      newOCI.OCIConfigSecret,
		OCIZoneName:     newOCI.DNSZoneName,
		CompartmentOCID: newCompartmentOCID,
	}
	if useIPInSecret {
		expectedIssuerTemplateData.UseInstancePrincipals = true
	}
	expectedIssuer, _ := createAcmeClusterIssuer(vzlog.DefaultLogger(), expectedIssuerTemplateData)

	client := fake.NewClientBuilder().WithScheme(testScheme).WithObjects(localvz, oldSecret, newSecret, existingIssuer).Build()
	ctx := spi.NewFakeContext(client, updatedVz, nil, false)

	var err error
	if upgrade {
		err = fakeComponent.PostUpgrade(ctx)
	} else {
		err = fakeComponent.PostInstall(ctx)
	}
	assert.NoError(t, err)

	actualIssuer, err := createACMEIssuerObject(ctx)
	assert.Equal(t, expectedIssuer.Object["spec"], actualIssuer.Object["spec"])
	assert.NoError(t, err)
}

// TestClusterIssuerUpdated tests the createOrUpdateClusterIssuer function
// GIVEN a call to createOrUpdateClusterIssuer
// WHEN the ClusterIssuer is updated and there are existing certificates with failed and successful CertificateRequests
// THEN the Cert status is updated to request a renewal, and any failed CertificateRequests are cleaned up beforehand
func TestClusterIssuerUpdated(t *testing.T) {
	asserts := assert.New(t)

	localvz := defaultVZConfig.DeepCopy()
	localvz.Spec.Components.CertManager.Certificate.Acme = acme
	// set OCI DNS secret value and create secret
	oci := &vzapi.OCI{
		OCIConfigSecret: testOCIDNSName,
		DNSZoneName:     testDNSDomain,
	}
	localvz.Spec.Components.DNS = &vzapi.DNSComponent{
		OCI: oci,
	}

	// The existing cluster issuer that will be updated
	existingClusterIssuer := &certv1.ClusterIssuer{
		ObjectMeta: metav1.ObjectMeta{
			Name: constants2.VerrazzanoClusterIssuerName,
		},
		Spec: certv1.IssuerSpec{
			IssuerConfig: certv1.IssuerConfig{
				CA: &certv1.CAIssuer{
					SecretName: ca.SecretName,
				},
			},
		},
	}

	// The certificate that we expect to be renewed
	certName := "mycert"
	certNamespace := "certns"
	certificate := &certv1.Certificate{
		ObjectMeta: metav1.ObjectMeta{Name: certName, Namespace: certNamespace},
		Spec: certv1.CertificateSpec{
			IssuerRef: cmmeta.ObjectReference{
				Name: constants2.VerrazzanoClusterIssuerName,
			},
			SecretName: certName,
		},
		Status: certv1.CertificateStatus{},
	}

	// A certificate request for the above cert that was successful
	certificateRequest1 := &certv1.CertificateRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foorequest1",
			Namespace: certificate.Namespace,
			Annotations: map[string]string{
				certRequestNameAnnotation: certificate.Name,
			},
		},
		Status: certv1.CertificateRequestStatus{
			Conditions: []certv1.CertificateRequestCondition{
				{Type: certv1.CertificateRequestConditionReady, Status: cmmeta.ConditionTrue, Reason: certv1.CertificateRequestReasonIssued},
			},
		},
	}

	// A certificate request for the above cert that is in a failed state; this should be deleted (or it will block an Issuing request)
	certificateRequest2 := &certv1.CertificateRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foorequest2",
			Namespace: certificate.Namespace,
			Annotations: map[string]string{
				certRequestNameAnnotation: certificate.Name,
			},
		},
		Status: certv1.CertificateRequestStatus{
			Conditions: []certv1.CertificateRequestCondition{
				{Type: certv1.CertificateRequestConditionReady, Status: cmmeta.ConditionFalse, Reason: certv1.CertificateRequestReasonFailed},
			},
		},
	}

	// An unrelated certificate request, for different certificate; this should be untouched
	otherCertRequest := &certv1.CertificateRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "barrequest",
			Namespace: certificate.Namespace,
			Annotations: map[string]string{
				certRequestNameAnnotation: "someothercert",
			},
		},
		Status: certv1.CertificateRequestStatus{
			Conditions: []certv1.CertificateRequestCondition{
				{Type: certv1.CertificateRequestConditionReady, Status: cmmeta.ConditionFalse, Reason: certv1.CertificateRequestReasonFailed},
			},
		},
	}

	// The OCI DNS secret is expected to be present for this configuration
	ociSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testOCIDNSName,
			Namespace: ComponentNamespace,
		},
	}

	// Fake controllerruntime client and ComponentContext for the call
	client := fake.NewClientBuilder().WithScheme(testScheme).WithObjects(existingClusterIssuer, certificate, ociSecret).Build()
	ctx := spi.NewFakeContext(client, localvz, nil, false)

	// Fake Go client for the CertManager clientset
	cmClient := certv1fake.NewSimpleClientset(certificate, certificateRequest1, certificateRequest2, otherCertRequest)

	defer func() { getCMClientFunc = GetCertManagerClientset }()
	getCMClientFunc = func() (certv1client.CertmanagerV1Interface, error) {
		return cmClient.CertmanagerV1(), nil
	}

	// Create an issuer
	issuerName := caCertCommonName + "-a23asdfa"
	fakeIssuerCert := createFakeCertificate(issuerName)
	fakeIssuerCertBytes, err := createFakeCertBytes(issuerName, nil)
	if err != nil {
		return
	}
	issuerSecret, err := createCertSecret(caCertificateName, ComponentNamespace, fakeIssuerCertBytes)
	if err != nil {
		return
	}
	// Create a leaf cert signed by the above issuer
	fakeCertBytes, err := createFakeCertBytes("common-name", fakeIssuerCert)
	if err != nil {
		return
	}
	certSecret, err := createCertSecret(certName, certNamespace, fakeCertBytes)
	if !asserts.NoError(err, "Error creating test cert secret") {
		return
	}
	defer func() { getClientFunc = k8sutil.GetCoreV1Client }()
	getClientFunc = func(log ...vzlog.VerrazzanoLogger) (v1.CoreV1Interface, error) {
		return k8sfake.NewSimpleClientset(certSecret, issuerSecret).CoreV1(), nil
	}

	// create the component and issue the call
	component := NewComponent().(certManagerConfigComponent)
	asserts.NoError(component.createOrUpdateClusterIssuer(ctx))

	// Verify the certificate status has an Issuing condition; this informs CertManager to renew the certificate
	updatedCert, err := cmClient.CertmanagerV1().Certificates(certificate.Namespace).Get(context.TODO(), certificate.Name, metav1.GetOptions{})
	asserts.NoError(err)
	asserts.True(cmutil.CertificateHasCondition(updatedCert, certv1.CertificateCondition{
		Type:   certv1.CertificateConditionIssuing,
		Status: cmmeta.ConditionTrue,
	}))

	// Verify the successful CertificateRequest was NOT deleted
	certReq1, err := cmClient.CertmanagerV1().CertificateRequests(certificate.Namespace).Get(context.TODO(), certificateRequest1.Name, metav1.GetOptions{})
	asserts.NoError(err)
	asserts.NotNil(certReq1)

	// Verify the failed CertificateRequest for the target certificate WAS deleted
	certReq2, err := cmClient.CertmanagerV1().CertificateRequests(certificate.Namespace).Get(context.TODO(), certificateRequest2.Name, metav1.GetOptions{})
	asserts.Error(err)
	asserts.True(errors.IsNotFound(err))
	asserts.Nil(certReq2)

	// Verify the unrelated CertificateRequest was NOT deleted
	otherReq, err := cmClient.CertmanagerV1().CertificateRequests(certificate.Namespace).Get(context.TODO(), otherCertRequest.Name, metav1.GetOptions{})
	asserts.NoError(err)
	asserts.NotNil(otherReq)
}

// TestUninstallNoCRDs tests the Uninstall function
// GIVEN a call to Uninstall
//
//	WHEN the CM are NOT CRDs are present
//	THEN no error is returned
func TestUninstallNoCRDs(t *testing.T) {
	runUninstallTest(t)
}

// TestUninstall tests the Uninstall function
// GIVEN a call to Uninstall
//
//	WHEN the CM CRDs are present
//	THEN no error is returned
func TestUninstall(t *testing.T) {
	runUninstallTest(t, createCertManagerCRDs()...)
}

func runUninstallTest(t *testing.T, objs ...clipkg.Object) {
	defer func() { common.ResetNewClientFunc() }()

	client := fake.NewClientBuilder().WithScheme(testScheme).WithObjects(objs...).Build()
	common.SetNewClientFunc(func(opts clipkg.Options) (clipkg.Client, error) {
		return client, nil
	})

	fakeContext := spi.NewFakeContext(client, &vzapi.Verrazzano{}, nil, false)

	// We do more exhaustive testing of uninstall in the tests for uninstallVerrazzanoCertManagerResources, so
	// we don't expect errors here
	err := fakeComponent.Uninstall(fakeContext)
	assert.NoError(t, err, "Got unexpected error")
}

// TestDryRun tests the behavior when DryRun is enabled, mainly for code coverage
// GIVEN a call to PostInstall/PostUpgrade/PreInstall
//
//	WHEN the ComponentContext has DryRun set to true
//	THEN no error is returned
func TestDryRun(t *testing.T) {
	client := fake.NewClientBuilder().WithScheme(testScheme).Build()
	ctx := spi.NewFakeContext(client, defaultVZConfig, nil, true)

	assert.NoError(t, fakeComponent.PreInstall(ctx))
	assert.NoError(t, fakeComponent.PostInstall(ctx))
	assert.True(t, fakeComponent.IsReady(ctx))

	installed, err := fakeComponent.IsInstalled(ctx)
	assert.True(t, installed)
	assert.NoError(t, err)

	assert.NoError(t, fakeComponent.Install(ctx))
	assert.NoError(t, fakeComponent.PreUpgrade(ctx))
	assert.NoError(t, fakeComponent.Upgrade(ctx))
	assert.NoError(t, fakeComponent.PostUpgrade(ctx))

	assert.NoError(t, fakeComponent.Uninstall(ctx))
}

func createCertSecret(name string, namespace string, fakeCertBytes []byte) (*corev1.Secret, error) {
	//fakeCertBytes, err := createFakeCertBytes(cn)
	//if err != nil {
	//	return nil, err
	//}
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			corev1.TLSCertKey: fakeCertBytes,
		},
		Type: corev1.SecretTypeTLS,
	}
	return secret, nil
}

func createCertSecretNoParent(name string, namespace string, cn string) (*corev1.Secret, error) {
	fakeCertBytes, err := createFakeCertBytes(cn, nil)
	if err != nil {
		return nil, err
	}
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			corev1.TLSCertKey: fakeCertBytes,
		},
		Type: corev1.SecretTypeTLS,
	}
	return secret, nil
}

var testRSAKey *rsa.PrivateKey

func getRSAKey() (*rsa.PrivateKey, error) {
	if testRSAKey == nil {
		var err error
		if testRSAKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
			return nil, err
		}
	}
	return testRSAKey, nil
}

func createFakeCertBytes(cn string, parent *x509.Certificate) ([]byte, error) {
	rsaKey, err := getRSAKey()
	if err != nil {
		return []byte{}, err
	}

	cert := createFakeCertificate(cn)
	if parent == nil {
		parent = cert
	}
	pubKey := &rsaKey.PublicKey
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, parent, pubKey, rsaKey)
	if err != nil {
		return []byte{}, err
	}
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	return certPem, nil
}

func createFakeCertificate(cn string) *x509.Certificate {
	duration30, _ := time.ParseDuration("-30h")
	notBefore := time.Now().Add(duration30) // valid 30 hours ago
	duration1Year, _ := time.ParseDuration("90h")
	notAfter := notBefore.Add(duration1Year) // for 90 hours
	serialNo := big.NewInt(int64(123123413123))
	cert := &x509.Certificate{
		Subject: pkix.Name{
			Country:      []string{"US"},
			Organization: []string{"BarOrg"},
			SerialNumber: "2234",
			CommonName:   cn,
		},
		SerialNumber:          serialNo,
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		SignatureAlgorithm:    x509.SHA256WithRSA,
	}

	cert.IPAddresses = append(cert.IPAddresses, net.ParseIP("127.0.0.1"))
	cert.IPAddresses = append(cert.IPAddresses, net.ParseIP("::"))
	cert.DNSNames = append(cert.DNSNames, "localhost")
	return cert
}

// All of this below is to make Sonar happy
type validationTestStruct struct {
	name      string
	old       *vzapi.Verrazzano
	new       *vzapi.Verrazzano
	coreV1Cli func(_ ...vzlog.VerrazzanoLogger) (v1.CoreV1Interface, error)
	wantErr   bool
}

var tests = []validationTestStruct{
	{
		name:    "No OCI DNS or CertManager present",
		old:     &vzapi.Verrazzano{},
		new:     &vzapi.Verrazzano{},
		wantErr: false,
	},
	{
		name: "CertManager and OCI DNS webhook enabled",
		old:  &vzapi.Verrazzano{},
		new: &vzapi.Verrazzano{
			Spec: vzapi.VerrazzanoSpec{
				Components: vzapi.ComponentSpec{
					CertManager: &vzapi.CertManagerComponent{
						Enabled: getBoolPtr(true),
					},
					DNS: &vzapi.DNSComponent{
						OCI: &vzapi.OCI{
							DNSScope:               "GLOBAL",
							DNSZoneCompartmentOCID: "ocid",
							DNSZoneOCID:            "zoneOcid",
							DNSZoneName:            "zoneName",
							OCIConfigSecret:        "oci",
						},
					},
				},
			},
		},
		wantErr: false,
	},
	{
		name: "CertManager Disabled and OCI DNS webhook enabled",
		old:  &vzapi.Verrazzano{},
		new: &vzapi.Verrazzano{
			Spec: vzapi.VerrazzanoSpec{
				Components: vzapi.ComponentSpec{
					CertManager: &vzapi.CertManagerComponent{
						Enabled: getBoolPtr(false),
					},
					DNS: &vzapi.DNSComponent{
						OCI: &vzapi.OCI{
							DNSScope:               "GLOBAL",
							DNSZoneCompartmentOCID: "ocid",
							DNSZoneOCID:            "zoneOcid",
							DNSZoneName:            "zoneName",
							OCIConfigSecret:        "oci",
						},
					},
				},
			},
		},
		wantErr: false,
	},
	{
		name: "CertManager Component enabled",
		old:  &vzapi.Verrazzano{},
		new: &vzapi.Verrazzano{
			Spec: vzapi.VerrazzanoSpec{
				Components: vzapi.ComponentSpec{
					CertManager: &vzapi.CertManagerComponent{
						Enabled: getBoolPtr(true),
					},
				},
			},
		},
		wantErr: false,
	},
}

func validationTests(t *testing.T, isUpdate bool) {
	defer func() { common.ResetNewClientFunc() }()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Cert Manager Namespace already exists" && isUpdate { // will throw error only during installation
				tt.wantErr = false
			}
			client := fake.NewClientBuilder().WithScheme(testScheme).WithObjects(createCertManagerCRDs()...).Build()
			common.SetNewClientFunc(func(opts clipkg.Options) (clipkg.Client, error) {
				return client, nil
			})
			c := NewComponent()
			runValidationTest(t, tt, isUpdate, c)
		})
	}
}

func runValidationTest(t *testing.T, tt validationTestStruct, isUpdate bool, c spi.Component) {
	//	defer func() { k8sutil.GetCoreV1Func = k8sutil.GetCoreV1Client }()
	if isUpdate {
		if err := c.ValidateUpdate(tt.old, tt.new); (err != nil) != tt.wantErr {
			t.Errorf("ValidateUpdate() error = %v, wantErr %v", err, tt.wantErr)
		}
		v1beta1New := &v1beta1.Verrazzano{}
		v1beta1Old := &v1beta1.Verrazzano{}
		err := tt.new.ConvertTo(v1beta1New)
		assert.NoError(t, err)
		err = tt.old.ConvertTo(v1beta1Old)
		assert.NoError(t, err)
		if err := c.ValidateUpdateV1Beta1(v1beta1Old, v1beta1New); (err != nil) != tt.wantErr {
			t.Errorf("ValidateUpdateV1Beta1() error = %v, wantErr %v", err, tt.wantErr)
		}

	} else {
		wantErr := tt.name != "disable" && tt.wantErr // hack for disable validation, allowed on initial install but not on update
		if tt.coreV1Cli != nil {
			k8sutil.GetCoreV1Func = tt.coreV1Cli
		} else {
			k8sutil.GetCoreV1Func = common.MockGetCoreV1()
		}
		if err := c.ValidateInstall(tt.new); (err != nil) != wantErr {
			t.Errorf("ValidateInstall() error = %v, wantErr %v", err, tt.wantErr)
		}
		v1beta1Vz := &v1beta1.Verrazzano{}
		err := tt.new.ConvertTo(v1beta1Vz)
		assert.NoError(t, err)
		if err := c.ValidateInstallV1Beta1(v1beta1Vz); (err != nil) != wantErr {
			t.Errorf("ValidateInstallV1Beta1() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}
