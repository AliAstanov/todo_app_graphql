package graph

import (
	"example/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Storage storage.StorageI
}

func NewResolwer(storage storage.StorageI)*Resolver{
	return &Resolver{Storage: storage}
}
