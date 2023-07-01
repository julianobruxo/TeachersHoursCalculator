/*package main

import (
	"fmt"
	"strings"
)

type professorASerPago struct {
	staffCode string
	nome      string
	valorHora float64
}

func main() {
	teachers := []professorASerPago{
		{staffCode: "001", nome: "Gilvane Cezar", valorHora: 30},
		{staffCode: "002", nome: "Carolina Paiva", valorHora: 25},
		{staffCode: "003", nome: "Fábio Henrique", valorHora: 30},
	}

	var staffCode string
	fmt.Print("Digite seu staffcode: ")
	fmt.Scan(&staffCode)

	var teacher *professorASerPago
	var tentativas int

	for tentativas < 3 {
		teacher = findTeacher(staffCode, teachers)
		if teacher != nil {
			break
		}
		tentativas++
		fmt.Println("Professor não encontrado. Tente novamente.")
		fmt.Print("Digite seu staffcode: ")
		fmt.Scan(&staffCode)
	}

	if tentativas == 3 {
		fmt.Println("Usuário não encontrado. Encerrando o programa.")
		return
	}

	if !confirmarIdentidade(teacher) {
		fmt.Println("Identidade não confirmada. Encerrando o programa.")
		return
	}

	fmt.Println("Professor encontrado:")
	fmt.Println("Nome:", teacher.nome)
	fmt.Println("Valor hora:", teacher.valorHora)

	var mesTrabalhado string
	fmt.Print("Digite o mês trabalhado: ")
	fmt.Scan(&mesTrabalhado)

	var numeroHorasTrabalhadas float64
	fmt.Print("Digite o número de horas trabalhadas: ")
	fmt.Scan(&numeroHorasTrabalhadas)

	// Realizar cálculos com as informações do professor

	fmt.Println("Cálculos concluídos. Resultado final:")
	fmt.Println("Mês trabalhado:", mesTrabalhado)
	fmt.Println("Número de horas trabalhadas:", numeroHorasTrabalhadas)
	fmt.Println("Total a pagar:", teacher.valorHora*numeroHorasTrabalhadas)
}

func findTeacher(staffCode string, teachers []professorASerPago) *professorASerPago {
	for i := range teachers {
		if teachers[i].staffCode == staffCode {
			return &teachers[i]
		}
	}
	return nil
}

func confirmarIdentidade(teacher *professorASerPago) bool {
	var confirmacao string
	fmt.Printf("Os dados do professor %s foram encontrados. Confirma a identidade? (S/N): ", teacher.nome)
	fmt.Scan(&confirmacao)
	confirmacao = strings.ToUpper(confirmacao)

	return confirmacao == "S"
}
/*