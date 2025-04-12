package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var speed int8 = 5
var position rl.Vector2 = rl.Vector2{300, 300}
var tileSize int8 = 20
var WIDTH int32 = 800
var HEIGHT int32 = 600
var width int32 = WIDTH / int32(tileSize)
var height int32 = HEIGHT / int32(tileSize)
var world [][]int32

func player() {
	if rl.IsKeyDown(rl.KeyW) {

		position.Y -= float32(speed)

	}
	if rl.IsKeyDown(rl.KeyS) {
		position.Y += float32(speed)
	}
	if rl.IsKeyDown(rl.KeyA) {
		position.X -= float32(speed)
	}
	if rl.IsKeyDown(rl.KeyD) {
		position.X += float32(speed)
	}

	rl.DrawRectangle(int32(position.X), int32(position.Y), int32(tileSize), int32(tileSize), rl.Black)
}

func initMap() {
	world = make([][]int32, width)
	for x := range world {
		world[x] = make([]int32, height)
	}
}

func generate_map() {

	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			random := rand.Intn(101)

			if random < 2 {
				world[x][y] = 2
			} else if random < 5 {
				world[x][y] = 1
			} else if random < 90 {
				world[x][y] = 0
			}
		}

	}
}

func draw_map() {
	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			var color rl.Color
			if world[x][y] == 0 {
				color = rl.Green
			}
			if world[x][y] == 1 {
				color = rl.Blue
			}
			if world[x][y] == 2 {
				color = rl.Gray
			}

			rl.DrawRectangle(int32(x)*int32(tileSize), int32(y)*int32(tileSize), int32(tileSize), int32(tileSize), color)
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	initMap()
	generate_map()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		draw_map()
		player()
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
