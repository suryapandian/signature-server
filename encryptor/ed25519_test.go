package encryptor

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Ed25519TestSuite struct {
	suite.Suite
	e *ed25519KeyPair
}

func (t *Ed25519TestSuite) SetupTest() {
	t.e = NewEd25519Encrytor([]byte(``))
}

func TestEd25519Suite(t *testing.T) {
	suite.Run(t, new(Ed25519TestSuite))
}

func (t *Ed25519TestSuite) TestGetPublicKey() {
	pubKey := t.e.GetPublicKey()
	t.NotEmpty(pubKey)
}

func (t *Ed25519TestSuite) TestGetTransactions() {
	result := t.e.Sign([]byte("/+ABAgM="))
	t.NotEmpty(result)
}
