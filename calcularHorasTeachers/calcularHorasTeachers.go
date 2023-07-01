package main

import (
	"bufio"
	"fmt"
	"os"
)

// Declaração dos Structs contendo os usuários e elementos a serem manipulados

type Teacher struct {
	nome                   string
	valorHora              float64
	numeroHorasTrabalhadas float64
	mesTrabalhado          string
	totalAPagar            float64
}

func main() {
	teacher := Teacher{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type teacher's name: ")
	nome, _ := reader.ReadString('\n')
	teacher.nome = nome

	fmt.Print("Type teacher's rate: ")
	fmt.Scan(&teacher.valorHora)

	fmt.Printf("Type the month to be paid: ")
	fmt.Scan(&teacher.mesTrabalhado)

	fmt.Printf("Type amount of worked hours during specified period: ")
	fmt.Scan(&teacher.numeroHorasTrabalhadas)

	teacher.totalAPagar = teacher.numeroHorasTrabalhadas * float64(teacher.valorHora)

	fmt.Println("Nome:", teacher.nome)
	fmt.Println("Valor hora mensal:", teacher.valorHora)
	fmt.Printf("No mês de %s trabalhou %.2f horas.\n", teacher.mesTrabalhado, teacher.numeroHorasTrabalhadas)
	fmt.Printf("Teacher %s, seu saldo a receber é R$ %.2f reais.\n", teacher.nome, teacher.totalAPagar)
}
