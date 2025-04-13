package main

import (
	"math/rand"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func initMap() {
	// Ініціалізуємо карту
	world = make([][]int32, width)
	for x := range world {
		world[x] = make([]int32, height)
	}
}

func generate_map() {
	// Параметри шуму
	alpha := 2.0  // плавність
	beta := 2.0   // вплив частоти
	nOctaves := 3 // кількість рівнів шуму
	seed := rand.Int63()

	// Створюємо Perlin noise генератор
	p := perlin.NewPerlin(alpha, beta, int32(nOctaves), seed)

	scale := 0.1 // масштаб шуму — впливає на "розмір" біомів

	// Генерація карти
	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			// Отримуємо шум у координаті (x, y)
			noise := p.Noise2D(float64(x)*scale, float64(y)*scale)

			// Нормалізуємо шум від -1..1 до 0..1
			normalized := (noise + 1) / 2

			// Присвоюємо тип біому за рівнем шуму
			switch {
			case normalized < 0.25:
				world[x][y] = 1 // Вода
			case normalized < 0.75:
				world[x][y] = 0 // Трава
			default:
				world[x][y] = 2 // Гори
			}
		}
	}
}

func draw_map() {
	// Малюємо карту
	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			var color rl.Color
			if world[x][y] == 0 {
				color = rl.Green // Трава
			}
			if world[x][y] == 1 {
				color = rl.Blue // Вода
			}
			if world[x][y] == 2 {
				color = rl.Gray // Гори
			}

			// Малюємо прямокутники на екрані
			rl.DrawRectangle(int32(x)*int32(tileSize), int32(y)*int32(tileSize), int32(tileSize), int32(tileSize), color)
		}
	}
}
