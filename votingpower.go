package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var (
	validatorNetworkAddress string
	ticker                  *time.Ticker
	network                 string

	validatorVotingPowerGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		// Namespace: "our_company",
		// Subsystem: "blob_storage",
		Name: "gaia_validator_voting_power",
		Help: "Voting power of configured validator",
	}, []string{"chainID"})
	totalVotingPowerGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		// Namespace: "our_company",
		// Subsystem: "blob_storage",
		Name: "gaia_total_voting_power",
		Help: "Total network voting power",
	}, []string{"chainID"})
	individualVotingPowerGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		// Namespace: "our_company",
		// Subsystem: "blob_storage",
		Name: "gaia_validators_voting_power",
		Help: "Voting power of each validator",
	}, []string{"address", "chainID"})
)

func startDataRetrieval() {
	baseURL := viper.GetString("baseURL")
	validatorsURLEndpoint := baseURL + "/validators"
	statusURLEndpoint := baseURL + "/status"

	validatorNetworkAddress = viper.GetString("validatorNetworkAddress")
	freq := time.Duration(viper.GetDuration("queryFrequency") * time.Second)
	fmt.Println("Query frequency:", freq)

	ticker = time.NewTicker(freq)

	// Retrieve chain-id
	chainID, err := getChainID(statusURLEndpoint)
	if err != nil {
		panic(err)
	}

	fmt.Println("Chain-ID:", chainID)

	go func() {
		for {
			retrieveValidatorData(validatorsURLEndpoint, chainID)

			<-ticker.C
		}
	}()
}

func getChainID(statusURLEndpoint string) (string, error) {
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}

	response, err := client.Get(statusURLEndpoint)
	if err != nil {
		return "", err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	value := gjson.GetBytes(responseBody, "result.node_info.network")
	return value.Str, nil
}

func retrieveValidatorData(validatorsURLEndpoint, chainID string) {
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}

	response, err := client.Get(validatorsURLEndpoint)
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

	labels := prometheus.Labels{"chainID": chainID}
	totalVotingPowerGauge.With(labels).Set(float64(totalVotingPower))
	validatorVotingPowerGauge.With(labels).Set(float64(validatorVotingPower))

	{
		value := gjson.GetBytes(responseBody, "result.validators")
		for _, v := range value.Array() {
			address := v.Get("address").String()
			votingPower := v.Get("voting_power").Uint()

			labels["address"] = address
			individualVotingPowerGauge.With(labels).Set(float64(votingPower))
		}
	}
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// TODO Ensure all configuration keys exist
	viper.SetDefault("queryFrequency", 30)
}

func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(validatorVotingPowerGauge)
	prometheus.MustRegister(totalVotingPowerGauge)
	prometheus.MustRegister(individualVotingPowerGauge)
}

func main() {
	readConfig()

	startDataRetrieval()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("localhost:8080", nil)
}
