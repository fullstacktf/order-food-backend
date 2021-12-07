package database

import (
	"comiditapp/api/env"
	"context"
	"time"
)

func (db *DB) DropDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.Client.Database(env.DB_NAME).Drop(ctx)
}
