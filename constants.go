package main

var SNK_PORT = "8989"
var READ_TIME = 5
var CONTAINERS = false
var DOCKER_SOCK = "/var/run/docker.sock"

var ALLOWED_DIRECTIVES = []string{"DIRS", "PORT", "READ_TIME", "CONTAINERS", "DOCKER_SOCK"}

const CONTAINER_URL = "http://unix/v1.47/containers/json"
const CONTAINER_LOGS_PATH = "/var/lib/docker/containers/"
