/*
Copyright © 2024 @schoolboybru
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// lockCmd represents the lock command
var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock your vault",
	Long:  `Lock your vault`,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("hdiutil", "detach", "/Volumes/Vault/")

		err := c.Run()

		if err != nil {
			fmt.Printf("Error locking vault: %v/n", err)
			return
		}
	},
}
