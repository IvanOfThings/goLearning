package main

// Enviamos los datos del canal de entrada por el primer canal de salida disponible (workers)
func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In {
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}

// Reparte los datos recibidos por los canales de entrada por el primer canal de salida disponible (workers)
func Turnout(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	var data int
	var more bool
	for {
		select {
		case data, more = <-InA:
		case data, more = <-InB:

		case <-Quit:
			close(InA)
			close(InB)
			Fanout(InA, OutA, OutB) //Flush the remanining data
			Fanout(InB, OutA, OutB)
			return
		}
	}
	//...
}

// Reparte los datos recibidos por los canales de entrada por el primer canal de salida disponible (workers)
func Turnout(InA <-chan int, InB <-chan int, OutA, OutB chan int) {
	var data int
	var more bool
	for {
		select {
		case data, more = <-InA:
		case data, more = <-InB:
		}
		if !more {
			//...?maybe hw have not to stop
			return
		}
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}
