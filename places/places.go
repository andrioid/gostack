package places

import (
	"fmt"

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
	PlaceTypes []PlaceType `gorm:"many2many:typesOfPlaces;"`
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
