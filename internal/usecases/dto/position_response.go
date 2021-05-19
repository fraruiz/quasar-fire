package dto

type PositionResponse struct {
	X float64 `json:"x" binding:"required"`
	Y float64 `json:"y" binding:"required"`
}

func NewPositionResponse(x, y float64) PositionResponse {
	return PositionResponse{
		X: x,
		Y: y,
	}
}
