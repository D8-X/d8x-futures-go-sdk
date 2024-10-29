package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/pkg/d8x_futures"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lgtable "github.com/charmbracelet/lipgloss/table"
	"github.com/ethereum/go-ethereum/common"
)

type Model struct {
	ChainConfig         utils.ChainConfig
	PriceConfig         utils.PriceFeedConfig
	BlockChainConnector d8x_futures.BlockChainConnector
	XchInfo             d8x_futures.StaticExchangeInfo
	PoolState           []d8x_futures.PoolState
	chainConfigNames    []string
	choices             []ScreenChoices
	selectedNetworkName string
	selectedPoolId      int32
	selectedPerpId      int32
	screen              int
	poolTable           table.Model
	perpTable           table.Model
	posTableStr         string
	perpState           d8x_futures.PerpetualState
	traderInput         textinput.Model
	traderAddr          common.Address
	PositionRisk        d8x_futures.PositionRisk
}

const LAST_PAGE int = 4
const (
	orange        = lipgloss.Color("#FCA43C")
	blue          = lipgloss.Color("#2786E4")
	blueHighlight = lipgloss.Color("#2353C8")
	red           = lipgloss.Color("#CC5150")
	green         = lipgloss.Color("#06BC42")
	gray          = lipgloss.Color("#CDCDCD")
	lightGray     = lipgloss.Color("#E6E6E6")
)

var re = lipgloss.NewRenderer(os.Stdout)
var (
	focusedStyle = lipgloss.NewStyle().Foreground(blueHighlight)
	cursorStyle  = focusedStyle.Copy()
	baseStyle    = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(blue)

	// HeaderStyle is the lipgloss style used for the table headers.
	HeaderStyle = re.NewStyle().Foreground(orange).Bold(true).Align(lipgloss.Center)
	// CellStyle is the base lipgloss style used for the table rows.
	CellStyle = re.NewStyle().Padding(0, 1).Width(14)
	// OddRowStyle is the lipgloss style used for odd-numbered table rows.
	OddRowStyle = CellStyle.Copy().Foreground(gray)
	// EvenRowStyle is the lipgloss style used for even-numbered table rows.
	EvenRowStyle = CellStyle.Copy().Foreground(lightGray)
	// BorderStyle is the lipgloss style used for the table border.
	BorderStyle = lipgloss.NewStyle().Foreground(orange)
)

type ScreenChoices struct {
	chooseOptions []string
	choiceName    string
	cursor        int
}

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
	// input for key/address
	m.traderInput = textinput.New()
	m.traderInput.Cursor.Style = cursorStyle
	m.traderInput.CharLimit = 66
	m.traderInput.Placeholder = "0xCaFe..."
	m.traderInput.Focus()
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// https://github.com/charmbracelet/bubbletea/blob/master/examples/views/main.go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil

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
			if m.screen == 1 {
				if m.choices[0].cursor > 0 {
					m.choices[0].cursor--
				}
				break
			}
			if m.screen == 2 {
				m.poolTable.MoveUp(1)
			}
			if m.screen == 3 {
				m.perpTable.MoveUp(1)
			}

		// The "down" and space " " keys move the cursor down
		case "down", " ":
			if m.screen == 1 {
				if m.choices[0].cursor < len(m.choices[0].chooseOptions)-1 {
					m.choices[0].cursor++
				}
			}
			if m.screen == 2 {
				m.poolTable.MoveDown(1)
			}
			if m.screen == 3 {
				m.perpTable.MoveDown(1)
			}
		// Spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "left":
			m.screen -= 1
			m.screen = max(m.screen, 0)
		case "right", "enter":
			m.screen += 1
			m.screen = min(m.screen, LAST_PAGE)
			if m.screen == 1 {
				hexStr := m.traderInput.Value()
				m.traderAddr = common.HexToAddress(hexStr)
				break
			}
			if m.screen == 2 {
				err := m.actionScreen12()
				if err != nil {
					m.screen--
					fmt.Println("An error occurred: " + err.Error())
				}
				break
			}
			if m.screen == 3 {
				// pool choice
				m.selectedPoolId = int32(m.poolTable.Cursor() + 1)
				err := m.actionScreen23()
				if err != nil {
					m.screen--
					fmt.Println("An error occurred: " + err.Error())
				}
				break
			}
			if m.screen == 4 {
				m.selectedPerpId = m.selectedPoolId*100000 + int32(m.perpTable.Cursor())
				err := m.actionScreen34()
				if err != nil {
					m.screen--
					fmt.Println("An error occurred: " + err.Error())
				}
				break
			}
		}
	}
	if m.screen == 0 {
		m.traderInput, cmd = m.traderInput.Update(msg)
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m Model) View() string {
	s := ""
	switch m.screen {
	case 0:
		s += m.setTraderView()
	case 1:
		s += m.selectNetworkView()
	case 2:
		s += m.poolView()
	case 3:
		s += m.perpView()
	case 4:
		s += m.perpDetailsView()
	}
	return s
}

func (m Model) setTraderView() string {
	s := topBarStatus("[0] Set trader address") + "\n\n"
	s = s + m.traderInput.View()
	s = s + "\n" + bottomBarStatus(0)
	return s
}

func (m Model) selectNetworkView() string {
	// The header
	s := topBarStatus("[1] Select network")

	s = s + "\n\n" + m.displayChoiceMenu()
	s = s + bottomBarStatus(1)

	// Send the UI for rendering
	return s
}

func (m *Model) poolView() string {
	screen := "[" + strconv.Itoa(m.screen) + "] "
	return topBarStatus(screen+"Connected to "+m.selectedNetworkName) + "\n" +
		baseStyle.Render(m.poolTable.View()) + "\n" + bottomBarStatus(2)
}

func (m *Model) perpView() string {
	screen := "[" + strconv.Itoa(m.screen) + "] "
	p := strconv.Itoa(int(m.selectedPoolId))
	return topBarStatus(screen+"Connected to "+m.selectedNetworkName+" - Pool "+p) + "\n" +
		baseStyle.Render(m.perpTable.View()) + "\n" +
		bottomBarStatus(3)
}

func (m *Model) perpDetailsView() string {
	screen := "[" + strconv.Itoa(m.screen) + "] "
	pool := " Pool " + strconv.Itoa(int(m.selectedPoolId))
	sym := m.XchInfo.PerpetualIdToSymbol[m.selectedPerpId]
	perp := " Perp " + strconv.Itoa(int(m.selectedPerpId)) + " " + sym

	s := topBarStatus(screen+"Connected to "+m.selectedNetworkName+pool+" "+m.traderAddr.Hex()) + "\n\n"
	var styleA = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#0A0A0A")).
		Background(green)
	var styleB = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#0A0A0A")).
		Background(red)
	var mkt string = perp + " "
	if m.perpState.IsMarketClosed {
		mkt = styleB.Render(mkt + "market closed")
	} else {
		mkt = styleA.Render(mkt + "market open")
	}
	s += mkt + "\n"
	const width = 78

	var style = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(blue)

	idx := style.Render(fmt.Sprintf("Index Price\n%.4f", m.perpState.IndexPrice))
	mark := style.Render(fmt.Sprintf("Mark Price\n%.4f", m.perpState.MarkPrice))
	mid := style.Render(fmt.Sprintf("Mid Price\n%.4f", m.perpState.MidPrice))
	fnd := style.Render(fmt.Sprintf("Funding Rate (bps)\n%.2f", m.perpState.CurrentFundingRateBps*10000))
	oi := style.Render(fmt.Sprintf("Open Interest\n%.4f", m.perpState.OpenInterestBC))
	s += lipgloss.JoinHorizontal(lipgloss.Top, mid, mark, idx, fnd, oi)
	s += "\n\n" + " Open Position"
	s += "\n" + m.posTableStr + "\n"

	s += "\n" + bottomBarStatus(4)
	return s
}

func topBarStatus(txt string) string {
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#0A0A0A")).Background(orange)
	return style.Render(txt)
}

func bottomBarStatus(pageNo int) string {
	activeDot := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	inactiveDot := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	var s string = "\n"
	for k := 0; k <= LAST_PAGE; k++ {
		if k == pageNo {
			s = s + activeDot
		} else {
			s = s + inactiveDot
		}
	}
	s = s + "\n\n←/→ page | ↓/↑/space: select | ctrl-c/q: quit\n"
	return s
}

func (m Model) displayChoiceMenu() string {
	var s string
	activeStyle := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"})
	inactiveStyle := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"})
	// Iterate over our choices
	for i, choice := range m.choices[0].chooseOptions {

		// Is the cursor pointing at this choice?
		cursor := "⚪" // not selected
		var choiceStr string
		if m.choices[0].cursor == i {
			cursor = lipgloss.NewStyle().Foreground(blue).Render("⚫") // selected!
			choiceStr = activeStyle.Render(choice)
		} else {
			choiceStr = inactiveStyle.Render(choice)
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choiceStr)
	}
	return s
}

func (m *Model) actionScreen12() error {
	// load selected chainconfig
	idx := m.choices[0].cursor
	network := m.choices[0].chooseOptions[idx]
	m.selectedNetworkName = network
	fmt.Print("selecting network " + network)
	chConf, err := config.GetDefaultChainConfig(network)
	if err != nil {
		return errors.New("connection failed")
	}
	m.ChainConfig = chConf
	pxConf, err := config.GetDefaultPriceConfig(m.ChainConfig.ChainId)
	if err != nil {
		return errors.New("connection failed")
	}
	conn, err := d8x_futures.CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		return errors.New("connection failed")
	}
	m.BlockChainConnector = conn
	nest, err := d8x_futures.QueryNestedPerpetualInfo(conn)
	if err != nil {
		return err
	}
	m.XchInfo, _ = d8x_futures.QueryExchangeStaticInfo(&conn, &chConf, &pxConf, &nest)
	m.PoolState, err = d8x_futures.RawQueryPoolStates(conn.Rpc, m.XchInfo)
	if err != nil {
		return err
	}
	m.poolTable = createPoolTable(m.XchInfo, m.PoolState)

	return nil
}

func (m *Model) actionScreen23() error {
	m.perpTable = createPerpTable(m.XchInfo, m.selectedPoolId)
	return nil
}

func (m *Model) actionScreen34() error {
	ids := []int32{m.selectedPerpId}
	s, err := d8x_futures.RawQueryPerpetualState(m.BlockChainConnector.Rpc,
		m.XchInfo,
		ids,
		m.ChainConfig.PriceFeedEndpoint,
		m.ChainConfig.PrdMktFeedEndpoint,
		m.ChainConfig.LowLiqFeedEndpoint,
	)
	if err != nil {
		return err
	}
	sym := m.XchInfo.PerpetualIdToSymbol[m.selectedPerpId]
	err = m.setPositionRisk(sym)
	if err != nil {
		return err
	}
	m.perpState = s[0]
	m.posTableStr = createPositionTable(m.PositionRisk)
	return nil
}

func (m *Model) setPositionRisk(symbol string) error {
	pRisk, err := d8x_futures.RawGetPositionRisk(
		m.XchInfo,
		m.BlockChainConnector.Rpc,
		&m.traderAddr,
		symbol,
		m.ChainConfig.PriceFeedEndpoint,
		m.ChainConfig.PrdMktFeedEndpoint,
		m.ChainConfig.LowLiqFeedEndpoint,
	)
	if err != nil {
		return err
	}
	m.PositionRisk = pRisk
	return nil
}

func createPositionTable(pos d8x_futures.PositionRisk) string {

	size := pos.PositionNotionalBaseCCY
	if size != 0 && pos.Side != d8x_futures.SIDE_BUY {
		size = size * -1
	}
	syms := strings.Split(pos.Symbol, "-")
	//margin := (size*(pos.MarkPrice-pos.EntryPrice)+pos.UnrealizedPnlQuoteCCY)*
	//	1/pos.CollToQuoteConversion + pos.CollateralCC
	margin := pos.CollateralCC
	rows := [][]string{
		{"Size", fmt.Sprintf("%.4f", size), syms[0]},
		{"Entry Price", fmt.Sprintf("%.4f", pos.EntryPrice), syms[1]},
		{"Liq. Price", fmt.Sprintf("%.4f", pos.LiquidationPrice[0]), syms[1]},
		{"Margin", fmt.Sprintf("%.4f", margin), syms[2]},
		{"Leverage", fmt.Sprintf("%.2f", pos.Leverage), "-"},
		{"Unreal.P&L", fmt.Sprintf("%.4f", pos.UnrealizedPnlQuoteCCY), syms[1]},
	}

	t := lgtable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(blue)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Headers("Attribute", "Value", "Unit").
		Rows(rows...)
	return fmt.Sprintln(t)
}

func createPoolTable(info d8x_futures.StaticExchangeInfo, poolSt []d8x_futures.PoolState) table.Model {
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Coll.Tkn", Width: 8},
		{Title: "#Perps", Width: 6},
		{Title: "Vault $Coll", Width: 15},
		{Title: "DF $Coll", Width: 15},
	}

	var rows []table.Row
	for k := 0; k < len(info.Pools); k++ {
		p := info.Pools[k]
		id := strconv.Itoa(int(p.PoolId))
		vaultCash := fmt.Sprintf("%.1f", poolSt[k].PnlParticipantCashCC)
		dfCash := fmt.Sprintf("%.1f", poolSt[k].DefaultFundCashCC)
		count := 0
		for j := 0; j < len(info.Perpetuals); j++ {
			if info.Perpetuals[j].PoolId == p.PoolId {
				count++
			}
		}
		noPerp := strconv.Itoa(count)
		rows = append(rows, table.Row{
			id, p.PoolMarginSymbol, noPerp, vaultCash, dfCash,
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
		Background(orange).
		Bold(false)
	t.SetStyles(s)
	return t
}

func createPerpTable(info d8x_futures.StaticExchangeInfo, poolId int32) table.Model {
	columns := []table.Column{
		{Title: "Id", Width: 10},
		{Title: "Symbol", Width: 20},
		{Title: "MaxLvg", Width: 6},
		{Title: "LotSize", Width: 8},
	}
	var rows []table.Row
	for k := 0; k < len(info.Perpetuals); k++ {
		p := info.Perpetuals[k]
		if info.Perpetuals[k].PoolId != poolId {
			continue
		}
		id := strconv.Itoa(int(p.Id))
		sym := info.PerpetualIdToSymbol[p.Id]
		lot := fmt.Sprintf("%.4f", p.LotSizeBC)
		lvg := fmt.Sprintf("%.0f", 1/p.InitialMarginRate)
		rows = append(rows, table.Row{
			id, sym, lvg, lot,
		})
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
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
		Background(orange).
		Bold(false)
	t.SetStyles(s)
	return t
}
