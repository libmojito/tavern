package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/libmojito/tavern/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type remoteCommand struct {
	Use     string
	Short   string
	Long    string
	Address string
}

func (rCmd *remoteCommand) run(cmd *cobra.Command, args []string) {
	conn, err := grpc.Dial(rCmd.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewCommandServerClient(conn)
	reply, err := client.Run(context.TODO(), &proto.CommandRequest{
		Args: args,
	})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	if reply.Stderr != "" {
		cmd.PrintErr(reply.Stderr)
	}
	if reply.Stdout != "" {
		cmd.Print(reply.Stdout)
	}
}

func init() {
	// read all command in configurations, this has to be done as part of init
	// instead of the preferred cobra.OnIntialize for the reason it alters cli
	// behavior and structure.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error %v", err)
	}
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".tavern")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("error parsing config %v", err)
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	var rCmds []remoteCommand
	viper.UnmarshalKey("tavern.cmds", &rCmds)

	for _, rCmd := range rCmds {
		cmdCopy := rCmd // make a copy so that run would run the correct command
		rootCmd.AddCommand(&cobra.Command{
			Use:                cmdCopy.Use,
			DisableFlagParsing: true, // flag parsing is handled by remote
			Short:              cmdCopy.Short,
			Long:               cmdCopy.Long,
			Run:                cmdCopy.run,
		})
	}

}
