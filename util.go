package gocaptcha

import (
	"io/ioutil"
	"fmt"
	"encoding/base64"
)

func File2base64str() {
	bytes, err := ioutil.ReadFile("/file/path/to/encode")
	if err != nil {
		panic(err)
	}
	str := string(base64.StdEncoding.EncodeToString(bytes))
	fmt.Println(str)
}
