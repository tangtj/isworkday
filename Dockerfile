FROM golang:1.16 as build

ENV CGO_ENABLED 0

COPY . /build
WORKDIR /build
RUN go build -o app

FROM scratch as run
COPY --from=build /build/app /
EXPOSE 8080
CMD ["/app"]