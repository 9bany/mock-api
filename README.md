# mock-api
Fast mock data in Rest API server.

## Build 
```shell
$ go build . 

# ouput: mapi
# move `mapi` to local bin or use directly
```


## How to use 

1. json file
- `data.json`
```json
{
    "data": "message"
}
```

2. Mock data.json on API
```shell
    $ cat data.json | mapi serve
```

3. Test

```shell
$ curl localhost:8080/mock
# ---- response
# {
#     "data": "message"
# }

```


