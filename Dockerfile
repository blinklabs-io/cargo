FROM golang:1.18 AS build

WORKDIR /code
COPY . .
RUN make build

FROM ubuntu:focal
COPY --from=build /code/cargo /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/cargo"]
