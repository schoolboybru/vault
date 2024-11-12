/*
Copyright © 2024 @schoolboybru
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create your vault",
	Long:  `Initialize your vault to store files securely on your device`,
	Run: func(cmd *cobra.Command, args []string) {

		size := "100m"
		volumeName := "Vault"
		outputPath := os.Getenv("HOME") + "/Vault.dmg"

		fmt.Print("Enter password for the encrypted disk image: ")
		password, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read password:", err)
			return
		}

		password = strings.TrimSpace(password)

		c := exec.Command("hdiutil", "create",
			"-encryption", "-stdinpass",
			"-size", size,
			"-fs", "APFS",
			"-volname", volumeName,
			outputPath,
		)

		c.Stdin = strings.NewReader(password)

		var stderr bytes.Buffer

		c.Stderr = &stderr

		err = c.Run()

		if err != nil {
			fmt.Printf("Error creating vault: %v\n", err)
			fmt.Printf("hdiutil stderr: %s\n", stderr.String())
			return
		}

		fmt.Println("Successfully created vault")

	},
}
