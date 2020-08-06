FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .

# Fetch dependencies
COPY . .
RUN make build-linux


# Second Stage, build smaller image
FROM scratch
COPY --from=builder /app/ltest_unix /app/

# Run the binary
ENTRYPOINT ["./app/ltest_unix"]
CMD ["dispatch"]