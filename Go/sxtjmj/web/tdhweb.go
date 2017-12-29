package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MDGSF/MJHuPai/Go/sxtjmj"
)

type reqHandCards struct {
	HandCards  []int `json:"handCards"`
	HuType     int   `json:"huType"`
	HuCard     int   `json:"huCard"`
	HeiSanFeng bool  `json:"heiSanFeng"`
	ZhongFaBai bool  `json:"zhongFaBai"`
	ZhongFaWu  bool  `json:"zhongFaWu"`
}

type resp struct {
	Hu      bool `json:"hu"`
	FengNum int  `json:"fengNum"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(body)

	msgBag := &reqHandCards{}
	err := json.Unmarshal(body, &msgBag)
	if err != nil {
		fmt.Println("json Unmarshal failed")
		return
	}

	fmt.Println(msgBag)

	ret, fengNum := sxtjmj.CanHu(msgBag.HandCards, msgBag.HuType, msgBag.HuCard,
		msgBag.HeiSanFeng, msgBag.ZhongFaBai, msgBag.ZhongFaWu)

	fmt.Println(ret, fengNum)

	respBag := &resp{}
	respBag.Hu = ret
	respBag.FengNum = fengNum

	b, err := json.Marshal(respBag)
	if err != nil {
		fmt.Println("json Marshal failed")
		return
	}

	fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("127.0.0.1:11189", nil)
}
