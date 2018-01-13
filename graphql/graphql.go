package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

// https://github.com/graphql-go/graphql/blob/master/examples/http/main.go

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data = map[string]user{
	"1": user{ID: "1", Name: "Cow"},
	"2": user{ID: "2", Name: "MOoOo"},
}

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
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

// HTTPHandler tells the http-server how to process GraphQL
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	var query string
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
	result := ExecuteQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

// ExecuteQuery does stuff
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v\n%v\n", result.Errors, query)
	}
	return result
}

func GetSchema() graphql.Schema {
	return schema
}
