// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/techtraits/go-hello-world/gateway"
)

var gateway_port *int32
var server_host *string

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Run the API gateway",
	Long:  `Run the API gateway to enable Http calls into the Hello World GRPC Server`,
	Run: func(cmd *cobra.Command, args []string) {
		gateway.RunGateway(*gateway_port, *server_host)
	},
}

func init() {
	RootCmd.AddCommand(gatewayCmd)
	gateway_port = gatewayCmd.Flags().Int32P("port", "p", 8080, "The port on which to serve requests")
	server_host = gatewayCmd.Flags().StringP("server", "s", "localhost:5000", "The hosts(s) where the backend server is running")

}
