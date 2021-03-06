// Incoming message builder.
package bearychat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Incoming struct {
	Text     string `json:"text"`
	Markdown bool   `json:"markdown"`
	Channel  string `json:"channel,omitempty"`

	Attachments []IncomingAttachment `json:"attachments,omitempty"`
}

type IncomingAttachment struct {
	Title  string          `json:"title,omitempty"`
	Text   string          `json:"text,omitempty"`
	Url    string          `json:"url,omitempty"`
	Color  string          `json:"color,omitempty"`
	Images []IncomingImage `json:"images,omitempty"`
}

type IncomingImage struct {
	Url string `json:"url"`
}

func (m Incoming) Build() (io.Reader, error) {
	if m.Text == "" {
		return nil, fmt.Errorf("text is required")
	}

	for _, attachment := range m.Attachments {
		if attachment.Title == "" && attachment.Text == "" {
			return nil, fmt.Errorf("title or text is required")
		}
	}

	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}
