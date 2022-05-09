package sequencer

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tachunwu/distpebble/pkg/config"
	corev1 "github.com/tachunwu/distpebble/pkg/proto/core/v1"
	"google.golang.org/protobuf/proto"
)

type Sequencer struct {
	corev1.UnimplementedSequencerServiceServer

	SequencerAddr        string
	SequencerCluster     []string
	LocalTxnLog          *nats.Conn
	SequencerClusterConn map[string]*nats.Conn
	TxnOrderedCh         map[string]chan *nats.Msg
}

func NewSequencer(conf *config.Config) *Sequencer {

	return &Sequencer{
		SequencerAddr:        conf.SequencerAddr,
		SequencerCluster:     conf.ClusterAddr,
		SequencerClusterConn: make(map[string]*nats.Conn),
		TxnOrderedCh:         make(map[string]chan *nats.Msg),
	}
}

func (s *Sequencer) Start() {
	// Connect local log
	localTxnLog, _ := nats.Connect(s.SequencerAddr)

	// Connect cluster log
	s.LocalTxnLog = localTxnLog
	for _, m := range s.SequencerCluster {
		conn, _ := nats.Connect(m)
		s.SequencerClusterConn[m] = conn
	}

	// Start sequencer ordering alogorithm
	s.GlobalTxnOrdered()
}

func (s *Sequencer) Txn(ctx context.Context, req *corev1.TxnRequest) (*corev1.TxnResponse, error) {
	// Enqueue to in-memoery NATS
	txn, _ := proto.Marshal(req.GetTxn())
	s.LocalTxnLog.Publish("Sequencer", txn)

	// Print enqueue txn
	fmt.Println(req.Txn.TxnId)
	return &corev1.TxnResponse{}, nil
}

func (s *Sequencer) GlobalTxnOrdered() {
	// Subscribe round-robin from NATS
	go func() {
		for {
			for partition, conn := range s.SequencerClusterConn {
				conn.ChanSubscribe(partition, s.TxnOrderedCh[partition])
			}
		}
	}()

	// Publish to self-partition's persist stream
	go func() {
		js, _ := s.LocalTxnLog.JetStream()
		for {
			for _, ch := range s.TxnOrderedCh {
				txn := <-ch
				js.Publish("GlobalOrderedTxn", txn.Data)
			}
		}
	}()
}
