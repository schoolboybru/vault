/*
Copyright © 2024 @schoolboybru
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var editorFlag bool

// unlockCmd represents the unlock command
var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock your vault",
	Long:  "Unlock your vault to access your files",
	Run: func(cmd *cobra.Command, args []string) {
		err := unlockHandler()

		if err != nil {
			fmt.Printf("Error unlocking vault: %v\n", err)
			return
		}

		fmt.Println("Vault unlocked")

		if editorFlag {
			err = openEditorHandler()

			if err != nil {
				fmt.Printf("Error opening editor: %s", err)
			}
		}
	},
}

func init() {
	unlockCmd.Flags().BoolVarP(&editorFlag, "editor", "e", false, "Open the vault with your default editor")
}

func unlockHandler() error {

	path := os.Getenv("HOME") + "/Vault.dmg"

	c := exec.Command("hdiutil", "attach", path)

	err := c.Run()

	if err != nil {
		fmt.Printf("Error unlocking vault: %v\n", err)
		return err
	}

	return nil
}

func openEditorHandler() error {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		return errors.New("EDITOR environment variable must be set")
	}

	dir := "/Volumes/Vault/"

	o := exec.Command(editor, dir)

	o.Stdin = os.Stdin
	o.Stdout = os.Stdout
	o.Stderr = os.Stderr

	err := o.Run()

	if err != nil {
		return err
	}

	return nil
}
