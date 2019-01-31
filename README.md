# tps tool

install [dep]() for go

```
dep ensure -update -v
go build
```

```
./forceiotps -h
Usage of ./forceiotps:
  -cfg string
        confg file path (default "./config.json")
  -g int
        gorountinue num to send (default 10)
  -s int
        time sleep between send request (ms) (default 10)
```