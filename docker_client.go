package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Container struct {
	ContainerId string   `json:"Id"`
	Names       []string `json:"Names"`
}

type DockerClient struct {
	UnixClient *http.Client
}

func NewDockerClient() (*DockerClient, error) {

	conn, err := net.Dial("unix", DOCKER_SOCK)
	if err != nil {
		return nil, fmt.Errorf("creating dockerclient: %s", err)

	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return conn, nil
			},
		},
	}

	return &DockerClient{
		UnixClient: &client,
	}, nil

}

func (d *DockerClient) FetchContainerIDs() ([]Container, error) {

	resp, err := d.UnixClient.Get(CONTAINER_URL)
	if err != nil {
		return nil, fmt.Errorf("fetching containers: %s", err)

	}
	defer resp.Body.Close()

	raw_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetching containers: %s", err)

	}

	var allContainers []Container

	err = json.Unmarshal(raw_data, &allContainers)
	if err != nil {
		return nil, fmt.Errorf("fetching containers: %s", err)
	}

	if CONTAINERS_LIST == "*" {
		return allContainers, nil
	}

	var requestedContainers []Container
	containerNames := strings.Split(CONTAINERS_LIST, ",")

	for _, container := range allContainers {
		for _, name := range container.Names {

			name = strings.TrimPrefix(name, "/")
			if isIncluded(name, containerNames) {

				requestedContainers = append(requestedContainers, container)
			}

		}
	}

	return requestedContainers, nil
}
