package main

import (
	"log"
	. "github.com/aerospike/aerospike-client-go"
)

func main() {

	client, err := NewClient("localhost", 3000)
	if err != nil {
		log.Fatal(err)
	}

	key, err := NewKey("test", "demo",
		"")
	if err != nil {
		log.Fatal(err)
	}

	bin1 := NewBin("bin1", "value1")
	bin2 := NewBin("bin2", "value2")

	// Write a record
	err = client.PutBins(nil, key, bin1, bin2)
	if err != nil {
		log.Fatal(err)
	}

	// Read a record
	record, err := client.Get(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	println(record)

	client.Close()
}
