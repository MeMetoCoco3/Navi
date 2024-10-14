package main

import (
	"github.com/charmbracelet/huh"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func getPaths(path string) ([]fs.DirEntry, error) {
	paths, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("(-) Error reading file: ", err)
	}

	return paths, nil
}

func main() {

	gotoPath, err := os.Getwd()
	if err != nil {
		log.Fatalln("(-) Error getting working directory", err)
	}

	// SIGNALS are syscalls , we create a channel that listen to them
	// and an event that listens to n or N
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	/*
		go func() {
			for {
				event := selectMenu.Run()
				if event.Key == "n" || event.Key == "N" {
					fmt.Println("Signal Recieved")
				}
			}

		}()
	*/
	for {
		select {
		case <-sigChan:

			//return gotoPath

		default:
			// TODO ABSOLUTE PATHS
			paths, err := getPaths(gotoPath)

			var options []string
			var huhOptions []huh.Option[string]

			// Format and creation of options
			for _, file := range paths {
				var name string
				if file.IsDir() {
					name = "[DIR]  " + file.Name()
				} else {
					name = "[FILE] " + file.Name()
				}
				new_option := name
				options = append(options, new_option)
				huhOption := huh.NewOption(new_option, file.Name())
				huhOptions = append(huhOptions, huhOption)
			}

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Let's Navigate").
						Options(
							huhOptions...,
						).
						Value(&gotoPath),
				),
			)
			err = form.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
