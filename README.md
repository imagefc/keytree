# keytree
Transfer consul key slice to cascade key tree.

- Before transfer (call consul key-value api to get keys)
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

- After transfer

KeyMap (unordered but have unique key index)
```
{
  "a": null,
  "a/": {               // use slash to separate item and dir
    "dir1/": {},
    "dir2/": {          // add missing parent(s)
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

KeyList (ordered but have no key index)
```
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
- key with 2 or more continous slash "a//b"


## License
This repository, also known as the following names (hereinafter referred to as "this repo"), is licensed under the GNU General Public License v3.0 (GPL v3.0).
- https://github.com/imagefc/keytree
- github.com/imagefc/keytree
- imagefc/keytree

Any operation of this repo (viewing, forking, cloning, modifying, redistributing, etc.) shall be deemed as your agreement to the following provisions:

1. Any file in this repo is licensed under GPL v3.0, even if it does not contain copyright information.
2. The author has the final right to interpret this repo.
