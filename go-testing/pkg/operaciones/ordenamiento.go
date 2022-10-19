package operaciones

import (
	"errors"
	"sort"
)

func Ordenar(numeros []int) ([]int, error) {
	if len(numeros) > 0 {
		sort.Slice(numeros, func(i, j int) bool {
			return numeros[i] < numeros[j]
		})
		return numeros, nil
	}
	return numeros, errors.New("el slice no puede estar vacio")

}
