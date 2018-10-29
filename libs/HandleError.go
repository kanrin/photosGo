package libs

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Errorf("[Error]: %s", err)
	os.Exit(-1)
}