package extensions

import (
	"fmt"
	"log"
	"name/enums"
)

func Logger(color, content string, v ...any ) {
	log.Printf(color + fmt.Sprintf(content, v...) + enums.Reset)
}