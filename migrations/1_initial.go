package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table log...")
		_, err := db.Exec(`CREATE TABLE logs(
			ID SERIAL PRIMARY KEY,
			CLIENT_IP VARCHAR(50),
			SERVER_IP VARCHAR(50),
			TAGS JSONB,
			MSG TEXT,
			CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table log...")
		_, err := db.Exec(`DROP TABLE log`)
		return err
	})
}
