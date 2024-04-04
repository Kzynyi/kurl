package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "GET request",
	Long:  "Sends a HTTP GET request",
	Run:   runGet,
}

func runGet(cmd *cobra.Command, args []string) {
	url := getStringFlag(cmd, "url")
	checkUrl(url)
	response := executeGet(url)
	printOutput(response)
}

func getStringFlag(cmd *cobra.Command, name string) string {
	url, err := cmd.Flags().GetString(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return url
}

func executeGet(url string) *http.Response {
	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res
}

func init() {
	getCmd.Flags().StringP("url", "u", "", "URL destination of the request")
	rootCmd.AddCommand(getCmd)
}
