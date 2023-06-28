FROM cgr.dev/chainguard/go:1.19 AS build

WORKDIR /code
COPY . .
RUN make build

FROM cgr.dev/chainguard/glibc-dynamic AS cargo
COPY --from=build /code/cargo /bin/
ENTRYPOINT ["cargo"]
