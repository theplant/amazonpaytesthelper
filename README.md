# Amazon Pay test helper

Amazon Pay test for Golang

## Install

```
go get github.com/theplant/amazonpaytesthelper
```

## Usage

```
package main

import (
	"fmt"

	"github.com/theplant/amazonpaytesthelper"
)

func main() {
	var account = amazonpaytesthelper.AmazonPayTestAccount{
		Email:         "123@123.com",
		EmailPassword: "123456",
	}

	var config = amazonpaytesthelper.AmazonPayConfig{
		MerchantID: "your amazonpay MerchantID",
		ClientID:   "your amazonpay ClientID",
	}

	tocken, id, err := amazonpaytesthelper.AmazonPayTestHelper(config, account)
	if err != nil {
		panic(err)
	}
	fmt.Print(tocken, id)

}

```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
