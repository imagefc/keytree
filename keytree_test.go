package keytree_test

import (
	"encoding/json"
	"testing"

	. "github.com/imagefc/keytree"
)

func TestBuildTree(t *testing.T) {
	keys := []string{
		"a",
		"a/",           // item and dir with same name
		"a/dir1/",      // empty dir
		"a/dir2/item3", // item without parent dir
		"a/item1",
		"a/item2",
		"a1/",
		"b/",
		"b/dir3/",
		"b/dir3/item4",
	}
	m := BuildMap(keys)
	t.Log("---keys in key map---")
	// logKeyMap(t, m)
	b, _ := json.MarshalIndent(m, "", "  ")
	t.Log(string(b))
	l := BuildKeyListFromMap(m)
	t.Log("---keys in key list---")
	// logKeyList(t, l)
	b, _ = json.MarshalIndent(l, "", "  ")
	t.Log(string(b))
}

func logKeyMap(t *testing.T, m KeyMap) {
	for k, v := range m {
		t.Log("name:", k)
		logKeyMap(t, v)
	}
}

func logKeyList(t *testing.T, l KeyList) {
	for _, key := range l {
		t.Logf("name: %s, type: %s, isLeaf: %t\n", key.Name, key.Type, key.IsLeaf)
		logKeyList(t, key.Children)
	}
}
