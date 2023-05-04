package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate the database",
	Long:  `migrate the database`,
	Run:   migrate,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) {
	fmt.Println("migrate command called! nothing is implemented")
}
