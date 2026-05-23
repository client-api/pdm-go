// Example: list cluster nodes.
//
// Run with:
//
//	PDM_HOST=https://pdm.example.com:8443 \
//	PDM_TOKEN='PDMAPIToken=root@pam!auto=...' \
//	go run ./examples/list_nodes
package main

import (
	"context"
	"fmt"
	"os"

	pdm "github.com/client-api/pdm-go"
)

func main() {
	host := os.Getenv("PDM_HOST")
	if host == "" {
		host = "https://localhost:8443"
	}
	cfg := pdm.NewConfiguration()
	cfg.Servers = append(pdm.ServerConfigurations{}, pdm.ServerConfiguration{URL: host + "/api2/json"})
	cfg.DefaultHeader["Authorization"] = os.Getenv("PDM_TOKEN")

	client := pdm.NewAPIClient(cfg)
	resp, _, err := client.NodesAPI.NodesGetNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "list nodes:", err)
		os.Exit(1)
	}
	nodes := resp.GetData()
	fmt.Printf("Found %d node(s):\n", len(nodes))
	for _, n := range nodes {
		// Other products expose a slimmer Node shape; print verbatim.
		fmt.Printf("  - %+v\n", n)
	}
}
