FROM node:10.23.1-slim AS npm-builder
WORKDIR /app
ADD cmd/ui/front/package.json .
ADD cmd/ui/front/package-lock.json .
RUN npm install
ADD cmd/ui/front/ .
RUN npm run build

FROM golang:1.15.6-alpine3.12 AS go-builder
WORKDIR /src/
RUN apk --no-cache add upx git
RUN go get -u github.com/kataras/bindata/cmd/bindata
ENV GO111MODULE on

COPY go.mod .
RUN go mod download

COPY pkg/client/go.mod pkg/client/
RUN (cd pkg/client && go mod download)

COPY cmd/client/go.mod cmd/client/
RUN (cd cmd/client && go mod download)

COPY cmd/ui/go.mod cmd/ui/
RUN (cd cmd/ui && go mod download)

COPY . .

WORKDIR /src/cmd/ui
COPY --from=npm-builder /dist dist
RUN bindata ./dist/...
RUN ls -la

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -ldflags "-s -w" -o /app/app .
RUN upx /app/app

FROM scratch
EXPOSE 8080
COPY --from=go-builder /app/app /

ENTRYPOINT [ "/app" ]