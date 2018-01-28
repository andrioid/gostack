package graphql

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql/gqlerrors"

	"firebase.google.com/go/auth"
	"github.com/andrioid/gostack/module"
	"github.com/graphql-go/graphql"
)

const bearer = "Bearer"

// https://github.com/graphql-go/graphql/blob/master/examples/http/main.go

// TODO: We need to be able to verify if the users token is valid
// - If the token is invalid we reject the GraphQL query
// - If the token is valid, we populate a "me" data object in the graph
// - This should be cached until it expires
// - https://firebase.google.com/docs/reference/rest/auth/#section-sign-in-with-oauth-credential

/*
   Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig
*/
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "User",
		Description: "User is a doofus",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "some stuff that will explain other stuff",
			},
		},
	},
)

/*
   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var schema graphql.Schema

// HTTPHandler tells the http-server how to process GraphQL
func HTTPHandler(a *auth.Client) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var query string
		tokenString, _ := authorizationFromHeader(r)

		token, _ := a.VerifyIDToken(tokenString) // token being nil, is fine

		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			var t struct {
				Query string
			}
			err := decoder.Decode(&t)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			query = t.Query
		} else {
			query = r.URL.Query().Get("query")
		}
		if query == "" {
			fmt.Println("Query is empty")
		}
		result := ExecuteQuery(query, schema, token)
		w.Header().Set("Content-Type", "application/json")
		// CORS is annoying, but necessary
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "ContentType, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(result)
		return
	})
}

func authorizationFromHeader(req *http.Request) (string, error) {
	header := req.Header.Get("Authorization")
	if header == "" {
		return "", fmt.Errorf("Authorization header not found")
	}

	l := len(bearer)
	if len(header) > l+1 && header[:l] == bearer {
		return header[l+1:], nil
	}

	return "", fmt.Errorf("Authorization header format must be 'Bearer {token}'")
}

func createErrorResult(es string) *graphql.Result {
	gqlErrors := make(gqlerrors.FormattedErrors, 0)
	err := errors.New(es)
	gqlErrors = append(gqlErrors, gqlerrors.FormatError(err))
	return &graphql.Result{
		Data:   nil,
		Errors: gqlErrors,
	}
}

// ExecuteQuery does stuff
func ExecuteQuery(query string, schema graphql.Schema, token *auth.Token) *graphql.Result {
	if token == nil {
		return createErrorResult("Not logged in")
	}
	fmt.Println("token", token)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v\n%v\n", result.Errors, query)
	}
	return result
}

func CreateSchema(m []module.Module) error {
	var queryTypeFields, mutationTypeFields graphql.Fields
	queryTypeFields = make(graphql.Fields)
	mutationTypeFields = make(graphql.Fields)

	for _, mod := range m {
		queryFields, err := mod.QueryTypes()
		if err != nil {
			panic(err)
		}
		for k, v := range queryFields {
			queryTypeFields[k] = v
		}
		mutationFields, err := mod.MutationTypes()
		if err != nil {
			panic(err)
		}
		for k, v := range mutationFields {
			mutationTypeFields[k] = v
		}
	}
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: queryTypeFields,
		},
	)

	schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	// iterate over m
	// make a new map with all values
	// warn if duplicates
	// create schema
	// return schema
	return nil
}
