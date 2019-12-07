package mailgo

import (
	"fmt"
	"net/http"
	"net/url"
)

type Sender string

type Messenger struct {
	apiKey       string
	senderDomain string
}

type Config struct {
	// APIKey should look like `api-<some hash>`
	APIKey, SenderDomain string
}

func New(conf Config) Messenger {
	return Messenger{
		senderDomain: conf.SenderDomain,
		apiKey:       conf.APIKey,
	}
}

func (m Messenger) GenerateSender(name, emailUser string) Sender {
	return Sender(fmt.Sprintf("%s <%s@%s>", name, emailUser, m.senderDomain))
}

func (m Messenger) Send(subject, to, text string, from Sender) error {
	fullURL := fmt.Sprintf(
		"https://api:%s@api.mailgun.net/v3/samples.mailgun.org/messages",
		m.apiKey,
	)

	_, err := http.PostForm(fullURL, url.Values{
		"from":    {string(from)},
		"to":      {to},
		"subject": {subject},
		"text":    {text},
	})

	return err
}

func (m Messenger) SendHTML(subject, from, to, html string) error {
	// key-<some hash>
	fullURL := fmt.Sprintf(
		"https://api:%s@api.mailgun.net/v3/samples.mailgun.org/messages",
		m.apiKey,
	)

	_, err := http.PostForm(fullURL, url.Values{
		"from":    {from},
		"to":      {to},
		"subject": {subject},
		"html":    {html},
	})

	return err
}
