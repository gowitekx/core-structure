package connection

import "github.com/jinzhu/gorm"

type ConnectionInterface interface {
	DBConnect() *gorm.DB
}
