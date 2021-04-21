package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gookit/color"
)

func Save_data_to_file(path, filename, data, timestramp string) {
	// 创建文件导出目录
	folderPath := filepath.Join(path, timestramp)
	if checkFilePathIsExist(path) { //如果文件路径存在
		color.Info.Println("文件路径已存在,日志信息将备份到：" + folderPath)
	} else {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}

	var f *os.File
	f, err_create := os.Create(filepath.Join(folderPath, filename+".log")) //创建文件
	if err_create != nil {
		panic(err_create)
	}
	defer f.Close()
	_, err := io.WriteString(f, data) //写入文件
	if err != nil {
		panic(err)
	}
}

// 检查文件目录是否存在
func checkFilePathIsExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func Int2Str(value int) string {
	return strconv.Itoa(value)
}
