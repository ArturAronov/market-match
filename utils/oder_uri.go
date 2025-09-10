package utils

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Order struct {
	TranactionType    int
	TransactionMethod int
	OrderType         int
	Ticker            string
	Quantity          int
	Price             int
	OrderDate         int
	GoodUntil         int
	TraderId          string
	ClientOrderId     string
}

var URI_LEN = 84
var ORDER_BYTE_LEN = 63

func ParseOrderUri(uri *string) (*Order, error) {
	*uri = strings.Replace(*uri, "/", "", 1)
	if len(*uri) != URI_LEN {
		return nil, fmt.Errorf("ParseOrderUri: Wrong URI length recieved - expected: %d, recieved: %d", URI_LEN, len(*uri))
	}

	orderBytes, orderBytesErr := base64.URLEncoding.DecodeString(*uri)
	if orderBytesErr != nil {
		return nil, orderBytesErr
	}

	if len(orderBytes) != ORDER_BYTE_LEN {
		return nil, fmt.Errorf("ParseOrderUri: Wrong byte length recieved - expected: %d, recieved: %d", ORDER_BYTE_LEN, len(orderBytes))
	}

	tranactionTypeByte := orderBytes[0]
	transactionMethodByte := orderBytes[1]
	orderTypeByte := orderBytes[2]
	tickerBytes := orderBytes[3:7]
	quantityBytes := orderBytes[7:11]
	priceBytes := orderBytes[11:15]
	orderDateBytes := orderBytes[15:23]
	goodUntilBytes := orderBytes[23:31]
	traderIdBytes := orderBytes[31:47]
	clientOrderIdBytes := orderBytes[47:63]

	order := Order{}

	if tranactionTypeByte < 1 || tranactionTypeByte > 2 {
		return nil, fmt.Errorf("ParseOrderUri: Wrong transaction type recieved: %v", tranactionTypeByte)
	} else {
		order.TranactionType = int(tranactionTypeByte)
	}

	if transactionMethodByte < 1 || transactionMethodByte > 3 {
		return nil, fmt.Errorf("ParseOrderUri: Wrong transaction method recieved: %v", transactionMethodByte)
	} else {
		order.TransactionMethod = int(transactionMethodByte)
	}

	if orderTypeByte < 1 || orderTypeByte > 2 {
		return nil, fmt.Errorf("ParseOrderUri: Wrong order type recieved: %v", orderTypeByte)
	} else {
		order.OrderType = int(orderTypeByte)
	}

	for _, byte := range tickerBytes {
		if byte > 0 && (byte < 65 || byte > 90) {
			return nil, fmt.Errorf("ParseOrderUri: Wrong ticker recieved: %s", string(tickerBytes))
		}
	}

	order.Ticker = string(tickerBytes)

	return nil, nil
}
