package ports

import "awesomeProject/pkg/domain"

type PlayerService interface {
	Create(player domain.Player) (id interface{}, err error)
}

type PlayerRepository interface {
	Insert(player domain.Player) (id interface{}, err error)
}
