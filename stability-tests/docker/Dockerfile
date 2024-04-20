ARG katd_IMAGE
ARG katminer_IMAGE

FROM ${katd_IMAGE} as katd
FROM ${katminer_IMAGE} as katminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/Katkoin/katd

WORKDIR /go/src/github.com/Katkoin/katd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=katd /app/ /app/
COPY --from=katminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/Katkoin/katd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]