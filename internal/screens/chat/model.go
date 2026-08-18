package chat

import (
	"strings"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	bifrost "github.com/maximhq/bifrost/core"
	"github.com/maximhq/bifrost/core/schemas"
	"github.com/terriblethinking/cli/internal/components/viewport"
	"github.com/terriblethinking/engine/agent"
)

type errMsg error

type Turn struct {
	Role    string
	Content string
}

type Focused int

const (
	Chat Focused = iota + 1
	Textarea
)

type Model struct {
	textarea textarea.Model

	agent *agent.Agent

	streamCh <-chan any

	turns []Turn

	incomingThinking *strings.Builder
	incomingText     *strings.Builder

	chatViewport viewport.Model

	state string

	height int
	width  int

	err error
}

func New(client bifrost.Bifrost) Model {

	//  NOTE: TEXTAREA SETUP

	ta := textarea.New()

	ta.ShowLineNumbers = false
	ta.DynamicHeight = true
	ta.MinHeight = 3
	ta.MaxHeight = 10

	// TODO doesn't yet work. Want to add a pretty round
	// border around it. Also would probably be nice to
	// move the styles to the styles package.
	//
	// textareaFocusedStyles := ta.Styles().Focused
	// textareaBlurredStyles := ta.Styles().Blurred
	//
	// newFocusedStyle := lipgloss.NewStyle().
	// 	Border(lipgloss.RoundedBorder()).
	// 	BorderForeground(lipgloss.Color("238")).
	// 	Inherit(textareaFocusedStyles.Base)
	//
	// newBlurredStyle := lipgloss.NewStyle().
	// 	Border(lipgloss.RoundedBorder()).
	// 	BorderForeground(lipgloss.Color("238")).
	// 	Inherit(textareaBlurredStyles.Base)
	//
	// textareaFocusedStyles.Base = newFocusedStyle
	// textareaBlurredStyles.Base = newBlurredStyle
	//
	// ta.SetStyles(textarea.Styles{Focused: textareaFocusedStyles, Blurred: textareaBlurredStyles, Cursor: ta.Styles().Cursor})

	// Here we take the default textarea keymap and modify
	// it. We do this to have `shift+enter` instead of
	// simply `enter`, which will be the key which submits
	// the user input.
	//
	// This is simply the way things have been done in the
	// AI UI feild, and this is how we will do ours.
	keymap := textarea.DefaultKeyMap()
	keymap.InsertNewline = key.NewBinding(key.WithKeys("shift+enter", "ctrl+m"), key.WithHelp("enter", "insert newline"))
	ta.KeyMap = keymap

	// Add some place holder text.
	ta.Placeholder = "today i want to..."

	ta.SetVirtualCursor(false)

	// Set the focus onto the textarea.
	ta.Focus()

	// Set the default textarea styles; adding true
	// will make it be the dark theme variant.

	ta.SetStyles(textarea.DefaultStyles(true))

	//  NOTE: CHAT VIEWPORT SETUP

	chatViewport := viewport.New()

	//  NOTE: AGENT SETUP

	agent01 := agent.Agent{
		Client:   client,
		Provider: "ollama",
		Model:    "ornith:9B",
		Messages: []schemas.ChatMessage{},
	}

	return Model{
		textarea:         ta,
		err:              nil,
		agent:            &agent01,
		chatViewport:     chatViewport,
		incomingThinking: &strings.Builder{},
		incomingText:     &strings.Builder{},
		state:            "idle",
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		tea.RequestBackgroundColor,
	)
}
