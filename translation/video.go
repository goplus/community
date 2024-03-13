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
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/client"
	language "golang.org/x/text/language"
)

var (
	ErrRequestASRFailed = errors.New("request ASR failed")
	ErrRequestTTSFailed = errors.New("request TTS failed")
	ErrPathNotExisted   = errors.New("requested path is not existed")
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

type ASRTaskData Video2TextCallbackResponse

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
	err := c.CredentialedCallWithJson(
		context.Background(),
		e.qiniuCred,
		auth.TokenQiniu,
		&res,
		"POST",
		qiniuASRAPI,
		nil,
		requestBody,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QueryVideo2TextTask
func (e *Engine) QueryVideo2TextTask(taskId string) (*ASRTaskData, error) {
	// Prepare request url
	requestURL := fmt.Sprintf("%s/%s", qiniuASRAPI, taskId)

	c := client.DefaultClient
	res := &ASRTaskData{}
	err := c.CredentialedCallWithJson(
		context.Background(),
		e.qiniuCred,
		auth.TokenQiniu,
		&res,
		"GET",
		requestURL,
		nil,
		nil,
	)
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

// GenerateWebVTTBytes generate caption for speech result
func (e *Engine) GenerateWebVTTBytes(asrTaskData ASRTaskData) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	webVTTBuffer := GenerateSimpleWebVTTFileFromASRTaskData(asrTaskData)
	_, err := webVTTBuffer.ToBuffer().WriteTo(&buffer)
	if err != nil {
		return buffer, err
	}

	return buffer, nil
}

// GenerateWebVTTFile generate caption for speech result
func (e *Engine) GenerateWebVTTFile(asrTaskData ASRTaskData, path string) error {
	buffer, err := e.GenerateWebVTTBytes(asrTaskData)
	if err != nil {
		return err
	}

	// Write buffer to file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = buffer.WriteTo(file)
	if err != nil {
		return err
	}

	return nil
}

// GenerateWebVTTBytesWithTranslation generate caption for speech result with translation
func (e *Engine) GenerateWebVTTBytesWithTranslation(asrTaskData ASRTaskData, from, to language.Tag) (bytes.Buffer, bytes.Buffer, error) {
	var originBuffer bytes.Buffer
	var translatedBuffer bytes.Buffer

	// Write head
	webVTTBuffer := GenerateSimpleWebVTTFileFromASRTaskData(asrTaskData)
	_, err := webVTTBuffer.ToBuffer().WriteTo(&originBuffer)
	if err != nil {
		return originBuffer, translatedBuffer, err
	}

	// Translate buffer
	err = e.TranslateWebVTT(&webVTTBuffer, from.String(), to)
	if err != nil {
		return originBuffer, translatedBuffer, err
	}

	_, err = webVTTBuffer.ToBuffer().WriteTo(&translatedBuffer)
	if err != nil {
		return originBuffer, translatedBuffer, err
	}

	return originBuffer, translatedBuffer, nil
}

// GenerateWebVTTFileWithTranslation generate caption for speech result with translation
func (e *Engine) GenerateWebVTTFileWithTranslation(asrTaskData ASRTaskData, path string, from, to language.Tag) error {
	originBuffer, translatedBuffer, err := e.GenerateWebVTTBytesWithTranslation(asrTaskData, from, to)
	if err != nil {
		return err
	}

	// Write buffer to file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = originBuffer.WriteTo(file)
	if err != nil {
		return err
	}

	// Write buffer to file
	transFile, err := os.Create(fmt.Sprintf("trans.%s", path))
	if err != nil {
		return err
	}
	defer transFile.Close()

	_, err = translatedBuffer.WriteTo(transFile)
	if err != nil {
		return err
	}

	return nil
}

const (
	WebVTTFileHead      = "WEBVTT FILE\n\n"
	WebVTTContentFormat = "%d\n%s --> %s\n%s\n\n"
)

// SimpleWebVTT is the simple format of WebVTT
type SimpleWebVTT struct {
	// Basic format
	// WEBVTT FILE
	//
	// 1
	// 00:00:09.500 --> 00:00:12.000
	// The ocean floor rises 5 miles to the shores

	WebVTTFileHead string
	Content        []SimpleWebVTTContent
}

// SimpleWebVTTContent is the content of SimpleWebVTT
type SimpleWebVTTContent struct {
	Index     int
	StartTime string
	EndTime   string
	Sentence  string
}

// NewSimpleWebVTT create a new SimpleWebVTT from io.Reader
//
// Example:
// Basic format
// WEBVTT FILE
//
// 1
// 00:00:09.500 --> 00:00:12.000
// The ocean floor rises 5 miles to the shores
func NewSimpleWebVTT(reader io.Reader) (SimpleWebVTT, error) {
	var simpleWebVTT SimpleWebVTT
	var scanner = bufio.NewScanner(reader)
	var line string
	var lineNum int

	// Skip first line
	for scanner.Scan() {
		line = scanner.Text()
		lineNum++

		// Check head
		if fs := strings.Fields(line); len(fs) > 0 && fs[0] == "WEBVTT" {
			simpleWebVTT.WebVTTFileHead = WebVTTFileHead
			break
		}
	}

	// Read content
	for scanner.Scan() {
		line = scanner.Text()
		lineNum++

		// Skip empty line
		if line == "" {
			continue
		}

		// Read batch content
		if fs := strings.Fields(line); len(fs) > 0 {
			// Read index
			index, err := strconv.Atoi(fs[0])
			if err != nil {
				return simpleWebVTT, err
			}

			// Read time
			scanner.Scan()
			line = scanner.Text()
			lineNum++
			fs = strings.Fields(line)
			if len(fs) < 3 {
				continue
			}
			startTime := fs[0]
			endTime := fs[2]

			// Read sentence
			scanner.Scan()
			line = scanner.Text()
			lineNum++
			sentence := line

			// Add content
			simpleWebVTT.AddContent(index, startTime, endTime, sentence)
		}
	}

	return simpleWebVTT, nil
}

// AddContent add content to SimpleWebVTT
func (w *SimpleWebVTT) AddContent(index int, startTime, endTime, sentence string) {
	w.Content = append(w.Content, SimpleWebVTTContent{
		Index:     index,
		StartTime: startTime,
		EndTime:   endTime,
		Sentence:  sentence,
	})
}

// ToBuffer convert SimpleWebVTT to buffer
func (w *SimpleWebVTT) ToBuffer() *bytes.Buffer {
	var buffer bytes.Buffer
	buffer.WriteString(w.WebVTTFileHead)
	for _, content := range w.Content {
		buffer.WriteString(fmt.Sprintf(WebVTTContentFormat, content.Index, content.StartTime, content.EndTime, content.Sentence))
	}
	return &buffer
}

// ToString convert SimpleWebVTT to string
func (w *SimpleWebVTT) ToString() string {
	return w.ToBuffer().String()
}

// GenerateSimpleWebVTTFileFromASRTaskData generate simple WebVTT file from ASRTaskData
func GenerateSimpleWebVTTFileFromASRTaskData(asrTaskData ASRTaskData) SimpleWebVTT {
	var simpleWebVTT SimpleWebVTT
	simpleWebVTT.WebVTTFileHead = WebVTTFileHead
	for i, detail := range asrTaskData.Data.SpeechResult.Detail {
		startTime := formatTime(detail.StartTime)
		endTime := formatTime(detail.EndTime)

		// Add content
		simpleWebVTT.AddContent(i+1, startTime, endTime, detail.Sentences)
	}
	return simpleWebVTT
}

func formatTime(timeInMs string) string {
	time, _ := strconv.Atoi(timeInMs)
	hours := time / 3600000
	minutes := (time % 3600000) / 60000
	seconds := (time % 60000) / 1000
	milliseconds := time % 1000
	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)
}
