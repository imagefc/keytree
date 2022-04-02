# keytree
Transfer consul key slice to cascade key tree.

(Archived and moved to [Gitee](https://gitee.com/FlyingOnion/keytree))

## Usage
1. Get keys from `localhost:8500/v1/kv/?keys` (string slice)

```go
resp, _ := http.Get("localhost:8500/v1/kv/?keys")
defer resp.Body.Close()
var keys []string
json.NewDecoder(resp.Body).Decode(&keys)
```

```
a
a/           // item and dir with same name
a/dir1/      // empty dir
a/dir2/item3 // item without parent dir
a/item
a/item2
a1/
b/
b/dir3/
b/dir3/item4
```

2. Transfer `[]string` to `KeyMap` (unordered but have unique key index)
```go
keyMap := keytree.BuildMap(keys)
```

Each key with a slash suffix (like `a/`) represents a "dir".

Keys without parent dir (like `a/dir2/item3`) will be fixed by adding missing "parent dirs".

The JSON form of `keyMap` is shown below (may be different because map is unordered).
```json
{
  "a": null,
  "a/": {
    "dir1/": {},
    "dir2/": {
      "item3": null
    },
    "item1": null,
    "item2": null
  },
  "a1/": {},
  "b/": {
    "dir3/": {
      "item4": null
    }
  }
}
```

3. Transfer `KeyMap` to `KeyList` (ordered but have no key index)
```go
keyList := keytree.BuildKeyListFromMap(keyMap)
```
The JSON form of `keyList` is shown below. It's ordered.
```json
[
  {
    "name": "a",
    "type": "dir",
    "children": [
      {
        "name": "dir1",
        "type": "dir",
        "isLeaf": false
      },
      {
        "name": "dir2",
        "type": "dir",
        "children": [
          {
            "name": "item3",
            "type": "item",
            "isLeaf": true
          }
        ],
        "isLeaf": false
      },
      {
        "name": "item1",
        "type": "item",
        "isLeaf": true
      },
      {
        "name": "item2",
        "type": "item",
        "isLeaf": true
      }
    ],
    "isLeaf": false
  },
  {
    "name": "a1",
    "type": "dir",
    "isLeaf": false
  },
  {
    "name": "b",
    "type": "dir",
    "children": [
      {
        "name": "dir3",
        "type": "dir",
        "children": [
          {
            "name": "item4",
            "type": "item",
            "isLeaf": true
          }
        ],
        "isLeaf": false
      }
    ],
    "isLeaf": false
  },
  {
    "name": "a",
    "type": "item",
    "isLeaf": true
  }
]
```

## Unsupported keys
Note that consul keys are not strictly cascaded. If you use this library you should exclude some special keys such as
- empty string ""
- slash "/"
- key with 2 or more continous slash (like "a//b")


## License
This repository, also known as the following names (hereinafter referred to as "this repo"), is licensed under the GNU General Public License v3.0 (GPL v3.0).
- https://github.com/imagefc/keytree
- github.com/imagefc/keytree
- imagefc/keytree

Any operation of this repo (viewing, forking, cloning, modifying, redistributing, etc.) shall be deemed as your agreement to the following provisions:

1. Any file in this repo is licensed under GPL v3.0, even if it does not contain copyright information.
2. The author has the final right to interpret this repo.
