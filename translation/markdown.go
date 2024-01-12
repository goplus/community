/*
 * Copyright (c) 2023 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package translation

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

var (
	RequestFailed = errors.New("request failed")
)

const (
	// APIURL is the url of translation api server
	niuTransAPI = "http://api.niutrans.com/NiuTransServer/translation"

	// RequestHeaders
	contentType = "application/x-www-form-urlencoded;charset=utf-8"
	userAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
)

type Language string

// Language type
const (
	English Language = "en"
	Chinese Language = "zh"
	Auto    Language = "auto"
)

// TranslateConfig is the config of translation server
type TranslateConfig struct {
	apiKey string // api key of translation server
}

// TranslateResponse is the response of translation server
type TranslateResponse struct {
	From    string `json:"from"`
	To      string `json:"to"`
	TgtText string `json:"tgt_text"`
}

// NewTranslateConfig create a new TranslateConfig
func NewTranslateConfig(apiKey string) *TranslateConfig {
	return &TranslateConfig{
		apiKey: apiKey,
	}
}

// TranslateMarkdown translate markdown text
func (c *TranslateConfig) TranslateMarkdown(src string, from, to Language) (string, error) {
	md := goldmark.New(goldmark.WithExtensions())
	reader := text.NewReader([]byte(src))
	doc := md.Parser().Parse(reader)

	pointer := doc.FirstChild()
	for pointer != nil {
		pointer = pointer.NextSibling()
		rawText := extractText(pointer, []byte(src))
		// Translate ast node
		toText, _ := c.TranslateSeq(rawText, "zh", "en")
		_ = toText
	}
	// TODO: Reconstrcut ast node to markdown text

	return "", nil
}

func extractText(n ast.Node, source []byte) string {
	return ""
}

func nodeToMarkdown(n ast.Node, source []byte) string {
	return ""
}

// Translate translate sequence of text
func (c *TranslateConfig) TranslateSeq(src string, from, to Language) (string, error) {
	// Get translate result from api server
	// Request form data
	formData := url.Values{
		"from":     {string(from)},
		"to":       {string(to)},
		"apikey":   {c.apiKey},
		"src_text": {src},
	}

	var req *http.Request
	req, err := http.NewRequest("POST", niuTransAPI, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", RequestFailed
	}

	// Set request headers
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", userAgent)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		return "", RequestFailed
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", RequestFailed
		}

		// Parse response body
		var raw []byte
		if raw, err = io.ReadAll(resp.Body); err != nil {
			return "", RequestFailed
		}

		// Parse json result
		var response TranslateResponse
		if err = json.Unmarshal(raw, &response); err != nil {
			return "", RequestFailed
		}

		return response.TgtText, nil
	}

	return "", RequestFailed
}

// TranslateBatchSeq translate a series of sequences of text
func (c *TranslateConfig) TranslateBatchSeq(src []string, from, to Language) ([]string, error) {
	// Max batch size is 50
	sem := make(chan struct{}, 50)
	defer close(sem)

	// TODO: Max tokens of src is 5000, need to split src
	var wg sync.WaitGroup
	result := make([]string, len(src))
	for i, s := range src {
		wg.Add(1)
		go func(i int, s string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			toText, err := c.TranslateSeq(s, from, to)
			if err != nil {
				result[i] = ""
			} else {
				result[i] = toText
			}
		}(i, s)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return result, nil
}
