FROM golang:1.17 AS build

COPY . /code

WORKDIR /code

RUN make build

FROM ubuntu:focal

COPY --from=build /code/cargo /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/cargo"]
