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
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/stretchr/testify/assert"
	language "golang.org/x/text/language"
)

var (
	// Your api key
	mockTranslationKey string
	mockAccessKey      string
	mockSecretKey      string
	transEngine        *Engine
)

// Mock API Client
type mockNiuTransClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Implement NiuTransClient interface
func (m *mockNiuTransClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestMain(m *testing.M) {
	mockTranslationKey = os.Getenv("NIUTRANS_API_KEY")
	mockAccessKey = os.Getenv("QINIU_ACCESS_KEY")
	mockSecretKey = os.Getenv("QINIU_SECRET_KEY")
	mockTranslationKey = os.Getenv("NIUTRANS_API_KEY")

	transEngine = New(mockTranslationKey, mockAccessKey, mockSecretKey)

	// Replace the default client with mock client
	transEngine.HTTPClient = &mockNiuTransClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// type translateResponse struct {
			// 	From    string `json:"from"`
			// 	To      string `json:"to"`
			// 	TgtText string `json:"tgt_text"`
			// }
			translateResponse := struct {
				From    string `json:"from"`
				To      string `json:"to"`
				TgtText string `json:"tgt_text"`
			}{
				From:    "zh",
				To:      "en",
				TgtText: "Hello",
			}

			// Mock response
			body, err := json.Marshal(translateResponse)
			if err != nil {
				return nil, err
			}

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(string(body))),
			}, nil
		},
	}

	// Replace the default client with mock client
	transEngine.QiNiuSDKClient = &mockQiNiuSDKClient{
		DoFunc: func(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
			method, reqUrl string, headers http.Header, param interface{}) error {
			return nil
		},
	}

	m.Run()
}

func TestTranslatePlainText(t *testing.T) {
	tests := []struct {
		src    string
		from   string
		to     language.Tag
		result string
	}{
		{"你好", "auto", language.English, "Hello"},
		{"hello", "auto", language.Chinese, "你好"},
	}

	for _, test := range tests {
		// Replace the default client with mock client
		transEngine.HTTPClient = &mockNiuTransClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				// type translateResponse struct {
				// 	From    string `json:"from"`
				// 	To      string `json:"to"`
				// 	TgtText string `json:"tgt_text"`
				// }
				translateResponse := struct {
					From    string `json:"from"`
					To      string `json:"to"`
					TgtText string `json:"tgt_text"`
				}{
					From:    "zh",
					To:      "en",
					TgtText: test.result,
				}

				// Mock response
				body, err := json.Marshal(translateResponse)
				if err != nil {
					return nil, err
				}

				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(string(body))),
				}, nil
			},
		}

		to, err := transEngine.TranslatePlain([]byte(test.src), test.from, test.to)
		assert.Nil(t, err)

		toStr := string(to)

		assert.EqualValues(t, test.result, toStr)
	}
}

func TestTranslateMarkdown(t *testing.T) {
	tests := []struct {
		src    string
		from   string
		to     language.Tag
		result string
	}{
		{`# Heading 

## 这是一个二级标题

这是一段普通的文本。

这是一段*粗体*的文本。

这是一段**粗体**的文本。

这是一段***粗斜体***的文本。

这是一段代码块：

这是一个列表：

- 列表项 1
- 列表项 2
	- 列表项 2.1
	- 列表项 2.2
- 列表项 3

这是一个有序列表：

1. 列表项 1
2. 列表项 2
	1. 列表项 2.1
	2. 列表项 2.2
3. 列表项 3

这是一个引用：

> 这是一段引用。

这里一个缩进代码块：

	func main() {
		fmt.Println("Hello, World!")
	}
`,
			"auto",
			language.English,
			`# Heading 

## This is a secondary title

This is an ordinary text.

This is a paragraph*Bold font*Gets or sets the text of the.

This is a paragraph**Bold font**Gets or sets the text of the.

This is a paragraph***Rough italics***Gets or sets the text of the.

This is a block of code:

Here's a list:

- List Item 1
- List Item 2
	- List item 2.1
	- Table item 2.2
- List Item 3

This is an ordered list:

1. List Item 1
2. List Item 2
	1. List item 2.1
	2. Table item 2.2
3. List Item 3

This is a reference:

> This is a quote.

Here is an indented code block:

	func main() {
		fmt.Println("Hello, World!")
	}
`,
		},
	}

	for _, test := range tests {
		// Replace the default client with mock client
		transEngine.HTTPClient = &mockNiuTransClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				// type translateResponse struct {
				// 	From    string `json:"from"`
				// 	To      string `json:"to"`
				// 	TgtText string `json:"tgt_text"`
				// }
				translateResponse := struct {
					From    string `json:"from"`
					To      string `json:"to"`
					TgtText string `json:"tgt_text"`
				}{
					From: "zh",
					To:   "en",
					TgtText: `Heading
9223372036854775807 
This is a secondary title 
9223372036854775807 
This is an ordinary text. 
9223372036854775807 
This is a paragraph 
9223372036854775807 
Bold font 
9223372036854775807 
Gets or sets the text of the. 
9223372036854775807 
This is a paragraph 
9223372036854775807 
Bold font 
9223372036854775807 
Gets or sets the text of the. 
9223372036854775807 
This is a paragraph 
9223372036854775807 
Rough italics 
9223372036854775807 
Gets or sets the text of the. 
9223372036854775807 
This is a block of code: 
9223372036854775807 
Here's a list: 
9223372036854775807 
List Item 1 
9223372036854775807 
List Item 2 
9223372036854775807 
List item 2.1 
9223372036854775807 
Table item 2.2 
9223372036854775807 
List Item 3 
9223372036854775807 
This is an ordered list: 
9223372036854775807 
List Item 1 
9223372036854775807 
List Item 2 
9223372036854775807 
List item 2.1 
9223372036854775807 
Table item 2.2 
9223372036854775807 
List Item 3 
9223372036854775807 
This is a reference: 
9223372036854775807 
This is a quote. 
9223372036854775807 
Here is an indented code block:
					`,
				}

				// Mock response
				body, err := json.Marshal(translateResponse)
				if err != nil {
					return nil, err
				}

				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(string(body))),
				}, nil
			},
		}

		translatedResult, err := transEngine.TranslateMarkdownText(test.src, test.from, test.to)
		assert.Nil(t, err)

		assert.EqualValues(t, test.result, translatedResult)
	}
}

func TestTranslateMarkdownLarge(t *testing.T) {
	// generate a large markdown text
	var largeMd strings.Builder
	for i := 1; i < 5000; i++ {
		largeMd.WriteString("你好")
		if i%500 == 0 {
			largeMd.WriteString("\n")
		}
	}

	tests := []struct {
		src    string
		from   string
		to     language.Tag
		result string
	}{
		{
			largeMd.String(),
			"auto",
			language.English,
			`Hello
Hello
Hello
Hello
Hello
Hello
Hello
Hello
Hello
Hello`,
		},
	}

	for _, test := range tests {
		// Replace the default client with mock client
		transEngine.HTTPClient = &mockNiuTransClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				// type translateResponse struct {
				// 	From    string `json:"from"`
				// 	To      string `json:"to"`
				// 	TgtText string `json:"tgt_text"`
				// }
				translateResponse := struct {
					From    string `json:"from"`
					To      string `json:"to"`
					TgtText string `json:"tgt_text"`
				}{
					From:    "zh",
					To:      "en",
					TgtText: "Hello",
				}

				// Mock response
				body, err := json.Marshal(translateResponse)
				if err != nil {
					return nil, err
				}

				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(string(body))),
				}, nil
			},
		}

		translatedResult, err := transEngine.TranslateMarkdownText(test.src, test.from, test.to)
		assert.Nil(t, err)

		assert.EqualValues(t, test.result, translatedResult)
	}
}

func TestTranslateBatch(t *testing.T) {
	tests := []struct {
		src    []string
		from   string
		to     language.Tag
		result []string
	}{
		{[]string{"你好"}, "auto", language.English, []string{"Hello"}},
		{[]string{"What 's your name"}, "auto", language.Chinese, []string{"你叫什么名字"}},
	}

	for _, test := range tests {
		// Replace the default client with mock client
		transEngine.HTTPClient = &mockNiuTransClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				// type translateResponse struct {
				// 	From    string `json:"from"`
				// 	To      string `json:"to"`
				// 	TgtText string `json:"tgt_text"`
				// }
				translateResponse := struct {
					From    string `json:"from"`
					To      string `json:"to"`
					TgtText string `json:"tgt_text"`
				}{
					From:    "zh",
					To:      "en",
					TgtText: test.result[0],
				}

				// Mock response
				body, err := json.Marshal(translateResponse)
				if err != nil {
					return nil, err
				}

				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(string(body))),
				}, nil
			},
		}
		translatedResult, err := transEngine.TranslateBatchPlain(test.src, test.from, test.to)
		assert.Nil(t, err)

		assert.EqualValues(t, test.result, translatedResult)
	}
}

func TestTranslateWebVTT(t *testing.T) {
	tests := []struct {
		src    string
		from   string
		to     language.Tag
		result string
	}{
		{"WEBVTT FILE\n\n1\n00:00:00.000 --> 00:00:03.000\n你好，世界！\n\n", "auto", language.English, "WEBVTT FILE\n\n1\n00:00:00.000 --> 00:00:03.000\nHello, world!\n\n"},
	}

	for _, test := range tests {
		// Replace the default client with mock client
		transEngine.HTTPClient = &mockNiuTransClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				// type translateResponse struct {
				// 	From    string `json:"from"`
				// 	To      string `json:"to"`
				// 	TgtText string `json:"tgt_text"`
				// }
				translateResponse := struct {
					From    string `json:"from"`
					To      string `json:"to"`
					TgtText string `json:"tgt_text"`
				}{
					From:    "zh",
					To:      "en",
					TgtText: "Hello, world!",
				}

				// Mock response
				body, err := json.Marshal(translateResponse)
				if err != nil {
					return nil, err
				}

				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(string(body))),
				}, nil
			},
		}

		vttFile, err := NewSimpleWebVTT(strings.NewReader(test.src))
		assert.Nil(t, err)

		err = transEngine.TranslateWebVTT(&vttFile, test.from, test.to)
		assert.Nil(t, err)

		translatedResult := vttFile.ToString()
		assert.EqualValues(t, test.result, translatedResult)
	}
}

func TestTranslateWebVTTLarge(t *testing.T) {
	// generate a large vtt file
	var largeVTT strings.Builder
	largeVTT.WriteString("WEBVTT FILE\n\n")
	largeVTT.WriteString("1\n00:00:00.000 --> 00:00:03.000\n")
	for i := 0; i < 2000; i++ {
		largeVTT.WriteString("你好，世界！")
	}
	largeVTT.WriteString("\n\n")

	tests := []struct {
		src    string
		from   string
		to     language.Tag
		result string
	}{
		{
			largeVTT.String(),
			"auto",
			language.English,
			"WEBVTT FILE\n\n1\n00:00:00.000 --> 00:00:03.000\nHello, world!\n\n",
		},
	}

	for _, test := range tests {
		// Replace the default client with mock client
		transEngine.HTTPClient = &mockNiuTransClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				// type translateResponse struct {
				// 	From    string `json:"from"`
				// 	To      string `json:"to"`
				// 	TgtText string `json:"tgt_text"`
				// }
				translateResponse := struct {
					From    string `json:"from"`
					To      string `json:"to"`
					TgtText string `json:"tgt_text"`
				}{
					From:    "zh",
					To:      "en",
					TgtText: "Hello, world!",
				}

				// Mock response
				body, err := json.Marshal(translateResponse)
				if err != nil {
					return nil, err
				}

				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(string(body))),
				}, nil
			},
		}

		vttFile, err := NewSimpleWebVTT(strings.NewReader(test.src))
		assert.Nil(t, err)

		err = transEngine.TranslateWebVTT(&vttFile, test.from, test.to)
		assert.Nil(t, err)

		translatedResult := vttFile.ToString()
		assert.EqualValues(t, test.result, translatedResult)
	}
}

func TestTranslateEmpty(t *testing.T) {
	tests := []struct {
		src    string
		from   string
		to     language.Tag
		result string
	}{
		{"", "auto", language.English, ""},
		{"", "auto", language.Chinese, ""},
	}

	for _, test := range tests {
		// Replace the default client with mock client
		transEngine.HTTPClient = &mockNiuTransClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				// type translateResponse struct {
				// 	From    string `json:"from"`
				// 	To      string `json:"to"`
				// 	TgtText string `json:"tgt_text"`
				// }
				translateResponse := struct {
					From    string `json:"from"`
					To      string `json:"to"`
					TgtText string `json:"tgt_text"`
				}{
					From:    "zh",
					To:      "en",
					TgtText: "",
				}

				// Mock response
				body, err := json.Marshal(translateResponse)
				if err != nil {
					return nil, err
				}

				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(string(body))),
				}, nil
			},
		}

		translatedResult, err := transEngine.TranslateMarkdownText(test.src, test.from, test.to)

		assert.Nil(t, err)
		assert.EqualValues(t, test.result, translatedResult)
	}
}
