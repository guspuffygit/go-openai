package openai

import (
	"context"
	"net/http"
)

const (
	tokenizeSuffix   = "/tokenize"
	detokenizeSuffix = "/detokenize"
)

type ChatTokenizeRequest struct {
	Model                string                  `json:"model"`
	Messages             []ChatCompletionMessage `json:"messages"`
	AddGenerationPrompt  bool                    `json:"add_generation_prompt,omitempty"`
	AddSpecialTokens     bool                    `json:"add_special_tokens,omitempty"`
	ContinueFinalMessage bool                    `json:"continue_final_message,omitempty"`
}

type TextTokenizeRequest struct {
	Model            string `json:"model"`
	Prompt           string `json:"prompt"`
	AddSpecialTokens bool   `json:"add_special_tokens,omitempty"`
}

type TokenizeResponse struct {
	Count          int   `json:"count"`
	MaxModelLength int   `json:"max_model_len"`
	Tokens         []int `json:"tokens"`

	httpHeader
}

type ChatDetokenizeRequest struct {
	Model  string `json:"model"`
	Tokens []int  `json:"tokens"`
}

type ChatDetokenizeResponse struct {
	Prompt string `json:"prompt"`

	httpHeader
}

func (c *Client) CreateChatTokenize(
	ctx context.Context,
	request ChatTokenizeRequest,
) (response TokenizeResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(tokenizeSuffix, request.Model), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) CreateTextTokenize(
	ctx context.Context,
	request TextTokenizeRequest,
) (response TokenizeResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(tokenizeSuffix, request.Model), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) CreateChatDetokenize(
	ctx context.Context,
	request ChatDetokenizeRequest,
) (response ChatDetokenizeResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(detokenizeSuffix, request.Model), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
