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
	"bytes"
	"context"
	"encoding/json"
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

func TestASR(t *testing.T) {
	// {
	// 	"audioUrl":"https://cdnfile.mp3",
	// 	"callback":"https://callbackurl.com",
	// 	"filedata":{
	// 		"aue":"wav",
	// 		"lang":"MANDARIN",
	// 		"sampleRateHertz":"16000",
	// 		"audioName":"这是一个示例文件.amr"
	// 	},
	// 	"speechConfig":{
	// 		"scene":"GERNERL",
	// 		"byWords":true,
	// 		"customWords":[
	// 			"七牛",
	// 			"信息技术有限公司"
	// 		],
	// 		"addPunctuation":true,
	// 		"wordsReplaceConfig":true,
	// 		"numOfSpeakers":1,
	// 		"convertNumber":true
	// 	},
	// 	"wordsReplaceConfig":[
	// 		{
	// 			"keywords":"他妈的",
	// 			"replace":"*"
	// 		},
	// 		{
	// 			"keywords":"破坏",
	// 			"replace":"xx"
	// 		}
	// 	]
	// }

	requestBody := map[string]interface{}{
		"audioUrl": "http://s7cjwtxc2.hn-bkt.clouddn.com/test.mp4",
		"callback": "https://callbackurl.com",
		"filedata": map[string]interface{}{
			"aue":             "wav",
			"lang":            "MANDARIN",
			"sampleRateHertz": "16000",
			"audioName":       "test.amr",
		},
		"speechConfig": map[string]interface{}{
			"scene":              "GERNERL",
			"byWords":            true,
			"customWords":        []string{"七牛", "信息技术有限公司"},
			"addPunctuation":     true,
			"wordsReplaceConfig": true,
			"numOfSpeakers":      1,
			"convertNumber":      true,
		},
		"wordsReplaceConfig": []map[string]interface{}{},
	}

	jsonPayload, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	reqURL := "https://yitu-audio.qiniuapi.com/v4/lasr"

	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	mac := auth.New(mockAccessKey, mockSecretKey)
	bearToken, err := mac.SignRequestV2(req)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearToken))

	client := http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
}
