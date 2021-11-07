package util

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/yanyiwu/gojieba"
)

type CONF map[string]interface{}

var Confs = make(CONF)

var JieBa = gojieba.NewJieba()

// GetPrevDir ...
func GetPrevDir(path string) string {
	latsindex := strings.LastIndex(path, "/")
	return path[:latsindex]
}

// InitConf ...
func InitConf() {
	if len(Confs) != 0 {
		fmt.Println(Confs)
		return
	}
	CurrentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fl, err := os.Open(CurrentDir + "/conf/config")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if n == 0 {
			break
		}
	}
	for _, v := range strings.Split(string(buf), "\n") {
		v = strings.Replace(v, " ", "", -1)
		index := strings.Index(v, ":")
		if index == -1 {
			continue
		}
		k := strings.Replace(v[:index], " ", "", -1)
		v := strings.Replace(v[index+1:], " ", "", -1)
		Confs[k] = v
	}
}

// GetConf ...
func GetConf(name string) interface{} {
	return Confs[name]
}

func GetConfStr(name string) string {
	if os.Getenv("GSC_DEBUG") == "true" {
		return fmt.Sprintf("%v", Confs[name])
	}
	return os.Getenv(name)
}

func SplitString(s string) []string {
	res := JieBa.CutForSearch(strings.TrimSpace(s), true)
	newRes := make([]string, 0)
	for _, ss := range res {
		ss = strings.TrimSpace(ss)
		// mysql 最少匹配2个字
		if utf8.RuneCountInString(ss) <= 1 {
			continue
		}
		newRes = append(newRes, ss)
	}
	return newRes
}

func init() {
	if os.Getenv("GSC_DEBUG") == "true" {
		InitConf()
	}
}
