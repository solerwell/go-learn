package yanyue_spider

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/solerwell/go-learn/logger"
	"github.com/solerwell/go-learn/yanyue_spider/yanyue"
	"github.com/valyala/fasthttp"
)

var log = logger.GetStdLogger()

func GetCig(url string) (*yanyue.ProductResp, error) {
	client := fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	req.SetHost(url)
	header := &req.Header
	header.SetMethod("POST")
	header.Add("Referer", "https://www.yanyue.cn/search?btype=DALU")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	header.Add("Cookie", "acw_tc=2760829515686154350617407e35e51e8d58f186cc59bfbbefdd7a88030b94; UM_distinctid=16d3f0c4b7588e-0d2711557007bd-5373e62-1fa400-16d3f0c4b767b1; __cfduid=dfca4f3b1540d15d8c43586c4c2b600071568720832; PHPSESSID=7tqec4lo1vrs1j2539no2n4hl1; CNZZDATA30058803=cnzz_eid%3D112472645-1568719884-https%253A%252F%252Fwww.google.com%252F%26ntime%3D1570781819; CNZZDATA30058806=cnzz_eid%3D56678817-1568719541-https%253A%252F%252Fwww.google.com%252F%26ntime%3D1570777408")
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	if err != nil {
		return nil, errors.WithMessage(err, "query yanyue website error")
	}
	if len(resp.Body()) == 0 {
		return &yanyue.ProductResp{}, nil
	}
	//bodyStr := string(resp.Body())
	pro := &yanyue.ProductResp{}
	err = json.Unmarshal(resp.Body(), pro)
	if err != nil {
		return nil, errors.WithMessage(err, "Deserialize yanyue resp json error")
	}
	return pro, nil
}
