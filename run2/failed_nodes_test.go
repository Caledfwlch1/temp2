package run2

import (
	"testing"
	"time"
)
// *****
func TestFailed_nodes(t *testing.T) {
	//Run2("http://35.190.167.185/custom?branch=dev&commit=e43c8b290fb575808b8aafb680a45eeeb2221282&roles=failed_nodes&wait=true")
	time.Sleep(time.Minute * 11)
	Run2("http://localhost/")
	return
}