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
	"errors"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/client"
	language "golang.org/x/text/language"
)

var (
	ErrRequestASRFailed = errors.New("request ASR failed")
	ErrRequestTTSFailed = errors.New("request TTS failed")
)

const (
	// APIURL is the url of ASR api server
	qiniuASRAPI = "https://yitu-audio.qiniuapi.com/v4/lasr"
	qiniuTTSAPI = "https://ap-gate-z0.qiniuapi.com/voice/v2/tts"
)

// ASRRequest
type ASRRequest struct {
	AudioUrl string `json:"audioUrl"`
	Callback string `json:"callback"`
}

// ASRResponse
type ASRResponse struct {
	Rtn      int    `json:"rtn"`
	Message  string `json:"message"`
	TaskId   string `json:"taskId"`
	AudioUrl string `json:"audio"`
}

// TTSRequest
type TTSRequest struct {
	Content string `json:"content"`
}

// TranslateVideo translates video, return the url of translated video
func (e *Engine) TranslateVideo(from, to language.Tag, src string, callback string) (*ASRResponse, error) {
	return e.Video2Text(src, callback)
}

// Video2Text
func (e *Engine) Video2Text(src string, callback string) (*ASRResponse, error) {
	// Prepare request body
	requestBody := ASRRequest{
		AudioUrl: src,
		Callback: callback,
	}

	c := client.DefaultClient
	res := &ASRResponse{}
	err := c.CredentialedCallWithJson(context.Background(), e.qiniuCred, auth.TokenQiniu, &res, "POST", qiniuASRAPI, nil, requestBody)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Video2TextCallbackResponse
type Video2TextCallbackResponse struct {
	Rtn       int    `json:"rtn"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	TaskId    string `json:"taskId"`
	Data      Data   `json:"data"`
}

// Data is the data of Video2TextCallbackResponse
type Data struct {
	StatusCode   int          `json:"statusCode"`
	StatusText   string       `json:"statusText"`
	SpeechResult SpeechResult `json:"speechResult"`
}

// SpeechResult is the speech result of Video2TextCallbackResponse
type SpeechResult struct {
	ResultText string   `json:"resultText"`
	Duration   int      `json:"duration"`
	Detail     []Detail `json:"detail"`
}

// Detail is the detail of SpeechResult
type Detail struct {
	Sentences  string       `json:"sentences"`
	StartTime  string       `json:"startTime"`
	EndTime    string       `json:"endTime"`
	SpeakerId  string       `json:"speakerId"`
	WordsPiece []WordsPiece `json:"wordsPiece"`
}

// WordsPiece is the words piece of Detail
type WordsPiece struct {
	Words           string `json:"words"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	TranscribedType string `json:"transcribedType"`
}

// Video2TextCallback
// doc: https://developer.qiniu.com/dora/11175/long-speech-recognition
func (e *Engine) Video2TextCallback(callbackData Video2TextCallbackResponse) (string, error) {
	if callbackData.Rtn == 0 {
		return "", ErrRequestTTSFailed
	}

	resp, err := e.Text2Audio(callbackData.Data.SpeechResult.ResultText)
	if err != nil {
		return "", ErrRequestTTSFailed
	}

	return resp.Result.AudioUrl, nil
}

// TTSResponse is the response of TTS
type TTSResponse struct {
	Code   string         `json:"code"`
	Msg    string         `json:"msg"`
	Result TTSAudioResult `json:"result"`
}

// TTSAudioResult is the result of TTS
type TTSAudioResult struct {
	AudioUrl string `json:"audioUrl"`
}

// Text2Audio
func (e *Engine) Text2Audio(content string) (*TTSResponse, error) {
	// Prepare request body
	requestBody := TTSRequest{
		Content: content,
	}

	c := client.DefaultClient
	res := &TTSResponse{}
	err := c.CredentialedCallWithJson(context.Background(), e.qiniuCred, auth.TokenQiniu, &res, "POST", qiniuTTSAPI, nil, requestBody)
	if err != nil {
		return nil, err
	}

	return res, nil
}
