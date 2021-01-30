FROM golang:latest

WORKDIR /go/src
COPY . .

# RUN go mod init github.com/csarnataro/test

# ENV GOROOT /usr/local/go/src

RUN go mod download

RUN go get -d -v ./...
RUN go install -v ./...

# RUN go build -o app


# CMD ["/bin/sh -c \"echo `pwd`\""]
# ./app 
CMD ["go", "run", "app.go"]