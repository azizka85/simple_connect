package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"simple_connect/services/greet"
	"simple_connect/services/greet/greetconnect"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port = flag.Int("port", 5000, "The server port")
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greet.GreetRequest],
) (*connect.Response[greet.GreetResponse], error) {
	log.Println("Request: ", req.Peer().Addr, req.Peer().Protocol)

	res := connect.NewResponse(
		&greet.GreetResponse{
			Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
		},
	)

	return res, nil
}

func (s *GreetServer) Send(
	ctx context.Context,
	stream *connect.ClientStream[greet.TimingInputRequest],
) (*connect.Response[greet.TimingInputResponse], error) {
	log.Println("Request: ", stream.Peer().Addr, stream.Peer().Protocol)

	res := connect.NewResponse(
		&greet.TimingInputResponse{
			Data: []*greet.MessageInfo{},
		},
	)

	for stream.Receive() {
		res.Msg.Data = append(
			res.Msg.Data,
			&greet.MessageInfo{
				Time:    timestamppb.Now(),
				Message: stream.Msg().Message,
			},
		)
	}

	if err := stream.Err(); err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	return res, nil
}

func (s *GreetServer) Accept(
	ctx context.Context,
	req *connect.Request[greet.TimingInputRequest],
	res *connect.ServerStream[greet.MessageInfo],
) error {
	log.Println("Request: ", req.Peer().Addr, req.Peer().Protocol)

	message := req.Msg.Message

	if strings.TrimSpace(message) == "" {
		message = "Server message"
	}

	for i := 0; i < 10; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}

		if err := res.Send(
			&greet.MessageInfo{
				Time:    timestamppb.Now(),
				Message: message,
			},
		); err != nil {
			return err
		}

		log.Println("print message:", message, i, "times")

		time.Sleep(5 * time.Second)
	}

	return nil
}

func main() {
	flag.Parse()

	greeter := &GreetServer{}
	mux := http.NewServeMux()

	path, handler := greetconnect.NewGreetServiceHandler(greeter)

	mux.Handle(path, handler)

	path, handler = greetconnect.NewTimingInputServiceHandler(greeter)

	mux.Handle(path, handler)

	path, handler = greetconnect.NewTimingOutputServiceHandler(greeter)

	mux.Handle(path, handler)

	fs := http.FileServer(http.Dir("./client/dist"))

	mux.Handle(
		"/",
		fs,
	)

	addr := fmt.Sprintf("localhost:%d", *port)

	log.Printf("Listening %s\n", addr)

	corsEntity := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://localhost:5000"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	/* http.ListenAndServeTLS(
		addr,
		"server.crt",
		"server.key",
		corsEntity.Handler(
			mux,
		),
	) */
	http.ListenAndServe(
		addr,
		corsEntity.Handler(
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
}
