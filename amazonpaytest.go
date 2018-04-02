package amazonpaytesthelper

import (
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

func AmazonPayTestHelper(account AmazonPayTestAccount) (token string, amazonOrderReferenceId string, err error) {

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

	err = page.Navigate("https://demo.getqor.com/")
	if err != nil {
		panic(err)
	}

	err = page.FirstByClass("products__list--title").Click()
	if err != nil {
		panic(err)
	}

	err = page.FindByButton("+ ADD TO CART").Click()
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

	token, err = page.FindByID("checkout-access-token").Attribute("value")
	if err != nil {
		panic(err)
	}
	amazonOrderReferenceId, err = page.FindByID("checkout-referenceid").Attribute("value")
	if err != nil {
		panic(err)
	}

	return
}
