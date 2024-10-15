package main

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
	Form     *huh.Form
	Select   *huh.Select[string]
	GotoPath string
}

func (m model) Init() tea.Cmd {
	return m.Form.Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q", "Q", "ctrl+c":
			form, _ := m.Form.Update(msg)
			if f, ok := form.(*huh.Form); ok {
				m.Form = f
			}
			return m, tea.Quit

		case "b", "B", "left":
			parentDir := filepath.Dir(m.GotoPath)
			m.GotoPath = parentDir
			updateFiles(m)
			if m.Form != nil {
				menuModel, cmd := m.Form.Update(msg)

				if menu, ok := menuModel.(*huh.Form); ok {
					m.Form = menu
				}
				return m, cmd
			}

		case "enter", "right":
			m.GotoPath = fmt.Sprintf("%s/%s", m.GotoPath, m.Select.GetValue())
			updateFiles(m)
			if m.Form != nil {
				menuModel, cmd := m.Form.Update(msg)

				if menu, ok := menuModel.(*huh.Form); ok {
					m.Form = menu
				}
				return m, cmd
			}
		}

		if m.Form != nil {
			menuModel, cmd := m.Form.Update(msg)

			if menu, ok := menuModel.(*huh.Form); ok {
				m.Form = menu
			}
			return m, cmd
		}

	}
	return m, nil
}

func (m *model) View() string {

	return m.Form.View()
}

func NewModel() model {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("(-) Error getting working directory: ", err)
	}

	paths := getPaths(path)

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
	selectComponent := huh.NewSelect[string]().Key("folder").
		Options(huhOptions...).
		Title("Start Navi!!")

	return model{
		Form:     huh.NewForm(huh.NewGroup(selectComponent)),
		Select:   selectComponent,
		GotoPath: path,
	}
}

func getPaths(path string) []fs.DirEntry {
	paths, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("(-) Error reading file: ", err)
	}

	return paths
}

func updateFiles(m *model) {
	paths := getPaths(m.GotoPath)

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
	m.Select.Options(huhOptions...)
}

func main() {
	m := NewModel()
	p := tea.NewProgram(&m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatalln("(-) Error starting the program: ", err)
	}
}
