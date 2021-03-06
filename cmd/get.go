package cmd
import "github.com/spf13/cobra"

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTodos()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

type Todos struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
    IsImportant bool `json:"isImportant"`
}

func getTodos () {
	url := "localhost:8080/api/todos"
	responseBytes := getTodosData(url)
	data := Todos{}

	if err := json.Unmarshal(responseBytes, &data); err != nil {
	    log.Printf("Something went wrong", err)
	}

	fmt.Println(string(title.Title))
}

func getTodosData (BaseAPI string) []byte {
    http.NewRequest(
        http.MethodGet,
        baseAPI,
        nil,
    )
    if err != nil {
        log.Printf("Can not make the request")
    }

    request.Header.Add("Accept", "application/json")

    response, err := http.DefaultClient.Do(request)

    responseBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Println("Can not make the request", err)
    }

    return responseBytes
}