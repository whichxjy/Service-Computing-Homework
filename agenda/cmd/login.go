/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whichxjy/Service-Computing-Homework/agenda/service"
)

// login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in Agenda system",
	Long: `Log in Agenda system with username and password`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		if username == "" {
			fmt.Println("Username cannot be empty")
			return
		}
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			fmt.Println("Password cannot be empty")
			return
		}
		if service.Login(username, password) {
			fmt.Printf("[%v] Login Successfully\n", username)
		} else {
			fmt.Printf("[%v] Fail to login\n", username)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	// flags
	loginCmd.Flags().StringP("user", "u", "", "username of the account")
	loginCmd.Flags().StringP("password", "p", "", "password of the account")
}
