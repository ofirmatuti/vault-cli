/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	cmdconfig "github.com/ofirmatuti/vt/cmd/config"
	cmdlogin "github.com/ofirmatuti/vt/cmd/login"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vt",
	Short: "vt helps navigating through vault secrets",

	Run: func(cmd *cobra.Command, args []string) {
		isTtl, _ := cmd.Flags().GetBool("ttl")

		if isTtl {
			fmt.Println("Hey! i am vt -t !") //TODO add logic
			os.Exit(0)
		}

		fmt.Println("hi! i am vt") //TODO

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("ttl", "t", false, "Prints ttl for vault okta authentication")
	rootCmd.AddCommand(cmdconfig.NewCmdConfig())
	rootCmd.AddCommand(cmdlogin.NewCmdLogin())
}
