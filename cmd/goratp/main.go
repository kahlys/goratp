package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/kahlys/goratp"
)

var (
	flagStation string
	flagType    string
	flagLine    string
)

func init() {
	flag.StringVar(&flagStation, "station", "dewoitine", "Slug name of the station (eg: bonne+nouvelle)")
	flag.StringVar(&flagType, "transport", "tramways", "Type of tansport (rers, metros, bus, tramways, noctiliens")
	flag.StringVar(&flagLine, "line", "6", "Transport line (eg: 8)")
}

func main() {
	flag.Parse()

	// get schedules datas
	sc, err := goratp.GetSchedules(flagStation, flagLine, flagType)
	if err != nil {
		log.Fatal(err)
	}

	// display datas
	var res string
	for k, v := range sc {
		res += fmt.Sprintf("\n%-20s", k)
		for _, t := range v {
			res += fmt.Sprintf(" %-20s", t)
		}
	}
	res = strings.Replace(res, "\n", "", 1)
	fmt.Println(res)
}
