package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "clientGRPC/client"
)

var ctx = context.Background()

type Data struct {
	Texto string `json:"texto"`
	Pais  string `json:"pais"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sendData", sendData).Methods("POST")
	fmt.Println("Server is running on port 3000...")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	log.Fatal(http.ListenAndServe(":3000", corsHandler))
}

func sendData(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tweet := map[string]string{
		"texto": data.Texto,
		"pais":  data.Pais,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweet)

	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	cl := pb.NewGetInfoClient(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Texto: tweet["texto"],
		Pais:  tweet["pais"],
	})

	if err != nil {
		log.Fatalf("could not get info: %v", err)
		return
	}

	log.Printf("Received from gRPC server: %v", ret)
}
