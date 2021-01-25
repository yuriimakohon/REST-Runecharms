# REST API for scandinavian charms
## Description
Service for CRUD operation through http-REST API with "Scandinavian charm" entity.  
Charms placed in in-memory storage.  
**Create**, **read** by `id` or `filters`, **delete** by `id`, **update** by `id`

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
### Filter usage

`/charm?key=value`

| Key | Value | Example |
|:---:|:-----:|:--------|
|**rune**|_string_|`rune=Mannaz`|
|**god**|_string_|`god=Odin`|
|**power**|_int_|`power=420`|
#### Example
**GET** `/charm?rune=Ansuz&power=200` returns:
```json5
[
  {
    "id" : 2,
    "rune" : "Ansuz",
    "god" : "Odin",
    "power" : 200
  },
  {
    "id" : 4,
    "rune" : "Ansuz",
    "god" : "Loki",
    "power" : 200
  },
  {
    "id" : 7,
    "rune" : "Ansuz",
    "god" : "Freya",
    "power" : 200
  }
]
```

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
