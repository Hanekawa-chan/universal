package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Hanekawa-chan/universal/protoc"
	"google.golang.org/grpc"
	"log"
	"os"
)

var sender string
var receiver string

func SignUp(ctx context.Context, in *protoc.User, client protoc.UniversalClient) (*protoc.Response, error) {
	response, _ := client.SignUp(ctx, in)
	return response, nil
}

func SignIn(ctx context.Context, in *protoc.User, client protoc.UniversalClient) (*protoc.Response, error) {
	response, _ := client.SignIn(ctx, in)
	return response, nil
}

func SignOut(ctx context.Context, in *protoc.User, client protoc.UniversalClient) (*protoc.Response, error) {
	response, _ := client.SignOut(ctx, in)
	return response, nil
}

func ReceiveMessage(ctx context.Context, client protoc.UniversalClient, user *protoc.User) {
	stream, err := client.ReceiveMessage(ctx, user)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}

}

func SendMessage(ctx context.Context, client protoc.UniversalClient, message string) {
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}

	msg := protoc.Message{
		Receiver: receiver,
		Text:     message,
		Sender:   sender,
	}
	err = stream.Send(&msg)
	if err != nil {
		return
	}

	ack, err := stream.CloseAndRecv()
	fmt.Printf("Message sent: %v \n", ack)
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

	for {
		if sender != "" {
			if receiver != "" {
				for scanner.Scan() {
					s := scanner.Text()
					switch s {
					case "/signout":
						out, _ := SignOut(ctx, &protoc.User{Name: sender})
						if out.Response == "OK" {
							sender = ""
							receiver = ""
						}
					case "/connect":
						receiver = scanner.Text()
					default:
						SendMessage(ctx, client, s)
					}
				}
			} else {
				for scanner.Scan() {
					s := scanner.Text()
					if s == "/connect" {
						receiver = scanner.Text()
					}
				}
			}
		} else {
			for scanner.Scan() {
				s := scanner.Text()
				switch s {
				case "/signin":
					name := scanner.Text()
					password := scanner.Text()
					in, _ := SignIn(ctx, &protoc.User{Name: name, Password: password})
					if in.Response == "OK" {
						sender = name
					}
				case "/signup":
					name := scanner.Text()
					password := scanner.Text()
					in, _ := SignUp(ctx, &protoc.User{Name: name, Password: password})
					if in.Response == "OK" {
						sender = name
					}
				}
			}
		}
	}
}
