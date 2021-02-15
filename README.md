# README

g.an - A tiny announcement server with Go.

Use for your small blog or as news distribution application.

Features:

* Manage announcements which are simple text snippets containing a `title` and plain text `content`
* Allow subscribing, so that users (via mail) get notified if there's a new announcement

There's also a [g.an frontend](https://github.com/v4rakh/gan-frontend).

## Install

1. Run `make clean dependencies` to fetch dependencies
2. Start `github.com/v4rakh/gan/cmd/gan-server` as Go application and ensure to have _required_ environment variables
   set

## Configuration

The following environment variables can be used to modify application behavior.

| Variable | Purpose | Required | Default/Description |
|:---|:---|:---|:---|
| DOMAIN | The domain | required | `http://localhost` (adapt and use full protocol, **no** trailing slash) |
| ADMIN_USER | Admin user name for login | required |  |
| ADMIN_PASSWORD | Admin password for login | required |  |
| DB_FILE | Path to the SQLITE file | optional | `<XDG_DATA_DIR>/gan/gan.db`, e.g. `~/.local/share/gan/gan.db` |
| SERVER_PORT | Port | optional | `8080` |
| SERVER_LISTEN | Server's listen address | optional | empty which equals `0.0.0.0` |
| CORS_ALLOW_ORIGIN | CORS configuration | optional | `*` |
| CORS_ALLOW_METHODS | CORS configuration | optional | `GET, POST, PUT, PATCH, DELETE, OPTIONS` |
| CORS_ALLOW_HEADERS | CORS configuration | optional | `Authorization, Content-Type` |
| GIN_MODE           | GIN mode, e.g. for debugging | optional | `debug` and `release` in docker |
|MAIL_ENABLED|Mails enabled?|required|`true` [true,false], otherwise subscription mails won't work|
|MAIL_FROM|From mail address|required if MAIL_ENABLED=true||
|MAIL_HOST|SMTP Setting|required if MAIL_ENABLED=true||
|MAIL_PORT|SMTP Setting|required if MAIL_ENABLED=true||
|MAIL_ENCRYPTION|SMTP Setting|required if MAIL_ENABLED=true|`SSL` [NONE, SSL, TLS]|
|MAIL_AUTH_USER|SMTP Setting|required if MAIL_ENABLED=true||
|MAIL_AUTH_PASSWORD|SMTP Setting|required if MAIL_ENABLED=true||
|MAIL_AUTH_TYPE|SMTP Setting|required if MAIL_ENABLED=true|`PLAIN` [PLAIN,LOGIN,CRAM_MD5]|

## Development

When language files under `lang/` changes, execute `make pkg-blobs` to generate a new blob bundle which
will be embedded into the Go binary.

## Release & deployment

### Native

Run `make clean build` and the binary will be placed into the `bin/` folder.

### Docker

To build docker images, do the following

```sh
export IMG_NAME="gan-server";
export IMG_TAG="latest";
sudo docker build --rm --no-cache -t ${IMG_NAME}:${IMG_TAG} .

# and/or

export IMG_NAME="gan-server";
export IMG_TAG="latest";
export REMOTE_REPO_URL="varakh";
sudo docker build --rm --no-cache -t ${REMOTE_REPO_URL}/${IMG_NAME}:${IMG_TAG} .
sudo docker push ${REMOTE_REPO_URL}/${IMG_NAME}:${IMG_TAG}
```

An example how to run with a persistent database file located on host system in `/my/host/data/folder/app.db`:

```sh
sudo docker run -p 8080:8080 \
    -v /my/host/data/folder:/data
    -e DB_FILE=/data/app.db \
    -e ADMIN_USER=admin \
    -e ADMIN_PASSWORD=changeit \
    -e MAIL_ENABLED=false \
    varakh/gan-server:latest
```  