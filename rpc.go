package boost

import (
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"strconv"
)

func getDefaultPort(port string) string {
	const defaultPort = "8081"

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return defaultPort
	}

	portNumber++
	return type_helper.ToString(portNumber)
}
