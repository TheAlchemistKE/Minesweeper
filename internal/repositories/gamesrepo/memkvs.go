package gamesrepo

import (
	"encoding/json"
	"errors"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (gamedomain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := gamedomain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return gamedomain.Game{}, errors.New("fail to get value from kvs")
		}

		return game, nil
	}

	return gamedomain.Game{}, errors.New("game not found in kvs")
}
