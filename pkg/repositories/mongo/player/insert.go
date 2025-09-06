package player

import (
	"awesomeProject/pkg/domain"
	"context"
	"fmt"
	"log"
)

func (r Repository) Insert(player domain.Player) (id interface{}, err error) {
	// Inicializamos un handler para trabajar con la collection players
	collection := r.client.Database("go-l").Collection("players")
	// Insertamos un documento en la collection. El contexto (como bien inicializamos arriba) indica el tiempo y cancelación de la operación.
	// El insertResult nos devolverá el ID que Mongo asignará al documento
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting player %w", err)
	}

	return insertResult.InsertedID, nil
}
