mkdir bin/ 2> /dev/null
go get github.com/shirou/gopsutil/mem

go build -o bin/service mem.go