package main

import (
	"flag"
)

var SNK_PORT = "8989"
var READ_TIME = 5

func main() {

	var harvesters []*Harvester

	configPath := flag.String("f", "", "Config File for CBeats")
	flag.Parse()

	if *configPath == "" {
		panic("Config File is necessary")
	}

	Cfg, err := ParseConfig(*configPath)
	if err != nil {
		panic(err)
	}

	globalDumper, err := NewDumper(SNK_PORT)
	if err != nil {
		panic(err)
	}

	for _, dir := range Cfg.Dirs {

		newHarvester, err := NewHarvester(HarvesterConfig{
			ReadDir:  dir,
			ReadTime: READ_TIME,
		}, globalDumper)
		if err != nil {
			panic(err)
		}

		harvesters = append(harvesters, newHarvester)
	}

	for _, harvester := range harvesters {
		go harvester.Start()
	}

	for {
	}

}
