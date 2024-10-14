/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/MeMetoCoco3/navi/favorites"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// gotoCmd represents the goto command
var gotoCmd = &cobra.Command{
	Use:   "goto",
	Short: "Goes to a directory",
	Long: `Recieves an index from the list of favorite directories
  and calls cd (change directory) to that path.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("(-) Error converting string to int:", err)
			os.Exit(1)
		}
		path, err := favorites.ChangeDir(index)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(path)
	},
}

func init() {
	rootCmd.AddCommand(gotoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gotoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gotoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
