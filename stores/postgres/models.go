// package Nadellain
package stores

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Client ScheNadella
type Clients struct {
	gorm.Model
	Name      string
	Equation  string
	Timestamp time.Time
}
