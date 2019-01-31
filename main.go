package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fanyang1988/force-go/config"

	"github.com/eosforce/goeosforce/ecc"

	"github.com/cihub/seelog"
	eos "github.com/eosforce/goeosforce"
	force "github.com/fanyang1988/force-go"
	"github.com/fanyang1988/force-go/action"
)

var configPath = flag.String("cfg", "./config.json", "confg file path")
var sleepms = flag.Int("s", 10, "time sleep between send request (ms)")
var threadNum = flag.Int("g", 10, "gorountinue num to send")

func init() {
	ecc.PublicKeyPrefixCompat = "EOS"
}

func send(num int, url string) {
	client, err := force.NewClientFromFile(url)
	if err != nil {
		seelog.Errorf("connect err by %s", err.Error())
		return
	}

	q, err := eos.NewAsset("0.0100 SYS")
	if err != nil {
		seelog.Errorf("asset err by %s", err.Error())
		return
	}

	for i := 0; i < 1000000000; i++ {
		_, err = client.PushActions(
			action.NewTransfer(eos.AN("v.test"), eos.AN("xgtjrtdask"), q, fmt.Sprintf("memo%d", i+num*100000)))

		if err != nil {
			seelog.Errorf("push action err by %s", err.Error())
		}

		time.Sleep(time.Duration(*sleepms) * time.Millisecond)
	}
}

type tpsCfg struct {
	Urls []string `json:"urls"`
}

func main() {
	defer seelog.Flush()
	flag.Parse()

	var tc tpsCfg
	config.LoadJSONFile("./tps.json", &tc)

	seelog.Infof("sned urls %v", tc.Urls)

	for i := 0; i < *threadNum; i++ {
		go send(i, tc.Urls[i%len(tc.Urls)])
	}

	for {
		seelog.Flush()
		time.Sleep(1 * time.Second)
	}
}
