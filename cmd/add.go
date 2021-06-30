package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"tasker/db"

	"github.com/spf13/cobra"
)

var AddCMD = &cobra.Command{
	Use:   "add",
	Short: "add is used to add commands",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		key, err := db.CreateTask(task)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		tasks, err := db.AllTasksFromDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, task := range tasks {
			if key == task.Key {
				fmt.Printf("Added id %d and key %d", i, key)
			}
		}
	},
}

func init() {
	RootCMD.AddCommand(AddCMD)
}
