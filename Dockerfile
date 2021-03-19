FROM alpine:latest

WORKDIR /app
COPY web /app/web
COPY app /app/clickbaiter

ENTRYPOINT ["/app/clickbaiter"]
