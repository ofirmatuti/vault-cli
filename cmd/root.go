/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	cmdConfig "github.com/ofirmatuti/vt/cmd/config"
	cmdLogin "github.com/ofirmatuti/vt/cmd/login"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vt",
	Short: "vt helps navigating through vault secrets",

	Run: func(cmd *cobra.Command, args []string) {
		isTtl, _ := cmd.Flags().GetBool("ttl")
		updateCache, _ := cmd.Flags().GetBool("update-cache")
		if isTtl {
			fmt.Println("Hey! i am vt -t !") //TODO add logic
			os.Exit(0)
		}
		if updateCache {
			fmt.Println("Hey! i am updating the cache now")
			os.Exit(0)
		}

		fmt.Println("hi! i am vt") //TODO

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			fmt.Println("error printing cmd stderr")
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("ttl", "t", false, "Prints authentication token ttl")
	rootCmd.Flags().BoolP("update-cache", "f", false, "Fetches all basedir subdirectories recursively to cache")
	rootCmd.AddCommand(cmdConfig.NewCmdConfig())
	rootCmd.AddCommand(cmdLogin.NewCmdLogin())
}
