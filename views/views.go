package maiaasad

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type model struct {
	GotoPath string
	Menu     *huh.Select[string]
	index    int
}

func New() *model {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("(-) Error getting working directory: ", err)
	}

	paths, err := getPaths(path)
	if err != nil {
		log.Fatalln("(-) Error getting paths: ", err)
	}

	var huhOptions []huh.Option[string]
	for _, file := range paths {
		var newOption string
		if file.IsDir() {
			newOption = "[DIR]  " + file.Name()
		} else {
			newOption = "[FILE] " + file.Name()
		}
		huhOption := huh.NewOption(newOption, file.Name())
		huhOptions = append(huhOptions, huhOption)
	}

	m := model{
		GotoPath: path,
		Menu:     huh.NewSelect[string]().Title("Let's navigate").Options(huhOptions...),
	}
	return &m
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		switch key {
		case "b", "B":
			parentDir := filepath.Dir(m.GotoPath)
			m.GotoPath = parentDir
			return m, nil
		case "q", "Q", "ctrl+c":
			return m, tea.Quit

		case "up", "k", "K":
			if m.Menu != nil {
			}
			return m, nil
		}
		if m.Menu != nil {
			menuModel, cmd := m.Menu.Update(msg)

			if menu, ok := menuModel.(*huh.Select[string]); ok {
				m.Menu = menu
			}
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	fmt.Println("View called:", m.GotoPath)
	updateFiles(&m)

	m.Menu.Select(m.index)

	return m.Menu.View()
}

func main() {
	m := New()
	program := tea.NewProgram(m)
	_, err := program.Run()
	if err != nil {
		log.Fatalln("(-) Error starting the program: ", err)
	}
}

func getPaths(path string) ([]fs.DirEntry, error) {
	paths, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("(-) Error reading file: ", err)
	}

	return paths, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func updateFiles(m *model) {
	paths, err := getPaths(m.GotoPath)
	if err != nil {
		log.Fatalln("(-) Error getting paths: ", err)
	}

	var huhOptions []huh.Option[string]
	for _, file := range paths {
		var newOption string
		if file.IsDir() {
			newOption = "[DIR]  " + file.Name()
		} else {
			newOption = "[FILE] " + file.Name()
		}
		huhOption := huh.NewOption(newOption, file.Name())
		huhOptions = append(huhOptions, huhOption)
	}
	m.Menu.Title("")
	m.Menu.Options(huhOptions...)
}
