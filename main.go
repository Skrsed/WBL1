package main

import (
	"L1/task"
	"fmt"
	"os"
	"strconv"

	// Your child packages get imported here.
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, os.Interrupt)
	// go func() {
	// 	fmt.Println("...ctrl + C to interrupt")
	// 	<-sigs

	// 	os.Exit(0)
	// }()

	switch os.Args[1] {
	case "task1":
		task.HumanActionEmbeded()
	case "task2":
		task.Concurency1()
	case "task3":
		task.Concurency2()
	case "task4":
		n, _ := strconv.Atoi(os.Args[2])
		task.ChanelDataStream(n)
	case "task7":
		task.ParallelMapWriter()
	case "task9":
		task.NumsConveyor()
	case "task10":
		task.Temperatures()
	case "task11":
		task.UGroup()
	case "task12":
		task.OwnSet()
	case "task14":
		task.CheckType()
	case "task15":
		task.SomeBadThing()
	case "task16":
		task.QuickSort()
	case "task17":
		task.BinarySearch()
	case "task18":
		task.ParallelIncrementation()
	case "task19":
		task.ReverseString()
	case "task20":
		task.SwapWords()
	case "task23":
		task.DeleteItem()
	case "task24":
		task.FindDistance()
	case "task25":
		task.MySleep()
	case "task26":
		task.AllUnique()
	}
}
