package param

import "app/app/entities"

type TeamCreateParameter struct {
	Team    *entities.TeamEntity     `json:"team" binding:"required"`
	Players []*entities.PlayerEntity `json:"players" binding:"required"`
}
