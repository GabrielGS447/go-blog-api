FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go-blog-api

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /go-blog-api /go-blog-api
EXPOSE 8080
USER nonroot:nonroot
ENV GO_ENV=production \
PORT=8080 \
MYSQL_URL=root:password@tcp(mysql)/go-blog-api?charset=utf8mb4&parseTime=True&loc=Local \
JWT_SECRET=secret \
RESET_DB=false
ENTRYPOINT ["/go-blog-api"]
