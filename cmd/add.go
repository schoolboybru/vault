/*
Copyright © 2024 @schoolboybru
*/
package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var filePath string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Quickly add a file into your vault",
	Long:  `Add a file into your vault to the specified location`,
	Run: func(cmd *cobra.Command, args []string) {

		c := exec.Command("hdiutil", "info")
		var out bytes.Buffer
		c.Stdout = &out

		err := c.Run()

		if err != nil {
			fmt.Printf("Failed to run hdiutil info: %s", err)
			return
		}

		dir := "/Volumes/Vault/"

		fmt.Println(strings.Contains(out.String(), dir))
		if strings.Contains(out.String(), dir) {
			err = addFileHandler(cmd, dir)

			if err != nil {
				fmt.Printf("Failed to add: %s", err)
				return
			}

			fmt.Println("Added file to vault")

		} else {

			err = unlockHandler()

			if err != nil {
				fmt.Printf("Failed to unlock: %s", err)
				return
			}

			err = addFileHandler(cmd, dir)

			if err != nil {
				fmt.Printf("Failed to add: %s", err)
				return
			}

			fmt.Println("Added file to vault")
		}

	},
}

func init() {
	addCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path for file to be added")
}

func addFileHandler(cmd *cobra.Command, dir string) error {
	file, _ := cmd.Flags().GetString("file")

	a := exec.Command("mv", file, dir)

	err := a.Run()

	return err
}
