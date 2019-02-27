package main

// avoid using relative import path, which Go version do you use?
// thats why glint failed
import "github.com/fehersanyi-bitrise/app-info/cmd"

func main() {
	cmd.Execute()
}
