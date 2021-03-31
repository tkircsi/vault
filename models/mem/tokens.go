package mem

import (
	"time"

	"github.com/tkircsi/vault/models"
	"github.com/tkircsi/vault/services"
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
		data, err := services.Decrypt(v)
		if err != nil {
			return nil, err
		}

		t := models.Token{
			Token: key,
			Value: data,
		}
		return &t, nil
	}
}

func (mv *MemVault) Put(value string, exp time.Duration) (*models.Token, error) {
	cipher, err := services.Encrpyt(value)
	if err != nil {
		return nil, err
	}
	t := models.Token{
		Token: cipher,
		Value: cipher,
	}
	tokens[t.Token] = t.Value
	return &t, nil
}
