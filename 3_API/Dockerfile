FROM docker.artifactory.kasikornbank.com:8443/alpine:3.11

RUN mkdir -p /go/exe
WORKDIR /go/exe

COPY ./simple_api /go/exe
RUN go build

CMD ["./simple_api"]