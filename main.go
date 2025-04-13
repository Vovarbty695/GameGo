package main

import (
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

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "Game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Ініціалізація карти
	initMap()

	// Генерація карти
	generate_map()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		// Малюємо карту
		draw_map()

		// Малюємо гравця
		player()

		// FPS
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
