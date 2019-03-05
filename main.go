package main

import "di-sample/di"

func main() {
	service := di.InjectSignUpService()
	service.SignUp("sam@ple.com", "sample")
}
