FROM --platform=$BUILDPLATFORM golang:1.23 AS builder

WORKDIR /app

# install dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY Makefile Makefile
COPY hathora hathora
COPY internal internal

ARG TARGETOS
ARG TARGETARCH
ARG BUILD_VERSION

RUN make build BUILD_VERSION=${BUILD_VERSION} TARGETOS=${TARGETOS} TARGETARCH=${TARGETARCH}

# final image
FROM gcr.io/distroless/static-debian12:nonroot

ARG TARGETOS
ARG TARGETARCH

COPY --from=builder /app/bin/hathora-${TARGETOS}-${TARGETARCH} /hathora

ENTRYPOINT ["/hathora"]
