package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

var ProxyURL *string

func Send(req *http.Request) (map[string]interface{}, error) {
	client := &http.Client{}
	if ProxyURL != nil {
		transport, err := constructProxy(ProxyURL)
		if err != nil {
			return nil, err
		}
		client.Transport = transport
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrTransmitFailed.Wrap(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := ErrRequestThirdServiceNotSuccess
		if resp.StatusCode > 200 && resp.StatusCode < 500 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, ErrReadResponseFailed.Wrap(err)
			}

			result := FormatResponse{}
			if err := json.Unmarshal(body, &result); err != nil {
				log.Println(string(body))
				return nil, ErrJsonMarshalerFailed.Wrap(err)
			}

			return nil, &CustomError{
				code: result.Code,
				msg:  result.Message,
			}
		}
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadResponseFailed.Wrap(err)
	}

	result := FormatResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(string(body))
		return nil, ErrJsonMarshalerFailed.Wrap(err)
	}

	return *result.Data, nil
}

func constructProxy(proxyURLStr *string) (*http.Transport, error) {
	// 创建一个代理 URL
	proxyURL, err := url.Parse(*proxyURLStr)
	if err != nil {
		return nil, ErrProxyFailed.Wrap(err)
	}

	// 创建一个自定义的 Transport，并设置代理
	return &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}, nil
}
