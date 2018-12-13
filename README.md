1.将代码git clone到$GOPATH下的src目录下
2.进入client目录,运行 go install nodiedGO/client, 将在$GOPATH/bin目录下得到client可执行文件
3.进入server, 运行 go run server.go

备注: 如果怕go代码死掉,可以将代码放入client.go中,其在死掉后将会由server.go不断启动!