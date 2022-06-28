package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"net"
	"os"
	pb "recsysProxyCacheMock/github.com/cjmcgraw/recsys-proxy-cache"

	farm "github.com/dgryski/go-farm"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
	pb.UnimplementedRecsysProxyCacheServer
}

func (s *server) GetScores(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResponse, error) {
	log.Printf("Received: %s", in)

	// initialize our buffer that will be used for each score
	buf := new(bytes.Buffer)

	// get context bytes and then write to buffer
	contextBytes, err := proto.Marshal(in.Context)
	if err != nil {
		log.Fatalln("Failed to turn request in bytes")
	}
	log.Printf("context bytes: %v", contextBytes)

	// tracking the context byte size allows us to truncate each item to save performance
	contextByteSize, err := buf.Write(contextBytes)
	if err != nil {
		log.Fatalln("failed to turn context into bytes for unknown reason!")
	}

	log.Printf("generating scores for %d items", len(in.GetItems()))
	scores := make([]float32, 0, len(in.GetItems()))
	for _, item := range in.GetItems() {

		// first we need to truncate to byte size
		buf.Truncate(contextByteSize)

		// next we extract the item into our binary buffer
		err = binary.Write(buf, binary.LittleEndian, item)
		if err != nil {
			log.Fatalf("failed to turn item into bytes!")
		}

		// hash it to a uint64
		hash := farm.Fingerprint64(buf.Bytes())

		// then we normalize it between 0 and 1
		score := float64(hash) / float64(math.MaxUint64)
		log.Printf("item:%v hash: %v score: %v", item, hash, score)
		scores = append(scores, float32(score))
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
