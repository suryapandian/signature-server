package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"signatureserver/config"
	txn "signatureserver/transactions"
)

type signRouter struct {
	transaction *txn.Transaction
}

func newSignRouter() *signRouter {
	return &signRouter{
		transaction: txn.NewTransaction([]byte(config.PRIVATE_KEY)),
	}
}

func (s *signRouter) setSignRoutes(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Get("/public_key", s.getPublicKey)
		r.Put("/transaction", s.saveTransaction)
		r.Post("/signature", s.getListOftransactions)
	})
}

func (s *signRouter) getPublicKey(w http.ResponseWriter, r *http.Request) {
	var publicKeyResp struct {
		PublicKey []byte `json:"public_key"`
	}

	publicKeyResp.PublicKey = s.transaction.GetPublicKey()
	writeJSONStruct(publicKeyResp, http.StatusOK, w)
}

func (s *signRouter) saveTransaction(w http.ResponseWriter, r *http.Request) {

	var transactionReq struct {
		Transaction *string `json:"txn"`
	}

	err := json.NewDecoder(r.Body).Decode(&transactionReq)
	if err != nil {
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
	}

	var resp struct {
		ID string `json:"id"`
	}

	resp.ID, err = s.transaction.Save(transactionReq.Transaction)
	switch err {
	case nil:
		writeJSONStruct(resp, http.StatusOK, w)
	case txn.InvalidTransaction:
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
	default:
		writeJSONMessage(err.Error(), http.StatusInternalServerError, w)
	}
}

func (s *signRouter) getListOftransactions(w http.ResponseWriter, r *http.Request) {

	var transactionIDs struct {
		IDs []string `json:"ids"`
	}

	err := json.NewDecoder(r.Body).Decode(&transactionIDs)
	if err != nil {
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
	}

	var resp struct {
		Message   []string `json:"message"`
		Signature string   `json:"Signature"`
	}

	resp.Message, resp.Signature, err = s.transaction.GetList(transactionIDs.IDs)
	switch err {
	case nil:
		writeJSONStruct(resp, http.StatusOK, w)
	case txn.TransactionNotFound:
		writeJSONMessage(err.Error(), http.StatusNotFound, w)
	default:
		writeJSONMessage(err.Error(), http.StatusInternalServerError, w)
	}
}
