FROM golang:1.16 as build
WORKDIR /src
COPY . .
RUN go build -o /out/vault ./cmd/web

FROM scratch as bin
COPY --from=build /out/vault /
ENTRYPOINT [ "/vault" ]