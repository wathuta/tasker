package cmd

import (
	"fmt"
	"log"
	"strconv"
	"tasker/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Removes the tasks from the task bar",
	Run: func(cmd *cobra.Command, args []string) {
		//parses all the input ids
		var ids []int
		for _, id := range args {
			newid, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal("invalid input id", err.Error())
			} else {
				ids = append(ids, newid)
			}
		}
		tasks, err := db.AllTasksFromDB()
		if err != nil {
			fmt.Println("Something Went Wrong", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid task number", id)
				continue
			}
			task := tasks[id-1]
			if err := db.DeleteTask(task.Key); err != nil {
				fmt.Printf("Failed to mark %d as completed.Error %s", id, err)
			} else {
				fmt.Printf("Marked %d as completed.", id)

			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCMD.AddCommand(doCmd)
}
