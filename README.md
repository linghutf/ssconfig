## ssconfig
- 这是一个便捷获取ss免费节点信息的工具，仅提供源码以供参考。
- 使用了goquery获取网页信息,ffjson快速格式化和反格式化json

# 准备
```
go get github.com/PuerkitoBio/goquery
go get github.com/pquerna/ffjson/ffjson
```
# 使用
```
1. go build main.go
2. ./main                 //当前目录出现最新配置文件
3. sslocl -c config*.json //enjoy fall out!

```
## 预编译文件
1. win64: main.exe

# ssconfig
- This is an util pacakge for scrap useful ss config node.
- It's a expriment for me to learn `goquery` and `ffjson`,and so simple.
- I will be happy if it can help you sometimes.

# prepare
```
go get github.com/PuerkitoBio/goquery
go get github.com/pquerna/ffjson/ffjson
```
# Usage:
## Download the package
```
1. go build main.go
2. ./main                 //config file disappear in current dir.
3. sslocl -c config*.json //enjoy fall out!
```
## or just modifie `main.go` like:
```
package main
import "<the package path>/ss"

func main(){
ss.ScrapUsefulNode(1080)
}
```

## Then`you will get three `config.json`,run it:
```
sslocal -c config[0/1/2].json
```
## precompiled
1. win64: main.go

---

## What you can get from this repository:
1. parse html by using `goquery`
2. json marshal and un marshal to files.

## attention
**I will provide static execute file just for daily use.
And the website is not *charitable*,just use it in emergency,
support them and it only cost you a bottle of water!**
