package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// File to store tasks
const taskFile = "tasks.json"

// Task represents a to-do task
type Task struct {
	Description string `json:"description"`
}

// Initialize tasks
var tasks []Task

func main() {

	rootCmd := &cobra.Command{
		Use:   "todo-app-cli",
		Short: "A CLI for managing todo lists",
	}

	//Create todo task

	rootCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create a new task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i, _ := range args {
				taskDescription := args[i]
				task := Task{Description: taskDescription}

				// Load existing tasks
				loadTasks()

				// Add the new task
				tasks = append(tasks, task)

				// Save tasks to file
				saveTasks()

				fmt.Printf("Task created: %s\n", taskDescription)

			}
		},
	})

	//List all todo tasks

	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			// Load tasks from file
			loadTasks()

			// Display tasks
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				fmt.Println("Tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task.Description)
				}
			}
		},
	})

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// Load tasks from the JSON file
func loadTasks() {
	file, err := os.Open(taskFile)
	if err != nil {
		// If file doesn't exist, assume no tasks
		if os.IsNotExist(err) {
			tasks = []Task{}
			return
		}
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Printf("Error decoding tasks: %v\n", err)
		os.Exit(1)
	}
}

// Save tasks to the JSON file
func saveTasks() {
	file, err := os.Create(taskFile)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Printf("Error encoding tasks: %v\n", err)
		os.Exit(1)
	}
}
