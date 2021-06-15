package main

import (
	"bufio"
	channel "cloud.google.com/go/channel/apiv1"
	"context"
	"fmt"
	"github.com/Hanekawa-chan/universal/protoc"
	"google.golang.org/grpc"
	"log"
	"os"
)

var sender string
var receiver string

func SignUp(ctx context.Context, in *protoc.User, opts ...grpc.CallOption) (*protoc.Response, error) {
	out := new(protoc.Response)
	err := c.cc.Invoke(ctx, "/protoc.Universal/signUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func SignIn(ctx context.Context, in *protoc.User, opts ...grpc.CallOption) (*protoc.Response, error) {
	out := new(protoc.Response)
	err := c.cc.Invoke(ctx, "/protoc.Universal/signIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func SignOut(ctx context.Context, in *protoc.User, opts ...grpc.CallOption) (*protoc.Response, error) {
	out := new(protoc.Response)
	err := c.cc.Invoke(ctx, "/protoc.Universal/signOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func sendMessage() {

}

func main() {
	var opts []grpc.DialOption
	conn, err := grpc.Dial("localhost:5565", opts...)
	if err != nil {
		log.Print(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Print(err)
		}
	}(conn)

	ctx := context.Background()
	client := protoc.NewUniversalClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	message := protoc.Message{}

	for {
		if sender != "" {
			for scanner.Scan() {
				if scanner.Text() == "/" {
					receiver = ""
					sender = ""
					break
				} else {
					go sendMessage(ctx, client, message)
				}
			}
		} else {
			fmt.Println("Enter channel name")
			for scanner.Scan() {
				channel.Name = scanner.Text()
				break
			}
			fmt.Println("Enter your name")
			for scanner.Scan() {
				channel.SendersName = scanner.Text()
				break
			}
			fmt.Println("Enter password for channel")
			for scanner.Scan() {
				channel.Password = scanner.Text()
				break
			}
			channelName = channel.Name
			senderName = channel.SendersName
			go joinChannel(ctx, client, &channel)
		}
	}
}
