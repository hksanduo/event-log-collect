package cmd

import (
	"event-log-collect/utils"

	"github.com/urfave/cli"
)

var Scan = cli.Command{
	Name:        "scan",
	Usage:       "start to collect system event log",
	Description: "start to collect system event log",
	Action:      utils.Scan,
	Flags:       []cli.Flag{
		// boolFlag("scan, s", "scan"),
		// intFlag("timeout, t", 5, "timeout"),
		// intFlag("scan_num, n", 5000, "thread num"),
		// stringFlag("ip_list, i", "iplist.txt", "iplist"),
		// stringFlag("user_dict, u", "user.dic", "user dict"),
		// stringFlag("pass_dict, p", "pass.dic", "password dict"),
		// stringFlag("outfile, o", "x_crack.txt", "scan result file"),
	},
}
