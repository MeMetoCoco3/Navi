/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MeMetoCoco3/navi/favorites"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

const favsRoute = "./favorites/favs.json"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add current directory to favourites.",
	Long: `Adds the current directory to the favoruites.json file for easy access in the future via 'fv' command.
  Keep track of your favourite directories, so you will
  be able to open them easly and efficiently`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dir string

		if len(args) > 0 {
			dir = args[0]
		} else {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalln("(-) Error getting current directory", err)
				return
			}
			dir = cwd
		}

		absDir, err := filepath.Abs(dir)
		if err != nil {
			log.Fatalln("(-) Error getting absolute path:", err)
		}

		favorites.AddFav(absDir)
		log.Fatalf("(+) %s was added to favorites!", absDir)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
