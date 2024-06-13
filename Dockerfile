FROM golang:1.22-alpine AS builder

#set the working directory
WORKDIR /app

# install depedencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source project
COPY . .

# install air
RUN go install github.com/air-verse/air@latest

# Release the binary
FROM builder AS release

COPY --from=builder /app /app
COPY --from=builder /go/bin/air /bin/air

CMD ["air"]