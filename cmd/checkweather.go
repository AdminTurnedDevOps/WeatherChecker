package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// checkweatherCmd represents the checkweather command
var Apikey string
var Location string

var checkweatherCmd = &cobra.Command{
	Use:   "checkweather",
	Short: "A brief description of your command",
	Long:  `Checks the weather forecast based on your location`,
	Run: func(cmd *cobra.Command, args []string) {

		response, err := http.Get("http://api.weatherstack.com/current?access_key=" + Apikey + "&query=" + Location + "&units=f")

		if err != nil {
			log.Println(err)
		}

		data, _ := ioutil.ReadAll(response.Body)

		_, js := strconv.Atoi(string(data))
		jsonOutput, _ := json.Marshal(js)

		fmt.Println(strconv.Atoi(string(jsonOutput)))
	},
}

func init() {
	rootCmd.AddCommand(checkweatherCmd)

	checkweatherCmd.PersistentFlags().StringVarP(&Apikey, "apikey", "a", "", "inserts an API key for Weatherstack at runtime")
	checkweatherCmd.PersistentFlags().StringVarP(&Location, "location", "l", "", "zip code that you want to check the weather for")
}
