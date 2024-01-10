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
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type Language string

// TranslateMarkdown translate markdown text
func TranslateMarkdown(src string, from, to Language) (string, error) {

	md := goldmark.New(goldmark.WithExtensions())
	reader := text.NewReader([]byte(src))
	doc := md.Parser().Parse(reader)

	pointer := doc.FirstChild()
	for pointer != nil {
		pointer = pointer.NextSibling()
		raw_text := extractText(pointer, []byte(src))
		// Translate ast node
		to_text, _ := TranslateSeq(raw_text, "zh", "en")
		_ = to_text
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
func TranslateSeq(src string, from, to Language) (string, error) {
	// Get translate result from api server

	return "", nil
}
