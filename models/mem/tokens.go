package mem

import (
	"time"

	"github.com/rs/xid"
	"github.com/tkircsi/vault/models"
)

type MemVault struct{}

var tokens = map[string]string{}

func NewMemVault() *MemVault {
	return &MemVault{}
}

func (mv *MemVault) Get(key string) (*models.Token, error) {
	if v, ok := tokens[key]; !ok {
		return nil, models.ErrNoRecord
	} else {
		t := models.Token{
			Token: key,
			Value: v,
		}
		return &t, nil
	}
}

func (mv *MemVault) Put(value string, exp time.Duration) (*models.Token, error) {
	t := models.Token{
		Token: xid.New().String(),
		Value: value,
	}
	tokens[t.Token] = t.Value
	return &t, nil
}
