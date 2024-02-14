package main

import (
	"context"
	"flag"
	"fmt"
	"log"
)

func main() {
	//test for client
	//client := client.New("http://localhost:3000")
	//
	//price, err := client.FetchPrice(context.Background(), "ETH")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%+v\n", price)
	//
	//return

	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))
	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()

	price, err := svc.FetchPrice(context.Background(), "eth")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(price)
}
