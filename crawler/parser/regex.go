package parser

import (
    "time"
    "bitbucket.org/zaphar/go-html-transform/h5"
)

type CrawlerSave struct {
    Name string
    Url string
    Author string
    Keywords []string
    Describe string
    Images []string
    Time time.Time
}


func (CrawlerSave *CrawlerSave) Init()  {
    CrawlerSave.Time = time.Now()
}

func (CrawlerSave *CrawlerSave) Parse(value string) (*h5.Tree, error)  {
    return h5.NewFromString(value)
}