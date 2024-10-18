package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/MeMetoCoco3/navi/favorites"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 40
const maxLenPaths = 20

type model struct {
	Form     *huh.Form
	Select   *huh.Select[string]
	GotoPath string
	Style    *Styles
	Width    int
}

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.ThickBorder()).Padding(1).Width(maxWidth)
	return s
}

func (m model) Init() tea.Cmd {
	return m.Form.Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Select.Title("GATOR")
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
		case "a", "A":
			checker := favorites.CheckFav(m.GotoPath)
			if checker == -1 {
				favorites.AddFav(m.GotoPath)
				m.Select.Title("Directory added.")
			} else {
				favorites.RemoveFav(m.GotoPath)
				m.Select.Title("Directory Removed")
			}
			if m.Form != nil {
				menuModel, cmd := m.Form.Update(msg)

				if menu, ok := menuModel.(*huh.Form); ok {
					m.Form = menu
				}
				return m, cmd
			}

		case "g", "G":
			return m, tea.Quit
		case "enter", "right":

			selectValue := m.Select.GetValue()
			paths := getPaths(fmt.Sprintf("%s/%s", m.GotoPath, selectValue))

			if selectValue == "" {
			} else if len(paths) == 0 {
				m.Select.Title("Empty Directory")
			} else {
				newPath := fmt.Sprintf("%s/%s", m.GotoPath, selectValue)

				newPathStats, err := os.Stat(newPath)
				if err != nil {
					log.Fatalln("Error getting path stats: ", err)
				}

				if !newPathStats.IsDir() {
					return m, nil
				}

				m.GotoPath = newPath
				updateFiles(m)
			}
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

	return lipgloss.JoinVertical(
		lipgloss.Center,
		m.Style.InputField.Render(m.Form.View()),
	)
}

func NewModel() model {

	style := DefaultStyles()

	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("(-) Error getting working directory: ", err)
	}

	paths := getPaths(path)
	var huhOptions []huh.Option[string]
	for _, file := range paths {
		var newOption string
		var fileName string

		if len(file.Name()) > maxWidth {
			fmt.Println(file.Name())
			fileName = file.Name()[:maxLenPaths] + "..."
			fmt.Println(fileName)
		} else {
			fileName = file.Name()
		}

		if file.IsDir() {
			newOption = "[DIR]  " + fileName
		} else {
			newOption = "[FILE] " + fileName
		}
		huhOption := huh.NewOption(newOption, file.Name())
		huhOptions = append(huhOptions, huhOption)
	}
	selectComponent := huh.NewSelect[string]().Key("folder").
		Options(huhOptions...).
		Title("GATOR")

	return model{
		Form:     huh.NewForm(huh.NewGroup(selectComponent)),
		Select:   selectComponent,
		GotoPath: path,
		Style:    style,
		Width:    maxWidth,
	}
}

func getPaths(path string) []fs.DirEntry {
	paths, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("(-) Error reading file: ", err)
	}

	var dirPaths []fs.DirEntry
	for _, path := range paths {
		// Check if the file is hidden (starts with a dot)
		if path.Name()[0] == '.' {
			dirPaths = append(dirPaths, path)
		} else {
			dirPaths = append(dirPaths, path)
		}
	}

	return dirPaths
}

func updateFiles(m *model) {
	paths := getPaths(m.GotoPath)

	var huhOptions []huh.Option[string]
	for _, file := range paths {
		var newOption string
		var fileName string

		if len(file.Name()) > maxWidth {
			fileName = file.Name()[:maxLenPaths] + "..."
		} else {
			fileName = file.Name()
		}

		if file.IsDir() {
			newOption = "[DIR]  " + fileName
		} else {
			newOption = "[FILE] " + fileName
		}
		huhOption := huh.NewOption(newOption, file.Name())
		huhOptions = append(huhOptions, huhOption)
	}
	if huhOptions == nil {

	} else {
		m.Select.Options(huhOptions...)
	}
}

func main() {
	m := NewModel()
	p := tea.NewProgram(&m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatalln("(-) Error starting the program: ", err)
	}
	fmt.Println(m.GotoPath)
}
