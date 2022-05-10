package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/tachunwu/distpebble/pkg/proto/distpebble/v1"
	"github.com/tachunwu/distpebble/pkg/utli/throttle"
	"google.golang.org/grpc"
)

func main() {

	connP1, _ := grpc.Dial("localhost:30001", grpc.WithInsecure(), grpc.WithBlock())
	connP2, _ := grpc.Dial("localhost:30002", grpc.WithInsecure(), grpc.WithBlock())
	connP3, _ := grpc.Dial("localhost:30003", grpc.WithInsecure(), grpc.WithBlock())

	defer connP1.Close()
	defer connP2.Close()
	defer connP3.Close()
	c1 := pb.NewSequencerServiceClient(connP1)
	c2 := pb.NewSequencerServiceClient(connP2)
	c3 := pb.NewSequencerServiceClient(connP3)

	go func() {
		th := throttle.NewThrottle(128)

		for i := 0; i < 333333; i++ {
			fmt.Println(i)
			th.Do()
			go func(th *throttle.Throttle) {
				defer th.Done(nil)
				ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
				_, err := c1.CreateTxn(ctx, &pb.CreateTxnRequest{
					Txn: &pb.Txn{
						TxnId: uint64(i),
					},
				})
				if err != nil {
					log.Println(err)
				}
			}(th)
		}

	}()

	go func() {
		th := throttle.NewThrottle(128)

		for i := 333334; i < 666666; i++ {
			fmt.Println(i)
			th.Do()
			go func(th *throttle.Throttle) {
				defer th.Done(nil)
				ctx, _ := context.WithTimeout(context.Background(), time.Second)
				_, err := c2.CreateTxn(ctx, &pb.CreateTxnRequest{
					Txn: &pb.Txn{
						TxnId: uint64(i),
					},
				})
				if err != nil {
					log.Println(err)
				}

			}(th)
		}

	}()
	go func() {
		th := throttle.NewThrottle(128)

		for i := 666667; i < 1000000; i++ {
			fmt.Println(i)
			th.Do()
			go func(th *throttle.Throttle) {
				defer th.Done(nil)
				ctx, _ := context.WithTimeout(context.Background(), time.Second)
				_, err := c3.CreateTxn(ctx, &pb.CreateTxnRequest{
					Txn: &pb.Txn{
						TxnId: uint64(i),
					},
				})
				if err != nil {
					log.Println(err)
				}
			}(th)
		}

	}()
	time.Sleep(120 * time.Second)
}
