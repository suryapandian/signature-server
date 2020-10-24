package server

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"signatureserver/encryptor"
)

type Transaction struct {
	encryptor    encryptor.Encryptor
	transactions transactionStore
}

func NewTransaction(privateKeyPem []byte) *Transaction {
	return &Transaction{
		encryptor:    encryptor.NewEd25519Encrytor(privateKeyPem),
		transactions: &mapStore{},
	}
}

func (t *Transaction) GetPublicKey() []byte {
	return t.encryptor.GetPublicKey()
}

var InvalidTransaction = errors.New("Empty data!")

func (t *Transaction) Save(data *string) (string, error) {
	if data == nil {
		return "", InvalidTransaction
	}

	id := uuid.NewV4().String()
	t.transactions.save(id, *data)

	return id, nil
}

var TransactionNotFound = errors.New("Transaction with the given ID not found")

func (t *Transaction) GetList(ids []string) (messages []string, signature string, err error) {

	for _, id := range ids {
		msg, ok := t.transactions.get(id)
		if !ok {
			err = TransactionNotFound
			return
		}

		messages = append(messages, msg)
	}

	messagesBytes, err := json.Marshal(messages)
	if err != nil {
		return
	}

	signature = base64.StdEncoding.EncodeToString(t.encryptor.Sign(messagesBytes))
	return
}
