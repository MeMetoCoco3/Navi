package views

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

const maxWidth = 50
const maxLenPaths = 30
const maxHeight = 30

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
	s.BorderColor = lipgloss.Color("126")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.ThickBorder()).Padding(1).Width(maxWidth).Height(maxHeight)
	return s
}

func (m model) Init() tea.Cmd {
	favorites.WriteOnTmp(m.GotoPath)
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

		case "g", "G":
			favorites.WriteOnTmp(m.GotoPath)
			return m, tea.Quit

		case "b", "B", "left":
			parentDir := filepath.Dir(m.GotoPath)
			m.GotoPath = parentDir
			m.Select.Title(BeautifulCD(m.GotoPath))

			UpdateFiles(m)
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

		case "enter", "right":
			selectValue := m.Select.GetValue()
			paths, err := GetPaths(fmt.Sprintf("%s/%s", m.GotoPath, selectValue))
			if err != nil {
				m.Select.Title("Not a Directory")
				return m, nil
			}
			if selectValue == "" {
				m.Select.Title("Not a Directory")
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
				UpdateFiles(m)
				m.Select.Title(BeautifulCD(m.GotoPath))
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
	UpdateFiles(m)
	return lipgloss.JoinVertical(
		lipgloss.Center,
		m.Style.InputField.Render(m.Form.View()),
	)
}

func NewModel() model {

	currentDirectory, err := os.Getwd()
	currentDirectory = BeautifulCD(currentDirectory)

	style := DefaultStyles()

	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("(-) Error getting working directory: ", err)
	}
	paths, _ := GetPaths(path)

	var huhOptions []huh.Option[string]
	var newOption string
	for _, file := range paths {
		var fileName string

		if len(file.Name()) > maxLenPaths-3 {
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
	heightSelect := len(huhOptions)
	if len(huhOptions) < maxHeight {
		heightSelect = maxHeight
	}
	selectComponent := huh.NewSelect[string]().Key("folder").
		Options(huhOptions...).
		Title(currentDirectory)

	return model{
		// This WithHeight kicked my ass
		Form:     huh.NewForm(huh.NewGroup(selectComponent).WithHeight(heightSelect)),
		Select:   selectComponent,
		GotoPath: path,
		Style:    style,
		Width:    maxWidth,
	}
}

func GetPaths(path string) ([]fs.DirEntry, error) {
	paths, err := os.ReadDir(path)
	if err != nil {
		return nil, err
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

	return dirPaths, nil
}
func UpdateFiles(m *model) {
	paths, _ := GetPaths(m.GotoPath)
	var newOption string
	var fileName string
	var currentPath string
	var huhOptions []huh.Option[string]
	for _, file := range paths {

		if len(file.Name()) > maxLenPaths {
			fileName = file.Name()[:maxLenPaths] + "..."
		} else {
			fileName = file.Name()
		}
		currentPath, _ = os.Getwd()

		currentPath = currentPath + "/" + file.Name()

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

func Run() {
	m := NewModel()
	p := tea.NewProgram(&m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatalln("(-) Error starting the program: ", err)
	}
}

func BeautifulCD(path string) string {
	var beautifulDir string
	if len(path) > maxLenPaths {
		beautifulDir = "..." + path[len(path)-maxLenPaths:]
	} else {
		beautifulDir = path
	}
	/*
		fmt.Println("Path length: ", len(path))
		fmt.Println("currentDirectory: ", path)
		fmt.Println("currentDirectory Beautiful: ", beautifulDir)
	*/
	return beautifulDir
}