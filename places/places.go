package places

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Module struct {
	db *gorm.DB
}

type PlaceType int64

const (
	PLACE_TYPE_PARKING_LOT PlaceType = iota
)

var types = [...]string {
	"Undefined",
	""
}

type Place struct {
	gorm.Model
	Name      string
	Latitude  float64
	Longitude float64
}

func New(db *gorm.DB) (*Module, error) {
	db.AutoMigrate(&Place{})
	return &Module{db: db}, nil
}

func (m Module) Hello() {
	fmt.Println("Hello")
}
