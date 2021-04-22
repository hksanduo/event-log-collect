package utils

import (
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gookit/color"
)

func Save_data_to_file(path, filename, data, timestramp string) {
	// 创建日志导出目录
	folderPath := filepath.Join(path, timestramp)
	if !checkFilePathIsExist(folderPath) { //如果文件路径存在
		// color.Info.Println("文件路径已存在,日志信息将备份到：" + folderPath)
		// } else {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			color.Error.Printf("创建目录失败![%v]", err)
		} else {
			color.Info.Println("创建目录成功!")
			color.Info.Println("日志信息将备份到：" + folderPath)
		}
	}
	// 导出日志到文件
	logFilePath := filepath.Join(folderPath, filename+".json")
	if checkFilePathIsExist(logFilePath) {
		// f, err := os.Open(logFilePath)
		f, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //os.O_TRUNC 覆盖写入，不加则追加写入
		if err != nil {
			color.Error.Printf("file create failed. err: " + err.Error())
		} else {
			_, err = f.Write([]byte(data))
			color.Info.Println("write succeed!")
			defer f.Close()
		}

	} else {
		f, err_create := os.Create(logFilePath) //创建文件
		if err_create != nil {
			panic(err_create)
		}
		// color.Info.Println("文件创建成功")
		_, err := io.WriteString(f, data) //写入文件
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}

}

// 检查文件目录是否存在
func checkFilePathIsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		color.Info.Println(err)
		return false
	}
	return true
}

func Int2Str(value int) string {
	return strconv.Itoa(value)
}
