package sendrequest

import (
	"net/http"
	"net/url"
)
func SendRequest(param1 string, param2 string) {
	// 定义请求的URL
	baseURL := "https://verify.wlh.lol/api"

	// 构建请求参数
	data := url.Values{}
	data.Set("param1", param1)
	data.Set("param2", param2)

	// 构建完整的URL
	reqURL := baseURL + "?" + data.Encode()

	// 发送GET请求
	resp, _ := http.Get(reqURL)
	if resp != nil {
		defer resp.Body.Close()
	}
}
