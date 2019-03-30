package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func jugador(nombre string, juego chan int) {
	defer wg.Done()

	for {
		pelota, existe := <-juego

		if !existe {
			fmt.Printf("¡%s ganó!\n", nombre)
			return
		}

		suerte := rand.Intn(100)
		if suerte%13 == 0 {
			fmt.Printf("%s perdió en la ronda: %d\n", nombre, pelota)
			close(juego)
			return
		}

		fmt.Printf("%s pegó en la ronda %d\n", nombre, pelota)
		pelota++
		juego <- pelota

	}
}

func main() {
	juego := make(chan int)

	wg.Add(2)

	go jugador("Nadal", juego)
	go jugador("Ivan", juego)

	//Lanzamos la pelota
	juego <- 1

	//Esperamos a que termine
	wg.Wait()
}
