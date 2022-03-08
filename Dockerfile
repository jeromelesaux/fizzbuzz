FROM golang:1.17

WORKDIR /opt/app/fizzbuzz

COPY  . . 

RUN go mod download && go mod verify

RUN make 

RUN cp  /opt/app/fizzbuzz/fizzbuzz /go/bin/

FROM debian:buster-slim

WORKDIR /fizzbuzz

COPY --from=0 /go/bin/fizzbuzz .
COPY --from=0 /opt/app/fizzbuzz/docs ./docs/


RUN ls -al 

ARG PORT=$PORT

ENV PORT ${PORT}

CMD ["/fizzbuzz/fizzbuzz"]