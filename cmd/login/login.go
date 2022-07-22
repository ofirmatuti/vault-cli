/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticates to vault using okta",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hi! i am vt login") //TODO
		},
	}
	return cmd
}
