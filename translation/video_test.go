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
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	// Your api key
	mockAccessKey = os.Getenv("QINIU_ACCESS_KEY")
	mockSecretKey = os.Getenv("QINIU_SECRET_KEY")
	mockBucket    = os.Getenv("QINIU_TEST_BUCKET")
	mockPipeline  = os.Getenv("QINIU_TEST_PIPELINE")
)

func TestUploadVideo(t *testing.T) {
	if mockAccessKey == "" {
		t.Skip("QINIU_ACCESS_KEY not set")
	}

	if mockSecretKey == "" {
		t.Skip("QINIU_SECRET_KEY not set")
	}

	localFile := "./test.mp4"
	key := "test.mp4"
	mac := auth.New(mockAccessKey, mockSecretKey)
	putPolicy := storage.PutPolicy{
		Scope: mockBucket + ":" + key,
	}
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	client := http.Client{}
	formUploader := storage.NewFormUploaderEx(&cfg, &storage.Client{Client: &client})
	ret := storage.PutRet{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestVideo2Text(t *testing.T) {
	if mockAccessKey == "" {
		t.Skip("QINIU_ACCESS_KEY not set")
	}

	if mockSecretKey == "" {
		t.Skip("QINIU_SECRET_KEY not set")
	}

	e := New("", mockAccessKey, mockSecretKey)
	resp, err := e.Video2Text("http://test.com/test.mp4", "http://test.com/callback")
	fmt.Printf("resp: %#v", resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestQueryTask(t *testing.T) {
	if mockAccessKey == "" {
		t.Skip("QINIU_ACCESS_KEY not set")
	}

	if mockSecretKey == "" {
		t.Skip("QINIU_SECRET_KEY not set")
	}

	e := New("", mockAccessKey, mockSecretKey)
	resp, err := e.QueryVideo2TextTask("mockTaskId")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGenerateWebVTTFile(t *testing.T) {
	if mockAccessKey == "" {
		t.Skip("QINIU_ACCESS_KEY not set")
	}

	if mockSecretKey == "" {
		t.Skip("QINIU_SECRET_KEY not set")
	}

	e := New("", mockAccessKey, mockSecretKey)
	resp, err := e.QueryVideo2TextTask("")
	if err != nil {
		t.Fatal(err)
	}

	err = e.GenerateWebVTTFile(*resp, "test.vtt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestText2Video(t *testing.T) {
	e := New("", mockAccessKey, mockSecretKey)
	resp, err := e.Text2Audio("hello world")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
