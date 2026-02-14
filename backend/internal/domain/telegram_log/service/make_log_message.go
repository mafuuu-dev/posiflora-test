package tgl_service

import "fmt"

func MakeLogMessage(number string, total int64, customerName string) string {
	return fmt.Sprintf(`Новый заказ %s на сумму %d ₽, клиент %s`, number, total, customerName)
}
