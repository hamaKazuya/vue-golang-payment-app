package main

import (
	"vue-golang-payment-app/backend-api/infrastructure"
)

func main() {
	infrastructure.Router.Run(":8888")
}
