# 编译基础镜像
FROM golang:alpine as builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -a -o cw_gateway_service .

FROM alpine as final

# 时区
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/cw_gateway_service /app/

EXPOSE 8020

ENTRYPOINT ["/app/cw_gateway_service"]