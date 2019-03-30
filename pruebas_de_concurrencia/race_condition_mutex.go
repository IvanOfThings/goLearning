package main

import (
	"fmt"
	"sync"
)

// Al lanzar el programa con go run -race ... (go nos informa de los puntos donde parece haber problemas de acceso a memoria compartida)

var wg sync.WaitGroup
var exluye sync.Mutex
var cuentaciclos int

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go ciclointerno(&wg, &exluye)
	}
	wg.Wait()
	fmt.Println(" El coneto final es: ", cuentaciclos)

}

func ciclointerno(espera *sync.WaitGroup, exum *sync.Mutex) {
	for i := 1; i <= 10; i++ {
		exum.Lock()
		x := cuentaciclos
		x++
		cuentaciclos = x
		exum.Unlock()
	}
	espera.Done()
}
