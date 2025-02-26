package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("15"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "-b":
			fmt.Println("\nReturning to main menu...")
			Start()
		case "-s":
			fmt.Println("\nStoring API keys...")
			StoreAPI()
		case "-r":
			fmt.Println("\nRetrieving API Keys...")
			retrieveAPIKeys()
		case "-v":
			fmt.Println("\nFetching Current Version...")
			version()
		case "-d":
			fmt.Println("\nDeleting API Keys...")
			deleteAPIKeys()
		case "-l":
			fmt.Println("\nListing stored API keys...")
			listEncryptedKeys()
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}

func Help() {
	ClearTerm()

	fmt.Println("════════════════════════════════════════════════════════════════════════")
	fmt.Println("             HELP MENU              ")
	fmt.Println("════════════════════════════════════════════════════════════════════════")
	fmt.Println("Usage: ./app [COMMAND]")

	items := []list.Item{
		item(`-s, --store      Encrypt and store a new API key`),
		item(`-r, --retrieve   Decrypt and display stored API keys`),
		item(`-d, --delete     Delete a stored API key`),
		item(`-l, --list       Show encrypted API keys`),
		item(`-h, --help       Display this help menu`),
		item(`-v, --version    Show application version`),
		item(`-b, --back       Return to the main menu`),
	}

	const defaultWidth = 25

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "AVAILABLE COMMANDS"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	fmt.Println("\nNotes:")
	fmt.Println(`
  - API keys are securely encrypted using AES-256-GCM.
  - A correct password is required to decrypt stored keys.
  - Encrypted keys are saved in 'api_keys.json' for easy retrieval.

  For more details, visit:
  https://github.com/Btylrob/APIKitten
	`)

	fmt.Println("════════════════════════════════════════════════════════════════════════")

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running help menu:", err)
		os.Exit(1)
	}
}
