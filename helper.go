package main

import (
	"fmt"
)


func addr() string {
  return fmt.Sprintf("%s:%s", APP_ADDRESS, APP_PORT)
} // addr


func checkParam(params... string) bool {
	
	for _, p := range params {
		if p == "" || len(p) == 0 {
			return false
		}
	}

	return true

} // checkParam
