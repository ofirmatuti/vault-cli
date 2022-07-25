/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/

package config

import (
	"github.com/spf13/cobra"
)

func NewCmdGlobalConfig() *cobra.Command {
	globalConfigCmd := &cobra.Command{
		Use:   "global",
		Short: "set persistent global configurations for vt",
		Run: func(cmd *cobra.Command, args []string) {
			baseDir, _ := cmd.Flags().GetString("base-dir")
			address, _ := cmd.Flags().GetString("address")
			if address != "" {
				println("hey i am adding address " + address + " to configuration!")
			}

			if baseDir != "" {
				println("hey i am adding basedir " + baseDir + " to configuration!")
			}
		},
	}

	globalConfigCmd.PersistentFlags().String("base-dir", "", "")
	globalConfigCmd.PersistentFlags().String("address", "", "sets your vault address to fetch secrets from")

	return globalConfigCmd
}
