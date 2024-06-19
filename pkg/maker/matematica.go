package maker

import (
	"fmt"
	"strconv"
)

func ArredondarCasasDecimais(numero interface{}, quantidade int) (float64, error) {
	qtdDecimais := strconv.Itoa(quantidade)

	num := fmt.Sprintf("%."+qtdDecimais+"f", numero)

	resultado, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}
	return resultado, nil
}
