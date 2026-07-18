package chat

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
)

type errMsg error

type Model struct {
	textarea textarea.Model

	height int
	width  int

	err error
}

func New() Model {
	ta := textarea.New()

	ta.ShowLineNumbers = false
	ta.DynamicHeight = true
	ta.MinHeight = 1
	ta.MaxHeight = 4

	// Set the default textarea styles; adding true
	// will make it be the dark theme alternative.
	ta.SetStyles(textarea.DefaultStyles(true))

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

	return Model{
		textarea: ta,
		err:      nil,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		tea.RequestBackgroundColor,
	)
}
