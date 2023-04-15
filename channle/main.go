// channel.go
package main

import (
	"fmt"
	"time"
)

// func main() {
// 	workpools()
// }
func workpools() {
	const number_of_jobs = 5
	const number_of_workers = 2
	jobs := make(chan int, number_of_jobs)
	results := make(chan string, number_of_jobs)
	// 向 任务队列写入任务
	for i := 1; i <= number_of_jobs; i++ {
		jobs <- i
	}
	fmt.Println("布置 job 后，关闭 jobs channel")
	close(jobs)
	// 控制并行度，每个 worker 函数都运行在单独的 goroutine 中
	for w := 1; w <= number_of_workers; w++ {
		go worker(w, jobs, results)
	}
	// 监听 results channel，只要有内容就会被取走
	for i := 1; i <= number_of_jobs; i++ {
		fmt.Printf("结果: %s\n", <-results)
	}
}

// worker 逻辑：一个不断从 jobs chan 中取任务的循环
// 并将结果放在 out channel 中待取
func worker(id int, jobs <-chan int, out chan<- string) {
	fmt.Printf("worker #%d 启动\n", id)
	for job := range jobs {
		fmt.Printf("worker #%d 开始 工作%d\n", id, job)
		// sleep 模拟 『正在处理任务』
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker #%d 结束 工作%d\n", id, job)
		out <- fmt.Sprintf("worker #%d 工作%d", id, job)
	}
	fmt.Printf("worker #%d 退出\n", id)
}
