FROM heroiclabs/nakama-pluginbuilder:3.22.0 AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend

COPY go.mod .
COPY main.go .
COPY ./database ./database
COPY ./service ./service
COPY ./util ./util

RUN go build --trimpath --mod=mod --buildmode=plugin -o ./backend.so

FROM registry.heroiclabs.com/heroiclabs/nakama:3.22.0

COPY --from=go-builder /backend/backend.so /nakama/data/modules/
COPY local.yml /nakama/data/
COPY ./migrations /nakama/data/migrations