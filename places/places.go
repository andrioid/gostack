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
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			idQuery, isOK := p.Args["id"].(int)
			fmt.Println("query id", idQuery)
			if isOK {
				var pl Place
				db.First(&pl, idQuery)
				fmt.Println("found place", pl)
				return pl, nil
			}
			fmt.Println("resolver not ok")
			return nil, nil
		},
	},
}

func New(db *gorm.DB) (*Module, error) {
	db = db
	db.AutoMigrate(&Place{}, &PlaceType{})
	bleh := Place{
		Name: "test place",
	}
	db.Create(&bleh)
	return &Module{}, nil
}

func (m Module) Hello() {
	fmt.Println("Hello")
}

func MutationTypes() (graphql.Fields, error) {
	return nil, nil
}

func (m Module) QueryTypes() (graphql.Fields, error) {
	return queryTypes, nil
}
