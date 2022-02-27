FROM golang:1.17 as builder
ENV GO111MODULE="on"
WORKDIR /app
COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build ./server/main.go 

FROM scratch
WORKDIR /
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["/main"]