package main

type incrementCmd struct{}

func (inc *incrementCmd) description() string {
	return "increments a version"
}

func (inc *incrementCmd) action() {
}
