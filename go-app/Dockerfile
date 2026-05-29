FROM golang:1.26-alpine

# 一旦PWDを/appにする
WORKDIR /app 
# その後に/appに/go-app(Dockerfileと同じ層)をコピーする
COPY . .
# そして、/app/cmd/myapp/main.goをビルドして、go-appという実行ファイルを作成する
RUN go build -o go-app cmd/myapp/main.go
CMD ["./go-app"]
