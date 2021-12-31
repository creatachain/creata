# Simple usage with a mounted data directory:
# > docker build -t creata .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.creata:/creata/.creata creata creatad init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.creata:/creata/.creata creata creatad start
FROM golang:1.15-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/creatachain/gaia

# Add source files
COPY . .

RUN go version

# Install minimum necessary dependencies, build Creata SDK, remove packages
RUN apk add --no-cache $PACKAGES && \
    make install

# Final image
FROM alpine:edge

ENV CREATA /creata

# Install ca-certificates
RUN apk add --update ca-certificates

RUN addgroup creata && \
    adduser -S -G creata creata -h "$CREATA"

USER creata

WORKDIR $CREATA

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/creatad /usr/bin/creatad

# Run creatad by default, omit entrypoint to ease using container with creatacli
CMD ["creatad"]
