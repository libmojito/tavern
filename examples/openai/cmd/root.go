package cmd

import (
	"fmt"
	"os"

	"github.com/libmojito/tavern/examples/openai/cmd/openai"
	"github.com/libmojito/tavern/examples/openai/cmd/serve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultCfgFileName = ".tavern-openai"
const defaultCfgFileType = "yaml"

var cfgFile string

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command-server",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		fmt.Sprintf(
			"config file (default is $HOME/%s.%s)",
			defaultCfgFileName,
			defaultCfgFileType,
		),
	)

	cmd.AddCommand(serve.Cmd)
	cmd.AddCommand(openai.NewCmd())

	return cmd
}

func Execute() {
	err := NewCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func InitConfig() {
	initConfig()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType(defaultCfgFileType)
		viper.SetConfigName(defaultCfgFileName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
