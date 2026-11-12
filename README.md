# HTTP Course

## Running the project

Run the listener (and output to a temporary file):

```sh
go run . | tee /tmp/tcp.txt
```

In another terminal:

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
