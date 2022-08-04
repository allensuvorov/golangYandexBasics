package main

type Response struct {
}

func ReadResponse(rawResp string) (Response, error) {

}

func main() {
	sample := `{"header": {"code": 0,"message": ""},	"data": [{"type": "user","id": 100, "attributes": {"email": "bob@yandex.ru", "article_ids": [10, 11, 12]}}]}`
	ReadResponse(sample)
}
