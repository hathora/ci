FROM --platform=$BUILDPLATFORM golang:1.22 AS builder

WORKDIR /app

# install dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY Makefile Makefile
COPY cmd cmd
COPY internal internal

ARG TARGETOS
ARG TARGETARCH
ARG BUILD_VERSION

RUN make build

# final image
FROM gcr.io/distroless/static-debian12:nonroot

ARG TARGETOS
ARG TARGETARCH

COPY --from=builder /app/bin/hathora-ci-${TARGETOS}-${TARGETARCH} /hathora-ci

ENTRYPOINT ["/hathora-ci"]
