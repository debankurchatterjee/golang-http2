# Golang - Http2

This app handles http2 request using golang,by default http2 will run TLS so we need to generate the certificates
here we are generating openssl certs

## Generate Certs
```text
  openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes
```

## Run App
```text
  make run
```

## Test

```text
  curl -v --http2 https://localhost:8080/stream --insecure
```