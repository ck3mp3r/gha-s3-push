FROM golang:latest as gopher
RUN mkdir /build
COPY src/golang /build
WORKDIR /build/s3push/cmd
RUN go test
WORKDIR /build/s3push
RUN go build s3push

FROM scratch as final
COPY --from=gopher ["/build/s3push", "/bin/s3push"]

ENV AWS_SDK_LOAD_CONFIG=1

WORKDIR /work

ENTRYPOINT ["/bin/s3push"]
