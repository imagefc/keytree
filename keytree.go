package keytree

import "sort"

// KeyMap is unordered, but has unique key index.
type KeyMap map[string]KeyMap

func BuildMap(keys []string) KeyMap {
	keyMap := KeyMap{}
	for _, k := range keys {
		keyLength := len(k)
		m := keyMap
		for pStart, pEnd := 0, 0; pStart < keyLength; pEnd++ {
			if k[pEnd] == '/' {
				name := k[pStart : pEnd+1]
				if _, ok := m[name]; !ok {
					// add the key with suffix '/' to represent a "directory"
					m[name] = KeyMap{}
				}
				m = m[name]
				pStart = pEnd + 1
				continue
			}
			if pEnd == keyLength-1 {
				name := k[pStart : pEnd+1]
				if _, ok := m[name]; !ok {
					// add the key without suffix '/' to represent an "item"
					m[name] = nil
				}
				pStart = pEnd + 1
				continue
			}
		}
	}
	return keyMap
}

// KeyList is ordered, but does not have unique key index.
type KeyList []*Key

type Key struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Children KeyList `json:"children,omitempty"`

	// IsLeaf is used in el-tree rendering.
	// Add other fields if needed.
	IsLeaf bool `json:"isLeaf"`
}

// Less function sets "dir"s before "item"s
func (k KeyList) Less(i, j int) bool {
	if k[i].Type < k[j].Type {
		return true
	}
	if k[i].Type == k[j].Type && k[i].Name < k[j].Name {
		return true
	}
	return false
}

func (k KeyList) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k KeyList) Len() int {
	return len(k)
}

func BuildKeyListFromMap(m KeyMap) KeyList {
	keyList := make(KeyList, 0, len(m))
	for k, v := range m {
		var name, t string
		leaf := k[len(k)-1] != '/'
		if leaf {
			name, t = k, "item"
		} else {
			name, t = k[:len(k)-1], "dir"
		}
		key := &Key{
			Name:   name,
			Type:   t,
			IsLeaf: leaf,
		}
		if len(v) > 0 {
			key.Children = BuildKeyListFromMap(v)
		}
		keyList = append(keyList, key)
	}
	sort.Sort(keyList)
	return keyList
}
