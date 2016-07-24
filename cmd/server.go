// Copyright © 2016 Usman Ismail usman@techtraits.com
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
	"github.com/techtraits/go-hello-world/server"
)

var service_port *int32
var db_connection *string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the hello world server",
	Long:  `Run the hello world server to accept Hello World Calls`,
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer(*service_port)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	service_port = serverCmd.Flags().Int32P("port", "p", 5000, "The port on which to serve requests")
	db_connection = serverCmd.Flags().StringP("db-connection-string", "d", "", "The hostname of the database server(s)")
}
