package redis

import "sync"

type Options struct {
	Host string
	Port uint16     //65535
}
var onceOption sync.Once
var options *Options
func GetOption() *Options{
	onceOption.Do(func(){
		options = &Options{}
	})
	return options
}
