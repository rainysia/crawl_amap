package main

import (
	"crawl_amap/conf"
	 "crawl_amap/fetcher"
	 "crawl_amap/handler"
	 "crawl_amap/parser"
	"fmt"
)

func main() {
	url := conf.URL
	fmt.Printf("url=%s", url)

	provinceUrl := fmt.Sprintf(url, "中国", 1, conf.Key)
	fmt.Println(provinceUrl)
	provinceData, err := fetcher.Fetch(provinceUrl)

	if err != nil {
		panic("获取省份数据失败")
	}

	respData, err := parser.Parse(provinceData)
	fmt.Println(respData)
	if err != nil {
		panic("数据解析失败")
	}

	error := handler.Handle(*respData)
	if error != nil {
		panic("数据处理失败")
	}
}
