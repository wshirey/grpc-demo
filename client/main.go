package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/wshirey/grpc-demo/addresses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var cmd string
	var id string
	flag.StringVar(&cmd, "cmd", "c", "the command to run")
	flag.StringVar(&id, "id", "0", "the id of address to modify")
	flag.Parse()
	rand.Seed(time.Now().UnixMicro())

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := addresses.NewAddressesClient(conn)
	ctx := context.Background()

	switch cmd {
	case "c":
		addr := addresses.CreateAddressRequest{
			Street: fmt.Sprintf("%d Sesame Street", rand.Int31n(1000)),
			City:   "Birmingham",
			Zip:    uint32(rand.Intn(99999)),
		}
		if resp, err := c.CreateAddress(ctx, &addr); err != nil {
			log.Fatalf("could not create address: %v", err)
		} else {
			log.Println(resp)
		}
	case "g":
		if id == "" {
			log.Fatalf("must provide id with -id param")
		}
		if resp, err := c.GetAddress(ctx, &addresses.GetAddressRequest{Id: id}); err != nil {
			log.Fatalf("could not get address: %v", err)
		} else {
			log.Println(resp)
		}
	case "d":
		if id == "" {
			log.Fatalf("must provide id with -id param")
		}
		if resp, err := c.DeleteAddress(ctx, &addresses.DeleteAddressRequest{Id: id}); err != nil {
			log.Fatalf("could not delete address: %v", err)
		} else {
			log.Println(resp)
		}
	default:
		log.Fatalf("unknown command: %v", cmd)
	}
}
