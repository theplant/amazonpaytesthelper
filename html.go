package amazonpaytesthelper

var amazonPayButtonHTML = `

<!doctype html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title></title>
  </head>

  <body>
    

<div id="amazon_login_widget" class="additional-checkout-buttons">
</div>


<div id="amazon_addressbook_widget" style="height: 500px;">
addressbook
</div>
<div id="amazon_wallet_widget" style="height: 500px;">
wallet
</div>

<input type="text" id="amazon_pay_order_reference_id" class="hiddenvalue">
<input type="text" id="amazon_pay_access_token" class="hiddenvalue">
<input type="text" id="amazon_pay_billing_id" class="hiddenvalue">

<style>
input.hiddenvalue {
    display: block;
    width: 500px;
}
</style>

<script type="text/javascript">
var amazonMerchantId = "%s";
var amazonPayClientId = "%s";
var amazonOrderReferenceId = null;

window.onAmazonLoginReady = function () {
    amazon.Login.setClientId(amazonPayClientId);
};

window.onAmazonPaymentsReady = function () {
        showAmazonButton();
        showAmazonAddress();
};

function showAmazonButton() {
    console.log("showAmazonButton")
    var authRequest;
    OffAmazonPayments.Button("amazon_login_widget", amazonMerchantId, {
        type: "PwA",
        color: "Gold",
        size: "medium",
        authorization: function () {
            loginOptions = {
                scope: "profile postal_code payments:widget payments:shipping_address",
                popup: "false"
            };
            authRequest = amazon.Login.authorize(loginOptions,"/amazon_pay_button");

        }
    });
}

function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}

function showAmazonAddress() {
    console.log("showAmazonAddress")
    new OffAmazonPayments.Widgets.AddressBook({
        sellerId: amazonMerchantId,
        agreementType: 'orderReference',
        onOrderReferenceCreate: function (orderReference) {
            console.log("onOrderReferenceCreate", orderReference);
            amazonOrderReferenceId = orderReference.getAmazonOrderReferenceId()
            document.getElementById("amazon_pay_order_reference_id").value = orderReference.getAmazonOrderReferenceId();
            document.getElementById("amazon_pay_access_token").value = getParameterByName('access_token');
            showAmazonWallet(amazonOrderReferenceId)
        },
        design: {
            designMode: 'responsive'
        },
        onReady: function(billingAgreement) {
            var billingAgreementId = billingAgreement.getAmazonBillingAgreementId(); 
            document.getElementById("amazon_pay_billing_id").value = billingAgreementId ;
        },
        onAddressSelect: function (billingAgreement) {
            console.log("billingAgreement", billingAgreement)
        },
        onError: function (error) {
        }
    }).bind("amazon_addressbook_widget");
}

function showAmazonWallet(amazonOrderReferenceId){
    console.log("showAmazonWallet")
    new OffAmazonPayments.Widgets.Wallet({
      sellerId: amazonMerchantId,
      amazonOrderReferenceId: amazonOrderReferenceId,
      onPaymentSelect: function(orderReference) {
      },
      design: {
        designMode: 'responsive'
      },
      onError: function(error) {
        console.log(error.getErrorCode() + ': ' + error.getErrorMessage());
      }
    }).bind("amazon_wallet_widget");
}

</script>

<script async="async" src='https://static-fe.payments-amazon.com/OffAmazonPayments/jp/sandbox/lpa/js/Widgets.js'></script>

  </body>
</html>

`
