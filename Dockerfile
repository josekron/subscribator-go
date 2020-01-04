FROM golang:alpine
RUN mkdir /go/src/subscribator-go
ADD . /go/src/subscribator-go/
WORKDIR /go/src/subscribator-go
COPY . /go/src/subscribator-go
EXPOSE 8080
CMD ["go", "run", "/go/src/subscribator-go/main.go"]
