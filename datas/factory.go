package datas

import "github.com/attic-labs/noms/chunks"

// Factory allows the creation of namespaced DataStore instances. The details of how namespaces are separated is left up to the particular implementation of Factory and DataStore.
type Factory interface {
	Create(string) (DataStore, bool)

	// Shutter shuts down the factory. Subsequent calls to Create() will fail.
	Shutter()
}

type localFactory struct {
	cf chunks.Factory
}

func (lf *localFactory) Create(ns string) (DataStore, bool) {
	if cs := lf.cf.CreateStore(ns); cs != nil {
		return newLocalDataStore(cs), true
	}
	return &LocalDataStore{}, false
}

func (lf *localFactory) Shutter() {
	lf.cf.Shutter()
}
