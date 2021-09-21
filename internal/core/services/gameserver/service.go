package gamesrv

import "errors"

type service struct {
	gamesRepository ports.GamesRepository
	uidGen          uidgen.UIDGen
}

func New(gamesRepository ports.GamesRepository, uidGen uidgen.UIDGen) *service {
	return &service{
		gamesRepository: gamesRepository,
		uidGen:          uidGen,
	}
}

//func (srv *service) Get(id string) (domain.Game, error) {}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New("the number of bombs is invalid")
	}

	game := domain.NewGame(srv.uidGen.New(), name, size, bombs)

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("create game into repository has failed")
	}

	return game, nil
}
