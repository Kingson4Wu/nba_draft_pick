# NBA Draft Pick

[![CodeSize](https://img.shields.io/github/languages/code-size/Kingson4Wu/nba_draft_pick)](https://github.com/Kingson4Wu/nba_draft_pick)
[![LICENSE](https://img.shields.io/badge/license-MIT-green)](https://mit-license.org/)

<details>
<summary><strong>README 中文版本</strong></summary>
<div>

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

</div>
</details>

## Local deployment
+ `go run cmd/main.go`
+ access http://127.0.0.1:8080/

## Project introduction
+ Simulate the draft lottery process and experience the results through a web page.
+ Use github.com/PuerkitoBio/goquery to grab image data of NBA teams.

## Experience address
+ https://cc24-120-230-98-139.ngrok.io/
+ If the address is unavailable, send "nba draft" to the WeChat public account "拉巴力不吃三文鱼" to obtain the latest experience address.

## Deployment status
+ `sh make.sh` to generate executable files.
+ `sh upload.sh` to upload.
+ `sh run.sh` to run.

+ Currently deployed on a small home network server, using ngrok to proxy to the external network.
