package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB boil.ContextExecutor
}
