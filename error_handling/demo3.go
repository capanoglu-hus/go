package errorhandling

import "fmt"

type BorderException struct {
	parameter int
	message   string
}

func (b BorderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)

}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return " ", &BorderException{tahmin, "sınırların dışında kaldı"}
	}

	return "bildiniz", nil
}
