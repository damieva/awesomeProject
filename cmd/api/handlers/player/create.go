package player

import (
	"awesomeProject/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreatePlayer(c *gin.Context) {
	// El handler ha de encargarse de:
	//   - Traducir el request
	//   - Validación (casos de uso)
	//   - Consumir el servicio
	//   - Traducir el response
	var playerCreateParms domain.Player
	if err := c.BindJSON(&playerCreateParms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// ========================

	// Definimos arriba una vble de tipo interfaz ports.PlayerService y aquí usamos el metodo create que no
	insertedId, err := h.PlayerService.Create(playerCreateParms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	// =======================
	// Envía una respuesta HTTP con el código 200 y el player_id
	c.JSON(http.StatusOK, gin.H{"player_id": insertedId})
}
