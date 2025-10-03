package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// O slice de inteiros é a estrutura de dados principal.
var numeros []int

// Variável para facilitar a leitura de entrada do usuário.
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("========================================")
	fmt.Println("  GERENCIADOR DE NÚMEROS")
	fmt.Println("========================================")

	for {
		exibirMenu()
		opcao, err := lerOpcao()

		if err != nil {
			fmt.Println("\nEntrada inválida. Por favor, digite um número inteiro.")
			continue
		}

		fmt.Println("----------------------------------------")

		switch opcao {
		case 1:
			adicionarNumero()
		case 2:
			listarNumeros()
		case 3:
			removerPorIndice()
		case 4:
			mostrarEstatisticas()
		case 5:
			realizarDivisaoSegura()
		case 6:
			limparLista()
		case 7:
			ordenarLista() // Funcionalidade Bônus
		case 8:
			exibirPares() // Funcionalidade Bônus
		case 9:
			exportarParaArquivo() // Funcionalidade Bônus
		case 0:
			fmt.Println("Encerrando a aplicação.")
			return
		default:
			fmt.Println("Opção inválida. Escolha uma opção de 0 a 9.")
		}

		fmt.Println("----------------------------------------")
	}
}

// exibeMenu mostra as opções disponíveis, incluindo os bônus.
func exibirMenu() {
	fmt.Println("\n-------------------- MENU --------------------")
	fmt.Println(" 1) Adicionar número")
	fmt.Println(" 2) Listar números")
	fmt.Println(" 3) Remover por índice")
	fmt.Println(" 4) Estatísticas (Mínimo, Máximo, Média)")
	fmt.Println(" 5) Divisão segura")
	fmt.Println(" 6) Limpar lista")
	fmt.Println("--- Funcionalidades Bônus ---")
	fmt.Println(" 7) Ordenar lista (Crescente/Decrescente)")
	fmt.Println(" 8) Exibir apenas números pares")
	fmt.Println(" 9) Exportar lista para arquivo texto")
	fmt.Println("-----------------------------")
	fmt.Println(" 0) Sair")
	fmt.Print(">> Escolha uma opção: ")
}

// ===============================================
// FUNÇÕES DE UTILIDADE E INPUT
// ===============================================

// lerOpcao lê a entrada do usuário e retorna a opção como inteiro.
func lerOpcao() (int, error) {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	opcao, err := strconv.Atoi(input)
	return opcao, err
}

// lerNumeroInt solicita e lê um número inteiro do usuário.
func lerNumeroInt(prompt string) (int, error) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	return num, err
}

// ===============================================
// FUNCIONALIDADES OBRIGATÓRIAS (1.1)
// ===============================================

// Funcionalidade 1: Adiciona um número ao slice, impedindo negativos (BÔNUS 1).
func adicionarNumero() {
	num, err := lerNumeroInt("Digite o número inteiro a adicionar: ")
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, digite um número inteiro.")
		return
	}

	// BÔNUS 1: Impedir que números negativos sejam adicionados.
	if num < 0 {
		fmt.Println("Não é permitido adicionar números negativos (Funcionalidade Bônus).")
		return
	}

	numeros = append(numeros, num)
	fmt.Printf("Número %d adicionado à lista.\n", num)
}

// Funcionalidade 2: Lista todos os números armazenados.
func listarNumeros() {
	if len(numeros) == 0 {
		fmt.Println("A lista de números está vazia.")
		return
	}
	fmt.Printf("Números na lista (Total: %d):\n", len(numeros))
	for i, num := range numeros {
		fmt.Printf("  [%d]: %d\n", i, num)
	}
}

// Funcionalidade 3: Remove um número pelo índice.
func removerPorIndice() {
	if len(numeros) == 0 {
		fmt.Println("A lista está vazia. Nada a remover.")
		return
	}

	indice, err := lerNumeroInt("Digite o índice do número a remover: ")
	if err != nil || indice < 0 || indice >= len(numeros) {
		fmt.Println("Índice inválido ou fora do intervalo da lista.")
		return
	}

	valorRemovido := numeros[indice]
	// Remove o elemento do slice: Junta os pedaços antes e depois do índice
	numeros = append(numeros[:indice], numeros[indice+1:]...)

	fmt.Printf("O número %d (índice %d) foi removido.\n", valorRemovido, indice)
}

// Funcionalidade 4: Calcula e retorna mínimo, máximo e média (MÚLTIPLOS RETORNOS).
func calcularEstatisticas(lista []int) (min int, max int, media float64) {
	if len(lista) == 0 {
		return 0, 0, 0.0
	}

	min = lista[0]
	max = lista[0]
	soma := 0

	for _, num := range lista {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		soma += num
	}

	media = float64(soma) / float64(len(lista))
	return // Retorno nomeado (min, max, media)
}

func mostrarEstatisticas() {
	if len(numeros) == 0 {
		fmt.Println("A lista está vazia. Não é possível calcular estatísticas.")
		return
	}

	min, max, media := calcularEstatisticas(numeros)

	fmt.Println("Estatísticas da Lista:")
	fmt.Printf("  Mínimo: %d\n", min)
	fmt.Printf("  Máximo: %d\n", max)
	fmt.Printf("  Média: %.2f\n", media)
}

// Funcionalidade 5: Divisão segura com TRATAMENTO DE ERRO (if err != nil).
func divisaoSegura(a int, b int) (resultado float64, err error) {
	if b == 0 {
		// Retorna erro se o divisor for zero, seguindo o padrão Go.
		return 0, errors.New("divisor não pode ser zero")
	}

	return float64(a) / float64(b), nil
}

func realizarDivisaoSegura() {
	a, errA := lerNumeroInt("Digite o dividendo (numerador): ")
	if errA != nil {
		fmt.Println("Dividendo inválido.")
		return
	}

	b, errB := lerNumeroInt("Digite o divisor (denominador): ")
	if errB != nil {
		fmt.Println("Divisor inválido.")
		return
	}

	resultado, err := divisaoSegura(a, b)

	// Tratamento de erro padrão Go: if err != nil
	if err != nil {
		fmt.Printf("Não foi possível realizar a divisão: %s\n", err.Error())
		return
	}

	fmt.Printf("O resultado da divisão de %d por %d é %.2f\n", a, b, resultado)
}

// Funcionalidade 6: Esvazia o slice de números.
func limparLista() {
	if len(numeros) == 0 {
		fmt.Println("A lista já está vazia.")
		return
	}

	numeros = nil // Limpa o slice, liberando a memória alocada.
	fmt.Println("A lista de números foi esvaziada.")
}

// ===============================================
// FUNCIONALIDADES BÔNUS (1.3)
// ===============================================

// BÔNUS 2: Implementar uma opção de ordenação crescente e decrescente.
func ordenarLista() {
	if len(numeros) < 2 {
		fmt.Println("A lista deve ter pelo menos dois números para ser ordenada.")
		return
	}

	fmt.Print("Escolha o tipo de ordenação (C=Crescente / D=Decrescente): ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(strings.TrimSpace(input))

	switch input {
	case "C":
		sort.Ints(numeros) // Usa o pacote sort para ordenação crescente
		fmt.Println("Lista ordenada em modo Crescente.")
	case "D":
		// Usa sort.Slice para definir uma ordenação decrescente personalizada
		sort.Slice(numeros, func(i, j int) bool {
			return numeros[i] > numeros[j] // i > j para decrescente
		})
		fmt.Println("Lista ordenada em modo Decrescente.")
	default:
		fmt.Println("Opção de ordenação inválida. Retornando ao menu.")
		return
	}
	listarNumeros()
}

// BÔNUS 3: Criar uma opção para exibir apenas números pares.
func exibirPares() {
	if len(numeros) == 0 {
		fmt.Println("A lista está vazia.")
		return
	}

	fmt.Println("Números Pares na Lista:")
	encontrouPar := false
	for _, num := range numeros {
		if num%2 == 0 { // Usa o operador módulo (%)
			fmt.Printf("  %d\n", num)
			encontrouPar = true
		}
	}

	if !encontrouPar {
		fmt.Println("Não há números pares na lista.")
	}
}

// BÔNUS 4: Implementar exportação da lista para um arquivo texto.
func exportarParaArquivo() {
	if len(numeros) == 0 {
		fmt.Println("A lista está vazia. Nada para exportar.")
		return
	}

	nomeArquivo := "lista_numeros.txt"

	// Cria o arquivo. Se ele existir, seu conteúdo é truncado (apagado).
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Printf("Não foi possível criar o arquivo: %s\n", err.Error())
		return
	}
	defer arquivo.Close() // Garante que o arquivo será fechado ao sair da função

	for _, num := range numeros {
		// Converte o número para string, adiciona uma quebra de linha ('\n') e escreve no arquivo.
		linha := strconv.Itoa(num) + "\n"
		_, err := arquivo.WriteString(linha)
		if err != nil {
			fmt.Printf("Erro ao escrever o número %d no arquivo. Exportação interrompida.\n", num)
			return
		}
	}

	fmt.Printf("Lista exportada para o arquivo %s com %d números.\n", nomeArquivo, len(numeros))
}
