package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

type reqHandCards struct {
	HandCards []int `json:"handCards"`
	Laizi     []int `json:"laizi"`
}

type resp struct {
	Hu      bool `json:"hu"`
	Dianshu int  `json:"dianshu"`
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

	ret, dianshu := sxtdhmj.CanHuWithLaiZi(msgBag.HandCards, msgBag.Laizi)

	fmt.Println(ret, dianshu)

	respBag := &resp{}
	respBag.Hu = ret
	respBag.Dianshu = dianshu

	b, err := json.Marshal(respBag)
	if err != nil {
		fmt.Println("json Marshal failed")
		return
	}

	fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("127.0.0.1:11188", nil)
}
