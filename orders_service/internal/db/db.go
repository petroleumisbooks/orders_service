package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DBUser     string
	DBPassword string
	DBName     string
}

func Connect(db *DB) *sqlx.DB {
	log.Println("Connecting to database")

	store, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			db.DBUser, db.DBPassword, db.DBName))

	if err != nil {
		log.Fatalf("Can't connect to database. Error: %s\n", err)
	}

	log.Println("Success connect to DB")

	return store
}

// Get cache from store
func GetCache(store *sqlx.DB, cache *[]Order) {
	var err error

	err = store.Select(cache, "SELECT * FROM orders")

	if err != nil {
		log.Fatalf("Can't select data from orders. Error: %s", err)
		// log.Fatal(err.Error())
	}

	for i, order := range *cache {
		err = store.Get((*cache)[i].Delivery, "SELECT * FROM delivery WHERE id = $1", order.DeliveryId)

		if err != nil {
			log.Fatalf("Can't select data from delivery")
		}

		err = store.Get((*cache)[i].Payment, "SELECT * FROM payment WHERE transaction = $1", order.PaymentId)

		if err != nil {
			log.Fatalf("Can't select data from payment")
		}

		err = store.Get((*cache)[i].Items, "SELECT * FROM item WHERE order_uid = $1", order.OrderUid)

		if err != nil {
			log.Fatalf("Can't select data from item")
		}
	}
}

// Add Data to store, return id for added delivery
func AddData(store *sqlx.DB, tx *sqlx.Tx, delivery Delivery, payment Payment, items []Item, order Order) (id int) {
	st, err := store.PrepareNamed("INSERT INTO delivery (name, phone, zip, city, address, region, email) " +
		"VALUES (:name, :phone, :zip, :city, :address, :region, :email) RETURNING id")

	if err != nil {
		log.Printf("Can't add data to delivery. Error: %s", err)
	}

	st.Get(&id, &delivery)

	_, err = tx.NamedExec("INSERT INTO payment (transaction, request_id, currency, provider, "+
		"amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES (:transaction, :request_id, "+
		":currency, :provider, :amount, :payment_dt, :bank, :delivery_cost, :goods_total, :custom_fee)", &payment)

	if err != nil {
		log.Printf("Can't add data to payment. Error: %s", err)
	}

	for _, item := range items {
		_, err = tx.NamedExec("INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, "+
			"total_price, nm_id, brand, status, order_uid) VALUES (:chrt_id, :track_number, :price, :rid, :name, :sale, :size, "+
			":total_price, :nm_id, :brand, :status, :order_uid)", &item)
		if err != nil {
			log.Printf("Can't add data to payment. Error: %s", err)
		}
	}

	_, err = tx.NamedExec("INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, "+
		"locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) "+
		"VALUES (:order_uid, :track_number, :entry, :delivery_id, :payment_id, :locale, :internal_signature, "+
		":customer_id, :delivery_service, :shardkey, :sm_id, :date_created, :oof_shard)", &order)
	if err != nil {
		log.Printf("Can't add data to payment. Error: %s", err)
	}

	return id
}
