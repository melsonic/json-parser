package main

//     _
//    (_)___  ___  _ __        _ __   __ _ _ __ ___  ___ _ __
//    | / __|/ _ \| '_ \ _____| '_ \ / _` | '__/ __|/ _ \ '__|
//    | \__ \ (_) | | | |_____| |_) | (_| | |  \__ \  __/ |
//   _/ |___/\___/|_| |_|     | .__/ \__,_|_|  |___/\___|_|
//  |__/                      |_|
//

import (
	"log"
	"os"

	"github.com/melsonic/json-parser/util"
)

func main() {
	var result bool = false
	jsonFileName := os.Args[1]
	content, err := os.ReadFile(jsonFileName)
	if err != nil {
		log.Fatal(err)
	}
  result = util.Validate(content)
	util.PrintResult(result)
}
