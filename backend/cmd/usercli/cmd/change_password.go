package cmd

import (
	"context"
	"fmt"
	"syscall"

	"github.com/nabidam/baaham/pkg/password"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Change a user's password",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		if username == "" {
			return fmt.Errorf("username required")
		}

		fmt.Print("New password: ")
		pass1, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			return err
		}

		fmt.Print("Confirm password: ")
		pass2, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			return err
		}

		if string(pass1) != string(pass2) {
			return fmt.Errorf("passwords do not match")
		}

		hash, err := password.HashPassword(string(pass1))
		if err != nil {
			return err
		}

		updateErr := repo.UpdatePassword(
			context.Background(),
			username,
			hash,
		)
		if updateErr != nil {
			return updateErr
		}

		fmt.Print("User password updated.")
		return nil
	},
}

func init() {
	changePasswordCmd.Flags().StringP("username", "u", "", "username")
	rootCmd.AddCommand(changePasswordCmd)
}
