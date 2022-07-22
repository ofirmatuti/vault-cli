/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package config

import (
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Modifies vtconfig file",
	}

	cmd.AddCommand(NewCmdConfigSetBaseDir())
	cmd.AddCommand(NewCmdConfigSetAddress())
	cmd.AddCommand(NewCmdConfigList())
	return cmd
}
