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
	"net/http"
	"testing"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/stretchr/testify/assert"
	language "golang.org/x/text/language"
)

// Mock API Client
type mockQiNiuSDKClient struct {
	DoFunc func(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
		method, reqUrl string, headers http.Header, param interface{}) error
}

// Implement NiuTransClient interface
func (m *mockQiNiuSDKClient) CredentialedCallWithJson(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
	method, reqUrl string, headers http.Header, param interface{}) error {
	return m.DoFunc(ctx, cred, tokenType, ret, method, reqUrl, headers, param)
}

func TestTranslateVideo(t *testing.T) {
	tests := []struct {
		from     language.Tag
		to       language.Tag
		src      string
		callback string
		result   *ASRResponse
	}{
		{
			from:     language.Chinese,
			to:       language.English,
			src:      "http://test.com/test.mp4",
			callback: "http://test.com/callback",
			result:   &ASRResponse{},
		},
	}

	for _, test := range tests {
		res, err := transEngine.TranslateVideo(test.from, test.to, test.src, test.callback)
		assert.Nil(t, err)
		assert.Equal(t, test.result, res)
	}
}

func TestQueryTask(t *testing.T) {
	tests := []struct {
		taskId string
		result *ASRTaskData
	}{
		{
			taskId: "8c5bd27079924fe884d4f67b512d6740",
			result: &ASRTaskData{},
		},
	}

	for _, test := range tests {
		resp, err := transEngine.QueryVideo2TextTask(test.taskId)
		assert.Nil(t, err)
		assert.Equal(t, test.result, resp)
	}
}

func TestGenerateWebVTTFileWithTranslation(t *testing.T) {
	tests := []struct {
		asrTaskData ASRTaskData
		path        string
		from        language.Tag
		to          language.Tag
	}{
		{
			asrTaskData: ASRTaskData{},
			path:        "test_en.vtt",
			from:        language.Chinese,
			to:          language.English,
		},
	}

	for _, test := range tests {
		err := transEngine.GenerateWebVTTFileWithTranslation(test.asrTaskData, test.path, test.from, test.to)
		assert.Nil(t, err)
	}
}

func TestVideo2TextCallback(t *testing.T) {
	tests := []struct {
		callbackData Video2TextCallbackResponse
		result       string
	}{
		{
			callbackData: Video2TextCallbackResponse{
				Rtn: 1,
			},
			result: "",
		},
	}

	for _, test := range tests {
		resp, err := transEngine.Video2TextCallback(test.callbackData)
		assert.Nil(t, err)
		assert.Equal(t, test.result, resp)
	}
}

func TestText2Video(t *testing.T) {
	tests := []struct {
		content string
		result  *TTSResponse
	}{
		{
			content: "test",
			result:  &TTSResponse{},
		},
	}

	for _, test := range tests {
		res, err := transEngine.Text2Audio(test.content)
		assert.Nil(t, err)
		assert.Equal(t, test.result, res)
	}
}

func TestGenerateWebVTT(t *testing.T) {
	tests := []struct {
		asrTaskData ASRTaskData
		path        string
	}{
		{
			asrTaskData: ASRTaskData{
				Data: Data{
					SpeechResult: SpeechResult{
						Detail: []Detail{
							{
								StartTime: "0.0",
								EndTime:   "10.0",
							},
						},
					},
				},
			},
			path: "test.vtt",
		},
	}

	for _, test := range tests {
		err := transEngine.GenerateWebVTTFile(test.asrTaskData, test.path)
		assert.Nil(t, err)
	}
}
