package maker

import (
	"fmt"
	"time"
)

/*
Essa função obtem a data atual e retorna.

Timezones:
Padrão (de acordo com o timezone do servidor)

Ex:

  - America/Sao_Paulo

  - America/Cuiaba
*/
func Hoje(timezone ...string) time.Time {

	if len(timezone) > 0 && timezone[0] != "" {
		var err error

		timezone, err := time.LoadLocation(timezone[0])
		if err != nil {
			fmt.Println("Erro ao carregar o fuso horário:", err)
			return time.Now()
		}

		return time.Now().In(timezone)
	}
	return time.Now()
}

// Retorna o ano como um inteiro de acordo com a data informada.
func Ano(data time.Time) int {
	return data.Year()
}

// Retorna o mês como um inteiro de acordo com a data informada.
func Mes(data time.Time) int {
	return int(data.Month())
}

// Retorna o dia como um inteiro de acordo com a data informada.
func Dia(data time.Time) int {
	return data.Day()
}

// Data retira as horas da data, retornando apenas a data com as horas zeradas.
func Data(data time.Time) time.Time {
	ano, mes, dia := data.Date()
	return time.Date(ano, mes, dia, 0, 0, 0, 0, data.Location())
}
