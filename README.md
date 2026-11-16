# HTTP Course

## Running the project

### TCP

**Run the listener (and output to a temporary file):**

```sh
go run ./cmd/tcplistener | tee /tmp/tcplistener.txt
```

To close the listener:

```sh
Ctrl + C
```

**In another terminal, send a GET request:**

```sh
curl http://localhost:42069/prime/agen
```

Expected output:

```text
Request line:
- Method: GET
- Target: /prime/agen
- Version: 1.1
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
