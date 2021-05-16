# Faasd Example function (golang)
This is a exampl e function for the faasd and JAMstack installion i recently build.

The idea of the function is tho get the status of different cloud services.

At the moment its only:

- jFrog
- STACKIT
- github
- digitalocean
- dropbox
- reddit
- scaleway

## faas-cli

There are several ways to download the faas-cli on your local machine. The most convinent way is to use arkade

###Get arkade

Note: you can also run without `sudo` and move the binary yourself

```shell
curl -sLS https://dl.get-arkade.dev | sudo sh

arkade --help
ark --help  # a handy alias
```

###Install faas-cli

```shell
ark get faas-cli
```

You should see following output.

```shell
Downloading faas-cli
https://github.com/openfaas/faas-cli/releases/download/0.13.9/faas-cli-darwin
8.69 MiB / 8.69 MiB [-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00%
Tool written to: /Users/dirien/.arkade/bin/faas-cli

# Add (faas-cli) to your PATH variable
export PATH=$PATH:$HOME/.arkade/bin/

# Test the binary:
/Users/dirien/.arkade/bin/faas-cli

# Or install with:
sudo mv /Users/dirien/.arkade/bin/faas-cli /usr/local/bin/
```

## Templates

```shell
faas-cli template store list

NAME                    SOURCE             DESCRIPTION
csharp                  openfaas           Official C# template
dockerfile              openfaas           Official Dockerfile template
...
golang-http              openfaas           Golang HTTP template
golang-middleware        openfaas           Golang Middleware template
...
```

Choose the golang-http template and retrieve it locally with the command:

```shell
faas-cli template store pull golang-http
```

Now you can create the function via templage with this commmand:

```shell
faas-cli new  faas-cloud-svc-status --lang golang-http
```

You will now see two files generate:

```shell
go-fn.yml
./faas-cloud-svc-status/
./faas-cloud-svc-status/handler.go
```

You can now edit handler.go and use the faas-cli to build and deploy your function.

Dependencies should be managed with a Go vendoring tool such as dep or Go modules.

I use go modules, so I created the `go.mod` and added 

```yaml
...
build_args:
  GO111MODULE: auto
...
```
to my faas-cloud-svc-status.yml

## Build and deploy

Build your function via

```shell
faas-cli build -f faas-cloud-svc-status.yml
```

Build your function via

```shell
faas-cli build -f faas-cloud-svc-status.yml
```

Push the image with following image to your registry

```shell
faas-cli push -f faas-cloud-svc-status.yml
```

Do deploy the function, we need to

```shell
faas-cli login --password xxxx
```

Don't forgett to set the URL for your openfaas server

```shell
export OPENFAAS_URL=https://faasd.ediri.online
```

Deploy via this command:

```shell
faas-cli deploy -f faas-cloud-svc-status.yml
```

Logs you can check via this command:

```shell
faas-cli logs faas-cloud-svc-status

2021-05-16T14:00:38Z 2021/05/16 14:00:38 SIGTERM received.. shutting down server in 10s
2021-05-16T14:00:38Z 2021/05/16 14:00:38 Removing lock-file : /tmp/.lock
2021-05-16T14:00:38Z 2021/05/16 14:00:38 [entrypoint] SIGTERM received.. shutting down server in 10s
2021-05-16T14:00:48Z 2021/05/16 14:00:48 No new connections allowed. Exiting in: 10s
2021-05-16T14:00:48Z 2021/05/16 14:00:48 [entrypoint] No new connections allowed. Exiting in: 10s
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Version: 0.8.4 SHA: bbd2e96214264d6b87cc97745ee9f604776dd80f
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Forking: ./handler, arguments: []
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Started logging: stderr from function.
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Started logging: stdout from function.
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Watchdog mode: http
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Timeouts: read: 10s, write: 10s hard: 10s.
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Listening on port: 8080
2021-05-16T14:00:59Z 2021/05/16 14:00:59 Writing lock-file to: /tmp/.lock
```

Current details under the [official documentation](https://docs.openfaas.com/cli/templates/)