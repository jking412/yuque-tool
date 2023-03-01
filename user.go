package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	UserAgent string
	Token     string
	BaseUrl   string
}

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"

func (u *User) Get(path string) {
	req, _ := http.NewRequest(http.MethodGet, u.BaseUrl+path, nil)
	req.Header.Add("User-Agent", u.UserAgent)
	req.Header.Add("X-Auth-Token", u.Token)

	resp, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(resp.Body)

	file, _ := os.OpenFile("out.json", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	fmt.Fprint(file, string(body))
}
