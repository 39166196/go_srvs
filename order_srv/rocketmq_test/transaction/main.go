package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type OrderListener struct{}

func (o *OrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("执行本地事务")
	time.Sleep(time.Second * 3)
	fmt.Println("执行本地事务失败")
	return primitive.UnknowState
}
func (o *OrderListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("回查事务状态")
	time.Sleep(time.Second * 15)
	return primitive.CommitMessageState
}
func main() {
	p, err := rocketmq.NewTransactionProducer(
		&OrderListener{},
		producer.WithNameServer([]string{"192.168.11.130:9876"}),
	)
	if err != nil {
		panic("生成producer失败")
	}
	if err = p.Start(); err != nil {
		panic("启动producer失败")
	}
	res, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("TransTopic", []byte("hello TransTopic  msg2")))
	if err != nil {
		fmt.Println("发送事务消息失败:", err)
	} else {
		fmt.Println("发送事务消息成功", res.String())
	}
	time.Sleep(time.Hour)
	if err := p.Shutdown(); err != nil {
		panic("关闭producer失败")
	}
}
