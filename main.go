package main

import (
	"golang/advanced"
	"fmt"
)

func main()  {
	type Tmpdata struct {
		Id		int64 `xml:"Id,attr"`
		Num		int64 `xml:"Age,attr"`
	}

	type TmpDataCfg struct {
		Data []Tmpdata `xml:"data"`
	}

	tmpData := &TmpDataCfg{}

	strPath := "D:/project/src/golang/advanced/data.xml"
	err := advanced.LoadConfig(strPath, tmpData)

	fmt.Println("tmpData.Data:", tmpData.Data)
	if err != nil {
		fmt.Println("读取配置错误", err.Error())
	}

	for _, v := range tmpData.Data {
		fmt.Println(v)
	}
}
