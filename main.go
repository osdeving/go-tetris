package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
	blockSize     = 30
	cols          = 10
	rows          = 20
)

var tetrominos = map[string][][]int{
	"I": {
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	"O": {
		{1, 1},
		{1, 1},
	},
	"T": {
		{0, 1, 0},
		{1, 1, 1},
		{0, 0, 0},
	},
	"L": {
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	},
}

type Tetromino struct {
	shape [][]int
	x, y  int
}

// Peça atual
var currentTetromino Tetromino

// Inicia uma peça aleatória no topo do tabuleiro
func spawnTetromino() {
	currentTetromino = Tetromino{
		shape: tetrominos["T"],
		x:     cols/2 - 1,
		y:     0,
	}
}

// Move a peça no eixo X
func moveTetromino(dx int) {
	currentTetromino.x += dx
}

// Desenha a peça na tela
func drawTetromino(renderer *sdl.Renderer) {
	for y, row := range currentTetromino.shape {
		for x, cell := range row {
			if cell != 0 {
				renderer.SetDrawColor(0, 255, 0, 255) // Cor da peça
				renderer.FillRect(&sdl.Rect{
					X: int32((currentTetromino.x + x) * blockSize),
					Y: int32((currentTetromino.y + y) * blockSize),
					W: blockSize,
					H: blockSize,
				})
			}
		}
	}
}

var board [rows][cols]int

func drawBoard(renderer *sdl.Renderer) {
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if board[y][x] == 0 {
				renderer.SetDrawColor(30, 30, 30, 255)
			} else {
				renderer.SetDrawColor(255, 255, 255, 255)
			}

			renderer.FillRect(&sdl.Rect{
				X: int32(x * blockSize),
				Y: int32(y * blockSize),
				W: blockSize,
				H: blockSize,
			})
		}
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "SDL Init error: %s\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Go Testris",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Window creation error: %s\n", err)
		os.Exit(1)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Renderer creation error: %s\n", err)
		os.Exit(1)
	}
	defer renderer.Destroy()

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_LEFT:
					moveTetromino(-1) // Move para a esquerda
				case sdl.K_RIGHT:
					moveTetromino(1) // Move para a direita
				case sdl.K_DOWN:
					currentTetromino.y++ // Move para baixo
				}
			}
		}

		// Limpa a tela
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		// Desenha o tabuleiro
		drawBoard(renderer)

		// Desenha a peça atual
		drawTetromino(renderer)

		// Atualiza a tela
		renderer.Present()
		sdl.Delay(16) // ~60 FPS

	}
}
