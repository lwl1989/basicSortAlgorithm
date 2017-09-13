package main

import (
	"container/list"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"github.com/fern4lvarez/go-metainspector/metainspector"
	"gopkg.in/mgo.v2"
	"time"
	"github.com/axgle/mahonia"
	"crawler/download"
	"crawler/parser"
)

//url队列
type urls struct {
	got  list.List
	want list.List
}

//爬虫任务
type crawlerUrl struct {
	Name     string   //任务名
	MaxLen   int      //最大长度
	MaxTime  int      //最大执行时间
	Interval int      //自动重启时间
	Inlet    *url.URL //入口
	Urls     urls
}

//json
type reqCrawler struct {
	Name     string `json:"name"`
	MaxLen   string `json:"maxLen"`
	MaxTime  string `json:"maxTime"`
	Interval string `json:"interval"`
	Url      string `json:"url"`
}

type CrawlerSave struct {
	Name string
	Url string
	Author string
	Keywords []string
	Describe string
	MetaImage string
	Images []string
	Time time.Time
}


func Crawler(w http.ResponseWriter, req *http.Request) {
	//var capacity int64
	//buf := bytes.NewBuffer(make([]byte, 0, capacity))
	//n , _ := buf.ReadFrom(req.Body)
	//println(n)
	//println(string(buf.Bytes()))

	body, _ := ioutil.ReadAll(req.Body)
	println("body")
	body_str := string(body)
	var reqCrawler reqCrawler
	var crawlerUrl crawlerUrl
	if "" != body_str {
		if err := json.Unmarshal(body, &reqCrawler); err == nil {
			println("crawler")

			crawlerUrl.Inlet, err = url.Parse(reqCrawler.Url)
			if nil != err {
				io.WriteString(w, "over!\n")
				panic("over")
			}
			crawlerUrl.MaxLen, _ = strconv.Atoi(reqCrawler.MaxLen)
			crawlerUrl.MaxTime, _ = strconv.Atoi(reqCrawler.MaxTime)
			crawlerUrl.Interval, _ = strconv.Atoi(reqCrawler.Interval)
			crawlerUrl.Name = reqCrawler.Name
			crawlerUrl.Urls.want.PushFront(crawlerUrl.Inlet)

			str := download.Get(crawlerUrl.Inlet.String())
			scraper , err := parser.NewScraper(crawlerUrl.Inlet,str)
			if nil == err {
				crawlerUrl.save(scraper)
			} else {
				fmt.Println(err)
				io.WriteString(w, "error");
			}
			//mi := crawlerUrl.run()
			//crawlerUrl.save(mi)
		} else {
			println("err")
			fmt.Println(err)
		}
	}

	println("over")
	io.WriteString(w, "hello, world!\n")
}

func (crawler *crawlerUrl) save(scraper *parser.Scraper)  {
	session, err := mgo.Dial("54.222.155.203:56790,54.222.182.136:56790,54.223.193.154:56790")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	keywords := scraper.Keywords
	description := scraper.Description
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("crawler").C("test")
	save := &CrawlerSave{
		crawler.Name,
		crawler.Inlet.String(),
		scraper.Author,
		keywords,
		description,
		scraper.MetaImage,
		scraper.Images,
		time.Now(),
	}
	fmt.Println("save",save)
	err = c.Insert(save)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
func (crawler *crawlerUrl) run() (*metainspector.MetaInspector) {

	MI, err := metainspector.New(crawler.Inlet.String())
	if err != nil {
		//fmt.Println(MI)
		fmt.Println("err:",err)
		return nil
	}
	return MI
}

func convert(value ,newCharset string) (string) {
	str := mahonia.NewEncoder(newCharset).ConvertString(value)
	return str
}
func main() {
	http.HandleFunc("/", Crawler)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
