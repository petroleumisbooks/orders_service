package main

import (
	"io/ioutil"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "pusher", stan.NatsURL("nats://localhost:4222"))

	dataFromFile, _ := ioutil.ReadFile("model.json")
	sc.Publish("foo1", dataFromFile)
}
