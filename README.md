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

Expected output (no matter the request right now):

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
< Content-Type: text/plain
* no chunk, no close, no size. Assume close to signal end
< 
Hello World!
* Recv failure: Connection reset by peer
* Closing connection
curl: (56) Recv failure: Connection reset by peer
```
