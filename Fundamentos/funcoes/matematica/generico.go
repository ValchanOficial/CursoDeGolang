package matematica

//Calculo : executa calculo baseado na função enviada
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador : multiplica um número x pelo número y
func Multiplicador(x, y int) int {
	return x * y
}

//Divisor : divide um número x pelo número y
func Divisor(x, y int) (resultado int) {
	resultado = x / y
	return
}

//DivisorComResto : divide um número x pelo número e retorna o resultado e o resto da divisão
func DivisorComResto(x, y int) (resultado, resto int) {
	resultado = x / y
	resto = x % y
	return
}
