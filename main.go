package main

import (
	// "command-line-arguments/home/samson/Desktop/GG/jsonApiGrpc/types/client"
	// "context"
	"flag"
	// "fmt"
	// "log"
)

func main() {
	// client := client.New("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)

	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()

}
