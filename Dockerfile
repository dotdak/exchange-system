# Build stage
FROM golang AS build-env
ENV CGO_ENABLED=0
WORKDIR /exchange-system
COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . .
RUN make build

# Production stage
FROM scratch
COPY --from=build-env /exchange-system/app /

ENTRYPOINT ["/app"]
