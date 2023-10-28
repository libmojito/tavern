package serve

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/libmojito/tavern/examples/openai/cmd/openai"
	"github.com/libmojito/tavern/examples/openai/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var port int

type commandServer struct {
	proto.UnimplementedCommandServerServer
}

func (s *commandServer) Run(ctx context.Context, req *proto.CommandRequest) (*proto.CommandReply, error) {

	ow := bytes.NewBufferString("")
	ew := bytes.NewBufferString("")

	cmd := openai.NewCmd(
		openai.WithOut(ow),
		openai.WithErr(ow),
	)
	cmd.SetArgs(req.Args)

	err := cmd.Execute()
	if err != nil {
		return &proto.CommandReply{
			Status: 1,
			Stdout: ow.String(),
			Stderr: ew.String(),
		}, nil
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

var Cmd = &cobra.Command{
	Use:   "serve",
	Short: "start the grpc server",
	Long:  "start the grpc server",
	Run: func(cmd *cobra.Command, args []string) {
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

func init() {
	Cmd.Flags().IntVar(&port, "port", 50051, "The server port")
}
