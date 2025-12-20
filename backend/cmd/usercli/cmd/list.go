package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List users",
	RunE: func(cmd *cobra.Command, args []string) error {
		users, err := repo.List(context.Background())
		if err != nil {
			return err
		}

		if len(users) == 0 {
			fmt.Print("There is no user in db.")
			return nil
		}
		for _, u := range users {
			fmt.Printf(
				"%s | admin=%v | created=%s\n",
				u.Username,
				u.IsAdmin,
				u.CreatedAt.Format("2006-01-02"),
			)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
