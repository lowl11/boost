package async

import (
	"context"
)

func WaitAll(ctx context.Context, tasks ...func(ctx context.Context) error) error {
	if len(tasks) == 0 {
		return nil
	}

	runTasks := make([]*Task, 0, len(tasks))

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
