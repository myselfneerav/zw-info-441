CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o gateway *.go
docker build -t zanewebb/zanewebbuw .