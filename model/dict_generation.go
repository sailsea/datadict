package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"strings"
	"time"
)

var (
	engine *xorm.Engine
)

func DictionaryGeneration(DriverName,GeneratedPath,Format,SubSql  string){
	if DriverName == "" {
		fmt.Println("Error DriverName cannot be empty.")
		return
	}

	var err error
	engine, err = xorm.NewEngine("mysql", DriverName)
	if err != nil {
		panic(err)
	}

	gres, gerr := engine.Query("select database();")
	if gerr != nil {
		fmt.Println("Error",gerr)
		return
	}

	databaseName := string(gres[0]["database()"])

	TitleList := SqlToTitleList(SubSql)
	b,err := PathExists(GeneratedPath)
	if err != nil {
		fmt.Println("Error",err)
		return
	}

	if !b {
		if err := os.MkdirAll(GeneratedPath, os.ModePerm);err != nil {
			fmt.Println("Error",err)
			return
		}
	}

	switch Format {
	case "html":
		fmt.Println("html")
		filePath := fmt.Sprintf("%s/%s_%d.html",GeneratedPath,databaseName,time.Now().Unix())
		dictFile, error := os.Create(filePath);
		if error != nil {
			fmt.Println(error);
			return
		}
		dictFile.WriteString(fmt.Sprintf(`<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <title>%s database data dictionary</title> </head> <body>`,databaseName))
		for _,v := range GetAllTable(){
			data := QueryTableStructToHtml(TitleList,SubSql,databaseName,v)
			dictFile.WriteString(data)
		}
		dictFile.WriteString(`</body> <style type="text/css"> table { font-family: verdana,arial,sans-serif; font-size:11px; color:#333333; border-width: 1px; border-color: #666666; border-collapse: collapse; } table th { border-width: 1px; padding: 8px; border-style: solid; border-color: #666666; background-color: #dedede; } table td { border-width: 1px; padding: 8px; border-style: solid; border-color: #666666; background-color: #ffffff; } </style> </html>`)
		dictFile.Sync()
		dictFile.Close()

	case "markdown":
		fmt.Println("markdown")
		filePath := fmt.Sprintf("%s/%s_%d.md",GeneratedPath,databaseName,time.Now().Unix())
		dictFile, error := os.Create(filePath);
		if error != nil {
			fmt.Println(error);
			return
		}

		for _,v := range GetAllTable(){
			data := QueryTableStructToMD(TitleList,SubSql,databaseName,v)
			dictFile.WriteString(data)
		}
		dictFile.Sync()
		dictFile.Close()

	default:
		fmt.Println(fmt.Sprintf("%s format output is not supported. The supported types are: html, markdown"),Format)
		return
	}

}


func QueryTableStructToMD(TitleList []string,subSql,databaseName,tableName string)string{

	sqlStr := fmt.Sprintf(`SELECT %s FROM information_schema. COLUMNS WHERE TABLE_SCHEMA = '%s' and TABLE_NAME='%s';`,subSql,databaseName,tableName)
	gres, gerr := engine.Query(sqlStr)

	if gerr != nil {
		fmt.Println("Error",gerr)
		return ""
	}

	str := fmt.Sprintf("\n#### TABLE:%s\n|",tableName)
	strs := "|"
	for _,key := range TitleList{
		str = str + key + "|"
		strs = strs + ":---|"
	}
	str = str + "\n" + strs + "\n"
	for _, v := range gres  {
		str = str + "|"
		for _,key := range TitleList{
			str = str + string(v[key]) + "|"
		}
		str = str + "\n"
	}

	return str
}


func QueryTableStructToHtml(TitleList []string,subSql,databaseName,tableName string)string{

	sqlStr := fmt.Sprintf(`SELECT %s FROM information_schema. COLUMNS WHERE TABLE_SCHEMA = '%s' and TABLE_NAME='%s';`,subSql,databaseName,tableName)
	gres, gerr := engine.Query(sqlStr)
	if gerr != nil {
		fmt.Println("Error",gerr)
		return ""
	}

	str := fmt.Sprintf("\n<h4>TABLE:%s</h4>\n<table>\n<tr>",tableName)
	for _,key := range TitleList{
		str = str + fmt.Sprintf("<th>%s</th>",key)
	}
	str = str + "</tr>"

	for _, v := range gres  {
		str = str + "<tr>"
		for _,key := range TitleList{
			str = str + fmt.Sprintf("<td>%s</td>",string(v[key]))
		}
		str = str + "</tr>\n"
	}
	str = str + "</table>"
	return str
}



func GetAllTable() []string{
	gres, gerr := engine.Query("show tables;")
	if gerr != nil {
		fmt.Println("Error",gerr)
		return nil
	}

	tableList := []string{}
	for _, v := range gres  {
		for _,tableName := range v {
			tableList = append(tableList,string(tableName))
		}
	}

	return tableList
}

func SqlToTitleList(str string) []string {
	//var digitsRegexp = regexp.MustCompile(`(SELECT|select).*?(from|FROM)`)

	//str := digitsRegexp.FindStringSubmatch(sqlStr)[0]
	//str = strings.Replace(str,"SELECT","",-1)
	//str = strings.Replace(str,"select","",-1)
	//str = strings.Replace(str,"FROM","",-1)
	//str = strings.Replace(str,"from","",-1)
	str = strings.Replace(str,"as","AS",-1)
	str = strings.Replace(str,`'`,"",-1)
	str = strings.Replace(str,`"`,"",-1)
	str = strings.Replace(str," ","",-1)

	var titleList []string

	for _,v := range strings.Split(str,","){
		titles := strings.Split(v,"AS")
		titleList = append(titleList,titles[len(titles)-1])
	}
	return titleList
}


func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}