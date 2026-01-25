package main

import (
	"fmt"
	"matrix-game/internal/board"
	"matrix-game/internal/engine"
	"matrix-game/internal/entity"
)

func main() {
	var size int
	fmt.Print("Tamaño de la matriz: ")
	fmt.Scan(&size)
	b := board.New(size)

	moveCh := make(chan entity.Move, 100)
	doneCh := make(chan entity.Done, 1)

	var numMuros int
	fmt.Print("Número de muros: ")
	fmt.Scan(&numMuros)
	for i := 0; i < numMuros; i++ {
		var x, y int
		fmt.Printf("Posición muro %d (x y): ", i+1)
		fmt.Scan(&x, &y)
		b.Set(x, y, '#')
	}

	var numPhones int
	fmt.Print("Número de teléfonos: ")
	fmt.Scan(&numPhones)
	for i := 0; i < numPhones; i++ {
		var x, y int
		fmt.Printf("Posición teléfono %d (x y): ", i+1)
		fmt.Scan(&x, &y)
		b.Set(x, y, 'T')
	}

	// Configurar protagonista
	var px, py int
	fmt.Print("Posición del protagonista (x y): ")
	fmt.Scan(&px, &py)
	prota := &entity.Protagonist{X: px, Y: py}
	b.Set(px, py, prota.Icon())

	// Configurar agentes
	var numAgents int
	fmt.Print("Número de agentes: ")
	fmt.Scan(&numAgents)
	agents := make([]*entity.Agent, numAgents)
	for i := 0; i < numAgents; i++ {
		var ax, ay int
		fmt.Printf("Posición agente %d (x y): ", i+1)
		fmt.Scan(&ax, &ay)
		a := &entity.Agent{X: ax, Y: ay, ID: i}
		b.Set(ax, ay, a.Icon())
		agents[i] = a
	}

	// Función para dar posición del protagonista
	target := func() (int, int) {
		return prota.X, prota.Y
	}

	// Lanzar goroutines con doneCh
	go prota.Run(b, moveCh, doneCh)
	for _, a := range agents {
		go a.Run(b, moveCh, doneCh, target)
	}

	// Ejecutar motor con fin de juego
	engine.Run(b, moveCh, doneCh)
}
