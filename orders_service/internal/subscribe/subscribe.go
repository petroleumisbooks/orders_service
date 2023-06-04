package subscribe

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"orders_service/internal/db"
)

type Subscriber struct {
	ClusterID string
	ClientID  string
	URL       string
}

// Subsctibe to streaming
func (s *Subscriber) Subscribe(store *sqlx.DB, cache *[]db.Order, mu *sync.Mutex) {
	sc, err := stan.Connect(s.ClusterID, s.ClientID, stan.NatsURL(s.URL))

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Succsess connect")
	}
	defer sc.Close()

	_, err = sc.Subscribe("foo1", func(m *stan.Msg) {
		order := db.Order{}
		err := json.Unmarshal(m.Data, &order)

		if err != nil {
			log.Printf("Invalid data from streaming. Error: %s", err)
		} else {
			log.Println("Successful data acquisition")
		}

		for i := range order.Items {
			order.Items[i].OrderUid = order.OrderUid
		}

		tx := store.MustBegin()

		deliveryId := db.AddData(store, tx, order.Delivery, order.Payment, order.Items, order)
		order.Delivery.Id = deliveryId
		order.DeliveryId = deliveryId
		order.PaymentId = order.Payment.Transaction

		if err := tx.Commit(); err != nil {
			log.Printf("Can't commit the transaction. Error: %s", err)
		} else {
			log.Println("Data successfully added to the database")
		}

		mu.Lock()
		*cache = append(*cache, order)
		mu.Unlock()
		log.Println("Message data successfully added to the memory")
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Println(err)
	}
}
