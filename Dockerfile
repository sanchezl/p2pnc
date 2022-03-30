FROM golang:1.17
WORKDIR /go/src/app
COPY . .
RUN go build -mod=vendor ./cmd/p2pnc

FROM quay.io/openshift/origin-base:4.5
COPY --from=0 /go/src/app/p2pnc /bin
