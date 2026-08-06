# =============================================================================
# Stage 1: Go toolchain + pre-downloaded module cache
# =============================================================================
FROM golang:1.26-bookworm AS go-builder

# Install golangci-lint via go install — more reliable than the curl script
# since the Go toolchain is already present in this image.
RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

# Pre-warm the Go module cache.
# Copy only module files first so Docker can cache this layer independently
# of source code changes.
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

# =============================================================================
# Stage 2: Node 24 + pre-installed npm dependencies
# =============================================================================
FROM node:24-bookworm-slim AS node-builder

WORKDIR /build/tests/e2e-shell-use
COPY tests/e2e-shell-use/package.json tests/e2e-shell-use/package-lock.json ./
RUN npm ci

# =============================================================================
# Stage 3: Combined development / CI image
# =============================================================================
FROM debian:bookworm-slim AS dev

# ── System packages ───────────────────────────────────────────────────────────
RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        git \
        make \
        bash \
        build-essential \
        procps \
    && rm -rf /var/lib/apt/lists/*

# ── Go runtime ────────────────────────────────────────────────────────────────
COPY --from=go-builder /usr/local/go /usr/local/go
# golangci-lint binary (go install writes to $GOPATH/bin = /go/bin)
COPY --from=go-builder /go/bin/golangci-lint /usr/local/bin/golangci-lint
# Pre-downloaded module cache (avoids network on every go build / go test)
COPY --from=go-builder /go/pkg/mod /go/pkg/mod

# ── Node runtime ──────────────────────────────────────────────────────────────
COPY --from=node-builder /usr/local/bin /usr/local/bin
COPY --from=node-builder /usr/local/lib/node_modules /usr/local/lib/node_modules

# Pre-installed project node_modules
COPY --from=node-builder /build/tests/e2e-shell-use/node_modules \
                         /workspace/tests/e2e-shell-use/node_modules

# ── shell-use CLI (required by @microsoft/shell-use Node package) ─────────────
# The @microsoft/shell-use npm package is a thin client; the actual automation
# engine is a native Rust binary that must be available on PATH.
RUN curl --proto '=https' --tlsv1.2 -LsSf \
        https://raw.githubusercontent.com/microsoft/shell-use/main/install/install.sh \
    | bash \
    && (mv "${HOME}/.cargo/bin/shell-use" /usr/local/bin/shell-use 2>/dev/null \
        || find /root /tmp -name "shell-use" -type f -exec mv {} /usr/local/bin/shell-use \; 2>/dev/null \
        || true)

# ── Environment ───────────────────────────────────────────────────────────────
ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GOMODCACHE="/go/pkg/mod"
# Allow go tooling to find modules in the pre-warmed cache
ENV GOFLAGS="-mod=mod"
# Standard color terminal — required for BubbleTea to render correctly
ENV TERM=xterm-256color

WORKDIR /workspace

# Default: drop into bash (overridden by Makefile docker-* run targets)
CMD ["bash"]
