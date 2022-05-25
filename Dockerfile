FROM golang:1.18 as build
WORKDIR /app
COPY go.mod go.sum Makefile ./
RUN make deps
COPY . .
RUN make build
RUN mv api bin/api

# ---

FROM scratch
COPY --from=build /app/bin /
EXPOSE 8080
CMD ["/gopin"]
