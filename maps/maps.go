//--Summary:
//  Write a program to display server status.
//
//--Requirements:

package main

import "fmt"

type Server map[string]int

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// Creating status map for making readable Statuses

var statusMap = map[int]string{0: "online", 1: "Offline", 2: "Maintenance", 3: "Retired"}

func serversByStatus(servers Server) {
	fmt.Println("Total Number of Servers: ", len(servers))

	statusCount := make(map[string]int)

	for _, status := range servers {
		readableStatus := statusMap[status]
		statusCount[readableStatus] += 1
	}
	fmt.Println(statusCount)
}
func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Store the existing slice of servers in a map
	//* Default all of the servers to `Online`
	serverMap := make(Server)
	for _, server := range servers {
		serverMap[server] = 0
	}

	// * Perform the following status changes and display server info:
	//   - display server info
	serversByStatus(serverMap)
	// - change `darkstar` to `Retired`
	serverMap["darkstar"] = 3
	// - change `aiur` to `Offline`
	serverMap["aiur"] = 1
	// - display server info
	serversByStatus(serverMap)
	// - change all servers to `Maintenance`
	for name := range serverMap {
		serverMap[name] = 2
	}
	// - display server info
	serversByStatus(serverMap)

}
