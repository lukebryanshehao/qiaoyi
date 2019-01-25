## 项目部署
安装go1.11或1.12版本<br>
go get 需要的包<br>
复制项目或打包复制到GOPATH目录下的src目录下<br>
## 项目启动
方法1:复制源码,cd到项目目录下后台运行<br>
nohup go run main.go &<br>
cd到项目目录下<br>
<br>
方法2:打包,复制包,cd到项目目录下后台运行<br>
(1)在windows上打包:<br>
SET GOOS=linux<br>
SET GOARCH=amd64<br>
go build main.go<br>
(2)linux后台运行<br>
nohup ./main &<br>
(3)查看进程<br>
ps aux | grep main
