package main

import (
	"context"
	"fmt"
	"go_srvs/goods_srv/proto"
	"google.golang.org/grpc"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetGoodsList() {
	rsp, err := brandClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 130361,
		PriceMin:    90,
		//KeyWords:    "深海速冻",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, value := range rsp.Data {
		fmt.Println(value.Name, value.ShopPrice)

	}

}
func TestbatchGetGoodsList() {
	rsp, err := brandClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{421, 422, 423, 424, 425},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, value := range rsp.Data {
		fmt.Println(value.Name, value.ShopPrice)

	}

}
func TestGetGoodsDetail() {
	rsp, err := brandClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: 421,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Name)

}
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}
func main() {
	Init()
	//TestCreateUser()
	//TestGetCategoryList()
	TestGetGoodsDetail()
	conn.Close()
}
