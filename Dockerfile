FROM golang:1.17
WORKDIR /go/src/app
COPY . .
RUN go build ./cmd/p2pnc

FROM quay.io/openshift/origin-base:4.10
COPY --from=0 /go/src/app/p2pnc /bin
