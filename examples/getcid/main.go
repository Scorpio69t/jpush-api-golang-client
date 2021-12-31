package main

import (
	"fmt"

	"github.com/Scorpio69t/jpush-api-golang-client"
)

func main() {
	cid := jpush.NewCidRequest(1, "")
	res, err := cid.GetCidList("2c741f8a5cee16580357e791", "860e76cd5ccb91f313463d4f") // 这里的 key 和 secret 需要替换成自己的
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", res.String())
	}
}
