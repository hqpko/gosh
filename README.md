# gosh
## 示例

```
package main

import "gosh"

func main() {
	//简单的运行sh
	gosh.Run("ls -l")
  
	//新建一个会话
	s := gosh.NewSession()
	//设置环境变量&运行目录
	gopath := "~/github/"
	s.SetDir(gopath + "~/github/src/hqpko/gosh/")
	s.SetEvn("GOPATH", gopath)
	//运行
	s.Run("go build")
	s.Run("mv", gopath+"/bin/sh", "~/Desktop/")
	
 	//设置输出（time,pkg,in,out)
 	//gosh.EchoTime=true
 	//gosh.EchoPkg=true
 	//gosh.EchoIn=true
 	//gosh.EchoOut=true
}


```
