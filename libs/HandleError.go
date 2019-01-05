package libs

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Errorf("[Error]: %s", err.Error())
	os.Exit(-1)
}

func HandleErrorF(err error)  {
	if err != nil {
		fmt.Printf("[ERROR]: %s", err.Error())
	}
}