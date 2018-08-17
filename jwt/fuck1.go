package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func parseUrl(auth_url string) (string, string) {
	u, err := url.Parse(auth_url)
	fmt.Println(err)
	return u.Scheme, u.Host
}

func main() {
	user := "user"
	key := user
	v := base64.StdEncoding.EncodeToString([]byte(user + ":" + key))
	fmt.Println(v)

	s := "www://user:pass@host.com:5432/path?k=v#f"
	sc, h := parseUrl(s)
	fmt.Println(sc)
	fmt.Println(h)
}
