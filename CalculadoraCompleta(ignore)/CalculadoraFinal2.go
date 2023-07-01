// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type employee struct {
// 	codigoUnico string
// 	teacherName string
// 	rate        int32
// }

// type empDetails struct {
// 	name string
// 	rate int32
// }

// func confirmID(codigoUnico string) (bool, empDetails) {
// 	var inputTeacher string
// 	fmt.Println("Bem vindo. Digite seu staffcode:")

// 	attempts := 3
// 	for attempts > 0 {
// 		fmt.Scan(&inputTeacher)
// 		if inputTeacher == codigoUnico {
// 			emp := getEmpDetails(codigoUnico)
// 			return true, emp
// 		} else {
// 			fmt.Println("Staffcode errado. Tente novamente.")
// 			attempts--
// 			fmt.Println("Tentativas restantes:", attempts)
// 			if attempts == 0 {
// 				fmt.Println("Número de tentativas excedido. Encerrando programa.")
// 				return false, empDetails{}
// 			}
// 			fmt.Println("Digite seu staffcode novamente:")
// 		}
// 	}
// 	return false, empDetails{}
// }

// func getEmpDetails(codigoUnico string) empDetails {
// 	funcionario := map[string]employee{
// 		"001": {codigoUnico: "001", teacherName: "Gilvane Cezar", rate: 30},
// 		"002": {codigoUnico: "002", teacherName: "Carolina Paiva", rate: 25},
// 		"003": {codigoUnico: "003", teacherName: "Fábio Henrique", rate: 30},
// 	}

// 	emp := funcionario[codigoUnico]
// 	return empDetails{
// 		name: emp.teacherName,
// 		rate: emp.rate,
// 	}
// }

// func calcularPagamento(codigoUnico string, mes string, horasTrabalhadas float64) float64 {
//
// 	for valorTentativas < 3 {
// 		if horasTrabalhadas <= 0 || horasTrabalhadas > 100 {
// 			fmt.Println("Número inválido. Digite novamente:")
// 			valorTentativas++
// 			if valorTentativas == 3 {
// 				fmt.Println("Número de tentativas excedido. Encerrando programa")
// 				break
// 			}
// 			var inputHorasTrabalhadas float64
// 			fmt.Println("Digite o total de horas trabalhadas no mês de", mes, ":")
// 			fmt.Scan(&inputHorasTrabalhadas)
// 			horasTrabalhadas = inputHorasTrabalhadas
// 		} else {
// 			emp := getEmpDetails(codigoUnico)
// 			totalAmount := float64(emp.rate) * horasTrabalhadas
// 			return totalAmount
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	funcionario := []employee{
// 		{codigoUnico: "001", teacherName: "Gilvane Cezar", rate: 30},
// 		{codigoUnico: "002", teacherName: "Carolina Paiva", rate: 25},
// 		{codigoUnico: "003", teacherName: "Fábio Henrique", rate: 30},
// 	}

// 	fmt.Println("Digite o mês trabalhado:")

// 	mesValido := false
// 	for tentativasMes := 3; tentativasMes > 0 && !mesValido; tentativasMes-- {
// 		var inputMes string
// 		fmt.Scan(&inputMes)
// 		mesesValidos := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
// 		for _, resultadoRange := range mesesValidos {
// 			if strings.ToLower(inputMes) == strings.ToLower(resultadoRange) {
// 				mesValido = true
// 				break
// 			}
// 		}

// 		if !mesValido {
// 			fmt.Println("Mês inválido. Digite novamente.")
// 			fmt.Println("Tentativas restantes:", tentativasMes-1)
// 		}
// 	}

// 	if !mesValido {
// 		fmt.Println("Número de tentativas excedido. Encerrando programa.")
// 		return
// 	}

// 	valid, emp := confirmID("001")
// 	if valid {
// 		var inputHorasTrabalhadas float64
// 		fmt.Println("Digite o total de horas trabalhadas no mês de", inputMes, ":")
// 		fmt.Scan(&inputHorasTrabalhadas)
// 		totalPagar := calcularPagamento(emp.codigoUnico, inputMes, inputHorasTrabalhadas)
// 		fmt.Println("Sr.(a)", emp.name, "seu saldo a receber no mês de", inputMes, "é de: R$", totalPagar)

// 	} else {
// 		fmt.Println("Acesso negado.")
// 	}
// }