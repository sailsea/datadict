
package main

import (
	"datadict/model"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "MysqlTool"
	//程序的用途描述
	oApp.Usage = "Generating a database data dictionary"
	//程序的版本号
	oApp.Version = "1.0.0"
	//设置多个命令处理函数
	oApp.Commands = []cli.Command{
		{
			//命令全称
			Name:"DataDictionary",
			//命令简写
			Aliases:[]string{"dd"},
			//命令详细描述
			Flags:[]cli.Flag {
				cli.StringFlag{
					Name: "DriverName",
					Value: "",
					Usage: "Database driver name, now supported : mysql",
				},
				cli.StringFlag{
					Name: "GeneratedPath",
					Value: "data_dict",
					Usage: "This parameter is optional, if blank, the default value is DataDict, then will",
				},
				cli.StringFlag{
					Name: "Format",
					Value: "markdown",
					Usage: "Conversion format, now supports markdown and html, default is markdown",
				},
				cli.StringFlag{
					Name: "SubSql",
					Value: "COLUMN_NAME AS 字段, COLUMN_TYPE AS 类型, IS_NULLABLE AS 是否为空, COLUMN_DEFAULT AS 默认, COLUMN_COMMENT AS 注释",
					Usage: "SubSql is a sql fragment, generally the display condition between select and from, the default is `COLUMN_NAME AS 字段, COLUMN_TYPE AS 类型, IS_NULLABLE AS 是否为空, COLUMN_DEFAULT AS 默认, COLUMN_COMMENT AS 注释`",
				},
			},
			Usage:"Generating a database data dictionary",
			Action: func(c *cli.Context) {
				DriverName := c.String("DriverName")
				GeneratedPath := c.String("GeneratedPath")
				Format := c.String("Format")
				SubSql := c.String("SubSql")
				fmt.Println(DriverName,"\n",GeneratedPath,"\n",Format,"\n",SubSql,"\n")
				model.DictionaryGeneration(DriverName,GeneratedPath,Format,SubSql)
			},
		},
	}

	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
