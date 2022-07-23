/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdConfigSetAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-address",
		Short: "sets vault address", //TODO
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Hey! you should specify a vault address!")
				return
			}
			fmt.Println("vault address config was set: " + args[0]) //TODO

		},
	}
	return cmd
}
