

## 本地启动
+ `go run cmd/main.go`
+ access http://127.0.0.1:8080/

## 项目介绍
+ 模拟选秀抽签流程，通过网页体验抽签结果
+ 使用github.com/PuerkitoBio/goquery抓取球队的图片数据

## 体验地址
+ https://cc24-120-230-98-139.ngrok.io/
+ 若地址失效，通过微信公众号（拉巴力不吃三文鱼），发送【nba选秀】即可获取最新体验地址

## 部署情况
+ `sh make.sh` 生成可执行文件
+ `sh upload.sh` 上传
+ `sh run.sh` 运行
+ 目前部署在家庭网络小型服务器，使用ngork代理到外网