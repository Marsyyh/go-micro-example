module carconf-api/auth

go 1.12

require (
	carconf-api/proto v0.0.0
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/go-micro v1.1.0
)

replace carconf-api/proto => ../proto
