package main

import (
	"log"
	"os"

	"github.com/wuyoushe/hyper-go/tool/hyper/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Hyper"
	app.Usage = "工具集"
	app.Version = commands.GetVersion()
	app.Authors = []*cli.Author{{
		Name:  "Felton",
		Email: "Wudadongfen@126.com",
	}}
	cli.HelpFlag = &cli.BoolFlag{
		Name:  "Help",
		Usage: "帮助查看",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "GoCMD Version",
	}
	//初始化命令行
	app.Commands = commands.InitCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Print(err)
	}

}
