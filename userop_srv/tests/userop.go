package main

import (
	"context"
	"go_srvs/userop_srv/proto"
	"google.golang.org/grpc"
)

var userFavClient proto.UserFavClient
var messageClient proto.MessageClient
var addressClient proto.AddressClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50054", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userFavClient = proto.NewUserFavClient(conn)
	messageClient = proto.NewMessageClient(conn)
	addressClient = proto.NewAddressClient(conn)
}

func testAddress() {
	_, err := addressClient.GetAddressList(context.Background(), &proto.AddressRequest{
		UserId: 54,
	})
	if err != nil {
		panic(err)

	}
}
func testMessage() {
	_, err := messageClient.MessageList(context.Background(), &proto.MessageRequest{
		UserId: 54,
	})
	if err != nil {
		panic(err)

	}
}
func testuserfav() {
	_, err := userFavClient.GetFavList(context.Background(), &proto.UserFavRequest{
		UserId: 54,
	})
	if err != nil {
		panic(err)

	}
}
func main() {
	Init()
	testAddress()
	testMessage()
	testuserfav()
	conn.Close()
}
