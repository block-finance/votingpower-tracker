package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	url                       string
	validatorNetworkAddress   string
	ticker                    *time.Ticker
	validatorVotingPowerGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		// Namespace: "our_company",
		// Subsystem: "blob_storage",
		Name: "validator_voting_power",
		Help: "Voting power of configured validator",
	})
	totalVotingPowerGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		// Namespace: "our_company",
		// Subsystem: "blob_storage",
		Name: "total_voting_power",
		Help: "Total network voting power",
	})
)

func startDataRetrieval() {
	url = viper.GetString("url")
	validatorNetworkAddress = viper.GetString("validatorNetworkAddress")
	freq := time.Duration(viper.GetDuration("queryFrequency") * time.Second)

	fmt.Println("Query frequency:", freq)
	ticker = time.NewTicker(freq)

	go func() {
		for {
			retrieveValidatorData()

			<-ticker.C
		}
	}()
}

func retrieveValidatorData() {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	totalVotingPowerGauge.Set(float64(totalVotingPower))
	validatorVotingPowerGauge.Set(float64(validatorVotingPower))
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

func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(validatorVotingPowerGauge)
	prometheus.MustRegister(totalVotingPowerGauge)
}

func main() {
	readConfig()

	startDataRetrieval()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
