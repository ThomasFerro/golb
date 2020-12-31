FROM golang:1.15.1 as back-builder

WORKDIR /src

ADD go.mod .
ADD go.sum .

ADD main.go .
ADD config ./config
ADD posts ./posts
ADD blog ./blog

RUN CGO_ENABLED=0 GOOS=linux go build -o /dist/golb

FROM alpine

WORKDIR /root

COPY --from=back-builder /dist .
COPY blog/homePageTemplate.go.html .
COPY blog/postPageTemplate.go.html .

CMD [ "./golb" ]