package cmd
import "github.com/spf13/cobra"

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: "A tool that helps you to receive your todos",
	Run: func(cmd *cobra.Command, args []string) {
		getTodos()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getTasks () {
    resp, err := http.Get("http://localhost:8080/api/todos")
      if err != nil {
        log.Fatalln(err)
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        log.Fatalln(err)
      }

      log.Println(string(body))
}

type Todos struct {
    ID string `json:"id"`
    Title string `json:"title"`
    IsImportant bool `json:"isImportant"`
    Completed bool `json:"completed"`
}

func getTodos () {
	url := "http://localhost:8080/api/todos"
	responseBytes := getTodosData(url)
	var data []Todos

	if err := json.Unmarshal(responseBytes, &data); err != nil {
	    log.Printf("Something went wrong", err)
	}

	fmt.Println(data)
}

func getTodosData (BaseAPI string) []byte {
    request, err := http.NewRequest(
        http.MethodGet,
        BaseAPI,
        nil,
    )

    request.Header.Add("Accept", "application/json")

    response, err := http.DefaultClient.Do(request)

    responseBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Println("Can not make the request", err)
    }

    return responseBytes
}