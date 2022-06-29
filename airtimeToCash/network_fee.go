package airtimetocash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type networkFee struct {
	Amount  int
	Network int
}

type networkFeeRes struct {
	Error   bool
	Message string
	Data    networkFeeDataBody
}

type networkFeeDataBody struct {
	Mobile string
	Amount string
	Value  int
	Charge int
}

func NetworkFee(amount, network_id int) (*networkFeeRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/airtime-cash/fee", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := networkFee{}
	payload.Amount = amount
	payload.Network = network_id

	jsonReq, jsonErr := json.Marshal(&payload)
	if jsonErr != nil {
		return nil, 0, jsonErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response networkFeeRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
