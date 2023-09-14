package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
	"time"
)

// Add your RPC definitions here.

// task status enumeration
type TaskStatus = string

const (
	Idle       TaskStatus = "idle"
	InProgress TaskStatus = "in-progress"
	Completed  TaskStatus = "completed"
)

// task type enumeration
type TaskType = string

const (
	Map    TaskType = "map"
	Reduce TaskType = "reduce"
)

// task
type Task struct {
	id         int
	taskStatus TaskStatus
	taskType   TaskType
	fileName   string
	startTime  time.Time
}

type heartBeatMessage struct {
	workerId int
}

type reportMessage struct {
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
