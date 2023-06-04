package main

import (
	"sync"

	"orders_service/internal/db"
	"orders_service/internal/server"
	"orders_service/internal/subscribe"

	"github.com/nats-io/stan.go"
)

func main() {
	database := &db.DB{
		DBUser:     "seymour",
		DBPassword: "seymour",
		DBName:     "store",
	}

	// Connection to database
	store := db.Connect(database)
	defer store.Close()

	// Init cache for Orders
	cache := make([]db.Order, 0)
	mu := sync.Mutex{}

	// Upload data from database to cache
	db.GetCache(store, &cache)

	sub := &subscribe.Subscriber{
		ClusterID: "test-cluster",
		ClientID:  "reader",
		URL:       stan.DefaultNatsURL,
	}

	go sub.Subscribe(store, &cache, &mu)

	// Start http server
	server.StartServer(&cache, &mu)
}
