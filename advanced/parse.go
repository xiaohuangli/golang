package advanced

import (
	"io/ioutil"
	"encoding/xml"
	"os/exec"
	"os"
	"path/filepath"
	"strings"
)

// 加载xml配置文件
func LoadConfig(filename string, v interface{}) error {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		if err = xml.Unmarshal(contents, v); err != nil {
			return err
		}
		return nil
	}
}

// 获取运行路径
func GetRuntimePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	i := strings.LastIndex(path, "/")

	if i < 0 {
		i = strings.LastIndex(path,"\\")
	}
	return string(path[0:i+1]) + "../"
}
