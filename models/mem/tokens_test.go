package mem

import "testing"

var (
	data = `{
    "name": "Tibcsi",
    "age": 45
}`
	mv = &MemVault{}
)

func TestPut(t *testing.T) {
	_, err := mv.Put(data, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	token, _ := mv.Put(data, 0)
	gt, err := mv.Get(token.Token)
	if err != nil {
		t.Fatal(err)
	}
	if gt.Token != token.Token {
		t.Fatalf("token mismatch. expect: %+v, got: %+v", token, gt)
	}
}
