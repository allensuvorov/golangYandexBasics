package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		Id         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIds []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	var res Response
	err := json.Unmarshal([]byte(rawResp), &res)
	if err != nil {
		fmt.Println("error:", err)
	}
	return res, err
}

func main() {
	//sample := `{"header": {"code": 0,"message": ""},"data": [{"type": "user","id": 100, "attributes": {"email": "bob@yandex.ru", "article_ids": [10, 11, 12]}}]}`
	sample := `
		{
			"header": {
				"code": 0,
				"message": ""
			},
			"data": [{
				"type": "user",
				"id": 100,
				"attributes": {
					"email": "bob@yandex.ru",
					"article_ids": [10, 11, 12]
				}
			}]
		} 
	`
	fmt.Println(ReadResponse(sample))
}
