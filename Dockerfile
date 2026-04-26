FROM quay.io/projectquay/golang:1.26 AS build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /app/kbot .

FROM quay.io/projectquay/golang:1.26

WORKDIR /app

COPY --from=build /app/kbot /app/kbot

ENTRYPOINT ["/app/kbot"]