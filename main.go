package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "tasks.txt"

func main() {
	fmt.Println("Добро пожаловать в менеджер задач!")
	fmt.Println("Доступные команды: add, list, delete, exit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := parts[1:]

		switch command {
		case "add":
			if len(args) == 0 {
				fmt.Println("Использование: add <текст задачи>")
				continue
			}
			taskText := strings.Join(args, " ")
			addTask(taskText)
			fmt.Println("Задача добавлена")

		case "list":
			listTasks()

		case "delete":
			if len(args) == 0 {
				fmt.Println("Использование: delete <номер задачи>")
				continue
			}
			num, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Ошибка: номер должен быть числом")
				continue
			}
			if deleteTask(num) {
				fmt.Println("Задача удалена")
			} else {
				fmt.Println("Ошибка: задача с таким номером не найдена")
			}

		case "exit":
			fmt.Println("Выход из программы")
			return

		default:
			fmt.Println("Неизвестная команда. Доступные команды: add, list, delete, exit")
		}
	}
}

func readTasks() []string {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		fmt.Println("Ошибка при чтении файла:", err)
		return []string{}
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return []string{}
	}

	return tasks
}

func saveTasks(tasks []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при сохранении задач:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		_, err := writer.WriteString(task + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи задачи:", err)
			return
		}
	}
	writer.Flush()
}

func addTask(text string) {
	tasks := readTasks()
	tasks = append(tasks, text)
	saveTasks(tasks)
}

func listTasks() {
	tasks := readTasks()
	if len(tasks) == 0 {
		fmt.Println("Нет задач")
		return
	}

	fmt.Println("Список задач:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func deleteTask(num int) bool {
	tasks := readTasks()
	if num < 1 || num > len(tasks) {
		return false
	}

	// Удаляем задачу (индексы начинаются с 0)
	tasks = append(tasks[:num-1], tasks[num:]...)
	saveTasks(tasks)
	return true
}
