package main

import (
	"context"
	"fmt"
	"go_srvs/order_srv/proto"
	"google.golang.org/grpc"
)

var orderClient proto.OrderClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	orderClient = proto.NewOrderClient(conn)
}
func TestSetInv(userId, nums, goodsId int32) {
	rsp, err := orderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		GoodsId: goodsId,
		Nums:    nums,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)

}

func TestCARTitemlist(userId int32) {
	rsp, err := orderClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: userId,
	})
	if err != nil {
		panic(err)
	}
	for _, itiem := range rsp.Data {
		fmt.Println(itiem.Id, itiem.GoodsId, itiem.Nums)
	}

}
func TestUpdate(id int32) {
	_, err := orderClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{
		Id:      id,
		Checked: true,
	})
	if err != nil {
		panic(err)
	}
}
func testcreateorder() {
	_, err := orderClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  54,
		Address: "四川神",
		Name:    "张三",
		Mobile:  "123456",
		Post:    "不要辣椒",
	})
	if err != nil {
		panic(err)
	}
}
func tgod(orderid int32) {
	rsp, err := orderClient.OrderDetail(context.Background(), &proto.OrderRequest{
		Id: orderid,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.OrderInfo.OrderSn)
	for _, good := range rsp.Goods {
		fmt.Println(good.GoodsId, good.GoodsName, good.GoodsPrice, good.Nums)
	}

}
func testorderlist() {
	rsp, err := orderClient.OrderList(context.Background(), &proto.OrderFilterRequest{
		UserId: 1,
	})
	if err != nil {
		panic(err)
	}
	for _, order := range rsp.Data {
		fmt.Println(order.OrderSn)
	}
}
func main() {
	Init()
	TestSetInv(54, 8, 424)
	//TestCARTitemlist(54)
	//TestUpdate(7)
	//testcreateorder()
	//tgod(2)
	//testorderlist()
	conn.Close()

}
