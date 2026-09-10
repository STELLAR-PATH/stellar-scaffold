package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "stellar-scaffold",
	Short:   "A project generator for STELLAR-PATH smart contracts",
	Long:    `stellar-scaffold scaffolds complete Stellar & Soroban smart contract monorepos.`,
	Version: "v0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// The --version flag is added by default when Version is set on rootCmd.
}
