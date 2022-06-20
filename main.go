package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

var uuu = flag.String("u", "", "Input anchor podcast url")

func downloadFile(url string, p string) {
	// defer wg.Done()
	// Get the data
	err := os.MkdirAll(path.Dir(p), 0744)
	if err != nil {
		panic(err)
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Create output file
	out, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// copy stream
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
func download(group, uu string) {
	u, err := url.Parse(uu)
	if err != nil {
		panic(err)
	}
	log.Println("download", u.Path)
	downloadFile(uu, fmt.Sprintf("rerouces/%s%s", group, u.Path))
	// log.Println(u.Path)
}

func main() {
	flag.Parse()
	// p := podcast.New()
	//https://anchor.fm/s/590ba140/podcast/rss
	log.Println("ready to do", *uuu)
	//https://d3t3ozftmdmh3i.cloudfront.net
	// if matched, err := regexp.MatchString(`https://anchor.fm/s/[A-Za-z0-9]{1,}/podcast/rss`, *uuu); err != nil || !matched {
	// 	panic("not auchor.fm")
	// }
	//https://anchor.fm/s/590ba140/podcast/play/49235017/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2022-2-18%2Fbf950af3-3925-d32a-20ca-c0304419276f.mp3

	// var p podcast.Podcast

	res, err := http.Get(fmt.Sprintf("https://anchor.fm/s/%s/podcast/rss", *uuu))
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// flysnowRegexp := regexp.MustCompile(`^https:\/\/anchor.fm\/\w.html$`)
	flysnowRegexp := regexp.MustCompile(`https://anchor.fm[^\s]*.mp3`)

	params := flysnowRegexp.FindAllStringSubmatch(string(b), -1)
	ans := string(b)
	// fmt.Println(len(params))
	// ress := make([]string, 0)
	for _, param := range params {
		// fmt.Println(param[0])
		// fmt.Println(regexp.MustCompile(`/https[^\s]*.mp3`).FindString(param[0]))
		mp3 := regexp.MustCompile(`/https[^\s]*.mp3`).FindString(param[0])
		if len(mp3) < 5 {
			panic("error url")
		}
		enEscapeUrl, err := url.QueryUnescape(mp3[1:])
		// fmt.Println(enEscapeUrl)
		if err != nil {
			panic(err)
		}
		ans = strings.Replace(ans, param[0], enEscapeUrl, 1)
	}
	// ali := ans
	rep := func(key string) {
		ress := regexp.MustCompile(`https://`+key+`.cloudfront.net/[^\s]*.[mp3|jpeg|jpg|png]`).FindAllString(ans, -1)

		for i := 0; i < len(ress); i++ {
			// u, err := url.Parse(ress[i])
			// if err != nil {
			// 	panic(err)
			// }
			// ali = strings.ReplaceAll(ali, ress[i], "https://anchoragent.oss-cn-hangzhou.aliyuncs.com/"+key+u.Path)
			download(key, ress[i])
		}
	}
	rep("d3ctxlq1ktw2nl")
	rep("d3t3ozftmdmh3i")
	ioutil.WriteFile(*uuu+"_anchor.xml", []byte(ans), 0600)
	// ioutil.WriteFile(*uuu+"_ali.xml", []byte(ali), 0600)
}
