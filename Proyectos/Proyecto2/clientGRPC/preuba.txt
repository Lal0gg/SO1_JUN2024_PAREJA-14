package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "clientGRPC/client"
)

var (
	ctx     = context.Background()
	mu      sync.Mutex
	counter = make(map[string]int)
)

type Data struct {
	Texto string
	Pais  string
}

type CountryCount struct {
	Pais     string
	Contador int
}

func sendToRust(data *Data) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post("http://localhost:8000/set", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}

func sendCountryCountToRust(countryCount *CountryCount) {
	jsonData, err := json.Marshal(countryCount)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post("http://localhost:8000/country_count", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}

func sendData(c *fiber.Ctx) error {
	/* API REST */
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	tweet := Data{
		Texto: data["texto"],
		Pais:  data["pais"],
	}

	go sendGrpcServer(tweet)
	go sendToRust(&tweet)
	go updateCountryCount(tweet.Pais)

	return nil
}

func sendGrpcServer(tweet Data) {
	/* GRPC Client */
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Texto: tweet.Texto,
		Pais:  tweet.Pais,
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Respuesta del servidor ", ret)
	}
}

func updateCountryCount(country string) {
	mu.Lock()
	defer mu.Unlock()

	counter[country]++
	countryCount := CountryCount{
		Pais:     country,
		Contador: counter[country],
	}

	go sendCountryCountToRust(&countryCount)
}

func main() {
	app := fiber.New()

	app.Post("/insert", sendData)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
