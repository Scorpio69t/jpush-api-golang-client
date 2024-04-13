package main

import (
	"fmt"

	"github.com/Scorpio69t/jpush-api-golang-client"
)

func main() {
	cid := jpush.NewCidRequest(1, "")
	fmt.Println(cid)
	res, err := cid.GetCidList("xxx", "xxx") // 这里的 key 和 secret 需要替换成自己的
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", res.String())
	}
}
