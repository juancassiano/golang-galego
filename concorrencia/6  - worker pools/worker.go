package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}
