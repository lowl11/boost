package async

import (
	"context"
	"github.com/lowl11/boost/data/interfaces"
)

func WaitAll(ctx context.Context, tasks ...func(ctx context.Context) error) error {
	if len(tasks) == 0 {
		return nil
	}

	runTasks := make([]interfaces.Task, 0, len(tasks))

	for _, f := range tasks {
		task := NewTask(ctx)
		task.Run(f)
		runTasks = append(runTasks, task)
	}

	for _, task := range runTasks {
		if err := task.Wait(); err != nil {
			return err
		}
	}

	return nil
}
