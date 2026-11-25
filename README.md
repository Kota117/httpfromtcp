# HTTP Course

## Running the project

### TCP

**Run the listener (and output to a temporary file):**

```sh
go run ./cmd/tcplistener | tee /tmp/body.txt
```

To close the listener:

```sh
Ctrl + C
```

**In another terminal, send a request:**

```sh
curl -X POST http://localhost:42069/coffee \
-H 'Content-Type: application/json' \
-d '{"type": "dark mode", "size": "medium"}'
```

Expected output:

```text
Request line:
- Method: GET
- Target: /use-neovim-btw
- Version: 1.1
Headers:
- user-agent: curl/8.5.0
- accept: */*
- host: localhost:42069
```

*Note: Currently, this terminal will need to closed with `Ctrl + C` after sending the request since the listener doesn't respond.*

### UDP

Run a netcat listener:

```sh
nc -u -l 42069
```

In another terminal, start the UDP sender:

```sh
go run ./cmd/udpsender
```

*Note: The listener does not need to be running to send the message; UDP doesn't require the handshake like TCP does.*

## Testing the project

From the project root:

```sh
go test ./...
```

### HTTP

**Run the server:**

```sh
go run ./cmd/httpserver/
```

**In another terminal, send a request (with verbose output):**

```sh
curl -v localhost:42069
```

Expected output:

```text
* Host localhost:42069 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:42069...
* Connected to localhost (::1) port 42069
> GET / HTTP/1.1
> Host: localhost:42069
> User-Agent: curl/8.5.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< content-type: text/html
< content-length: 129
< connection: close
< 
<html>
<head>
<title>200 OK</title>
</head>
<body>
<h1>Success!</h1>
<p>Your request was an absolute banger.</p>
</body>
</html>
* Closing connection
```

More examples:

```sh
curl -v localhost:42069/yourproblem
```

Expected output:

```text
* Host localhost:42069 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:42069...
* Connected to localhost (::1) port 42069
> GET /yourproblem HTTP/1.1
> Host: localhost:42069
> User-Agent: curl/8.5.0
> Accept: */*
> 
< HTTP/1.1 400 Bad Request
< content-length: 140
< connection: close
< content-type: text/html
< 
<html>
<head>
<title>400 Bad Request</title>
</head>
<body>
<h1>Bad Request</h1>
<p>Your request honestly kinda sucked.</p>
</body>
</html>
* Closing connection
```

```sh
curl -v localhost:42069/myproblem
```

Expected output:

```text
* Host localhost:42069 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:42069...
* Connected to localhost (::1) port 42069
> GET /myproblem HTTP/1.1
> Host: localhost:42069
> User-Agent: curl/8.5.0
> Accept: */*
> 
< HTTP/1.1 500 Internal Server Error
< content-length: 164
< connection: close
< content-type: text/html
< 
<html>
<head>
<title>500 Internal Server Error</title>
</head>
<body>
<h1>Internal Server Error</h1>
<p>Okay, you know what? This one is on me.</p>
</body>
</html>
* Closing connection
```

**To see chunked encoding, use netcat**:

```sh
echo -e "GET /httpbin/stream/100 HTTP/1.1\r\nHost: localhost:42069\r\nConnection: close\r\n\r\n" | nc localhost 42069
```

Full example output, including once where the entire chunked response was not read at once (id=86):

```text
HTTP/1.1 200 OK
connection: close
content-type: text/plain
transfer-encoding: chunked

17b
{"id":0,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":1,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":2,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":3,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":4,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":5,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":6,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":7,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":8,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17b
{"id":9,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":10,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":11,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":12,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":13,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":14,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":15,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":16,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":17,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":18,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":19,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":20,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":21,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":22,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":23,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":24,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":25,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":26,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":27,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":28,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":29,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":30,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":31,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":32,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":33,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":34,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":35,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":36,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":37,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":38,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":39,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":40,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":41,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":42,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":43,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":44,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":45,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":46,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":47,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":48,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":49,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":50,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":51,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":52,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":53,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":54,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":55,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":56,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":57,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":58,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":59,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":60,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":61,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":62,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":63,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":64,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":65,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":66,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":67,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":68,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":69,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":70,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":71,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":72,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":73,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":74,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":75,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":76,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":77,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":78,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":79,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":80,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":81,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":82,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":83,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":84,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":85,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

62
{"id":86,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["G
11a
o-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":87,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":88,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":89,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":90,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":91,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":92,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":93,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":94,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":95,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":96,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":97,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":98,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

17c
{"id":99,"args":{},"headers":{"Accept-Encoding":["gzip"],"Host":["httpbingo.org"],"User-Agent":["Go-http-client/2.0"],"Via":["2 fly.io"],"X-Forwarded-For":["67.176.82.120, 66.241.125.232"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Ssl":["on"],"X-Request-Start":["t=1795646098329765"]},"origin":"67.176.82.120","url":"https://httpbingo.org/stream/100"}

0

```
