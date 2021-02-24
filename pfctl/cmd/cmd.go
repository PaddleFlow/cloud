package cmd

import (
	"os"
	"path/filepath"

	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/paddleflow/pfctl/pkg/client"
)

func cmd() *cobra.Command {
	viper.SetDefault("deadline", 18000)

	viper.SetConfigName("config")
	var conf string
	if len(conf) == 0 {
		if v, ok := os.LookupEnv("CONFIG_DIR"); ok {
			conf = v + "/ctl/config.yml"
		} else {
			conf = os.Getenv("HOME") + "/.config/pfctl/config.yml"
		}
	}

	viper.SetConfigFile(conf)
	if err := viper.ReadInConfig(); err != nil {
		err = Create(conf)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if err = viper.ReadInConfig(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	var c *client.Client
	if k := viper.GetString("kubeconfig"); len(k) > 0 {
		c = client.GetConfigClient(k)
	} else {
		fmt.Errorf("no kubeconfig !")
	}

	cmd := &cobra.Command{
		Use:          "pfctl",
		Short:        "PaddleFlow CLI tool for multi-clusters management.",
		Long:         "pfctl is a CLI tool for managing PaddleFlow related resources across multiple clusters/namespaces.",
		SilenceUsage: true,
	}

	cmd.AddCommand(contextCmd(c))


	cmd.PersistentFlags().StringSliceP("context", "x", nil, "Specify cluster context.")
	cmd.PersistentFlags().StringP("namespace", "n", "", "Specify the namespace.")

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := cmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func Create(file string) error {
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		return err
	}
	_, err := os.Create(file)
	return err
}