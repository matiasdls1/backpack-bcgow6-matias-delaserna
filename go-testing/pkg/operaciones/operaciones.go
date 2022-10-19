package operaciones

import "fmt"

func Add(num1, num2 int) (int, error) {
	if num1 == 0 {
		return 0, fmt.Errorf("num1 no puede ser: %d", num1)
		//panic("num1 no puede ser: 0")
	}
	return num1 + num2, nil
}

func Sub(num1, num2 int) (int, error) {
	if num1 < num2 {
		return 0, fmt.Errorf("num1 debe ser mayor que num2")
	}
	return num1 - num2, nil
}

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("El denomimindador no puede ser 0")
	}
	return num / den, nil
}
