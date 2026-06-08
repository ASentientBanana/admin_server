package dto

type PositionPair struct {
	ID       uint `json:"id"`
	Position int  `json:"position"`
}

type UpdateProjectPositionsDto struct {
	Positions []PositionPair `json:"positions"`
}
