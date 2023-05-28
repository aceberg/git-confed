FROM golang:alpine AS builder

RUN apk add build-base
COPY cmd /src/cmd
COPY internal /src/internal
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src/cmd/git-confed/ && CGO_ENABLED=0 go build -o /git-confed .


FROM scratch

WORKDIR /data/git-confed
WORKDIR /app

COPY --from=builder /git-confed /app/

ENTRYPOINT ["./git-confed"]