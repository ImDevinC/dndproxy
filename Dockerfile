ARG GO_VERSION=1.22

FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:${GO_VERSION} AS builder

WORKDIR /src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o /proxy ./main.go

FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/static

USER nonroot:nonroot

COPY --from=builder --chown=nonroot:nonroot /proxy /proxy

ENTRYPOINT [ "/proxy" ]
