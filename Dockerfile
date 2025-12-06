ARG GO_VERSION=1.25
ARG GOLANGCI_LINT_VERSION=v2.5

# ---------- Dev-образ с инструментами ----------
FROM golang:${GO_VERSION}-bookworm AS dev
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl git make bash gcc g++ gdb && rm -rf /var/lib/apt/lists/*

# Инструменты разработчика
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
  | sh -s -- -b /usr/local/bin ${GOLANGCI_LINT_VERSION}

# Непривилегированный пользователь
ARG UID=1000
ARG GID=1000
RUN groupadd -g ${GID} dev && useradd -m -u ${UID} -g ${GID} dev

RUN mkdir -p /home/dev/.air && chown dev:dev /home/dev/.air
RUN chown -R dev:dev /home/dev/.air

# GOPATH и кэши
RUN mkdir -p /go/pkg/mod /go/build-cache && chown -R dev:dev /go
ENV GOMODCACHE=/go/pkg/mod \
    GOCACHE=/go/build-cache \
    CGO_ENABLED=1

WORKDIR /workspace
USER dev

# ---------- Prod-сборка (многостадийная) ----------
FROM golang:${GO_VERSION}-bookworm AS build-prod
WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
ARG TARGETOS=TARGETARCH
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/app

FROM gcr.io/distroless/static:nonroot AS prod
WORKDIR /app
COPY --from=build-prod /out/app /app/app
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]