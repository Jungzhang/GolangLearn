package main

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strings"
	"fmt"
	"log"
	"time"
	"strconv"
)

func requestQuestionnaire(id string) (*gorequest.SuperAgent, string, string) {
	request := gorequest.New().Proxy("http://127.0.0.1:6666")
	request.Url = "https://www.wjx.cn/jq/" + id + ".aspx"
	request.Method = http.MethodGet
	rsp, body, err := request.End()
	if err != nil {
		e := fmt.Errorf("%v", err)
		panic(e)
	}

	if rsp.StatusCode != 200 {
		panic(fmt.Errorf("http status code if not 200, code: %v", rsp.StatusCode))
	}

	idx := strings.Index(body, "rndnum")
	flag := false
	rndnum := ""

	for i := idx + 7; i < idx+100; i++ {
		if body[i] == '"' {
			if !flag {
				flag = true
				continue
			} else {
				break
			}
		}
		rndnum += string(body[i])
	}

	starttime := ""
	flag = false
	idx = strings.Index(body, "starttime")
	for i := idx + 10; i < idx+100; i++ {
		if body[i] == '"' {
			if !flag {
				flag = true
				continue
			} else {
				break
			}
		}
		starttime += string(body[i])
	}

	return request, rndnum, starttime
}

func postAnswer(id string) {
	request, rndnum, starttime := requestQuestionnaire(id)
	request.Method = http.MethodPost
	url := "joinnew/processjq.ashx?submittype=1&curID=" + id +
		"&t=" + strconv.FormatInt(time.Now().UnixNano()/1000000, 10) + "&starttime=" + starttime + "&rndnum=" + rndnum
	request.Url = "https://www.wjx.cn/" + url
	request.Set("authority", "www.wjx.cn")
	request.Set("method", "POST")
	request.Set("path", url)
	request.Set("scheme", "https")
	request.Set("accept", "*/*")
	request.Set("accept-encoding", "gzip, deflate, br")
	request.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	request.Set("dnt", "1")
	request.Set("referer", "https://www.wjx.cn/jq/23173552.aspx")
	request.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")

	rsp, body, err := request.Send("submitdata=1$2}2$1}3$2}4$1}5$1}6$1}7$1}8$1}9$2}10$1|2}11$2}12$1}13$5}14$3}15$4}16$3}17$4}18$5}19$2}20$5}21$4}22$4}23$4}24$5}25$5}26$5}27$4}28$5}29$5}30$3}31$4}32$4}33$5}34$3}35$5}36$5}37$5}38$4").End()
	if err != nil {
		e := fmt.Errorf("%v", err)
		panic(e)
	}
	fmt.Println(body)
	log.Println(rsp)
}

func main() {
	postAnswer("23173552")
}
