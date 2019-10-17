package main

import (
	"github.com/solerwell/go-learn/logger"
	"github.com/solerwell/go-learn/yanyue_spider"
	"time"

	//"fmt"
	//"github.com/valyala/fasthttp"
	"regexp"
	"strconv"
)

var log = logger.GetStdLogger()

func main() {
	cigUrl := "https://www.yanyue.cn/api/rc/product/yanlist?brandid=1&typeid=0&tar=&nicotine&co=&price=&page={PAGE}&pagenum=12"
	var page int32 = 1
	var totalNum int32 = 0
	for {
		if totalNum > 0 && page*12 > totalNum {
			break
		}
		re, _ := regexp.Compile("\\{PAGE\\}")
		re.ReplaceAllLiteralString(cigUrl, strconv.FormatInt(int64(page), 10))
		products, err := yanyue_spider.GetCig(cigUrl)
		if err != nil {
			log.Errorf("Query yanyue no.%d page's data error, the error is %s", page, err.Error())
			continue
		}
		if products != nil {
			totalNum = products.TotalNum
			yanyue_spider.Write2Sqlite(products.ProductList)
		}
		page++
		time.Sleep(time.Duration(1) * time.Second)
	}
}
