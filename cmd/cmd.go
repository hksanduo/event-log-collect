package cmd

import (
	"event-log-collect/plugins"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
)

// 构建scan命令及其参数
var Scan = cli.Command{
	Name:        "scan",
	Usage:       "start to collect system event log",
	Description: "start to collect system event log",
	Action:      ScanTask,
	Flags: []cli.Flag{
		// outfile参数，默认参数为test.json
		stringFlag("path, p", "./_out", "scan result path"),
		// boolFlag("scan, s", "scan"),
		// intFlag("timeout, t", 5, "timeout"),
		// intFlag("scan_num, n", 5000, "thread num"),
		// stringFlag("ip_list, i", "iplist.txt", "iplist"),
		// stringFlag("user_dict, u", "user.dic", "user dict"),
		// stringFlag("pass_dict, p", "pass.dic", "password dict"),
		// stringFlag("outfile, o", "x_crack.txt", "scan result file"),
	},
}

// 处理字符串数据
func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

// 处理布尔型数据
func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

// 处理整形数据
func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

/*
新建扫描任务，目前想做成框架直接加载组件，但是受限于go框架，目前没什么好的想法。
*/
func ScanTask(ctx *cli.Context) (err error) {
	// color.Info.Println("测试")
	timestramp := time.Now().Format("2006-01-02-15-04-05")
	path := getfilepath(ctx.String("path"))
	// go plugins.NetInterfaces(path, timestramp)
	// go plugins.NetConnections(path, timestramp)
	plugins.NetInterfaces(path, timestramp)
	plugins.NetConnections(path, timestramp)
	// color.Info.Println(ctx.String("outfile"))
	// color.Info.Println(ctx.String("outfile"))
	return err
}

func getfilepath(path string) string {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	path = filepath.Join(exPath, path)
	return path
}
