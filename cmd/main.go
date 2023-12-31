package main

import (
	"aone-qc/cmd/server"

	"log"
	"os"
	"path"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(log.Default().Flags() | log.Llongfile)

	var rootCmd = &cobra.Command{Use: path.Base(os.Args[0])}
	rootCmd.PersistentFlags().StringP("config", "c", "./config.yaml", "apiserver config file path.")

	server.Register(rootCmd)
	rootCmd.Execute()
}
