package places

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Module struct {
	db *gorm.DB
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

func New(db *gorm.DB) (*Module, error) {
	db.AutoMigrate(&Place{}, &PlaceType{})
	return &Module{db: db}, nil
}

func (m Module) Hello() {
	fmt.Println("Hello")
}

func MutationTypes() (graphql.Fields, error) {
	return nil, nil
}

func QueryTypes() (graphql.Fields, error) {
	return nil, nil
}
