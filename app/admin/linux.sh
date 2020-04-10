# linux build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./main.go -o main

// linux run
nohup ./main -port=5000 -env=development &