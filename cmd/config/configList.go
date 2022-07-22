/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdConfigList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "prints vtconfig file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hi! i am vt config list") //TODO
		},
	}

	return cmd
}
