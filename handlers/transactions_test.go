package handlers

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPublicKey(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/public_key", nil)
	w := httptest.NewRecorder()
	GetRouter().ServeHTTP(w, r)
	response := w.Result()
	assert.New(t).Equal(http.StatusOK, response.StatusCode)
}

func TestSaveTransaction(t *testing.T) {
	var testCases = []struct {
		desc               string
		payload            []byte
		expectedStatusCode int
	}{
		{
			"no payload",
			nil,
			http.StatusBadRequest,
		},
		{
			"invalid json",
			[]byte(`{"key":}`),
			http.StatusBadRequest,
		},
		{
			"valid payload",
			[]byte(fmt.Sprintf(`{
				"txn": "%s"
			}`, "transaction data")),
			http.StatusOK,
		},
		{
			"valid payload",
			[]byte(fmt.Sprintf(`{
				"txn": "%s"
			}`, ``)),
			http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			a := assert.New(t)
			r := httptest.NewRequest(http.MethodPut, "/transaction", bytes.NewReader(testCase.payload))
			w := httptest.NewRecorder()
			GetRouter().ServeHTTP(w, r)
			response := w.Result()
			a.Equal(testCase.expectedStatusCode, response.StatusCode)
		})
	}

}

func TestGetListOfTransactions(t *testing.T) {
	var testCases = []struct {
		desc               string
		payload            []byte
		expectedStatusCode int
	}{
		{
			"no payload",
			nil,
			http.StatusBadRequest,
		},
		{
			"invalid json",
			[]byte(`{"key":}`),
			http.StatusBadRequest,
		},
		{
			"valid payload",
			[]byte(fmt.Sprintf(`{
				 "ids": ["b23d6390-0d8c-4942-920f-f8f6c7635370"]
			}`)),
			http.StatusNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			a := assert.New(t)
			r := httptest.NewRequest(http.MethodPost, "/signature", bytes.NewReader(testCase.payload))
			w := httptest.NewRecorder()
			GetRouter().ServeHTTP(w, r)
			response := w.Result()
			a.Equal(testCase.expectedStatusCode, response.StatusCode)
		})
	}

}
