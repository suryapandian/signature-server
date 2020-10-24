package server

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionTestSuite struct {
	suite.Suite
	txn *Transaction
}

func (t *TransactionTestSuite) SetupTest() {
	t.txn = NewTransaction([]byte(``))
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}

func (t *TransactionTestSuite) TestGetPublicKey() {
	pubKey := t.txn.GetPublicKey()
	t.NotEmpty(pubKey)
}

func (t *TransactionTestSuite) TestSave() {
	var data *string
	id, err := t.txn.Save(data)
	t.Error(err)
	t.Equal(InvalidTransaction, err)
	t.Empty(id)

	validData := "/+ABAgM="
	id, err = t.txn.Save(&validData)
	t.NoError(err)
	t.NotEmpty(id)

}

func (t *TransactionTestSuite) TestGetList() {
	data := "/+ABAgM="
	tID, err := t.txn.Save(&data)
	t.NoError(err)
	t.NotEmpty(tID)

	messages, signature, err := t.txn.GetList([]string{tID})
	t.NoError(err)
	t.Equal(data, messages[0])
	t.NotEmpty(signature)

	messages, signature, err = t.txn.GetList([]string{"InvalidID"})
	t.Error(err)
	t.Empty(messages)
	t.Equal(TransactionNotFound, err)
}
