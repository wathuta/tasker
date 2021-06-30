package cmd

import (
	"fmt"
	"log"
	"tasker/db"

	"github.com/spf13/cobra"
)

var ListCMD = &cobra.Command{
	Use:   "list",
	Short: "a list of all the tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.AllTasksFromDB()
		if err != nil {
			log.Println("[ERROR] ", err.Error())
		}
		if len(tasks) == 0 {
			fmt.Println("you have no tasks")
			return
		}
		for i, task := range tasks {
			fmt.Println(i+1, task.Value)
		}
	},
}

func init() {
	RootCMD.AddCommand(ListCMD)
}
