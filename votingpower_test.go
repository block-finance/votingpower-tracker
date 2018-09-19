package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/tidwall/gjson"
	"strconv"
	"testing"
	"time"
)

func TestTickers(t *testing.T) {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func TestGJSON(t *testing.T) {
	{
		value := gjson.GetBytes(testJSON, "result.block_height")
		fmt.Println(value)
	}
	{
		value := gjson.GetBytes(testJSON, "result.validators.#.voting_power")

		var totalVotingPower uint64
		for _, v := range value.Array() {
			totalVotingPower += v.Uint()
		}

		fmt.Println("Total voting power:", totalVotingPower)
	}

	const validatorNetworkAddress = "01F78669F9515FD83DF9250F5C0EE143D3DAD65C"
	{
		query := fmt.Sprintf(`result.validators.#[address="%v"].voting_power`, validatorNetworkAddress)
		value := gjson.GetBytes(testJSON, query)
		fmt.Println(value)
	}

}

func TestJSON(t *testing.T) {
	blockHeight, err := jsonparser.GetString(testJSON, "result", "block_height")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Block height: " + blockHeight)

	var totalVotingPower, validatorVotingPower uint64
	const validatorNetworkAddress = "01F78669F9515FD83DF9250F5C0EE143D3DAD65C"

	var handler = func(data []byte, dataType jsonparser.ValueType, offset int, err error) {
		votingPower, _ := jsonparser.GetString(data, "voting_power")

		votingPoweri, _ := strconv.ParseUint(votingPower, 10, 64)
		totalVotingPower += votingPoweri

		address, _ := jsonparser.GetString(data, "address")
		if address == validatorNetworkAddress {
			validatorVotingPower = votingPoweri
			fmt.Println("Found the validator. Voting power: ", validatorVotingPower)
		}
	}

	jsonparser.ArrayEach(testJSON, handler, "result", "validators")
	fmt.Println("Total voting power found: ", totalVotingPower)
}

var testJSON = []byte(`
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "block_height": "214603",
    "validators": [
      {
        "address": "015B9D526E2332DEF04168B5C95C6BE15577CC15",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "2RYn/WYgNvvDb0NLcdJ70bDiuAVH+tLH7VWH2M59i1c="
        },
        "voting_power": "132",
        "accum": "132"
      },
      {
        "address": "01F78669F9515FD83DF9250F5C0EE143D3DAD65C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "1K9/Z+oiK380Obf+LYVN6P7ydFIWCFvNQpT+ToRUURg="
        },
        "voting_power": "868",
        "accum": "13004"
      },
      {
        "address": "02FE4B97A8387F39D1AC19E0C1702B3CC614812F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "RwPRoiY5C0covekqbr3VrQwxWGHioUUIf2+TOq8LIC0="
        },
        "voting_power": "43",
        "accum": "2222"
      },
      {
        "address": "0548261C50222FD710EB5EBDF03A0E6567B21D20",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "8K3clCjVU33BTIpUhdahGmu++WxHj4NUE9krCRkk++s="
        },
        "voting_power": "1383",
        "accum": "-2910"
      },
      {
        "address": "0989D0355A104B67AE0AE028C68FB610690E174B",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "BB4a/Xh5z+dkGCRlF+pSGC3iDOoDrFse/xzQAtmxMF4="
        },
        "voting_power": "365",
        "accum": "-10953"
      },
      {
        "address": "0A692DBFFE9E5DD28E7DECC33CED1AEB4C0D014E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "N3K5kDdfcKJurfaa6s2zfKgtYvz1Pagz7VWi9ZfX8yM="
        },
        "voting_power": "931",
        "accum": "13603"
      },
      {
        "address": "0B1D3A322EE4A5A0FBDDCC1B249DA6B7B8DDF29F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "cVE8I8bbrOcczxsFhvEzDeVczHUPPKXZ/YpOeILus3M="
        },
        "voting_power": "3",
        "accum": "4359"
      },
      {
        "address": "0CAA37F0CEAA8A65F6A8AA39B9161FF4C671B2AF",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "1UyFY3YJja8vrmDc/ZqQFo7ORvLdLK2X2wYeCUBZzJc="
        },
        "voting_power": "61",
        "accum": "-1500"
      },
      {
        "address": "0DDF97111D73DB825138EC15EC27DE5FE8536901",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "H0SIA/BU6Xj8oT5bQkvLpEITN3CqFLbMeBcQ72NZrAE="
        },
        "voting_power": "43",
        "accum": "2222"
      },
      {
        "address": "12CC487B78CEF97C8DB6AB7B7692A45C33527161",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "eNKen3yMRu2SDLEu3V+U2LXsyka3Na7sp99V7AuDP+s="
        },
        "voting_power": "21",
        "accum": "12289"
      },
      {
        "address": "1751786EFC085D5331C61168FFFC8202EA151B4E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "xlO2cnii42KisAn8OcstC/3XV5+I0FlcSbWuyy5MVA8="
        },
        "voting_power": "38",
        "accum": "-4793"
      },
      {
        "address": "1ECD8C5FE1341D74CDC633A245DD56858DE2F9A2",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "lUMCRAeu47BsOhNvCQTQJQeB68z0/VaElC9j5gDt9y8="
        },
        "voting_power": "725",
        "accum": "955"
      },
      {
        "address": "1F3A68CAAD483768AA2BF667DEF10F7104C89928",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "Go9GXHI6SCQo2QKMxkAkgYLhfo3XrVjWLR2nE2AvYyk="
        },
        "voting_power": "30",
        "accum": "13464"
      },
      {
        "address": "28770D17E8530825E7310A0659308654B7EC9682",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "Xi7nIgj4PqVXrpKLfJhcyxyVY1d3HRo72sKKPDmuU78="
        },
        "voting_power": "48",
        "accum": "9492"
      },
      {
        "address": "29B93E4EECD8C591AB76C2D52FD7C43CBEEEC50B",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "GTwvM3OOWuXdwH7suTJIWdYdtwUU1hLAHIXmpdzGfl0="
        },
        "voting_power": "2",
        "accum": "2906"
      },
      {
        "address": "2B409FC3A80631C06CAB4317D06B2EFD566838F7",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "zJDJtbSuhYZlRoobVzkJ5h5uDOv6h2zzRgdGsoGaCGA="
        },
        "voting_power": "3",
        "accum": "4359"
      },
      {
        "address": "2E6AF0D5B1A85E173F5EBEDA93D7E7D97A88D06C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "G3YJxcM2oMCW8R5s/HSqmnudsQjcj+e/vpZBtMyob1w="
        },
        "voting_power": "68",
        "accum": "-6357"
      },
      {
        "address": "30BB6A4DAE2249126862536D7E0CD97AA1A36765",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "s9acFRanTXWN1gkfD99onpxzN3JTMpnlp4g9+4Ev6cs="
        },
        "voting_power": "235",
        "accum": "10554"
      },
      {
        "address": "33412C9A69C60CB0FAA24A5C5B3A906DE24B47C2",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4DEDoU/RsHMPS54GzgwkWnW5zPQfMt9aInFFc3GyfA8="
        },
        "voting_power": "148",
        "accum": "4402"
      },
      {
        "address": "33516EF7B2CF68ECDD0EE4EA5F674C0A8175F095",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "bZSEBDNIOr0xJ/PxaAScJIyG6hqFtryBAMNwghAOTTU="
        },
        "voting_power": "113",
        "accum": "13804"
      },
      {
        "address": "3363E8F97B02ECC00289E72173D827543047ACDA",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "mPnu910hOOa1tAQ7pbOLFDxvllbQUmrbtGjqQrYg1nM="
        },
        "voting_power": "88",
        "accum": "2475"
      },
      {
        "address": "36E9B7FAEC964F97D84C028ED62493E373AD0B38",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ViJzrw9Unk1XyDg4f0UQLFE4zfUaFh23wZZY0f31Qg0="
        },
        "voting_power": "356",
        "accum": "6096"
      },
      {
        "address": "38A8D85D36BEA2E5FCF8A1E778F53628E450EF13",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "gJDxhwAE6GeGCKQeVaNZ5is7+7MFHXtOG0UsnguKdoA="
        },
        "voting_power": "163",
        "accum": "-3684"
      },
      {
        "address": "398610C4CF11C84C89AC6975470972EE75DA17E4",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "5GAxIHzu06l0n3MU8wNQrCcrrnpZR9EKe5qZMsM2h/E="
        },
        "voting_power": "91",
        "accum": "12841"
      },
      {
        "address": "39DFDA0C5DB84B269642188B23253508537798E6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "omAzuJps8KX3/iOC1LjwkMPMH3c6tjfLXwCNWXRBdWw="
        },
        "voting_power": "450",
        "accum": "-7712"
      },
      {
        "address": "3AAD96F8B01F73344079C1D030B7AF5351BF9C11",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "WnV5CjfHHyMrAPbX796iXgnfAwTP/cFL+kzz99bW+kA="
        },
        "voting_power": "155",
        "accum": "-7098"
      },
      {
        "address": "3AC99B708404452366E0953A4BE849226DC86CBC",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ISIM341M+EUYOMpwhn9gyCvRUZ06xomp/+lOwa50meQ="
        },
        "voting_power": "581",
        "accum": "1528"
      },
      {
        "address": "3DCB822A10ED319EC7C4393D8E3FD4A41B15AAF2",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "O95l52FybUIUBkaYeGHCLRNkSI5zAewb4SxzS0LOb50="
        },
        "voting_power": "34",
        "accum": "-10605"
      },
      {
        "address": "3E899BA36A93C7370347BD5FE85909F2ECD93F8C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "M2Iu/A8RQcmoiIEar/iV4wBjVj0I0gcd/nNnvwmnT2g="
        },
        "voting_power": "296",
        "accum": "9049"
      },
      {
        "address": "45427C23D3E0922D097C99597DE76935D4B69724",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "yG0obCbA7Gqv8coqrxcC0MYz59a2LFb2u44vUDcsiGQ="
        },
        "voting_power": "95",
        "accum": "-12355"
      },
      {
        "address": "462B656869DB02A8FD1077E255F8CC7401DC37C6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "PxJbo5FKA6mXtgwclRQVNIjOCQK3Q7WkLQrvM9lYbGI="
        },
        "voting_power": "46",
        "accum": "6581"
      },
      {
        "address": "4944F8BBBB5303E17BC0607DF332320EDBC750EA",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "A6GzeXUM3vsXaDAEYMSDgSKkqn9AoUYjs8empH46MGY="
        },
        "voting_power": "1326",
        "accum": "-2748"
      },
      {
        "address": "4FFD769178F58B29475F33488ED0E3EBAF24D849",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4FESXNnABiMpKtaOGIGczAZuBiMWLBl4YT2ZndV9beM="
        },
        "voting_power": "4",
        "accum": "-16053"
      },
      {
        "address": "5104F488EA6B1E2C33CBDA3C18389C88104DBC72",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "JREG9eji9gfJR2b61qMu8jNZ835ezPY2A4o+zjYMfDg="
        },
        "voting_power": "134",
        "accum": "-15827"
      },
      {
        "address": "570B6A755ED339EAC56BB3390A48069F6AC9449B",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "5XCMLGNTGkMh9/JsUzi/8Xl7Y0cdDe38stX+zvBLXuE="
        },
        "voting_power": "184",
        "accum": "-10642"
      },
      {
        "address": "5A6B67F68591404C5079BEBD7A6B7E046779D3B6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "I2ILaY31gonxYQhlIk4PFUIN+Pk7+9dDTK1C/s+Vcb0="
        },
        "voting_power": "1",
        "accum": "1453"
      },
      {
        "address": "5B9628695CDFC6025857578114E3D3126687EAE1",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "NCwjL8K9R2CLOMYub3MhAkkB0fktT0mZC75TEtcXyQ4="
        },
        "voting_power": "65",
        "accum": "13399"
      },
      {
        "address": "5BDF8EA41B1835FBAD865B181A2A4B23F511B5E4",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "RBuRRNEzA9RA1Wrdi9PPFQJ29/n/bqN9O2tQv9Gq248="
        },
        "voting_power": "655",
        "accum": "-10372"
      },
      {
        "address": "5E8673673E37450F01B64138FBF4B172BDB52D68",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "CuPin/hLM2tK6KvcUlOi/Brm5xNi3zq6UPgpLFPSzt8="
        },
        "voting_power": "246",
        "accum": "-3344"
      },
      {
        "address": "60229E2E273297FE00219E46DFFB8F76844CDEF9",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "oL/QCr7LEOivyTqpGrmwVd1r+hYI2WB5+kSVzpDMxx4="
        },
        "voting_power": "207",
        "accum": "-4"
      },
      {
        "address": "62410D6D1CD98278136E4D0738AA0EA273CDFC54",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "K4kLogLtZxqrYSqRVJfrFm9tUG+Tc3QWXWIewnAgI9w="
        },
        "voting_power": "92",
        "accum": "-13833"
      },
      {
        "address": "64F57F8E796A9B7B590C1EFA4123988C86631D39",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "DsTbM0AgHfhSUKvOGkxudDOY3ojYT6bifhpelqHs8+s="
        },
        "voting_power": "30",
        "accum": "13464"
      },
      {
        "address": "65E49617D972BA6CFED171DF6D6037DBCE64BD3A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "LiI4RvDpOfr++QHDIS2K8zgjGoyzcN8ixmQNbecMtUA="
        },
        "voting_power": "1",
        "accum": "1453"
      },
      {
        "address": "69D99B2C66043ACBEAA8447525C356AFC6408E0C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "8pfpbIxBBiu88hpxS3CeRpv7kClEjl8SwVgckDNBGlE="
        },
        "voting_power": "557",
        "accum": "-5687"
      },
      {
        "address": "69EDADFA4A803E9DBF548C900D4A465E7D154262",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "LRK6flgMsHrBYkIIahegEkzlwC7gQ3dAyw0jcAd1u/k="
        },
        "voting_power": "160",
        "accum": "-8043"
      },
      {
        "address": "6B275FFA0571EEFC4482CD7A6CAE836F5785CB4E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "pggq0WUTlsiA8cBiptRgxlw4WUG2GmXfLQxkbMbCBPk="
        },
        "voting_power": "62",
        "accum": "-47"
      },
      {
        "address": "6D3AD380E398D8592B03AA124BAFED5F125DF863",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ENAVynNXVpj/IdYx9kCPKaPs4bWSxRIHNlmS9QiDuZQ="
        },
        "voting_power": "334",
        "accum": "4261"
      },
      {
        "address": "6F6C93B8BF3495063FEC9B5C771158607F4AE9D6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "S8s6fdAQNQ3bN9SNVAsHB/j8uv1CM1roxeLesL+fh4g="
        },
        "voting_power": "198",
        "accum": "-12836"
      },
      {
        "address": "732CEEF54C374DDC6ADECBFD707AEFD07FEDC143",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "HjSC7VkhKih6xMhudlqfaFE8ZZnP8RKJPv4iqR7RhcE="
        },
        "voting_power": "257",
        "accum": "8064"
      },
      {
        "address": "78C38CD2B41499966683A0F61A9AAE2AA64E7254",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "qNMe+4TjuopyX62Og3wYZUbcJOgXZyrPKgG7U6PgVL0="
        },
        "voting_power": "835",
        "accum": "10400"
      },
      {
        "address": "807CF675BA8479FC4FEC1047215DF07970040F1F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "2p8s/pRBZPjYWKKMlR7AOXypDzDmPo762iXlKpCwtco="
        },
        "voting_power": "94",
        "accum": "-13558"
      },
      {
        "address": "82F68C22CDFE170628672BAD532CD967E2E60136",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "GKRc6rG/r7dbbCi5+rTkTOn++LFpHi2kkAQHpfWunek="
        },
        "voting_power": "3",
        "accum": "4359"
      },
      {
        "address": "83AB321E012C57746689FED3A6064DF331F81007",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "aUViYC2znC55sleHfmsIN9hZ45SbYPbDcYA0gVzglsc="
        },
        "voting_power": "743",
        "accum": "3497"
      },
      {
        "address": "86DDBEEF35D35786A68834B34869174AC314BB74",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ilDFePqWkNRNWUEoWKq42jZC5CD4rglfs7+gCuq6Qc8="
        },
        "voting_power": "818",
        "accum": "-14056"
      },
      {
        "address": "8F06FE1FD042683C426137FBC4BBAA4288924217",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "lB+CgSbTzpHXJHUU15fJIPaPl6XF0nSfdNxBnlKJqTk="
        },
        "voting_power": "2350",
        "accum": "11968"
      },
      {
        "address": "927CC759310F4DF5C8936EF6D666ED3135D609B7",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "tkdOxR2CIOctyc5e9IMGwnMwB420OTN3Rfb6fUMVhr8="
        },
        "voting_power": "1",
        "accum": "1453"
      },
      {
        "address": "95FAFBC1E0579C83D290A063020CE87C5334BD9B",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "UBTju7UZfXLVPPYb1a8gPZ69BeCv2Fho7YVo2EUbxKc="
        },
        "voting_power": "43",
        "accum": "2222"
      },
      {
        "address": "96BDE45EC52606CF5CB59853097871AD60680549",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4JoJuRfaANhdM1x3AWRo1/Cj9DH3VA+fi1SynzknV+w="
        },
        "voting_power": "48",
        "accum": "9492"
      },
      {
        "address": "97DBB3D4B3D6C29287244864448F651DD5B07390",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "XQDVMXja3kFk5Jb47BsqJmzcDsM4lE9+r+f/J3O5Jms="
        },
        "voting_power": "88",
        "accum": "2475"
      },
      {
        "address": "9C29C9B76D27586E4D9AD2AED0E943162645E55C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "WRsXnLz3gf8o4lYYeCZjAXgPU1cdmOOYPdy7aY63iIA="
        },
        "voting_power": "29",
        "accum": "12011"
      },
      {
        "address": "9D93AD02FE1BD7E76B9E58D7EB461088EAD011E5",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "jj0Y/Fy8JSJR3g+PHU6Ce0ecYwHGUVJ4bVyR7WwcyLI="
        },
        "voting_power": "703",
        "accum": "-13967"
      },
      {
        "address": "A07EC59210E1228559EB2C6E40B529CB78263A7E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "LOufg1uwcE3MkcDFGzaqRVhkN5lnnI73STTjD8Jzs+4="
        },
        "voting_power": "26",
        "accum": "7652"
      },
      {
        "address": "A0A6466F4ABEA09C3B55E3709AFDE217A103209F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "Abws3eXrUFAH8LeZJIcECakPL945TTmFsBlXONOUeII="
        },
        "voting_power": "46",
        "accum": "6581"
      },
      {
        "address": "A0ED8A257D7862094FE8015BCE2135D111722DF6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "dvLH3HPQ3YH0oDefmGTPvpwn/gTFhvBAQiYNV4DVs5k="
        },
        "voting_power": "18",
        "accum": "1918"
      },
      {
        "address": "A2EF42C4B7A4D452D9425C07A3F31D7FB0F29CED",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "o5vCLTsctDc5awWHRfN5iVyP7oCrk+vFNf58joCWpBw="
        },
        "voting_power": "9",
        "accum": "-16804"
      },
      {
        "address": "A5868971300F155020634D7BB49C7634F7AACCD3",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "cCoFsZzKZ9SQZbHe4NueVObIezP6ts0tRTZ/aN96dig="
        },
        "voting_power": "48",
        "accum": "9492"
      },
      {
        "address": "AC96444A9083890B7C7B430DB17BF1D3EA5EED04",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "QzxYamaSVNx+Idlb9bCcYsPI4Oa/zQ+rx4MN+XA2xLI="
        },
        "voting_power": "698",
        "accum": "-1075"
      },
      {
        "address": "B051AF0D7327EEAFAF6450DB698B7B92499273AA",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4db00EtiqFCrgYP0C93Hkm9RPC7PLlTZ5pZhmwKEvLQ="
        },
        "voting_power": "1087",
        "accum": "-14107"
      },
      {
        "address": "B197425B3210375F211661DE02EF2AE36580EB0F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "j9be+ddLChInrFz6/820/uYh4WZBzlp61klyJBDy/ZY="
        },
        "voting_power": "119",
        "accum": "-7359"
      },
      {
        "address": "C1134A2FF8F7338E7CDB3FE57DDD2A655DEB1E22",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "L0I4JoDfktbDWe0fCDL/nQlBPkF5mNgqamnM5JKJ1Uc="
        },
        "voting_power": "40",
        "accum": "-2137"
      },
      {
        "address": "C39A8A2B85198DCB8F23AD1F6C198E2BA7AB5D60",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "R/3f7VruxWpu+2hiHlVpplTwoOou5kfQI1k/6/9H/y8="
        },
        "voting_power": "94",
        "accum": "-6919"
      },
      {
        "address": "CAC50BD244845AAEE52A3850A0C363698348100B",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "UQHY8OYNqCwDkAm7bn+X/VqV6p3PHTRqkQGEsqnpdws="
        },
        "voting_power": "33",
        "accum": "-12058"
      },
      {
        "address": "CEB97DB5CB18270247ABDE6CF2F84F9AEE8A2904",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "wdUJrATbIN7C9pgzDzaBZ4WANQzQL+/r+36yXqDdeCQ="
        },
        "voting_power": "3",
        "accum": "4359"
      },
      {
        "address": "D19CB89DB5AD9DEBA4C8BF88AFFDF10C33104E1A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "PLF4LU9I+jHHO/HibSjtkafkL4ZHQFvyC3N/pIG7A7Q="
        },
        "voting_power": "5",
        "accum": "2460"
      },
      {
        "address": "D2134B826A1E25A955905CB79A4D8731CAFA0CCE",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "k3YLQYEN2QMP6XITRsBmgb+pNGhJ5Jbg0bzUW977kK0="
        },
        "voting_power": "27",
        "accum": "9105"
      },
      {
        "address": "D6285C3422C8886445D53757CBE12D2B91B5A9D0",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "NQX4yKpOztKrmgBhGIC5WOALOLOq3LTpbzsN4ZLXGec="
        },
        "voting_power": "293",
        "accum": "4690"
      },
      {
        "address": "D77FB84A729972CD3E757C9BEB61561855832B92",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "UjTvuOew2EaooduJBiYmBWeF5ai0yFJG8uio5YXpJgg="
        },
        "voting_power": "85",
        "accum": "-7901"
      },
      {
        "address": "DB7E6732471D3EA5C90C17A9974D9F88AB45F2FB",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "P9RgE4RMQT/aHap2oICpwpgKeBAwxPUwuU9zIffKFNM="
        },
        "voting_power": "367",
        "accum": "-8047"
      },
      {
        "address": "DBA70FA7E9D55E035AD87B41C4DC0C38511FD09A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "3O/S1+hdkUi5BxAohHHkxytw7S+tukjby46dRr6VhPE="
        },
        "voting_power": "4",
        "accum": "5812"
      },
      {
        "address": "DDA3062E760B84CE8217A1A7550AC1E8F8BC3DD6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ncrdAm41M/iC+6ROlnGvtwYZ2UHioDeQl2Hc6okEDew="
        },
        "voting_power": "1",
        "accum": "1453"
      },
      {
        "address": "DF54B862410BCD9B5C516FBB6C56BDC826ED6FA7",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "FiEl7a96NaaDaVH5mV9EK/6tOUWL5mkkT46ARRKAWIg="
        },
        "voting_power": "1292",
        "accum": "-10690"
      },
      {
        "address": "E161D3FC5A61E381D68CE244FBEC27913930B37D",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "0ZPHmftStU+DSqvnIleNyaFmBQsgWJwPC7vk4uF5u00="
        },
        "voting_power": "743",
        "accum": "-3017"
      },
      {
        "address": "E300E9D484374C1D3E97C1D93CAC3B160A056298",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "F2wCn9rKafNZsYZwoLGkSQIpr3rk86cjYyuhSjsjRaE="
        },
        "voting_power": "77",
        "accum": "-5375"
      },
      {
        "address": "E39B5778E0D4297A46F45CF6940B0A163A430FF8",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "0HqB2x6x5HzeozpHatePECw07x1UcDdSz8kQGNznnA8="
        },
        "voting_power": "1816",
        "accum": "1272"
      },
      {
        "address": "E3B19738F7132F950668504D326C7660F159D170",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "tOEqjO2t51PEgO9Tv0B7qM0yPmy1n5tMa3Beg0tp3ns="
        },
        "voting_power": "246",
        "accum": "-3344"
      },
      {
        "address": "E42F8AEABB685219F7027F54D2B0DBC7846F633A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "eeImG09hOPo1W7j7lKepN/Lx6I9GGHqVBVEKmznxACc="
        },
        "voting_power": "1177",
        "accum": "-3846"
      },
      {
        "address": "E4E6C2E81269DFFA196C17F98240E491479A6008",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ydjx2ea+PVuChrny6X2dluJwyXta+BsNQRsgHXp8fXw="
        },
        "voting_power": "48",
        "accum": "9492"
      },
      {
        "address": "E50ADECD5FD27FA2CD610C07F8AED36A2FC4A6A6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "BaaCxmYHKJ6obIzTCdRtjw1cc8d2mUJcMbLWCjf1aLo="
        },
        "voting_power": "1013",
        "accum": "-1620"
      },
      {
        "address": "EC3401441726C40ECC7E51F744483B05BC6DE9BD",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "3wRufybSUsTMnUeQkP74uJNDRKeM8jBLAS64T0BRfpY="
        },
        "voting_power": "53",
        "accum": "-13124"
      },
      {
        "address": "EC67CD8E0F5ECD9FAE10918EE950D803F9324F1D",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ePxcUeuPx3UiqfCPLioJjYwDgk/E7HGhatxfzl4/vI0="
        },
        "voting_power": "426",
        "accum": "-5810"
      },
      {
        "address": "EC83E2F9859837767A926223B96F93EF12659D95",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "e/yHTwpzu+69G/uXg7swLKYs60Xehf7pBNZjO5RZytk="
        },
        "voting_power": "441",
        "accum": "2641"
      },
      {
        "address": "EC95B111D032181E1EC3B5E8875C5BB8E4AE7140",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "66j9af4xDJSblMLS+mFbp7d8TaFGu0FOo+0MwEYm2lE="
        },
        "voting_power": "359",
        "accum": "10455"
      },
      {
        "address": "F13C1805E8131C9B47E2A46818F57D8A2841E578",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "G5xqaI1M/ckd6k5tLzWkGF/MdGZ4/xmJThP4AjksrFk="
        },
        "voting_power": "8",
        "accum": "11624"
      },
      {
        "address": "F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="
        },
        "voting_power": "349",
        "accum": "-4075"
      },
      {
        "address": "F4268597DE93E3E2CED38845010F8AA34E2A879E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "Epd2FDKZwDybzT38Z7WB0y2jiLn9/2OLzmY3Zu18l6I="
        },
        "voting_power": "1035",
        "accum": "470"
      },
      {
        "address": "F6738260186D33D9C14FC6E7017AFE6BB952A63D",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "2qtEBT+Tc+SD2wJsdrVMHXrBKfvesxtmtSKDK5fXwA0="
        },
        "voting_power": "44",
        "accum": "3675"
      }
    ]
  }
}`)
