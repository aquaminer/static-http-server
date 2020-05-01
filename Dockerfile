FROM golang:alpine

WORKDIR /usr/src/app
CMD ["go", "run", "main.go"]
EXPOSE 8080

COPY . /usr/src/app
RUN go build main.go