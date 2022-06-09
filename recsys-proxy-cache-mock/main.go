package main

import (
	"context"
	"flag"
	"log"
    "io/ioutil"
	"math/rand"
	"net"
    "os"
	pb "recsysProxyCacheMock/github.com/cjmcgraw/recsys-proxy-cache"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRecsysProxyCacheServer
}

func (s *server) GetScores(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResponse, error) {
	log.Printf("Received: %s", in)
	scores := make([]float32, 0, len(in.GetItems()))
	log.Printf("generating scores for %d items", len(in.GetItems()))

	log.Print("generating random scores")
	for item := range in.GetItems() {
		score := rand.Float32()
		log.Printf("item:%v score: %v", item, score)
		scores = append(scores, score)
	}

	response := &pb.ScoreResponse{Scores: scores}
	log.Printf("Response: %s", response)
	return response, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRecsysProxyCacheServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

    if len(os.Getenv("LOGGING")) <= 0 {
        log.Printf("set LOGGING=1 to enable logging")
        log.SetOutput(ioutil.Discard)
    }

	if err := s.Serve(lis); err != nil {
        log.SetOutput(os.Stderr)
		log.Fatalf("failed to serve: %v", err)
	}
}
