
FROM golang:alpine
RUN apk update && apk add --no-cache git
ENV USER=appuser
ENV UID=10001 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"
WORKDIR $GOPATH/src/zippo/
COPY . .

RUN go mod download
RUN go mod verify
RUN export CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/zippo ./cmd/main.go

USER appuser:appuser
# Run the hello binary.
ENTRYPOINT ["/go/bin/zippo"]

