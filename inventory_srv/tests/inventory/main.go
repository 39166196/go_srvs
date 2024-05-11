package main

import (
	"context"
	"fmt"
	"go_srvs/inventory_srv/proto"
	"google.golang.org/grpc"
	"sync"
)

var invClient proto.InventoryClient
var conn *grpc.ClientConn

func TestSetInv(goodsId, Num int32) {
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")

}

func TestInvDetail(goodsId int32) {
	rsp, err := invClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	//打印

	fmt.Println(rsp.Num)
}
func TestSell(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 1},
			//{GoodsId: 423, Num: 1},
		},
	})
	if err != nil {
		panic(err)
	}
	//打印

	fmt.Println("库存扣减成功")
}
func TestReback() {
	_, err := invClient.Reback(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 1},
			{GoodsId: 423, Num: 1},
		},
	})
	if err != nil {
		panic(err)
	}
	//打印

	fmt.Println("归还成功")
}

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	invClient = proto.NewInventoryClient(conn)
}
func main() {
	Init()
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go TestSell(&wg)
	}
	wg.Wait()
	//for i := 421; i < 840; i++ {
	//	TestSetInv(int32(i), 100)
	//}
	//TestSetInv(423, 100)
	//TestInvDetail(421)

	//TestReback()
	conn.Close()
}
