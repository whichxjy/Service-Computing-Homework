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

// register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Create a new account",
	Long: "Create a new account with username, password and email",
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
		email, _ := cmd.Flags().GetString("email")
		if email == "" {
			fmt.Println("Email cannot be empty")
			return
		}
		if service.Register(username, password, email) {
			fmt.Printf("[%v] Registered Successfully\n", username)
		} else {
			fmt.Printf("[%v] Fail to register\n", username)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	// flags
	registerCmd.Flags().StringP("user", "u", "", "username of the new account")
	registerCmd.Flags().StringP("password", "p", "", "password of the new account")
	registerCmd.Flags().StringP("email", "e", "", "email of the new account")
}
