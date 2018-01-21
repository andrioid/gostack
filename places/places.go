package places

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Module struct {
}

type Place struct {
	gorm.Model
	Name      string
	Latitude  float64
	Longitude float64
	// Remember to update this if you change the PlaceType
	PlaceTypes []PlaceType `gorm:"many2many:place_type_pairs;"`
}

type PlaceType struct {
	gorm.Model
	Name string
}

var placeSchemaType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Place",
	Description: "place description",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "some stuff that will explain other stuff",
		},
	},
})

var queryTypes = graphql.Fields{
	"place": &graphql.Field{
		Type: placeSchemaType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.ID,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			idQuery, isOK := p.Args["id"].(int64)
			if isOK {
				var pl Place
				db.First(&pl, idQuery)
				return pl, nil
			}
			return nil, nil
		},
	},
}

func New(db *gorm.DB) (*Module, error) {
	db = db
	db.AutoMigrate(&Place{}, &PlaceType{})
	return &Module{}, nil
}

func (m Module) Hello() {
	fmt.Println("Hello")
}

func MutationTypes() (graphql.Fields, error) {
	return nil, nil
}

func QueryTypes() (graphql.Fields, error) {
	return queryTypes, nil
}
