package handler

type Error struct {
	Message string `json:"message" binding:"required"`
}
