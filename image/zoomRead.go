package image

import (
	"regexp"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
)

/**
根据请求进行正则判断
 */
func getExp(file string) ([]string,uint8) {
	var result []string
	reg := `(?P<path>.+)_(?P<width>[0-9]+)x(?P<height>[0-9]+)_(?P<func>[a-z]+)?\.(?P<ext>[a-z]{3,4})`
	if has,_ := regexp.MatchString(reg,file); !has {
		reg = `(?P<path>.+)_(?P<width>[0-9]+)x(?P<height>[0-9]+)?\.(?P<ext>[a-z]{3,4})`
		if has,_ = regexp.MatchString(reg,file); !has {
			reg = `(?P<path>.+)?\.(?P<ext>[a-z]{3,4})`
			if has,_ = regexp.MatchString(reg,file); !has {
				fmt.Println("no reg")
				return result,0
			}
		}
	}

	if exp,err := regexp.Compile(reg); err == nil {
		result = exp.FindStringSubmatch(file)
	}

	return result,uint8(len(result))
}
/**
从s3读取图片
 */
func getImageBuf(output *s3.GetObjectOutput) []byte {
	defer output.Body.Close()
	buf := make([]byte, *output.ContentLength)
	n,err := output.Body.Read(buf)
	if err != nil  || n == 0 {
		if err.Error() == "EOF" {
			return buf
		}
		return buf
	}
	return buf
}