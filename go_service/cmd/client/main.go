package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"go_service/pkg/api"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("Not enough args!")
	}

	host := flag.Arg(0)

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("x = ")
	sx, _ := reader.ReadString('\n')
	x, err := strconv.Atoi(sx[:len(sx)-1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("y = ")
	sy, _ := reader.ReadString('\n')
	y, err := strconv.Atoi(sy[:len(sx)-1])
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
