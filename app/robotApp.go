package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/TheCoy/toolkit/routinepool"
	"go.uber.org/ratelimit"
)

type RobotApp struct {
	Times  int64
	QPS    int64
	Worker int

	limiter ratelimit.Limiter
	wg      sync.WaitGroup
}

func (app *RobotApp) initApp() error {
	app.limiter = ratelimit.New(int(app.QPS))

	return nil
}

func (app *RobotApp) run() error {
	workPool := routinepool.NewPool(app.Worker)
	go func() {
		defer close(workPool.EntryChannel)
		for i := 0; i < int(app.Times); i++ {
			app.limiter.Take()
			app.wg.Add(1)
			workPool.EntryChannel <- app.buildTask(int64(i))
		}
	}()

	workPool.Run()

	app.wg.Wait()

	return nil
}

func (app *RobotApp) buildTask(i int64) *routinepool.Task {
	task := routinepool.NewTask(func() error {
		defer app.wg.Done()

		maxTimes := 1000000
		qps := int(app.QPS)
		if qps > 1000 {
			qps = 1000
		}

		sleep := maxTimes / qps / app.Worker
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		fmt.Println("=======REQUESTED!=======", i)

		return nil
	})

	return task
}

//SetNewLimiter changes built-in rateLimiter
func (app *RobotApp) SetNewLimiter(qps int) {
	if qps != int(app.QPS) {
		app.QPS = int64(qps)
	}
	app.limiter = ratelimit.New(qps)
}
