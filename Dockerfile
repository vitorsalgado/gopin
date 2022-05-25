FROM golang:1.18 as build
WORKDIR /app
COPY go.mod go.sum Makefile ./
RUN go mod download && make deps
COPY . .
RUN make build && mv api bin/api

# ---

FROM scratch
COPY --from=build /app/bin /
EXPOSE 8080
CMD ["/app"]
