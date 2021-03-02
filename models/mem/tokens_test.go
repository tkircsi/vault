package mem

import "testing"

var (
	data = `{ "id": 12 }`
	mv   = &MemVault{}
)

func TestPut(t *testing.T) {
	token, err := mv.Put(data, 0)
	if err != nil {
		t.Fatal(err)
	}
	if token.Value != data {
		t.Fatal("token mismatch")
	}
}

func TestGet(t *testing.T) {
	token, _ := mv.Put(data, 0)
	gt, err := mv.Get(token.Token)
	if err != nil {
		t.Fatal(err)
	}
	if (gt.Token != token.Token) || (gt.Value != token.Value) {
		t.Fatal("token mismatch")
	}
}
