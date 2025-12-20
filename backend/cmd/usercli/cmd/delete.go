package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		if username == "" {
			return fmt.Errorf("username required")
		}

		err := repo.Delete(context.Background(), username)
		if err != nil {
			return err
		}

		fmt.Printf("User deleted.")

		return nil
	},
}

func init() {
	deleteCmd.Flags().StringP("username", "u", "", "username")
	rootCmd.AddCommand(deleteCmd)
}
