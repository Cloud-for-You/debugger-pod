# Build binary
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.25.5 as builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-X main.Commit=$(git rev-parse HEAD)" -a -o app cmd/main.go

# Build distroless container from binary
FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/static:nonroot
LABEL org.opencontainers.image.source="https://github.com/Cloud-for-You/debugger-pod"
WORKDIR /
COPY --from=builder /src/app /
ENTRYPOINT ["/app"]