// Package module is an interface to domain packages, and accompanied utils. Sorry for the generic name.
// - GraphQL: Module can add types and additions to the query and mutation type
// - HTTP: Module can add routes and middleware
// Inspired by: https://gist.github.com/abdullin/3e3fd199674255e4d206

package module

import (
	"github.com/graphql-go/graphql"
)

type Module interface {
	MutationTypes() (mutations graphql.Fields, err error) // GraphQL mutation to expose
	QueryTypes() (queries graphql.Fields, err error)      // GraphQL Queries to expose
	// Routes() string // Which HTTP routes to expose, maybe layer
}

// Takes a collection of modules and returns a GraphQL schema
func CreateSchema(m []Module) (graphql.Schema, error) {
	return graphql.Schema{}, nil
}
