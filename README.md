# HTTP Course

## Running the project

Run the listener (and output to a temporary file):

```sh
go run ./cmd/tcplistener | tee /tmp/tcplistener.txt
```

In another terminal:

- Create an interactable conncection:

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

Close the server:

```sh
Ctrl + C
```
