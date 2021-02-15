#
# Build image
#
FROM alpine:3 AS builder
LABEL maintainer="Varakh <varakh@varakh.de>"

RUN apk --update upgrade && \
    apk add go gcc make sqlite && \
    # See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
    mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    rm -rf /var/cache/apk/*

# compile
WORKDIR /app
COPY . .
RUN rm -rf bin/ && \
    GO111MODULE=on go mod download && \
    CC=gcc CGO_ENABLED=1 GO111MODULE=on go build -o bin/gan-server cmd/gan-server/main.go

#
# Actual image
#
FROM alpine:3
LABEL maintainer="Varakh <varakh@varakh.de>"

ENV USER appuser
ENV GROUP appuser
ENV UID 1000
ENV GID 1000

RUN apk --update upgrade && \
    apk add sqlite && \
    rm -rf /var/cache/apk/* && \
    addgroup -S ${GROUP} -g ${GID} && \
    adduser -S ${USER} -G ${GROUP} -u ${UID}

COPY --from=builder /app/bin/gan-server /usr/bin/gan-server

USER ${USER}

ENV SERVER_PORT 8080
ENV GIN_MODE release
EXPOSE ${SERVER_PORT}
CMD ["/usr/bin/gan-server"]