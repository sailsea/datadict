# datadict
一个自动生成数据字典的mysql工具

#### 参数说明

|参数|必选|默认|注释|
|:---|:---|:---|:---|
|DriverName|是| |mysql的连接信息|
|GeneratedPath|否|data_dict|生成文件保存路径|
|Format|否|markdown|选择生成文件的类型 支持 markdown/html|
|SubSql|否|COLUMN_NAME AS 字段, COLUMN_TYPE AS 类型, IS_NULLABLE AS 是否为空, COLUMN_DEFAULT AS 默认, COLUMN_COMMENT AS 注释|自定义要显示的内容|

#### 使用方法

打开链接https://github.com/SailSea/datadict/releases/tag/v1.0.0
```
datadict-cli-darwin-386.zip
datadict-cli-darwin-amd64.zip
datadict-cli-linux-386.zip
datadict-cli-linux-amd64.zip
datadict-cli-windows-386.zip
datadict-cli-windows-amd64.zip
```

解压后放在环境变量的目录中

使用例子

指定数据库信息其他参数使用默认来生成数据字典
```bash
./datadict dd -DriverName="root:555666@tcp(47.98.218.146:3306)/mall_user?charset=utf8"
```


生成效果
#### html
![html](https://github.com/SailSea/datadict/blob/master/image/html.png)
#### markdown
![markdown](https://github.com/SailSea/datadict/blob/master/image/markdown.png)