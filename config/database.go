package config

import (
	"github.com/goravel/framework/support/facades"
)

func init() {
	config := facades.Config
	config.Add("database", map[string]interface{}{
		//Default database connection name, only support Mysql now.
		"default": config.Env("DB_CONNECTION", "mysql"),

		//Database connections
		"connections": map[string]interface{}{
			"mysql": map[string]interface{}{
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "nft"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",
			},
		},

		//Migration Repository Table
		//This table keeps track of all the migrations that have already run for
		//your application. Using this information, we can determine which of
		//the migrations on disk haven't actually been run in the database.
		"migrations": "migrations",
	})
}
