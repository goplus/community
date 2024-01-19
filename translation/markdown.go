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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	language "golang.org/x/text/language"
)

var (
	ErrRequestFailed       = errors.New("request failed")
	ErrTranslationNotMatch = errors.New("translation result length not match")
)

const (
	// APIURL is the url of translation api server
	niuTransAPI = "http://api.niutrans.com/NiuTransServer/translation"

	// Max content length of request
	maxContentLength = 5000

	// RequestHeaders
	contentType = "application/x-www-form-urlencoded;charset=utf-8"
	userAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
)

// Engine is the config of translation server
type Engine struct {
	translationAPIKey string // api key of translation server
	qiniuAccessKey    string // access key of qiniu
	qiniuSecretKey    string // secret key of qiniu
}

// translateResponse is the response of translation server
type translateResponse struct {
	From    string `json:"from"`
	To      string `json:"to"`
	TgtText string `json:"tgt_text"`
}

// New create a new TranslateConfig
func New(translationAPIKey string, qiniuAccessKey string, qiniuSecretKey string) *Engine {
	return &Engine{
		translationAPIKey: translationAPIKey,
		qiniuAccessKey:    qiniuAccessKey,
		qiniuSecretKey:    qiniuSecretKey,
	}
}

// TranslateMarkdown translate markdown with bytes
func (c *Engine) TranslateMarkdown(src []byte, from, to language.Tag) (ret []byte, err error) {
	// Init markdown parser
	md := goldmark.New(goldmark.WithExtensions())
	reader := text.NewReader(src)
	doc := md.Parser().Parse(reader)

	// Prepare translation
	translationSep := generateSeparator()
	translationVec := make([]string, 0)
	translationSeg := make([]text.Segment, 0) // location of text in src

	// Walk through the AST
	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		switch n := node.(type) {
		case *ast.Text:
			if entering {
				translationVec = append(translationVec, string(n.Segment.Value(src)))
				translationSeg = append(translationSeg, n.Segment)
			}
		}
		// TODO: Skip for custom node in goplus markdown

		return ast.WalkContinue, nil
	})

	// Translate
	// toBeTranslatedStr := strings.Join(translationVec, fmt.Sprintf("\n%s\n", translationSep))
	var translatedStr string
	// Judge whether to use batch translate(over 5000 characters)
	toBeTranslatedStrList := joinWithMaxLength(translationVec, fmt.Sprintf("\n%s\n", translationSep), maxContentLength)
	if len(toBeTranslatedStrList) == 1 {
		translatedStr, err = c.TranslatePlainText(toBeTranslatedStrList[0], from, to)
		if err != nil {
			return nil, err
		}
	} else {
		translatedStrList, err := c.TranslateBatchPlain(toBeTranslatedStrList, from, to)
		if err != nil {
			return nil, err
		}

		translatedStr = strings.Join(translatedStrList, fmt.Sprintf("\n%s\n", translationSep))
	}

	// Replace text
	resultVec := strings.Split(translatedStr, translationSep)
	if len(resultVec) != len(translationVec) {
		return nil, ErrTranslationNotMatch
	}

	result := make([]byte, 0)
	result = append(result, src[:translationSeg[0].Start]...)
	result = append(result, []byte(resultVec[0])...)
	for idx := 1; idx < len(resultVec); idx++ {
		result = append(result, src[translationSeg[idx-1].Stop:translationSeg[idx].Start]...)
		result = append(result, []byte(strings.TrimSpace(resultVec[idx]))...)
	}
	result = append(result, src[translationSeg[len(translationSeg)-1].Stop:]...)

	return result, nil
}

// TranslateMarkdown translate markdown text
func (c *Engine) TranslateMarkdownText(src string, from, to language.Tag) (ret string, err error) {
	retByte, err := c.TranslateMarkdown([]byte(src), from, to)

	return string(retByte), err
}

// joinWithMaxLength join []string with sep, and make sure the length of joined string is less than maxLen
// if the length of joined string is greater than maxLen, create a new string with sep
// and append it to the result
func joinWithMaxLength(src []string, sep string, maxLen int) []string {
	ret := make([]string, 0)
	cur := ""
	for _, s := range src {
		if len(cur)+len(s)+len(sep) > maxLen {
			// start to create new string
			ret = append(ret, cur)
			cur = ""
		}

		if cur == "" {
			cur = s
		} else {
			cur = fmt.Sprintf("%s%s%s", cur, sep, s)
		}
	}

	ret = append(ret, cur)

	return ret
}

func generateSeparator() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// Translate translate sequence of bytes
func (c *Engine) TranslatePlain(src []byte, from, to language.Tag) (ret []byte, err error) {
	retString, err := c.TranslatePlainText(string(src), from, to)

	return []byte(retString), err
}

// Translate translate sequence of text
func (c *Engine) TranslatePlainText(src string, from, to language.Tag) (ret string, err error) {
	// Get translate result from api server
	// Request form data
	formData := url.Values{
		"from":              {from.String()},
		"to":                {to.String()},
		"translationAPIKey": {c.translationAPIKey},
		"src_text":          {src},
	}

	var req *http.Request
	req, err = http.NewRequest("POST", niuTransAPI, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}

	// Set request headers
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrRequestFailed
	}

	// Parse response body
	var raw []byte
	if raw, err = io.ReadAll(resp.Body); err != nil {
		return "", err
	}

	// Parse json result
	var response translateResponse
	if err = json.Unmarshal(raw, &response); err != nil {
		return "", err
	}

	return response.TgtText, nil
}

// TranslateBatchSeq translate a series of sequences of text
func (c *Engine) TranslateBatchPlain(src []string, from, to language.Tag) ([]string, error) {
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

			toText, err := c.TranslatePlainText(s, from, to)
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
