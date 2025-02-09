package main

const SNK_PORT = "8989"

func main() {
	dumper, err := NewDumper()
	if err != nil {
		panic(err)

	}

	sample_data := []byte("Foo Completed\nBar Completed\n")
	sample_logs := Logs{
		batch: 0,
		data:  &sample_data,
	}

	dumper.dumpLogs(&sample_logs)

}
