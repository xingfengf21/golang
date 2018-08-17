package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const EXPIRE_TIME = 0
const BD_API_VERSION = "2.0"
const BD_API_TOKEN = "sRe0ca140nvSe"
const BD_API_CID = "njkJen433oiure434"
const BD_API_URL = "https://lookup.api.bsb.baidu.com/urlquery"

func bdapi(domain string) (string, error) {
	url := fmt.Sprintf("%s?url=%s&ver=%s&key=%s&cid=%s",
		BD_API_URL, domain, BD_API_VERSION, BD_API_TOKEN, BD_API_CID)
	ret, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if ret.StatusCode != 200 {
		return "", fmt.Errorf("%d", ret.StatusCode)
	}
	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	res, err := bdapi("www.baidu.com")
	fmt.Println(res)
	fmt.Println(err)
}
