package mlservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"book-recommendation-system/recommendations-api/internal/domain/entities"
)

type Client struct {
	baseURL string
}

func New(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

func (c *Client) EmbedBook(book entities.Book) {
	_, err := c.call("/embed/books", book)
	if err != nil {
		log.Printf("ml call failed /embed/books: %v", err)
	}
}

func (c *Client) EmbedUser(user entities.User) {
	_, err := c.call("/embed/users", user)
	if err != nil {
		log.Printf("ml call failed /embed/users: %v", err)
	}
}

func (c *Client) Train(payload map[string]any) (map[string]any, error) {
	return c.call("/train", payload)
}

func (c *Client) Recommend(userID string) (map[string]any, error) {
	return c.call("/recommend", map[string]string{"user_id": userID})
}

func (c *Client) call(path string, payload any) (map[string]any, error) {
	body, _ := json.Marshal(payload)
	resp, err := http.Post(c.baseURL+path, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ml status %d: %s", resp.StatusCode, string(responseBody))
	}

	if len(responseBody) == 0 {
		return map[string]any{"status": "ok"}, nil
	}

	var decoded map[string]any
	if err := json.Unmarshal(responseBody, &decoded); err != nil {
		return map[string]any{"raw": string(responseBody)}, nil
	}
	return decoded, nil
}
