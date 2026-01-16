package main

import (
	"fmt"
	"time"
)

func main() {
	const tasksNumber = 5
	const workersNumber = 3
	
	// Crea canales con buffer para tareas y resultados.
	tasks := make(chan int, tasksNumber)
	results := make(chan int, tasksNumber)

	// Lanza los workers que consumiran tareas en paralelo.
	for worker_id := 1; worker_id <= workersNumber; worker_id++ {
		go worker(worker_id, tasks, results)
	}

	// Envia las tareas al canal para que los workers las procesen.
	for task := 1; task <= tasksNumber; task++ {
		tasks <- task
	}	
	// Cierra el canal de tareas para indicar que no hay mas trabajo.
	close(tasks)
	
	// Espera todos los resultados antes de terminar el programa.
	for result := 1; result <= tasksNumber; result++ {
		<- results
	}	
}

func worker (id int, tasks <- chan int, results chan <- int) {
	// Recorre el canal hasta que se cierre y procesa cada tarea.
	for task := range tasks {
		fmt.Printf("worker %d: Processing task %d\n", id, task)
		// Simula trabajo costoso.
		time.Sleep(time.Second)
		fmt.Printf("worker %d: Working on task %d\n", id, task)
		// Devuelve el resultado procesado.
		results <- task * 2
	}
}
