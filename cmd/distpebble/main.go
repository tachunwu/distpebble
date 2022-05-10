package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/tachunwu/distpebble/pkg/config"
	pb "github.com/tachunwu/distpebble/pkg/proto/distpebble/v1"
	"github.com/tachunwu/distpebble/pkg/sequencer"
	"google.golang.org/grpc"
)

var (
	port    = flag.String("port", "", "distpebble serve port")
	cluster = flag.String("cluster", "", "cluster ports")
)

func main() {
	// Config
	flag.Parse()
	conf := config.NewDefaultConfig()
	conf.ClusterAddr = *cluster
	grpcServer := grpc.NewServer()

	// Regist sequencer
	sequencer := sequencer.NewSequencer(conf)
	sequencer.Start()
	pb.RegisterSequencerServiceServer(grpcServer, sequencer)

	// Start grpc server
	l, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	// Listen singnals
	handleSignal(grpcServer)

	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSignal(grpcServer *grpc.Server) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<-sigCh
		grpcServer.Stop()
	}()
}
