package places

import "github.com/jinzhu/gorm"

type Module struct {
	db *gorm.DB
}

func (p Module) New(db *gorm.DB) (*Module, error) {
	return nil, nil
}
