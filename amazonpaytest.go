package amazonpaytesthelper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/sclevine/agouti"
)

var (
	driver         *agouti.WebDriver
	page           *agouti.Page
	err            error
	requestTimeout = 3 * time.Second
)

func mux(config AmazonPayConfig) http.Handler {

	m := http.NewServeMux()

	m.HandleFunc("/amazon_pay_button", func(w http.ResponseWriter, r *http.Request) {
		MID := config.MerchantID
		CID := config.ClientID
		fmt.Fprintf(w, amazonPayButtonHTML, MID, CID)
	})

	return m
}

func AmazonPayTestHelper(config AmazonPayConfig, account AmazonPayTestAccount) (token string, amazonOrderReferenceId string, err error) {

	go func() {
		err = http.ListenAndServe(":50203", mux(config))
		if err != nil {
			panic(err)
		}
	}()

	// driver = agouti.ChromeDriver(agouti.ChromeOptions("headless", "true"))
	driver = agouti.ChromeDriver()

	driver.Start()

	page, err = driver.NewPage()
	if err != nil {
		panic(errors.Wrap(err, "driver.NewPage failed"))
	}

	err = page.ClearCookies()
	if err != nil {
		panic(err)
	}

	err = page.Navigate("http://127.0.0.1:50203/amazon_pay_button")
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	err = page.FindByClass("amazon-button-image-2017").Click()
	if err != nil {
		panic(err)
	}

	rootWindow, err := page.Session().GetWindow()
	if err != nil {
		panic(err)
	}

	err = page.NextWindow()

	err = page.FindByID("ap_email").Fill(account.Email)
	if err != nil {
		panic(err)
	}
	err = page.AllByName("password").Fill(account.EmailPassword)
	if err != nil {
		panic(err)
	}
	err = page.FindByButton("サインイン（セキュリティシステムを使う）").Click()
	if err != nil {
		panic(err)
	}

	err = page.Session().SetWindow(rootWindow)
	if err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)

	token, err = page.FindByID("amazon_pay_access_token").Attribute("value")
	if err != nil {
		panic(err)
	}
	amazonOrderReferenceId, err = page.FindByID("amazon_pay_order_reference_id").Attribute("value")
	if err != nil {
		panic(err)
	}

	return
}
