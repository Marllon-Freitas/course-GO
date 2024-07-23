// post com http client
package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"key":"value"}`)) // precisar passar o json pelo buffer
	res, err := c.Post("http://google.com", "aplication/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}
