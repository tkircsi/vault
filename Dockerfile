FROM golang:1.16 as build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/vault ./cmd/web

FROM scratch as bin
COPY --from=build /out/vault /

EXPOSE 5000
ENTRYPOINT [ "/vault" ]