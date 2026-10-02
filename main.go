package main

func main() {
	cfg := &config{
		commands: getCommands(),
		next:     21,
		previous: 1,
	}
	startRepl(cfg)
}
