package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.11.130:9876"}),
		consumer.WithGroupName("testGroup"),
	)
	err := c.Subscribe("qinjun1", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("msgId:%s, body:%s\n", msgs[i], string(msgs[i].Body))
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic("订阅失败")
	}
	_ = c.Start()
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
