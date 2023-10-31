package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/pkg/d8x_futures"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ChainConfig         utils.ChainConfig
	PriceConfig         utils.PriceFeedConfig
	BlockChainConnector d8x_futures.BlockChainConnector
	XchInfo             d8x_futures.StaticExchangeInfo
	chainConfigNames    []string
	choices             []ScreenChoices
	selectedNetworkName string
	screen              int
	poolTable           table.Model
}

type ScreenChoices struct {
	chooseOptions []string
	choiceName    string
	cursor        int
	selected      int
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func initialModel() Model {
	names, err := config.GetDefaultChainConfigNames()
	if err != nil {
		panic(err)
	}
	var m = Model{
		chainConfigNames: names,
		screen:           0,
		choices:          make([]ScreenChoices, 2),
	}
	m.choices[0].chooseOptions = names
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

// https://github.com/charmbracelet/bubbletea/blob/master/examples/views/main.go
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
			if m.screen == 0 {
				if m.choices[m.screen].cursor > 0 {
					m.choices[m.screen].cursor--
				}
				break
			}
			if m.screen == 1 {
				m.poolTable.MoveUp(1)
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.screen == 0 {
				if m.choices[m.screen].cursor < len(m.choices[m.screen].chooseOptions)-1 {
					m.choices[m.screen].cursor++
				}
			}
			if m.screen == 1 {
				m.poolTable.MoveDown(1)
			}
		// Spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case " ":
			m.choices[m.screen].selected = m.choices[m.screen].cursor
		case "enter":
			m.screen += 1
			m.screen = max(m.screen, 1)
			f, _ := tea.LogToFile("debug.log", "screen="+strconv.Itoa(m.screen))
			defer f.Close()
			if m.screen == 1 {
				err := m.actionScreen01()
				if err != nil {
					m.screen--
					fmt.Println("An error occurred: " + err.Error())
				}
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

func (m *Model) initialView() string {
	return "Connected to " + m.selectedNetworkName + "\n" +
		baseStyle.Render(m.poolTable.View()) + "\n"
}

func (m Model) selectNetworkView() string {
	// The header
	s := "Which Network?\n\n"
	s = s + m.displayChoiceMenu()
	// Send the UI for rendering
	return s
}

func (m Model) displayChoiceMenu() string {
	var s string
	// Iterate over our choices
	for i, choice := range m.choices[m.screen].chooseOptions {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.choices[m.screen].cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if i == m.choices[m.screen].selected {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"
	return s
}

func (m *Model) actionScreen01() error {
	// load selected chainconfig
	idx := m.choices[m.screen-1].selected
	network := m.choices[m.screen-1].chooseOptions[idx]
	log := fmt.Sprintf("screen no %d, network='%s' options=%s", m.screen, network, m.choices[m.screen-1].chooseOptions[idx])
	f, err := tea.LogToFile("debug.log", log)
	defer f.Close()
	m.selectedNetworkName = network
	fmt.Print("selecting network " + network)
	chConf, err := config.GetDefaultChainConfig(network)
	if err != nil {
		return errors.New("connection failed")
	}
	m.ChainConfig = chConf
	pxConf, err := config.GetDefaultPriceConfig(m.ChainConfig.PriceFeedNetwork)
	if err != nil {
		return errors.New("connection failed")
	}
	conn := d8x_futures.CreateBlockChainConnector(pxConf, chConf)
	m.BlockChainConnector = conn
	nest, err := d8x_futures.QueryNestedPerpetualInfo(conn)
	if err != nil {
		return err
	}
	m.XchInfo = d8x_futures.QueryExchangeStaticInfo(conn, chConf, nest)
	m.poolTable = createPoolTable(m.XchInfo)
	return nil
}

func createPoolTable(info d8x_futures.StaticExchangeInfo) table.Model {
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Coll.Tkn", Width: 8},
		{Title: "#Perps", Width: 6},
	}
	var rows []table.Row
	for k := 0; k < len(info.Pools); k++ {
		p := info.Pools[k]
		id := strconv.Itoa(int(p.PoolId))
		count := 0
		for j := 0; j < len(info.Perpetuals); j++ {
			if info.Perpetuals[j].PoolId == p.PoolId {
				count++
			}
		}
		noPerp := strconv.Itoa(count)
		rows = append(rows, table.Row{
			id, p.PoolMarginSymbol, noPerp,
		})
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}
