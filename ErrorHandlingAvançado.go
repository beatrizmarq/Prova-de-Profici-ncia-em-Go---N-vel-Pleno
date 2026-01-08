package main

import (
	"errors"
	"fmt"
	"regexp"
)

// Erros sentinela para identificação via errors.Is
var (
	ErrCPFInvalidFormat     = errors.New("cpf com formato inválido")
	ErrCPFRepeatedDigits    = errors.New("cpf com todos os dígitos iguais")
	ErrCPFInvalidCheckDigit = errors.New("cpf com dígitos verificadores inválidos")
)

// ValidateCPF valida um CPF brasileiro com ou sem formatação
func ValidateCPF(cpf string) error {
	cleanCPF, err := normalizeCPF(cpf)
	if err != nil {
		return fmt.Errorf("falha ao normalizar cpf: %w", err)
	}

	if hasRepeatedDigits(cleanCPF) {
		return ErrCPFRepeatedDigits
	}

	if !validateCheckDigits(cleanCPF) {
		return ErrCPFInvalidCheckDigit
	}

	return nil
}

// Funções auxiliares
func normalizeCPF(cpf string) (string, error) {
	re := regexp.MustCompile(`\D`)
	clean := re.ReplaceAllString(cpf, "")

	if len(clean) != 11 {
		return "", ErrCPFInvalidFormat
	}

	return clean, nil
}

func hasRepeatedDigits(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func validateCheckDigits(cpf string) bool {
	// Primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	d1 := (sum * 10) % 11
	if d1 == 10 {
		d1 = 0
	}
	if d1 != int(cpf[9]-'0') {
		return false
	}
	// Segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	d2 := (sum * 10) % 11
	if d2 == 10 {
		d2 = 0
	}
	if d2 != int(cpf[10]-'0') {
		return false
	}

	return true
}
