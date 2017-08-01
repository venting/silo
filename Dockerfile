FROM golang:1.8-alpine as builder

RUN apk --update add ca-certificates \
    && apk --update add --virtual build-deps git linux-headers build-base

COPY ./ /go/src/github.com/venting/silo
WORKDIR /go/src/github.com/venting/silo

RUN go get -u -v ./... \
    && go test ./... \
    && go build -o /bin/silo-agent

FROM alpine
LABEL maintainer "venting"

RUN addgroup venting \
     && adduser -S -G venting venting \
     && apk --update --no-cache add ca-certificates

USER venting

EXPOSE 8080

COPY --from=builder /bin/silo-agent /bin/silo-agent

CMD [ "/bin/silo-agent" ]
