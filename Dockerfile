ARG GO_VERSION=latest
FROM golang:${GO_VERSION} as builder
ENV CGO_ENABLED=0
WORKDIR /work
COPY . .
RUN go build -o test ./testdata

FROM scratch
COPY --from=builder /work/test /test
ENTRYPOINT ["/test"]
