package printer

import "fmt"

func PrintGreeting() {
	fmt.Println(
		`

  | |__   ___   ___  ___| |_ 
  | '_ \ / _ \ / _ \/ __| __|
  | |_) | (_) | (_) \__ \ |_ 
  |_.__/ \___/ \___/|___/\__|
  Minimalist Go framework based on FastHTTP
  https://github.com/lowl11/boost
--------------------------------------------`,
	)
}
