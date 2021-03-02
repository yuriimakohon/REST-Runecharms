# REST API for scandinavian charms
## Description
Service for CRUD operation through http-REST API with "Scandinavian charm" entity.  
Charms placed in in-memory storage.  
**Create**, **read**, **delete** and **update** by `id`

## Run
1. Set environment variable `CHARM_STORAGE` with one of these:
- `redis`
- `inmem`
- `postgres`
- `mongo`
- `grpc`
- `elastic`
2. (Only if you choosed `grpc` in a previous step) Set `CHARM_GRPC_STORAGE` with one of these:
- `redis`
- `inmem`
- `postgres`
- `mongo`
- `elastic`  
After that run gRPC server: `go run cmd/grpc/main.go`
3. Run http server: `go run cmd/server/main.go`

## HTTP API
### CREATE
**POST**  

`/charm`
Content in **_json_** format:  
```json5
{
    "rune" : "Mannaz",
    "god" : "Odin",
    "power" : 270
}
```
Returns created entity with generated id in _**json**_:
```json5
{
    "id" : 0,
    "rune" : "Mannaz",
    "god" : "Odin",
    "power" : 270
}
```

---

### READ
**GET** `/charm`  
Returns all charms in _**json**_:
```json5
[
  {
    "id" : 0,
    "rune" : "Mannaz",
    "god" : "Odin",
    "power" : 270
  },
  {
    "id" : 1,
    "rune" : "Ansuz",
    "god" : "Tyr",
    "power" : 320
  }
]
```

[comment]: <> (### Filter usage)

[comment]: <> (`/charm?key=value`)

[comment]: <> (| Key | Value | Example |)

[comment]: <> (|:---:|:-----:|:--------|)

[comment]: <> (|**rune**|_string_|`rune=Mannaz`|)

[comment]: <> (|**god**|_string_|`god=Odin`|)

[comment]: <> (|**power**|_int_|`power=420`|)

[comment]: <> (#### Example)

[comment]: <> (**GET** `/charm?rune=Ansuz&power=200` returns:)

[comment]: <> (```json5)

[comment]: <> ([)

[comment]: <> (  {)

[comment]: <> (    "id" : 2,)

[comment]: <> (    "rune" : "Ansuz",)

[comment]: <> (    "god" : "Odin",)

[comment]: <> (    "power" : 200)

[comment]: <> (  },)

[comment]: <> (  {)

[comment]: <> (    "id" : 4,)

[comment]: <> (    "rune" : "Ansuz",)

[comment]: <> (    "god" : "Loki",)

[comment]: <> (    "power" : 200)

[comment]: <> (  },)

[comment]: <> (  {)

[comment]: <> (    "id" : 7,)

[comment]: <> (    "rune" : "Ansuz",)

[comment]: <> (    "god" : "Freya",)

[comment]: <> (    "power" : 200)

[comment]: <> (  })

[comment]: <> (])

[comment]: <> (```)

**GET** `/charm/{id}`  
Returns one entity with `id`

---

### DELETE

**DELETE**  `/charm/{id}`  
Delete and returns deleted entity from storage

---

### UPDATE
**PUT** `/charm/{id}`
Updates entity with `id` according to content in _**json**_(each field is optional):
```json5
{
  "rune" : "Mannaz",
  "god" : "Loki",
  "power" : 300
}
```
Returns updated entity
