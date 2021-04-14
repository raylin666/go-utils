#! /bin/sh

# 系统服务 - 项目根目录下执行生成GO文件命令
protoc -I $(pwd)/grpc/system_services/ $(pwd)/grpc/system_services/proto/system_services.proto --go_out=plugins=grpc:grpc
