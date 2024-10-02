ARG IMAGE_PREFIX
ARG GO_IMAGE=golang:1.22.0

FROM ${IMAGE_PREFIX}${GO_IMAGE} as builder

WORKDIR /workspace

COPY Makefile Makefile
COPY go.mod go.mod
COPY go.sum go.sum

RUN make mod-download

COPY dev.yaml dev.yaml
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

FROM scratch
WORKDIR /
COPY --from=builder /workspace/bin/server .
COPY --from=builder /workspace/dev.yaml .

# Run as non-root
USER 65532:65532

ENTRYPOINT ["/server", "api", "-c", "dev.yaml"]
