package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"todo_app/config"
	"todo_app/db"
	"todo_app/todo"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDatabase(cfg)

	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		input, err := getInput(reader, "Выберите команду: ")
		if err != nil {
			log.Fatalf("Ошибка ввода: %v", err)
		}

		switch input {
		case "1":
			task, err := getInput(reader, "Введите задачу: ")
			if err != nil {
				log.Fatalf("Ошибка ввода задачи: %v", err)
			}
			todo.AddTask(task)
			fmt.Println("Задача добавлена:", task)
		case "2":
			tasks := todo.GetTasks()
			printTasks(tasks)
		case "3":
			idStr, err := getInput(reader, "Введите ID задачи для удаления: ")
			if err != nil {
				log.Fatalf("Ошибка ввода ID: %v", err)
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Fatalf("Неверный ID: %v", err)
			}
			todo.DeleteTask(id)
			fmt.Println("Задача с ID удалена:", id)
		case "4":
			confirmation, err := getInput(reader, "Вы уверены, что хотите удалить все задачи? (y/n): ")
			if err != nil {
				log.Fatalf("Ошибка ввода: %v", err)
			}
			if strings.ToLower(confirmation) == "y" {
				todo.DeleteAllTasks()
				fmt.Println("Все задачи удалены")
			} else {
				fmt.Println("Отмена удаления всех задач")
			}
		case "5":
			fmt.Println("Выход из приложения")
			return
		default:
			fmt.Println("Неверная команда, попробуйте снова.")
		}
	}
}

func printMenu() {
	fmt.Println("\nПриложение для управления задачами")
	fmt.Println("1. Добавить новую задачу")
	fmt.Println("2. Показать все задачи")
	fmt.Println("3. Удалить задачу по ID")
	fmt.Println("4. Удалить все задачи")
	fmt.Println("5. Выход")
}

func getInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func printTasks(tasks []todo.Todo) {
	if len(tasks) == 0 {
		fmt.Println("На данный момент у вас нет запланированных задач")
		return
	}
	fmt.Println("Список задач:")
	fmt.Println("ID\tЗадача")
	fmt.Println(strings.Repeat("-", 30))
	for _, t := range tasks {
		fmt.Printf("%d\t%s\n", t.ID, t.Task)
	}
	fmt.Println(strings.Repeat("-", 30))
}
