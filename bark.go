package bark

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// 发送消息
func (b *Bark) SendMessage() (*BarkRes, error) {
	client := &http.Client{}
	d, err := b.format()
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(d)
	req, err := http.NewRequest(http.MethodPost, b.ServerUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, _ := io.ReadAll(resp.Body)
	r := &BarkRes{}
	err = json.Unmarshal(respBody, r)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return r, errors.New("发送的请求有误")
	}
	return r, nil

}

func (b *Bark) format() ([]byte, error) {
	res, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return res, nil
}
