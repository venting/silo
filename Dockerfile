FROM golang:1.8.0-alpine
LABEL maintainer "Venting"

EXPOSE 8080

ENV GOPATH=/go

RUN addgroup venting \
     && adduser -S -G venting venting \
     && apk --update add ca-certificates \
     && apk --update add --virtual build-deps git

COPY ./ /go/src/github.com/venting/silo
WORKDIR /go/src/github.com/venting/silo

RUN go get \
 && go test ./... \
 && go build -o /bin/main

USER venting

CMD [ "/bin/main" ]