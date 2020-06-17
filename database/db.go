package database

import (
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)
var (	pool *pgx.ConnPool
)
func Connect(dbConfig string) {
	connectionString, err := pgx.ParseConnectionString(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	config := pgx.ConnPoolConfig{
		ConnConfig: connectionString,
	}
	pool, err = pgx.NewConnPool(config)
	if err != nil {
		log.Panic(err)
	}

	_, err = pool.Exec(DDLUsers)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pool.Exec(DDLNotifications)
	if err != nil {
		log.Fatal(err)
	}
}
