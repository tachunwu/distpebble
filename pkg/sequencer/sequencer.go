package sequencer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tachunwu/distpebble/pkg/config"
	pb "github.com/tachunwu/distpebble/pkg/proto/distpebble/v1"
	"github.com/tachunwu/distpebble/pkg/utli/throttle"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/proto"
)

type Sequencer struct {
	pb.UnimplementedSequencerServiceServer

	SeqCluster     string
	SeqClusterConn *nats.Conn
}

func NewSequencer(conf *config.Config) *Sequencer {
	return &Sequencer{
		SeqCluster: conf.ClusterAddr,
	}
}

func (s *Sequencer) Start() {
	// Connect local log
	nc, _ := nats.Connect(s.SeqCluster)
	s.SeqClusterConn = nc
	// Create sequencer stream
	js, _ := nc.JetStream()
	js.AddStream(&nats.StreamConfig{
		Name:     "Sequencer",
		Subjects: []string{"Sequencer.*"},
	})
	// Start sequencer ordering alogorithm
	go s.GetOrderedTxn()
}

func (s *Sequencer) CreateTxn(ctx context.Context, req *pb.CreateTxnRequest) (*pb.CreateTxnResponse, error) {
	// Enqueue to in-memoery NATS
	txn, _ := proto.Marshal(req.GetTxn())

	// JetStream
	// txnQueue, _ := s.SeqClusterConn.JetStream()
	// txnQueue.Publish("Sequencer.*", txn)
	// fmt.Println("Enqueue txn to JetStream: ", req.Txn.TxnId)

	// NATS
	s.SeqClusterConn.Publish("Sequencer", txn)
	fmt.Println("Enqueue txn to NATS: ", req.Txn.TxnId)

	return &pb.CreateTxnResponse{}, nil
}

func (s *Sequencer) GetOrderedTxn() {

	// Impl 1: Use JetStream sync sub
	// conn, _ := s.SeqClusterConn.JetStream()
	// sub, _ := conn.SubscribeSync("Sequencer.*")

	// Impl 2: Use JetStream pull sub
	// sub, _ := conn.PullSubscribe("Sequencer.*", "MONITOR")

	// Impl 3: Use NATS
	ch := make(chan *nats.Msg, 256)
	s.SeqClusterConn.ChanSubscribe("Sequencer", ch)

	// Singleflight
	var group singleflight.Group
	th := throttle.NewThrottle(512)
	for {
		// Impl 1: Use JetStream sync sub
		// m, err := sub.NextMsg(7 * time.Second)
		// if err != nil {
		// 	log.Println(err)
		// } else {
		// 	txn := &pb.Txn{}
		// 	proto.Unmarshal(m.Data, txn)
		// 	fmt.Println("Ready to process txn: ", txn)
		// }

		// Impl 2: Use JetStream pull sub
		// msgs, err := sub.Fetch(4096)
		// if err != nil {
		// 	log.Println(err)
		// } else {
		// 	txn := &pb.Txn{}
		// 	for i := range msgs {
		// 		proto.Unmarshal(msgs[i].Data, txn)
		// 		fmt.Println("Ready to process txn: ", txn)
		// 	}
		// }

		// Impl 3: NATS in-memory

		txn := &pb.Txn{}
		msg := <-ch
		proto.Unmarshal(msg.Data, txn)
		th.Do()
		go func() {
			defer th.Done(nil)
			group.Do(strconv.FormatUint(txn.TxnId, 10), func() (interface{}, error) {
				fmt.Println("Ready to process txn: ", txn)
				time.Sleep(1 * time.Millisecond)
				return nil, nil
			})
		}()
	}
}
