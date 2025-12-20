package cmd

import (
	"context"
	"fmt"
	"syscall"

	"github.com/nabidam/baaham/pkg/password"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		isAdmin, _ := cmd.Flags().GetBool("admin")

		if username == "" {
			return fmt.Errorf("username required")
		}

		fmt.Print("Password: ")
		pass, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			return err
		}

		hash, err := password.HashPassword(string(pass))
		if err != nil {
			return err
		}

		created, err := repo.Create(
			context.Background(),
			username,
			hash,
			isAdmin,
		)
		if err != nil {
			return err
		}

		fmt.Printf(
			"User created: %s (admin=%v, id=%s)\n",
			created.Username,
			created.IsAdmin,
			created.ID,
		)

		return nil
	},
}

func init() {
	createCmd.Flags().StringP("username", "u", "", "username")
	createCmd.Flags().Bool("admin", false, "is admin")
	rootCmd.AddCommand(createCmd)
}
