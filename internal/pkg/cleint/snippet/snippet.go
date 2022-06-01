package snippet

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Snippet struct {
	SnippetID string `json:"snippet_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
type List struct {
	Data []*Snippet `json:"data"`
}

type Ping struct {
	Status string `json:"status"`
}

type Client interface {
	CreateSnippet(title string, content string) (string, error)
	GetSnippets() ([]*Snippet, error)
	GetSnippet(snippetID string) (*Snippet, error)
	Ping() (string, error)
}

type clientImpl struct {
	Scheme     string
	Host       string
	Port       int
	SigningKey string
}

func (c clientImpl) Addr() string {
	return fmt.Sprintf("%s://%s:%d", c.Scheme, c.Host, c.Port)
}

func NewClient(host string, port int, signingKey string) Client {
	return &clientImpl{
		Scheme:     "http",
		Host:       host,
		Port:       port,
		SigningKey: signingKey,
	}
}

func NewHTTPSClient(host string, port int, signingKey string) Client {
	return &clientImpl{
		Scheme:     "https",
		Host:       host,
		Port:       port,
		SigningKey: signingKey,
	}
}

func (c clientImpl) Ping() (string, error) {
	url := fmt.Sprintf("%s/api/v1/ping", c.Addr())
	resp, err := resty.New().R().Get(url)

	if err != nil {
		return "", err
	}

	respBody := resp.Body()
	var ping Ping
	err = json.Unmarshal(respBody, &ping)

	if err != nil {
		return "", err
	}

	return ping.Status, nil
}

func (c clientImpl) CreateSnippet(title string, content string) (string, error) {
	url := fmt.Sprintf("%s/api/v1/snippets", c.Addr())

	body := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		Title:   title,
		Content: content,
	}

	token, err := createToken(c.SigningKey)

	if err != nil {
		return "", err
	}

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", token)).
		SetBody(body).
		SetResult(&Snippet{}).
		Post(url)

	if err != nil {
		return "", err
	}

	respBody := resp.Body()
	var snippet Snippet
	err = json.Unmarshal(respBody, &snippet)

	if err != nil {
		return "", err
	}

	return snippet.SnippetID, nil
}

func (c clientImpl) GetSnippets() ([]*Snippet, error) {
	url := fmt.Sprintf("%s/api/v1/snippets", c.Addr())
	token, err := createToken(c.SigningKey)

	if err != nil {
		return nil, err
	}

	resp, err := resty.New().R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", token)).
		Get(url)

	if err != nil {
		return nil, err
	}

	respBody := resp.Body()
	var snippets List
	err = json.Unmarshal(respBody, &snippets)

	if err != nil {
		return nil, err
	}

	return snippets.Data, nil
}

func (c clientImpl) GetSnippet(snippetID string) (*Snippet, error) {

	url := fmt.Sprintf("%s/api/v1/snippets/{snippetId}", c.Addr())
	token, err := createToken(c.SigningKey)

	if err != nil {
		return nil, err
	}

	resp, err := resty.New().R().SetPathParams(map[string]string{
		"snippetId": snippetID,
	}).
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", token)).
		Get(url)

	if err != nil {
		return nil, err
	}

	respBody := resp.Body()
	var snippet Snippet
	err = json.Unmarshal(respBody, &snippet)

	if err != nil {
		return nil, err
	}

	return &snippet, nil
}

func createToken(signingKey string) (string, error) {
	signingKeyBytes := []byte(signingKey)

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Issuer:    "private",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKeyBytes)
}
