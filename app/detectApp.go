package app

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/TheCoy/razor/util"
	"github.com/TheCoy/toolkit/routinepool"
)

type RegexDetectApp struct {
	InputFileName  string
	OutputFileName string
	LogFileName    string
	DetectType     string
	Worker         int

	*log.Logger
	wg            sync.WaitGroup
	inputPointer  *os.File
	outputPointer *os.File
}

func (app *RegexDetectApp) String() string {
	str := fmt.Sprintf("InputFileName=%s\nOutputFileName=%s\nLogFileName=%s\nDetectType=%s\nWorker=%d\n",
		app.InputFileName,
		app.OutputFileName,
		app.LogFileName,
		app.DetectType,
		app.Worker,
	)
	return str
}

func (app *RegexDetectApp) initApp() error {
	logOutput, err := os.OpenFile(app.LogFileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	app.Logger = log.New(logOutput, "", 1)

	app.outputPointer, err = os.OpenFile(app.OutputFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	app.inputPointer, err = os.Open(app.InputFileName)
	if err != nil {
		fmt.Println(app.InputFileName)
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return errors.New("open file failed")
	}

	return nil
}

func (app *RegexDetectApp) run() error {
	defer app.inputPointer.Close()
	defer app.outputPointer.Close()

	workPool := routinepool.NewPool(app.Worker)
	go func() {
		defer close(workPool.EntryChannel)
		inputReader := bufio.NewReader(app.inputPointer)
		for {
			inputLine, inputError := inputReader.ReadString('\n')
			if inputError == io.EOF {
				break
			}
			app.wg.Add(1)
			workPool.EntryChannel <- app.buildDetectTask(inputLine)
		}
	}()

	workPool.Run()

	app.wg.Wait()

	return nil
}

func (app *RegexDetectApp) buildDetectTask(line string) *routinepool.Task {
	task := routinepool.NewTask(func() error {
		app.wg.Done()
		app.Logger.Println("detectType:", app.DetectType)

		switch app.DetectType {
		case "html":
			find, matchStr := util.HtmlDetect(string(line))
			if find {
				fmt.Fprintln(app.outputPointer, "[html detected]", line, "\n", matchStr)
			}
		case "latex":
			find, matchStr := util.LatexDetect(string(line))
			if find {
				fmt.Fprintln(app.outputPointer, "[latex detected]", line, "\n", matchStr)
			}
		default:
			find, matchStr := util.MultiDetect(string(line))
			if find {
				fmt.Fprintln(app.outputPointer, "[multi detected]", line, "\n", matchStr)
			}
		}
		return nil
	})

	return task
}
