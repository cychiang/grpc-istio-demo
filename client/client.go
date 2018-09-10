package main

import (
	"context"
	"github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"log"
	"time"
)

func initGrpcClientPool() (*grpcpool.Pool, error) {
	var factory grpcpool.Factory
	factory = func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to start gRPC connection: %v", err)
		}
		return conn, err
	}
	pool, err := grpcpool.New(
		factory,
		10,             /* initial connections */
		50,             /* maximum connections */
		10*time.Second, /* idle timeout        */
		time.Minute,    /* maxLifeDuration     */
	)
	if err != nil {
		log.Printf("Failed to create gRPC pool: %v", err)
		return nil, err
	} else {
		return pool, nil
	}
}

func main() {
	// create grpc pool
	pool, err := initGrpcClientPool()
	if err != nil {
		log.Fatalf("Failed to connect gRPC server: %v", err)
	}
	defer pool.Close()
	// init context, it's good to set an timeout
	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()
	defer ctx.Done()
}
