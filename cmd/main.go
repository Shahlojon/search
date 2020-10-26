package main

import (
	// "github.com/Shahlojon/search/pkg/search"
	// "fmt"
	// "context"
	// "log"
	// "sync"
	// "time"
)

func main() {
	// ctx :=context.Background()
	// log.Print(ctx.Deadline())
	// log.Print(ctx.Err())

	// root:=context.Background()
	// ctx, _ := context.WithDeadline(root, time.Now().Add(time.Second*10))
	// log.Print(ctx.Deadline())
	// <-ctx.Done()
	// log.Print(ctx.Err())
	// log.Print(ctx.Deadline())
	// log.Print(ctx.Err()==context.DeadlineExceeded)

	// log.Print(ctx.Err())
	// cancel()
	// log.Print(ctx.Err())
	// log.Print(ctx.Err() == context.Canceled)
	// log.Print(root.Err())

	// go func(ctx context.Context) {
	// 	for {
	// 		select {
	// 		case<-ctx.Done():
	// 			log.Print("done")
	// 			return
	// 		case <- time.After(time.Second):
	// 			log.Print("tick")
	// 		}
	// 	}
	// }(ctx)
	// <-time.After(time.Second*10)
	// log.Print("cancel")
	// cancel()

	// root := context.Background()
	// ctx, _ := context.WithDeadline(root, time.Now().Add(time.Second*10))
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func(ctx context.Context) {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			log.Print("done")
	// 			wg.Done()
	// 			return
	// 		case <-time.After(time.Second):
	// 			log.Print("tick")
	// 		}
	// 	}
	// }(ctx)
	// wg.Wait()
	// log.Print("main done")

	// res := search.Any(context.Background(),"HTTP", []string{"./test.txt","./test copy.txt"})


	// r, ok := <- res
	// if !ok {
	// 	fmt.Println("error ok =>", ok)
	// }

   //for _, r := range res {

	//    fmt.Println("---------------")
	//    fmt.Println("res.Phrase) => ", r.Phrase)
	//    fmt.Println("res.Line) => ", r.Line)
	//    fmt.Println("res.LineNum) => ", r.LineNum)
	//    fmt.Println("res.ColNum) => ", r.ColNum)
	//    fmt.Println("---------------")
   //}

}
