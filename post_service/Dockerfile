# 编译基础镜像
FROM golang:alpine as builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -a -o cw_post_service .

FROM scratch as final

# 时区
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/cw_post_service /app/

EXPOSE 8019

ENTRYPOINT ["/app/cw_post_service"]
