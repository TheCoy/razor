package app

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/TheCoy/toolkit/routinepool"
)

type DemoApp struct {
	LogFileName string
	MaxWorker   int
	Times       int64
	*log.Logger
	wg sync.WaitGroup
}

func (app *DemoApp) initApp() error {
	logOutput, err := os.OpenFile(app.LogFileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	app.Logger = log.New(logOutput, app.LogFileName, 1)

	return nil
}

func (app *DemoApp) run() error {
	workPool := routinepool.NewPool(app.MaxWorker)
	go func() {
		defer close(workPool.EntryChannel)
		var i int64
		for i = 0; i < app.Times; i++ {
			app.wg.Add(1)
			workPool.EntryChannel <- app.buildTask(i)
			app.Logger.Printf("task[%d] added", i)
		}
	}()
	workPool.Run()
	app.wg.Wait()

	return nil
}

func (app *DemoApp) buildTask(seq int64) *routinepool.Task {
	task := routinepool.NewTask(func() error {
		time.Sleep(1000 * time.Microsecond)
		fmt.Printf("No[%d] started at %s\n", seq, time.Now().Format("2006-01-02 03:04:05"))
		app.wg.Done()
		return nil
	})

	return task
}
