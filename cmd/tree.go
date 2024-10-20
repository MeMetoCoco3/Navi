/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"os"
)

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Prints tree of current directory.",
	Long:  `Prints a tree showing all directories and files from the current directory, with labels indicating type.`,

	Run: func(cmd *cobra.Command, args []string) {

		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("PWD ERROR")
			return
		}

		files, err := os.ReadDir(dir)

		if err != nil {
			fmt.Println("ReadDir ERROR")
		}
		for _, file := range files {
			if file.IsDir() {
				fmt.Println("[DIR] ", file.Name())
			} else {
				fmt.Println("[FILE]", file.Name())
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// treeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// treeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
