# use image 'golang:1.19.1' for the application
FROM golang:1.19.1-alpine

# Setup folders
RUN mkdir /app

WORKDIR /app

COPY go.mod .

RUN  go mod download && go mod verify

# copy all file to image
COPY . .

RUN go mod tidy

# obtain the package needed to run code. Alternatively use GO Modules.
RUN go get github.com/lib/pq

# compile application
RUN go build ./cmd/served/main.go

# container listens on specified network ports at runtime
EXPOSE 8080

CMD ["./main"]
