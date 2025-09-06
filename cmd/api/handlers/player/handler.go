package player

import (
	"awesomeProject/pkg/ports"
)

type Handler struct {
	PlayerService ports.PlayerService
}
