package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// bfs

func main() {
	u := &User{
		BaseUrl:   "https://www.yuque.com/api/v2",
		Token:     viper.GetString("token"),
		UserAgent: userAgent,
	}
	u.Get("/repos/skynesser/cpcmrw/docs/pk27c68itt0471gi")
	fmt.Println("hel")
}
