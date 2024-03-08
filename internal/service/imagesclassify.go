package service

import (
	"ElectronicsRecycle/internal/util"
	"bytes"
	"encoding/json"
	"net/http"
)

func ClassifyImage(base64Image string) (map[string]interface{}, error) {
	accessToken, err := util.GetAccessToken()
	if err != nil {
		return nil, err
	}

	requestBody, _ := json.Marshal(map[string]string{
		"image":   base64Image,
		"top_num": "6",
	})

	req, _ := http.NewRequest("POST", "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/classification/dianzi2?access_token="+accessToken, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
