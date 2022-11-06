package create

import "fmt"

func CreateScript(name string) string {
	message := fmt.Sprintf("The script name is %v", name)
	return message
}
