package back

import (
	"fmt"
	"time"
)

type TimedTask struct {
	*TaskBasic
	Deadline time.Time
}

func (tt *TimedTask) String() string {
	taskStr := tt.TaskBasic.String()
	return fmt.Sprintf("%s, Дедлайн: %s", taskStr, tt.Deadline.Format("2006-01-02 15:04"))
}
