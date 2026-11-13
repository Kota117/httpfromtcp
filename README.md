# HTTP Course

## Running the project

### TCP

Run the listener (and output to a temporary file):

```sh
go run ./cmd/tcplistener | tee /tmp/tcplistener.txt
```

To close the listener:

```sh
Ctrl + C
```

In another terminal:

- Create a persistent conncection:

  ```sh
  nc -v localhost 42069
  ```

  *Note: The listener must be running or else there will be a `Connection refused` error. This is becuase TCP requires the handshake; UDP does not.*

  Typing any text after this will be logged by the listener. `Ctrl+C` to close the connection.

- Run lesson checks:

  ```sh
  bootdev run
  ```

- Manually test with one message:

  ```sh
  printf "Do you have what it takes to be an engineer at TheStartup™?\r\n" | nc -w 1 127.0.0.1 42069
  ```

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
