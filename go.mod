module go-server

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/raylin666/go-gin-api v0.0.0-20210508094758-6072bb3c4711
	github.com/shirou/gopsutil v3.21.4+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	google.golang.org/grpc v1.37.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/raylin666/go-gin-api => ../gin-api
