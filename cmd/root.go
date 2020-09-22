package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const (
	version = "1.3.0"
)

var Port int
var Tls bool
var Verbose bool

var rootCmd = &cobra.Command{
	Use:     "ktunnel",
	Short:   "Ktunnel is a network tunneling tool for kubernetes",
	Long:    `Built to ease development on kubernetes clusters and allow connectivity between dev machines and clusters`,
	Version: version,
	Args:    cobra.MinimumNArgs(1),
}

func Execute() {
	if genDoc := os.Getenv("GEN_DOC"); genDoc == "true" {
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			log.Errorf("Failed generating docs: %v", err)
		}
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if Verbose {
		log.SetLevel(log.DebugLevel)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 28688, "The port to use to establish the tunnel")
	rootCmd.PersistentFlags().BoolVarP(&Tls, "tls", "t", false, "Connection uses TLS if true, else plain TCP")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose mode")
	_ = rootCmd.MarkFlagRequired("port")
}
