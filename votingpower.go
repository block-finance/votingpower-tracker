package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	url                     string
	validatorNetworkAddress string
	ticker                  *time.Ticker
)

func startDataRetrieval() {
	url = viper.GetString("url")
	validatorNetworkAddress = viper.GetString("validatorNetworkAddress")
	freq := time.Duration(viper.GetDuration("queryFrequency") * time.Second)

	fmt.Println("Query frequency:", freq)
	ticker = time.NewTicker(freq)

	go func() {
		for {
			fmt.Println(retrieveValidatorData())

			<-ticker.C
		}
	}()
}

func retrieveValidatorData() (uint64, uint64) {
	response, _ := http.Get(url)
	responseBody, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var totalVotingPower, validatorVotingPower uint64

	{
		value := gjson.GetBytes(responseBody, "result.validators.#.voting_power")
		for _, v := range value.Array() {
			totalVotingPower += v.Uint()
		}
	}

	{
		query := fmt.Sprintf(`result.validators.#[address="%v"].voting_power`, validatorNetworkAddress)
		value := gjson.GetBytes(responseBody, query)
		validatorVotingPower = value.Uint()
	}

	return validatorVotingPower, totalVotingPower
}

// Keep the program open until killed by user.
// https://gobyexample.com/signals
func awaitTermination() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	ticker.Stop()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// TODO Ensure all configuration keys exist
	viper.SetDefault("queryFrequency", 30)
}

func main() {
	readConfig()

	startDataRetrieval()
	awaitTermination()
}
