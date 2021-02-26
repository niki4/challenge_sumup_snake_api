# SumUp Backend Challenge

Hello and welcome to the SumUp backend technical challenge. This document will
outline the challenge and provide you with the information needed in order to
complete your task.

This challenge can be performed in the language of your choice unless previously
specified during the interview process. 

## Overview

The concept of [Snake](https://en.wikipedia.org/wiki/Snake_(video_game_genre))
(the video game) has been around since 1976–when the video game Blockade
invented the genre. However, it saw a new popularity and ultimately rose to
prominence after the launch of Nokia's 6110 mobile phone in 1997. Its popularity
exploded, leading to many spin-offs and coming preloaded on the majority of
Nokia's subsequent phones. This ultimately cemented Snake's place in the arcade
hall-of-fame, where it resides to this day.

In order to continue the legacy of this arcade marvel, we would like to
implement the snake game with a global leaderboard on our SumUp Solo card
reader.

In this challenge, we would like you to build a simple validator for the game.
This validator should have a simple HTTP API allowing clients to implement their
own snake games, and allowing our hardware team to implement it on the Solo
reader.

Your API will be used to prevent cheating, and to make sure that a given set of
moves leads to a valid outcome; thus maintaining trust in our highly competitive
global leaderboard.

We will intentionally be simplifying the game of snake in order to make it more
achievable during this challenge. Here are some prior considerations.

- Your snake will never be longer than length 1. It will *not* grow after eating
    a fruit.
- You will only have 1 single fruit at a time on the game board.
- If the snake hits the edge of the game bounds, the game is over.
- The snake will always start at position `(0, 0)`, with a velocity of `(1, 0)`.
    - i.e. at `x=0`, `y=0` moving to the *right*.

The client will make an initial call to the validator to start a new game with
the provided parameters. It will then collect a series of "ticks", containing
moves that will lead the snake to the position of the fruit. A tick is a single
movement of the snake in a given (x, y) direction. Once the moves lead the snake
to the fruit, the game state and ticks are sent back to your validator in order
to confirm that the move-set is valid. If they are valid, the validator will
increment the game's score, generate a new position for the fruit and send back
the new game state, and so on and so forth.

## Requirements

The validator should expose a web server with two routes. The server itself
should be stateless. We also appreciate thorough testing.

- `GET /new?w=[width]&h=[height]`:
    - Query Params:
        - `w: int`
        - `h: int`
    - Response:
        - *200*: JSON marshalled `state`, with randomly generated fruit
            position.
        - *400*: Invalid request.
        - *405*: Invalid method.
        - *500*: Internal server error.

<br />

- `POST /validate`: 
    - JSON Body:
    ```go
    {
        ...state // all fields in state type
        ticks: [
            { velX: int, velY: int }, // velocity at that tick
            ...
        ]
    }
    ```
    - Response:
        - *200*: Valid state & ticks. Returns JSON marshalled `state` with new
            randomly generated fruit position and a score incremented by 1.
        - *400*: Invalid request.
        - *404*: Fruit not found, the ticks do not lead the snake to the fruit
            position.
        - *405*: Invalid method.
        - *418*: Game is over, snake went out of bounds or made an invalid move. 
        - *500*: Internal server error.

#### Types
```go
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
```

## Hints

#### Velocity

The snake maintains its movement in a given direction until it is given an
instruction to change it. The format of this velocity is `(x, y)`, where `x` is
the velocity in the `x` direction and `y` is the velocity in the `y` direction.
Velocity values can be one of `-1`, moving left/up, `0`, no movement for axis,
`1`, down/right. 

#### Movement

In the rules of Snake, only certain movements are possible at each time. For
example, when the snake is moving to the right, velocity `(1, 0)`, it is
impossible for it to do an immediate 180-degree turn to velocity `(-1, 0)`.
Calls to `/validate` with these sorts of invalid moves should return the invalid
response. 

## Submission

Complete the problem in this repository on a branch called `submission`. When
complete, open a pull-request against the master branch in order to receive
feedback in the form of comments.

We have included a small acceptance testing binary at path `/run-test*` to help
you validate your solution. In order for us to be sure that your solution passes
our acceptance tests, when run, the tests will output a code. The binary takes a
flag `-port` for you to select which port your validation server is running on.
The default port is `8080`.

**Note:** Make sure you run the binary that is applicable for your machine's OS
and architecture.

e.g. If you're using a Mac with an Intel CPU.
```bash
$ ./run-test_darwin_amd64 -port 8081
```

or Windows:
```bash
$ ./run-test_windows_amd64 -port 8081
```

If the tests pass, you will see something similar to:

```bash
$ ./run-test_darwin_amd64
PASS

== PLEASE INCLUDE THIS WITH YOUR SUBMISSION. ==
Well done! Your code is: [your_code_here] 
== PLEASE INCLUDE THIS WITH YOUR SUBMISSION. ==
```

**Please include the acceptance test verification code in your pull-request
description**.

## What we're looking for...

When we assess your submission, we are looking at the following criteria.

- **Simplicity**: Did you solve the problem in a convoluted manner, or did you
    keep it simple? This is a big one for us. We *greatly* value simple
    solutions over complex ones.

- **Maintainability**: Is the code going to be easy to maintain in the future?
    Will it act predictably to changes?

- **Consistency**: Is the code consistent within itself?

- **Testability**: How easy is your code to test? Does the structure lend itself
    to being easily tested?

## Extra for experts

1. Consider that the state returned may be maliciously manipulated by the
   client. How can we prevent this? Add this to your implementation.

2. Highlight a deficiency with the solution. What would a solution be?
