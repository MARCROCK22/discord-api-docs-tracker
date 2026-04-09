package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Fetch(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	return response, nil
}

func FetchWithBody(method string, url string, body map[string]any, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	return response, nil
}

func SendEmbed(embed map[string]any, pullRequestType string, webhooks []string) error {
	for _, webhook := range webhooks {
		_, err := FetchWithBody("POST", webhook, map[string]any{
			"embeds":   []map[string]any{embed},
			"username": "api-docs",
		}, nil)
		if err != nil {
			log.Printf("failed to send embed to webhook %s: %v\n", webhook, err)
		}
	}

	return nil
}
