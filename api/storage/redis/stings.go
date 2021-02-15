package redis

import "fmt"

func sRune(id int32) string {
	return fmt.Sprintf("charm.%d.rune", id)
}
func sGod(id int32) string {
	return fmt.Sprintf("charm.%d.god", id)
}
func sPower(id int32) string {
	return fmt.Sprintf("charm.%d.power", id)
}
