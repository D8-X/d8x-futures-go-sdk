package cli

import (
	"fmt"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/pkg/d8x_futures"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ChainConfig         utils.ChainConfig
	PriceConfig         utils.PriceFeedConfig
	BlockChainConnector d8x_futures.BlockChainConnector
	chainConfigNames    []string
	cursor              int
	selected            int
	screen              int
}

func initialModel() Model {
	names, err := config.GetDefaultChainConfigNames()
	if err != nil {
		panic(err)
	}
	return Model{
		chainConfigNames: names,
		screen:           0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.chainConfigNames)-1 {
				m.cursor++
			}

		// Spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case " ":

			if m.selected == m.cursor {
				m.selected = 0
			} else {
				m.selected = m.cursor
			}
		case "enter":
			m.screen += 1
			if m.screen == 1 {
				m.actionScreen01()
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Model) View() string {
	s := ""
	switch m.screen {
	case 0:
		s += m.selectNetworkView()
	case 1:
		s += m.initialView()
	}
	return s
}

func (m Model) initialView() string {
	return "Connected to " + m.chainConfigNames[m.cursor]
}

func (m Model) selectNetworkView() string {
	// The header
	s := "Which Network?\n\n"

	// Iterate over our choices
	for i, choice := range m.chainConfigNames {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if i == m.selected {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func (m Model) actionScreen01() string {
	// load selected chainconfig
	network := m.chainConfigNames[m.selected]
	chConf, err := config.GetDefaultChainConfig(network)
	if err != nil {
		return "connection failed"
	}
	m.ChainConfig = chConf
	pxConf, err := config.GetDefaultPriceConfig(m.ChainConfig.PriceFeedNetwork)
	if err != nil {
		return "connection failed"
	}
	conn := d8x_futures.CreateBlockChainConnector(pxConf, chConf)
	m.BlockChainConnector = conn
	return "Connected to " + network
}
