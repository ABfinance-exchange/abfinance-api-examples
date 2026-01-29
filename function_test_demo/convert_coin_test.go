package main

import (
	"context"
	"errors"
	"fmt"
	abfinance "github.com/abfinance-exchange/abfinance.go.api"
)

func main() {
	client := abfinance.NewABFinanceHttpClient("xxx", "xxx", abfinance.WithBaseURL(abfinance.TESTNET))
	GetConvertCoinList(client)
	response, err := RequestConvertQuote(client)
	if err != nil {
		fmt.Println("Error requesting convert quote:", err)
		return
	}
	quoteTxId, err := GetQuoteTxId(response)
	if err != nil {
		fmt.Println("Error getting quote Tx ID:", err)
		return
	}
	ConfirmConvertQuote(client, quoteTxId)
	GetConvertStatus(client, quoteTxId)
	GetConvertHistory(client)
}

func GetQuoteTxId(response *abfinance.ServerResponse) (quoteTxId string, err error) {
	result, ok := response.Result.(map[string]interface{})
	if !ok {
		errMsg := "Conversion of response.Result to map[string]interface{} failed"
		fmt.Println(errMsg)
		return "", errors.New(errMsg)
	}

	// Now you can safely retrieve the quoteTxId
	quoteTxId, ok = result["quoteTxId"].(string)
	if !ok {
		errMsg := "Retrieval of quoteTxId failed"
		fmt.Println(errMsg)
		return "", errors.New(errMsg)
	}

	fmt.Println("quoteTxId: ", quoteTxId)
	return quoteTxId, nil
}

func GetConvertCoinList(client *abfinance.Client) {
	params := map[string]interface{}{"coin": "USDT", "accountType": "eb_convert_uta"}
	response, err := client.NewUtaABFinanceServiceWithParams(params).GetConvertCoinList(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(abfinance.PrettyPrint(response))
}

func GetConvertStatus(client *abfinance.Client, quoteTxId string) {
	params := map[string]interface{}{"quoteTxId": quoteTxId, "accountType": "eb_convert_uta"}
	response, err := client.NewUtaABFinanceServiceWithParams(params).GetConvertStatus(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(abfinance.PrettyPrint(response))
}

func GetConvertHistory(client *abfinance.Client) {
	params := map[string]interface{}{"accountType": "eb_convert_uta"}
	response, err := client.NewUtaABFinanceServiceWithParams(params).GetConvertHistory(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(abfinance.PrettyPrint(response))
}

func RequestConvertQuote(client *abfinance.Client) (response *abfinance.ServerResponse, err error) {
	params := map[string]interface{}{"fromCoin": "BTC", "toCoin": "ETH", "requestCoin": "BTC", "requestAmount": "1", "accountType": "eb_convert_uta"}
	return client.NewUtaABFinanceServiceWithParams(params).RequestConvertQuote(context.Background())
}

func ConfirmConvertQuote(client *abfinance.Client, quoteTxId string) {
	params := map[string]interface{}{"quoteTxId": quoteTxId}
	response, err := client.NewUtaABFinanceServiceWithParams(params).ConfirmConvertQuote(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(abfinance.PrettyPrint(response))
}
