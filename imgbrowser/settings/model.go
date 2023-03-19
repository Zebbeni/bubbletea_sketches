package settings

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/characters"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/sampling"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/size"
)

type Model struct {
	active State

	Colors     colors.Model
	Characters characters.Model
	Size       size.Model
	Sampling   sampling.Model

	ShouldUnfocus bool
	ShouldClose   bool
}

func New() Model {
	return Model{
		active: Colors,

		Colors:     colors.New(),
		Characters: characters.New(),
		Size:       size.New(),
		Sampling:   sampling.New(),

		ShouldUnfocus: false,
		ShouldClose:   false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch m.active {
	case Colors:
		return m.handleColorsUpdate(msg)
	case Characters:
		return m.handleCharactersUpdate(msg)
	case Size:
		return m.handleSizeUpdate(msg)
	case Sampling:
		return m.handleSamplingUpdate(msg)
	}
	return m, nil
}

func (m Model) View() string {
	col := m.Colors.View()
	char := m.Characters.View()
	siz := m.Size.View()
	sam := m.Sampling.View()
	return lipgloss.JoinVertical(lipgloss.Top, col, char, siz, sam)
}
