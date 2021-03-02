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
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"` // X velocity of the snake (-1, 0, 1)
	VelY int `json:"velY"` // Y velocity of the snake (-1, 0, 1)
}

type velocity struct {
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

type gameStates struct {
	state
	Ticks []velocity `json:"ticks"`
}
