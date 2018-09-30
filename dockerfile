FROM golang:alphine

COPY . /go/src/go-server 

RUN go install go-server


EXPOSE 8080

CMD ["go-server"] 