package main

import (
	"github.com/Caledfwlch1/temp2/run2"
)

const getString = "http://35.190.167.185/custom?branch=dev&commit=6c910dffc96d6a4d0c5a32fc6f53bc4f9d050570&roles=builds&wait=true"

func main() {
	run2.Run2(getString)
	return
}