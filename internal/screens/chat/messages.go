package chat

import (
	"github.com/maximhq/bifrost/core/schemas"
	"strings"
)

// This function converts all of the agent turns we store in
// m.turns into a format which we can pass to our agent and
// which it can pass to bifrost AI gateway.

func (m Model) GetAgentReadableMessages() []schemas.ChatMessage {

	// First, we get all of the turns.

	turns := m.turns

	// Initialize the slice that will hold the final messages.

	var messages []schemas.ChatMessage

	// Here we create a var agentMessageBuilder of type strings.Builder. What this
	// will be used for is the assistant's messages, as inside the
	// m.turns we store the reasoning and text output seperately.
	// So when we loop over the turns, we will need to get them back
	// together.

	var agentMessageBuilder strings.Builder

	for _, turn := range turns {

		// Here we check whether the current turn is from the user.
		// If it is, that means that the agent has finished
		// it's output. This could also happen when we are in a new
		// conversation, therefore we add the second condition to check
		// whether there _is_ any agent content.

		if turn.Role == "person" && agentMessageBuilder.String() != "" {

			// We get the text from the strings.Builder var

			agentMessageText := agentMessageBuilder.String()

			// Create the actual schemas.ChatMessage struct

			agentMessage := schemas.ChatMessage{
				Role: "assistant",
				Content: &schemas.ChatMessageContent{
					ContentStr: &agentMessageText,
				},
			}

			// Append to messages

			messages = append(messages, agentMessage)

		} else if turn.Role == "model-thinking" {

			// Here we add <thinking></thinking> to the start and end of the
			// model reasoning. This is because the model is accustomed to
			// seeing them in its messages, and the thinking we recieve doesn't
			// have them.

			agentMessageBuilder.WriteString("<thinking>")
			agentMessageBuilder.WriteString(turn.Content)
			agentMessageBuilder.WriteString("</thinking>")

		} else if turn.Role == "model-text" {

			// If the message is text, we simply append it to the
			// messages slice.

			agentMessageBuilder.WriteString(turn.Content)

		}
	}

	// This is here in case the last message sent isn't a user message,
	// as that would mean that that last message isn't added to the
	// messages slice

	if agentMessageBuilder.String() != "" {
		agentMessageText := agentMessageBuilder.String()

		agentMessage := schemas.ChatMessage{
			Role: "assistant",
			Content: &schemas.ChatMessageContent{
				ContentStr: &agentMessageText,
			},
		}

		messages = append(messages, agentMessage)
	}

	return messages
}
