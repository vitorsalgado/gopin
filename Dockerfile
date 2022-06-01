FROM golang:1.18 as build
WORKDIR /app
COPY go.mod go.sum Makefile ./
RUN make download
COPY . .
RUN make swagger && \
    mkdir -p bin/docs && \
    mv docs/openapi bin/docs/openapi
RUN make build

# ---

FROM scratch
COPY --from=build /app/bin /
EXPOSE 8080
CMD ["/gopin"]
