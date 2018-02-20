package files

import (
	"os"
	"bufio"
	"io"
	"strings"
)

type Constant struct {
	maps map[string]string //map
	Prefix string   //string  类文件名
	Path string     //位置
}

func (c *Constant) ReadConstant() {
	path := c.Path
	c.maps = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//每次读一行
	r := bufio.NewReader(f)
	for {
		lineBytes, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		//读取一行并转化为字符串

		s := strings.TrimSpace(string(lineBytes))
		//如果为空
		if len(s) == 0 {
			continue
		}
		//获取到类的名字
		if strings.Index(s, "class") != -1 {
			c.Prefix = strings.TrimSpace(s[5:len(s)-1])
			continue
		}

		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		key := strings.TrimSpace(s[:index])
		if strings.Index(s, "//") != -1 {
			continue
		}
		if len(key) == 0 {
			continue
		}
		key = strings.TrimSpace(strings.Trim(key,"const"))

		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		value = strings.TrimLeft(value,"'")
		value = strings.TrimRight(value,"';")

		c.maps[key] = value
	}
}

func (c *Constant) GetConstant(key string) string  {
	return c.maps[key]
}
