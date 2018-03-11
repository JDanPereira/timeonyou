package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"gopkg.in/cheggaaa/pb.v2"
)

type timedTask struct {
	Title string
	Time  int
}

var tasks []timedTask

func main() {
	fmt.Println(`=========== TimeOnYou! =============`)
	showMenu()
}

func showMenu() {
	opt := 0
	for opt != 2 {
		fmt.Printf("You have %d task(s) \n", len(tasks))
		fmt.Println("Select an option: ")
		fmt.Println("1 - add tasks")
		fmt.Println("2 - die")
		if len(tasks) > 0 {
			fmt.Println("3 - start")
		}
		fmt.Print("======>  ")
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1:
			getTasks()
			break
		case 2:
			fmt.Println("I'll be back.")
			break
		case 3:
			startTasks()
			break
		}
	}
}

func getTasks() {
	name := ""
	time := ""

	printHelp()

	for {
		fmt.Print("\nEnter task (name and time): ")
		fmt.Scanf("%s %s", &name, &time)
		if name == "b" {
			break
		}

		tasks = append(tasks, timedTask{Title: name, Time: parseTime(time)})
	}
}

func printHelp() {
	fmt.Println("Usage: hit 'b' to go back")
	fmt.Println("Example: 'MyTask 1:h' or 'MyTask 60:m'")
	fmt.Println("I'll run the tasks by 2/4 for a MVP, 1/4 for tests, 1/4 for fine tune/refactor")
}

func parseTime(time string) int {
	t := strings.Split(time, ":")

	if len(t) < 2 || !strings.Contains(t[1], "m") && !strings.Contains(t[1], "h") {
		panic("Invalid time format")
	}

	res, err := strconv.Atoi(t[0])
	if err != nil {
		panic("Invalid time")
	}
	if t[1] == "h" {
		res *= 60
	}

	return res
}

func startTasks() {
	fmt.Println("Starting timer...")

	for i := 0; i < len(tasks); i++ {
		fmt.Printf("You have %d minutes", tasks[i].Time)
		showProgBar(tasks[i].Title, (tasks[i].Time*60)/2, 1)
		notifyFirst(tasks[i])
		showProgBar(tasks[i].Title, ((tasks[i].Time*60)/2)/2, 2)
		notifySecond(tasks[i])
		showProgBar(tasks[i].Title, ((tasks[i].Time*60)/2)/2, 3)
		notifyThird(tasks[i])
	}
	tasks = []timedTask{}
}

func showProgBar(title string, count int, part int) {
	t := fmt.Sprintf(`{{ red "%s [%d/3]:" }}`, title, part)
	tmpl := t + ` {{bar . | green}} {{percent . | blue}}`
	bar := pb.ProgressBarTemplate(tmpl).Start(count)
	defer bar.Finish()
	for i := 0; i < count; i++ {
		bar.Add(1)
		time.Sleep(time.Second)
	}
}

func notifyFirst(task timedTask) {
	showNotification(
		`First part of the time has passed. Do you have a MVP?`)
}

func notifySecond(task timedTask) {
	showNotification(
		`Second part of the time has passed. It's tested?`)
}

func notifyThird(task timedTask) {
	showNotification(
		`Last part of the time has passed. It's done?`)
}

func showNotification(message string) {
	beeep.DefaultDuration = 10
	err := beeep.Notify("TimeOnYou!", message, "dialog-information")
	if err != nil {
		panic(err)
	}
}
