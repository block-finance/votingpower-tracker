package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO Make configurable
const (
	queryFrequency          = 1 // Number of seconds between data retrieval.
	url                     = "https://gaia.validator.network/validators"
	validatorNetworkAddress = "01F78669F9515FD83DF9250F5C0EE143D3DAD65C"
)

var (
	ticker = time.NewTicker(queryFrequency * time.Second)
)

func startDataRetrieval() {
	// TODO Make configurable
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			fmt.Println(retrieveValidatorData())
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

func main() {
	fmt.Println(" --- Voting power tracker")

	// response, _ := http.Get(url)
	// buf, _ := ioutil.ReadAll(response.Body)
	// defer response.Body.Close()

	// fmt.Println(string(buf))

	startDataRetrieval()
	awaitTermination()
}
