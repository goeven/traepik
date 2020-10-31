# simple-plugin

## Build & Run

```bash
go build -o traefik . && ./traefik
```

The above command should start a traefik instance with a web server listening on port `8081`.

To test that the logging plugin is attached, open a new terminal and run:
```bash
curl -L localhost:8081/
```

You will get a `BadGateway` as a response from the `curl` command, but the interesting part happens in the logs of the `./traefik` command. You should get something like:
```
$ go build -o traefik . && ./traefik
INFO[0000] Configuration loaded from file: traepik/examples/simple-plugin/traefik.yaml
2020/10/31 14:54:20 main.go:32: traepik: hi-ya
```
