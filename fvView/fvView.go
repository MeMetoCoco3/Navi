package fvViews

import (
	"errors"
	"github.com/MeMetoCoco3/navi/favorites"
	//"github.com/MeMetoCoco3/navi/views"
	"log"
	"os"
	"strings"

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
			return m, tea.Quit

		case "g", "G":
			favorites.WriteOnTmp(m.GotoPath)
			return m, tea.Quit

		case "a", "A":
			dirPath := m.Select.GetValue()
			favorites.RemoveFav(dirPath.(string))
			err := UpdateFiles(m)

			if err != nil {
				return m, tea.Quit

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

func UpdateFiles(m *model) error {
	paths := favorites.ListFavs()
	if len(paths) <= 0 {
		return errors.New("(-) Favourites list is empty.")
	}
	var huhOptions []huh.Option[string]
	for _, file := range paths {
		var fileName string

		if len(file) >= maxLenPaths {
			fileName = string(0x2724) + "   " + "..." + file[len(file)-maxLenPaths:]
		} else {
			// Calculate how many spaces to add
			spaces := strings.Repeat(" ", maxLenPaths-len(file)+6)
			fileName = string(0x2724) + spaces + file
		}

		huhOption := huh.NewOption(fileName, file)
		huhOptions = append(huhOptions, huhOption)
	}
	if huhOptions == nil {
		return nil
	} else {
		m.Select.Options(huhOptions...)
		return nil
	}

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

	paths := favorites.ListFavs()
	if len(paths) <= 0 {
		log.Fatalln("(-) Empty favorites list.")
	}
	var huhOptions []huh.Option[string]
	for _, file := range paths {
		var fileName string

		if len(file) >= maxLenPaths {
			fileName = string(0x2724) + "   " + "..." + file[len(file)-maxLenPaths:]
		} else {
			// Calculate how many spaces to add
			spaces := strings.Repeat(" ", maxLenPaths-len(file)+6)
			fileName = string(0x2724) + spaces + file
		}

		huhOption := huh.NewOption(fileName, file)
		huhOptions = append(huhOptions, huhOption)
	}
	heightSelect := len(huhOptions)
	if len(huhOptions) < maxHeight {
		heightSelect = maxHeight
	}
	selectComponent := huh.NewSelect[string]().Key("folder").
		Options(huhOptions...).
		Title("My Favorites")

	return model{
		// This WithHeight kicked my ass
		Form:     huh.NewForm(huh.NewGroup(selectComponent).WithHeight(heightSelect)),
		Select:   selectComponent,
		GotoPath: path,
		Style:    style,
		Width:    maxWidth,
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
