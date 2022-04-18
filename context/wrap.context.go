package main

import (
	"context"
	"log"
	"os"
	"sync"
)

var w sync.WaitGroup

type Context struct {
	baseCtx context.Context
	logger  log.Logger
	isTest  bool
}

func (c Context) Context() context.Context { return c.baseCtx }
func (c Context) Logger() log.Logger       { return c.logger }
func (c Context) IsTest() bool             { return c.isTest }

func NewContext(logger log.Logger, isTest bool) Context {
	return Context{
		baseCtx: context.Background(),
		logger:  logger,
		isTest:  isTest,
	}
}

func goroutineFuncWithLog(ctx Context) {
	logger := ctx.Logger()

	if ctx.IsTest() {
		logger.Println("start test go routine !")

		/* 테스트 코드 로직 */

		w.Done()
		return
	}

	logger.Println("start go routine !")
	
	/* 코드 로직 */

	w.Done()
}

func main() {
	count := 2

	w.Add(count)

	ctx1 := NewContext(*log.New(os.Stdout, "ctx 0 : ", log.LstdFlags), true)
	ctx2 := NewContext(*log.New(os.Stdout, "ctx 1 : ", log.LstdFlags), false)

	go goroutineFuncWithLog(ctx1)
	go goroutineFuncWithLog(ctx2)

	w.Wait()
}

/*
	ctx 1 : 2022/04/17 22:17:32 start test go routine !
	ctx 0 : 2022/04/17 22:17:32 start go routine !
	
	reference : https://github.com/cosmos/cosmos-sdk/blob/1fe72ccd3af9f559d29bbf7053e1df57b6f8f6fb/types/context.go
*/
