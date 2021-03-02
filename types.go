package main

type state struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  fruit  `json:"fruit"`
	Snake  snake  `json:"snake"`
}

type fruit struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type snake struct {
	X    int `json:"x"`    // horizontal pos
	Y    int `json:"y"`    // vertical pos
	VelX int `json:"velX"` // X velocity of the snake (-1, 0, 1) where -1 is left, 1 is right
	VelY int `json:"velY"` // Y velocity of the snake (-1, 0, 1) where -1 is up, 1 is down
}

type velocity struct {
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

type gameStates struct {
	state
	Ticks []velocity `json:"ticks"`
}
