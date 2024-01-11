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
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

const (
	// Your api key
	mockKey = "xxx"
)

func TestTranslateSeq(t *testing.T) {
	tests := []struct {
		src  string
		from Language
		to   Language
	}{
		{"你好", "zh", "en"},
		{"hello", "en", "zh"},
	}

	trans := NewTranslateConfig(mockKey)
	for _, test := range tests {
		_, err := trans.TranslateSeq(test.src, test.from, test.to)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestTranslateMarkdown(t *testing.T) {
	tests := []struct {
		src  string
		from Language
		to   Language
	}{
		{`# Hello`, "en", "zh"},
		{`# 你好`, "zh", "en"},
	}

	trans := NewTranslateConfig(mockKey)
	for _, test := range tests {
		_, err := trans.TranslateMarkdown(test.src, test.from, test.to)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestExtractText(t *testing.T) {
	tests := []struct {
		src string
	}{
		{`# Title

		This is a paragraph.
		This is a paragraph.
		This is a paragraph.
		This is a paragraph.
		
		- This is a list item.`,
		},
		{`# Title

		This is a paragraph.
		
		- This is a list item.`,
		},
	}

	for _, test := range tests {
		markdown := goldmark.New(goldmark.WithExtensions())
		reader := text.NewReader([]byte(test.src))
		doc := markdown.Parser().Parse(reader)

		fmt.Printf("%#v\n", doc)
	}
}
