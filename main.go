package main

import (
	"golang/advanced"
	"fmt"
	"log"
)

func main()  {
	type Tmpdata struct {
		Id		int64 `xml:"Id,attr"`
		Num		float32 `xml:"Age,attr"`
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

	if func(a uint32) bool {
		fmt.Println(a + 3)
		return true
		}(uint32(2)) {
		fmt.Println("true")
	}

	log.Printf("testetesttest",0)

	var arr []uintptr
	arr = make([]uintptr, 5)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4

	fmt.Println(arr[2:])
	param := ParamStruct{
		m:15,
		n:45,
	}
	PrintInfo(3, 12, "param", 123, "sss", param)

	a := 3
	b := 5
	a, b = b, a
	fmt.Println(a,b)

	var user [2<<3]map[uint64]string
	t := make(map[uint64]string)
	t[1] = "sdf"
	user[1] = t

	fmt.Println(user)
}

type ParamStruct struct {
	m uint32
	n uint32
}

func PrintInfo(id uint32, args ...interface{}) {
	fmt.Println(id)
	for _, param := range args{
		switch param.(type) {
		case string:
			m := param.(string)
			fmt.Print(m)
		case int:
			fmt.Print(param)
		default:
			t := ParamStruct{}
			fmt.Println(t)
		}
	}
}
