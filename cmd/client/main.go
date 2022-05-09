package main

import (
	"context"
	"fmt"
	"log"
	"time"

	corev1 "github.com/tachunwu/distpebble/pkg/proto/core/v1"
	"google.golang.org/grpc"
)

func main() {

	connP1, _ := grpc.Dial("localhost:30001", grpc.WithInsecure(), grpc.WithBlock())
	connP2, _ := grpc.Dial("localhost:30002", grpc.WithInsecure(), grpc.WithBlock())
	connP3, _ := grpc.Dial("localhost:30003", grpc.WithInsecure(), grpc.WithBlock())

	defer connP1.Close()
	defer connP2.Close()
	defer connP3.Close()
	c1 := corev1.NewSequencerServiceClient(connP1)
	// c2 := corev1.NewSequencerServiceClient(connP2)
	// c3 := corev1.NewSequencerServiceClient(connP3)
	go func() {
		for i := 0; i < 333333; i++ {
			fmt.Println(i)
			go func() {
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				_, err := c1.Txn(ctx, &corev1.TxnRequest{
					Txn: &corev1.Txn{
						TxnId: uint64(i),
					},
				})
				if err != nil {
					log.Println(err)
				}
			}()

		}
	}()

	// go func() {
	// 	for i := 333334; i < 666666; i++ {
	// 		ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// 		_, err := c2.Txn(ctx, &corev1.TxnRequest{
	// 			Txn: &corev1.Txn{
	// 				TxnId: uint64(i),
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 	}
	// }()
	// go func() {
	// 	for i := 666667; i < 1000000; i++ {
	// 		ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// 		_, err := c3.Txn(ctx, &corev1.TxnRequest{
	// 			Txn: &corev1.Txn{
	// 				TxnId: uint64(i),
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 	}
	// }()
	time.Sleep(120 * time.Second)
}
