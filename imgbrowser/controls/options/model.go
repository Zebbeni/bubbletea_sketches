package options

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/characters"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls/options/colors"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls/options/sampling"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/size"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/style"
)

type Model struct {
	active State
	focus  State

	Colors     colors.Model
	Characters characters.Model
	Size       size.Model
	Sampling   sampling.Model

	ShouldUnfocus bool
	ShouldClose   bool
}

func New() Model {
	return Model{
		active: None,
		focus:  Colors,

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
	return m.handleMenuUpdate(msg)
}

func (m Model) View() string {
	col := m.renderWithBorder(m.Colors.View(), Colors)
	char := m.renderWithBorder(m.Characters.View(), Characters)
	siz := m.renderWithBorder(m.Size.View(), Size)
	sam := m.renderWithBorder(m.Sampling.View(), Sampling)

	return lipgloss.JoinVertical(lipgloss.Top, col, char, siz, sam)
}

var (
	activeColor = lipgloss.Color("#aaaaaa")
	focusColor  = lipgloss.Color("#ffffff")
	normalColor = lipgloss.Color("#555555")
)

func (m Model) renderWithBorder(content string, state State) string {
	renderColor := normalColor
	if m.active == state {
		renderColor = activeColor
	} else if m.focus == state {
		renderColor = focusColor
	}

	textStyle := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center).
		Padding(0, 1, 0, 1).
		Foreground(renderColor)
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(renderColor)

	renderer := style.BoxWithLabel{
		BoxStyle:   borderStyle,
		LabelStyle: textStyle,
	}

	return renderer.Render(stateTitles[state], content, 30)
}
