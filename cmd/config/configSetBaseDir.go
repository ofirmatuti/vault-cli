/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdConfigSetBaseDir() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-base-dir",
		Short: "sets the vault base dir to list secret from", //TODO
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Hey! you should specify a base dir!")
				return
			}
			fmt.Println("base dir config was set: " + args[0]) //TODO
		},
	}
	return cmd
}
