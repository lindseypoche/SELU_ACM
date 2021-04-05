FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git && \ 
    apk --no-cache add ca-certificates

ENV GOPATH=/go
ENV APIPATH=$GOPATH/src/github.com/cmd-ctrl-q/SELU_ACM/api

WORKDIR $APIPATH 

COPY ./api/go.mod . 
COPY ./api/go.sum .
RUN go mod download 

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# # copy main file into container root
COPY ./api/cmd/bot/main.go .

# copy everything else 
COPY ./api/internal ./internal

RUN go build -o bot .

RUN addgroup gopher && adduser -D -G gopher gopher \
    && chown -R gopher:gopher /go && \
    chmod +x ./bot

USER gopher

CMD ["./bot"]