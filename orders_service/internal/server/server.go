package server

import (
	"log"
	"net/http"
	"sync"
	"text/template"

	"orders_service/internal/db"

	"github.com/gorilla/mux"
)

type Server struct {
	IP   string
	Port string
}

func StartServer(cache *[]db.Order, mu *sync.Mutex) {
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/start.html"))
		tmpl.Execute(writer, nil)
	})

	r.HandleFunc("/order", func(writer http.ResponseWriter, request *http.Request) {
		orderUid := request.FormValue("order_uid")
		order := db.Order{}

		mu.Lock()
		for _, tempOrder := range *cache {
			if tempOrder.OrderUid == orderUid {
				order = tempOrder
				break
			}
		}
		mu.Unlock()

		if order.OrderUid == "" {
			tempStruct := struct {
				UidInput string
			}{
				UidInput: orderUid,
			}
			tmpl := template.Must(template.ParseFiles("templates/error.html"))
			tmpl.Execute(writer, tempStruct)
		} else {
			tmpl := template.Must(template.ParseFiles("templates/order.html"))
			tmpl.Execute(writer, order)
		}
	})

	http.Handle("/", r)

	log.Println("Succsess start server")
	if err := http.ListenAndServe("127.0.0.1:8888", nil); err != nil {
		log.Fatalf("Server failed to start. Error: %s\n", err)
	}
}
