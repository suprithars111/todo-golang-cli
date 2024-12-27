# To-Do CLI Application

This is a simple Command-Line Interface (CLI) application written in Go to manage a to-do list. Tasks are stored in a `tasks.json` file, ensuring persistence across multiple runs.

---

## Features

1. **Create a Task**
   - Add a new task to your to-do list.
   - Tasks are saved in a `tasks.json` file located in the same directory as the program.
2. **List Tasks**
   - View all tasks in your to-do list.
   - Displays a numbered list or a message if no tasks exist.
3. **JSON Storage**
   - Tasks are stored in a JSON file for easy portability and extensibility.

---

## Usage

### **Commands**

1. **Create a Task**
   - Command: `create`
   - Adds a task to the list.
   - Example:
     ```bash
     go run main.go create "Buy groceries"
     ```
   - Output:
     ```
     Task created: Buy groceries
     ```

2. **List Tasks**
   - Command: `list`
   - Displays all tasks in the to-do list.
   - Example:
     ```bash
     go run main.go list
     ```
   - Output:
     ```
     Tasks:
     1. Buy groceries
     ```
   - If no tasks exist, the output will be:
     ```
     No tasks available.
     ```

---

## Example `tasks.json`

When you create tasks, they are saved in a `tasks.json` file. Example content:

```json
[
    {
        "description": "Buy groceries"
    }
]
```

---

## How It Works

1. **Create Task (create command):**
   - Accepts a task description as an argument.
   - Loads existing tasks from `tasks.json`.
   - Appends the new task to the task list.
   - Saves the updated task list back to `tasks.json`.

   Example:
   ```bash
   go run main.go create "Buy groceries"
   ```
   Output:
   ```
   Task created: Buy groceries
   ```

2. **List Tasks (list command):**
   - Loads tasks from `tasks.json`.
   - Displays all tasks or a "No tasks available" message if the file is empty.

   Example:
   ```bash
   go run main.go list
   ```
   Output:
   ```
   Tasks:
   1. Buy groceries
   ```

---

## Benefits

1. **Persistence:** Tasks remain stored in `tasks.json` between runs.
2. **Portability:** The JSON file is easy to move or share.
3. **Extensibility:** Easily add more fields to the `Task` struct, such as priority or due dates.

---

## Improvements

1. **Error Handling:**
   - Add more detailed error messages for file read/write issues.

2. **Custom File Location:**
   - Use a flag to specify a custom file path for task storage.
   - Example:
     ```go
     rootCmd.PersistentFlags().StringVar(&taskFile, "file", "tasks.json", "File to store tasks")
     ```

3. **Task Deletion:**
   - Add a `delete` command to remove tasks by their index.

---

## Installation

1. Build the binary:
   ```bash
   go build -o todo-app-cli
   ```

2. Add the binary directory to your `PATH` for global usage:
   ```bash
   export PATH=$PATH:$(pwd)
   ```

3. Run the application:
   ```bash
   todo-app-cli create "Buy groceries"
   todo-app-cli list
   ```

---

