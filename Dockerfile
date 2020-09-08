FROM golang:1.15

LABEL maintainer="Varakh <varakh@varakh.de>"

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN make clean build

ENV SERVER_PORT 8080
ENV GIN_MODE release
EXPOSE ${SERVER_PORT}
CMD ["/app/bin/gan-server"]