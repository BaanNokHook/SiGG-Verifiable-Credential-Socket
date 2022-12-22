/*
Go-lang-Vcs-NEXTClan
*/

//go:generate mockgen -destination mocks/kms_mocks.go -self_package mocks -package mocks -source=kms.go -mock_names VCSKeyManager=MockVCSKeyManager

package kms

import (
	"net/http"

	"github.com/hyperledger/aries-framework-go/pkg/doc/jose/jwk"
	"github.com/hyperledger/aries-framework-go/pkg/kms"

	"github.com/trustbloc/vcs/pkg/doc/vc"
	vcsverifiable "github.com/trustbloc/vcs/pkg/doc/verifiable"
)

type Type string

const (
	AWS   Type = "aws"
	Local Type = "local"
	Web   Type = "web"
)

// Config configure kms that stores signing keys.
type Config struct {
	KMSType    Type
	Endpoint   string
	Region     string
	HTTPClient *http.Client

	SecretLockKeyPath string
	DBType            string
	DBURL             string
	DBPrefix          string
}

type VCSKeyManager interface {
	SupportedKeyTypes() []kms.KeyType
	CreateJWKKey(keyType kms.KeyType) (string, *jwk.JWK, error)
	CreateCryptoKey(keyType kms.KeyType) (string, interface{}, error)
	NewVCSigner(creator string, signatureType vcsverifiable.SignatureType) (vc.SignerAlgorithm, error)
}
