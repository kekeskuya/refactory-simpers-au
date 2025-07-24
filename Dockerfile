# Stage 1: Build the Go binary
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .

RUN go mod tidy
RUN go build -o main . 

# Stage 2: Run Binary
FROM alpine:edge AS runner

WORKDIR /app

#Create unprivileged user for security
RUN addgroup kecilin && adduser -G kecilin -D -s /bin/sh kecilin
RUN chown -R kecilin:kecilin /app 
USER kecilin

COPY --chown=kecilin:kecilin entrypoint.sh .
RUN chmod +x entrypoint.sh

# Ensure the env directory exists
RUN mkdir -p /app/env && chown -R kecilin:kecilin /app/env

#Copy Binary from Builder Stage
COPY --from=builder --chown=kecilin:kecilin /app/main .

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["/app/main"]