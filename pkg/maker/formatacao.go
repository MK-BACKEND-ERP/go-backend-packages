package maker

import (
	"strings"
	"time"
)

/*
Função para formatar a data conforme o formato especificado

Parâmetros:
- data: Objeto time.Time contendo a data e hora que serão formatadas.
- formato: String contendo o formato desejado para a formatação da data e hora.

Explicação do Map "replacements":

  - "yyyy": Representa o ano com quatro dígitos,
  - "MM": Representa o mês com dois dígitos, como "01" para janeiro.
  - "dd": Representa o dia do mês com dois dígitos, como "02".
  - "H": Representa a hora no formato 24 horas, como "15" para 3 horas da tarde.
  - "k": Similar ao "H", representa a hora no formato 24 horas, como "15".
  - "K": Representa a hora no formato 12 horas, como "3" para 3 horas da tarde.
  - "h": Similar ao "K", representa a hora no formato 12 horas, como "03".
  - "a": Representa "AM" ou "PM" para indicar manhã ou tarde.
  - "mm": Representa os minutos com dois dígitos, como "04".
  - "ss": Representa os segundos com dois dígitos, como "05".
  - "SSS": Representa os milissegundos, como ".000".

Retorna:

  - String contendo a data e hora formatada de acordo com o formato especificado.
*/
func FormatacaoDataHora(data time.Time, formato string) string {
	replacements := map[string]string{
		"yyyy": "2006",
		"MM":   "01",
		"dd":   "02",
		"H":    "15",
		"k":    "15",
		"K":    "3",
		"h":    "03",
		"a":    "PM",
		"mm":   "04",
		"ss":   "05",
		"SSS":  ".000",
	}

	for key, value := range replacements {
		formato = strings.Replace(formato, key, value, -1)
	}

	return data.Format(formato)
}
