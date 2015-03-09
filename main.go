/*
   This file is part of Usermanager.

   Usermanager is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Usermanager is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	um "github.com/alfonsodev/usermanager/usermanager"
	"github.com/spf13/cobra"
)

func addUser() {
	fmt.Println("user can't be added yet")
}

func listUsers() {
	um.List()
}

func main() {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display version",
		Long:  `Display version of this software`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Usermanager -- HEAD")
		},
	}
	var listCmd = &cobra.Command{
		Use:   "users",
		Short: "list users",
		Long:  `list all users in the system`,
		Run: func(cmd *cobra.Command, args []string) {
			listUsers()
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(versionCmd, listCmd)
	rootCmd.Execute()

}
