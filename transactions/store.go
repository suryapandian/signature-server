package server

import "sync"

// An abstraction to store and retrive the transactions
type transactionStore interface {
	save(string, string)
	get(string) (string, bool)
}

type mapStore struct {
	transactions sync.Map
}

func (ms *mapStore) save(id, data string) {
	ms.transactions.Store(id, data)
}

func (ms *mapStore) get(id string) (string, bool) {
	result, ok := ms.transactions.Load(id)
	if !ok {
		return "", ok
	}
	return result.(string), ok
}
