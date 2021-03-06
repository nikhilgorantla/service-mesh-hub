package csr_generator

import (
	"context"

	securityv1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	kubernetes_core "github.com/solo-io/service-mesh-hub/pkg/clients/kubernetes/core"
	"github.com/solo-io/service-mesh-hub/pkg/security/certgen"
	cert_secrets "github.com/solo-io/service-mesh-hub/pkg/security/secrets"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
)

const (
	PrivateKeyNameSuffix = "-private-key"
	PrivateKeySizeBytes  = 4096
)

type certClient struct {
	secretClient        kubernetes_core.SecretsClient
	signer              certgen.Signer
	privateKeyGenerator PrivateKeyGenerator
}

func NewCertClient(
	secretClient kubernetes_core.SecretsClient,
	signer certgen.Signer,
	privateKeyGenerator PrivateKeyGenerator,
) CertClient {
	return &certClient{
		secretClient:        secretClient,
		signer:              signer,
		privateKeyGenerator: privateKeyGenerator,
	}
}

// Persist the intermediate cert's private key as a secret of type cert_secrets.IntermediateCertSecretType
func (c *certClient) EnsureSecretKey(
	ctx context.Context,
	obj *securityv1alpha1.VirtualMeshCertificateSigningRequest,
) (*cert_secrets.IntermediateCAData, error) {
	secret, err := c.secretClient.Get(ctx, buildSecretName(obj), obj.GetNamespace())
	if err != nil {
		if !kubeerrs.IsNotFound(err) {
			return nil, err
		}
		privateKey, err := c.privateKeyGenerator.GenerateRSA(PrivateKeySizeBytes)
		if err != nil {
			return nil, err
		}
		certData := &cert_secrets.IntermediateCAData{
			CaPrivateKey: privateKey,
		}
		newSecret := certData.BuildSecret(buildSecretName(obj), obj.GetNamespace())
		if err = c.secretClient.Create(ctx, newSecret); err != nil {
			return nil, err
		}
		return certData, nil

	}
	return cert_secrets.IntermediateCADataFromSecret(secret)
}

// suffix the name of the CSR with "-private-key" to avoid confusion, since we're reusing the
// cert_secrets.IntermediateCertSecretType secret type
func buildSecretName(obj *securityv1alpha1.VirtualMeshCertificateSigningRequest) string {
	return obj.GetName() + PrivateKeyNameSuffix
}
