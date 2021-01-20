FROM golang:1.15.6-alpine3.12 AS go-builder
WORKDIR /src/
RUN apk --no-cache add upx git
ENV GO111MODULE on

COPY go.mod .
RUN go mod download

COPY . .

RUN cd cmd/server && CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -ldflags "-s -w" -o /app/app .
RUN upx /app/app

FROM scratch
COPY --from=go-builder /app/app /

ENTRYPOINT [ "/app" ]