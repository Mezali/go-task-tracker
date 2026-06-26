package commands

import "os"

func InitDb(JsonFileName string) {
	os.WriteFile(JsonFileName, []byte("[]"), 0644)
}
