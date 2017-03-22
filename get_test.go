package xhttp

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	body, err := Get("http://www.baidu.com")

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(string(body))
}
