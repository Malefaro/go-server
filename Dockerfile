# # STEP 1 build executable binary

# FROM golang:alpine as builder
# COPY . $GOPATH/src/go-server
# WORKDIR $GOPATH/src/go-server

# #get dependancies
# #you can also use dep
# RUN go get -d -v

# #build the binary
# RUN go build -o /go/bin/go-server

# # STEP 2 build a small image

# # start from scratch
# FROM scratch

# # Copy our static executable
# COPY --from=builder /go/bin/go-server /go/bin/go-server

# EXPOSE 8080

# CMD ["/go/bin/go-server"]

FROM golang:1.10.3-alpine

COPY . /go/src/go-server

RUN go install go-server

EXPOSE 8080

CMD ["go-server"]