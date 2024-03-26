package templates

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func HTMLfromString(str string, name string, path string) {
	newFile, err := os.Create(fmt.Sprintf("%s.html", name))
	if err != nil {
		log.Fatal("Error creating html file", err)
	}
	defer newFile.Close()

	io.Copy(newFile, strings.NewReader(str))

}
