package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	file, err := os.Open("base_teste[802].txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	db, err := sql.Open("postgres", "postgres://neoway2023:neoway2023@host.docker.internal:5432/clients?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS clients (id SERIAL PRIMARY KEY,
															  CPF varchar(255),
															  PRIVATE varchar(255),
															  INCOMPLETO varchar(255),
															  DATA_DA_ULTIMA_COMPRA varchar(255),
															  TICKET_MEDIO varchar(255),
															  TICKET_DA_ULTIMA_COMPRA varchar(255),
															  LOJA_MAIS_FREQUENTE varchar(255),
															  LOJA_DA_ULTIMA_COMPRA varchar(255),
															  CPF_VALIDO varchar(255),
															  CNPJ_LOJA_MAIS_FREQUENTE_VALIDO varchar(255),
															  CNPJ_LOJA_DA_ULTIMA_COMPRA_VALIDO varchar(255))`); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if count > 0 {

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			line = strings.TrimSpace(line)

			fields := strings.Split(line, " ")

			fields_values := []string{}
			for _, r := range fields {
				if r != "" {
					fields_values = append(fields_values, r)
				}
			}

			cpf_valido := validateCPF(fields_values[0])
			cnpj_valido_loja_frequente := validarCNPJ(fields_values[6])
			cnpj_valido_ultima_loja := validarCNPJ(fields_values[7])

			if _, err := tx.Exec(`INSERT INTO clients (CPF, 
													   PRIVATE,
													   INCOMPLETO,
													   DATA_DA_ULTIMA_COMPRA,
													   TICKET_MEDIO,
													   TICKET_DA_ULTIMA_COMPRA,
													   LOJA_MAIS_FREQUENTE,
													   LOJA_DA_ULTIMA_COMPRA,
													   CPF_VALIDO,
													   CNPJ_LOJA_MAIS_FREQUENTE_VALIDO,
													   CNPJ_LOJA_DA_ULTIMA_COMPRA_VALIDO) 
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
				fields_values[0],
				fields_values[1],
				fields_values[2],
				fields_values[3],
				fields_values[4],
				fields_values[5],
				fields_values[6],
				fields_values[7],
				cpf_valido,
				cnpj_valido_loja_frequente,
				cnpj_valido_ultima_loja); err != nil {
				log.Fatal(err)
			}
		}

		count++
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dados inseridos com sucesso!")
}

func validateCPF(cpf string) string {
	re := regexp.MustCompile("\\D")
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return "Inválido"
	}
	if cpf == "NULL" {
		return "Não preenchido"
	}

	cpfArray := [11]int{}
	for i := 0; i < 11; i++ {
		cpfArray[i], _ = strconv.Atoi(string(cpf[i]))
	}

	var soma int
	var peso int = 10
	for i := 0; i < 9; i++ {
		soma += cpfArray[i] * peso
		peso--
	}
	digito1 := 11 - (soma % 11)
	if digito1 > 9 {
		digito1 = 0
	}

	soma = 0
	peso = 11
	for i := 0; i < 10; i++ {
		soma += cpfArray[i] * peso
		peso--
	}
	digito2 := 11 - (soma % 11)
	if digito2 > 9 {
		digito2 = 0
	}

	if cpfArray[9] == digito1 && cpfArray[10] == digito2 {
		return "Válido"
	}

	return "Inválido"
}

func validarCNPJ(cnpj string) string {
	re := regexp.MustCompile("\\D")
	cnpj = re.ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return "Inválido"
	}

	if cnpj == "NULL" {
		return "Não preenchido"
	}

	cnpjArray := [14]int{}
	for i := 0; i < 14; i++ {
		cnpjArray[i], _ = strconv.Atoi(string(cnpj[i]))
	}

	var soma int
	var peso int = 2
	for i := 11; i >= 0; i-- {
		soma += cnpjArray[i] * peso
		peso++
		if peso > 9 {
			peso = 2
		}
	}
	digito1 := 11 - (soma % 11)
	if digito1 > 9 {
		digito1 = 0
	}

	soma = 0
	peso = 2
	for i := 12; i >= 0; i-- {
		soma += cnpjArray[i] * peso
		peso++
		if peso > 9 {
			peso = 2
		}
	}
	digito2 := 11 - (soma % 11)
	if digito2 > 9 {
		digito2 = 0
	}

	if cnpjArray[12] == digito1 && cnpjArray[13] == digito2 {
		return "Válido"
	}

	return "Inválido"
}
