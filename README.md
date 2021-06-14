# Is Workday ?

今天是工作日吗?

一个简单用来用判断今天是不是工作日的WEB项目.

### 数据源
数据源来自政府公告

编写成`json`文件存放于 `config` 目录下。按年分文件，`2021.json`即代表2021年的数据
### example
```shell
curl --request GET -sL --url 'https://day.tangtj.cn/today'
```
response
```json
{"ok":true,"year":2021,"month":6,"day":14,"weekday":1,"Date":"","work":false,"remark":"端午节"}
```
