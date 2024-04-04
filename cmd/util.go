package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func formatJson(jsonObject []byte) string {
	buffer := new(bytes.Buffer)
	err := json.Indent(buffer, jsonObject, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(buffer.Bytes())
}

func checkUrl(url string) {
	if url == "" {
		fmt.Println("This is not a valid URL")
		os.Exit(1)
	}
}

func printOutput(response *http.Response) {
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error Status Code: ", response.StatusCode)
	}
	contentType := strings.Split(response.Header.Get("Content-Type"), ";")[0]
	if contentType == "application/json" {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			println(err)
		}
		fmt.Println(formatJson(body))
	} else {
		reader := bufio.NewReader(response.Body)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(str)
		}
	}
}
