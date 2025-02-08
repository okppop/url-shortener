FROM alpine:3.19

WORKDIR /app
COPY url-shortener config.yaml /app/

HEALTHCHECK --interval=5s --timeout=1s CMD ["nc", "-z","127.0.0.1:8080"]
EXPOSE 8080/tcp
CMD ["./url-shortener"]