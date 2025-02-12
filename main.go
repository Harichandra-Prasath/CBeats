package main

import (
	"flag"
)

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

	if CONTAINERS {

		// Create the Client
		dockerClient, err := NewDockerClient()
		if err != nil {
			panic(err)
		}

		containers, err := dockerClient.FetchContainerIDs()
		if err != nil {
			panic(err)
		}

		for _, container := range containers {
			newHarvester, err := NewHarvester(HarvesterConfig{
				ReadDir:  CONTAINER_LOGS_PATH + container.ContainerId + "/",
				ReadTime: READ_TIME,
			}, globalDumper)
			if err != nil {
				panic(err)
			}

			harvesters = append(harvesters, newHarvester)
		}
	}

	for _, harvester := range harvesters {
		go harvester.Start()
	}

	for {
	}

}
