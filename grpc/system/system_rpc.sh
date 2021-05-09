#! /bin/sh

# 系统服务 - 该文件目录下执行生成GO RPC文件命令
protoc -I $(pwd)/ $(pwd)/system.proto --go_out=plugins=grpc:./rpc