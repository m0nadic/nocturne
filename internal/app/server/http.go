package server

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nocturne/internal/app/router"
)

func InitHttpServer(host string, port int, dbPath string) error {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	return router.NewRouter(db).Init().Run(addr)
}
