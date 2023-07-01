package main

import (
	"fmt"
	"strings"
)

// Structs contendo os tipos a serem usados
type professorASerPago struct {
	staffCode string
	nome      string
	valorHora float64
}
type MesASerPago struct {
	dia   int
	mes   string
	carga float64
}

// Função para verificar se o teacher existe ou não
func verificarTeacher(staffCode string, teachers []professorASerPago) *professorASerPago {
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
func verificarMes() string {
	fmt.Println("Digite o mês trabalhado:")

	mesValido := false
	var inputMes string

	for tentativasMes := 3; tentativasMes > 0 && !mesValido; tentativasMes-- {
		fmt.Scan(&inputMes)
		mesesValidos := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
		for _, m := range mesesValidos {
			if strings.ToLower(inputMes) == strings.ToLower(m) {
				mesValido = true
				break
			}
		}

		if !mesValido {
			fmt.Println("Mês inválido. Digite novamente.")
			fmt.Println("Tentativas restantes:", tentativasMes-1)
		}
	}

	if !mesValido {
		fmt.Println("Número de tentativas excedido. Encerrando programa.")
		return ""
	}

	return inputMes
}

// Verifica as horas inseridas
func verificarNumeroHorasTrabalhadas() float64 {
	fmt.Print("Digite o número de horas trabalhadas: ")

	var numeroHorasTrabalhadas float64
	var tentativasHoras int

	for tentativasHoras < 3 {
		fmt.Scan(&numeroHorasTrabalhadas)
		if numeroHorasTrabalhadas <= 0 && numeroHorasTrabalhadas < 100 {
			return numeroHorasTrabalhadas
		}
		tentativasHoras++
		fmt.Println("Número de horas inválido. Tente novamente.")
		fmt.Print("Digite o número de horas trabalhadas: ")
	}

	fmt.Println("Número de tentativas excedido. Encerrando o programa.")
	return 0
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
		teacher = verificarTeacher(staffCode, teachers)
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

	var mesTrabalhado = verificarMes()
	if mesTrabalhado == "" {
		return
	}

	var numeroHorasTrabalhadas = verificarNumeroHorasTrabalhadas()
	if numeroHorasTrabalhadas == 0 {
		return
	}

	// Realizar cálculos com as informações do professor

	fmt.Println("Cálculos concluídos. Resultado final:")
	fmt.Println("Mês trabalhado:", mesTrabalhado)
	fmt.Println("Número de horas trabalhadas:", numeroHorasTrabalhadas)
	fmt.Println("Total a pagar:", teacher.valorHora*numeroHorasTrabalhadas)
}
