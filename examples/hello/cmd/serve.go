package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/libmojito/tavern/examples/hello/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const ServeOptPort = "port"

type commandServer struct {
	proto.UnimplementedCommandServerServer
}

func (s *commandServer) Run(ctx context.Context, req *proto.CommandRequest) (*proto.CommandReply, error) {

	ow := bytes.NewBufferString("")
	ew := bytes.NewBufferString("")

	cmd := NewHelloCmd(
		WithOut(ow),
		WithErr(ow),
	)
	cmd.SetArgs(req.Args)

	err := cmd.Execute()
	if err != nil {
		return nil, err
	}

	return &proto.CommandReply{
		Status: 0,
		Stdout: ow.String(),
		Stderr: ew.String(),
	}, nil

}

func newServer() *commandServer {
	return &commandServer{}
}

func NewServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve the hello command",
		Long:  `serve the hello command`,
		Run: func(cmd *cobra.Command, args []string) {
			port, err := cmd.Flags().GetInt(ServeOptPort)
			if err != nil {
				log.Fatal("failed to retrieve port flag")
			}
			address := fmt.Sprintf("localhost:%d", port)
			lis, err := net.Listen("tcp", address)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			var opts []grpc.ServerOption
			grpcServer := grpc.NewServer(opts...)
			proto.RegisterCommandServerServer(grpcServer, newServer())
			log.Printf("starting server on %s...\n", address)
			grpcServer.Serve(lis)
		},
	}

	cmd.Flags().Int(ServeOptPort, 50051, "The server port")

	return cmd
}
