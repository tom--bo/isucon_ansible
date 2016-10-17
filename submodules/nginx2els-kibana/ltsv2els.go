package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type nginxLog struct {
	Time         string `json:"time"`
	Host         string `json:"host"`
	Forwardedfor string `json:"forwardedfor"`
	Req          string `json:"req"`
	Method       string `json:"method"`
	URI          string `json:"uri"`
	Status       string `json:"status"`
	Size         string `json:"size"`
	Referer      string `json:"referer"`
	UA           string `json:"ua"`
	Reqtime      string `json:"reqtime"`
	Runtime      string `json:"runtime"`
	Apptime      string `json:"apptime"`
	Cache        string `json:"cache"`
	Vhost        string `json:"vhost"`
	Protocol     string `json:"protocol"`
	Handler      string `json:"handler"`
}

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin

	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "\t")
		logmap := make(map[string]string)
		checkEmptyValue := map[string]string{
			"time":         "null",
			"host":         "null",
			"forwardedfor": "null",
			"req":          "null",
			"size":         "null",
			"reqtime":      "0",
			"runtime":      "0",
			"apptime":      "0",
		}
		for _, item := range items {
			val := strings.SplitN(item, ":", 2)
			if fill, ok := checkEmptyValue[val[0]]; ok {
				if val[1] == "-" {
					val[1] = fill
				}
			}
			logmap[val[0]] = val[1]
		}
		tmpstr := strings.Split(logmap["req"], " ")
		logmap["protocol"] = tmpstr[2]
		logmap["handler"] = makeHandlerPart(tmpstr[1])
		log := nginxLog{
			Time:         logmap["time"],
			Host:         logmap["host"],
			Forwardedfor: logmap["forwardedfor"],
			Req:          logmap["req"],
			Method:       logmap["method"],
			URI:          logmap["uri"],
			Status:       logmap["status"],
			Size:         logmap["size"],
			Referer:      logmap["referer"],
			UA:           logmap["ua"],
			Reqtime:      logmap["reqtime"],
			Runtime:      logmap["runtime"],
			Apptime:      logmap["apptime"],
			Cache:        logmap["cache"],
			Vhost:        logmap["vhost"],
			Protocol:     logmap["protocol"],
			Handler:      logmap["handler"],
		}
		data, err := json.Marshal(log)
		if err != nil {
			fmt.Print("JSON marshaling failed: %s", err)
		}

		postJSON(os.Args[2], data)
		time.Sleep(10 * time.Millisecond)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func postJSON(url string, data []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		panic(err2)
	}

	defer resp.Body.Close()
}

func makeHandlerPart(uristr string) string {
	tmp := strings.Split(uristr, "?")
	tmp = strings.Split(tmp[0], "#")
	tmp = strings.Split(tmp[0], "/")
	if len(tmp) < 2 {
		return uristr
	}
	return "/" + tmp[1]
}
