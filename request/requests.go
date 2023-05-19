/*
 @author: lynn
 @date: 2023/5/18
 @time: 23:43
*/

package request

import (
	"bytes"
	"encoding/json"
	"github.com/zerocodehero/BOKALPHA/models"
	"io"
	"log"
	"net/http"
)

func RAW(closer io.ReadCloser) []byte {
	//读取body
	resBody, err := io.ReadAll(closer) //把  body 内容读入字符串
	if err != nil {
		log.Println(err)
		return nil
	}
	return resBody
}

func Fetch(config models.RequestData) []byte {
	res, err := Send(config)
	if err != nil {
		return nil
	}
	if res.StatusCode == 200 {
		return RAW(res.Body)
	} else {
		return nil
	}

}

func Send(config models.RequestData) (*http.Response, error) {
	//add post body
	var bodyJson []byte
	if len(config.Body) != 0 {
		var err error
		bodyJson, err = json.Marshal(config.Body)
		if err != nil {
			log.Println(err)
		}
	} else {
		bodyJson = nil
	}

	req, err := http.NewRequest(config.Method, config.Url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36")
	//add params
	q := req.URL.Query()
	for key, val := range config.Params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()
	//add headers
	for key, val := range config.Headers {
		req.Header.Add(key, val)
	}
	//http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())

	//发送请求
	return client.Do(req)

}
