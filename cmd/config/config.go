/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package config

import (
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Modifies vtconfig file",
		Long:  "config",
		Run: func(cmd *cobra.Command, args []string) {
			baseDir, _ := cmd.Flags().GetString("base-dir")
			address, _ := cmd.Flags().GetString("address")
			if baseDir != "" {
				println("hey i am adding basedir to configuration!")
			}
			if address != "" {
				println("hey i am adding address to configuration!")
			}
		},
	}

	configCmd.AddCommand(NewCmdGlobalConfig())

	return configCmd
}
