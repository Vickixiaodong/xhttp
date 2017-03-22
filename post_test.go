package xhttp

import (
	"log"
	"testing"
)

func TestPost(t *testing.T) {
	postBody := []byte("{\"name\":\"张三\",\"timestamp\":1490163194}")
	header := make(map[string]string)

	header["Content-Type"] = "application/json"

	body, err := Post(postBody, "http://www.baidu.com", header)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(string(body))
}
