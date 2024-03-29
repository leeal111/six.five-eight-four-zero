package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.

type GetWorkeeIDArgs struct {
}
type GetWorkeeIDReply struct {
	WorkeeID int
}

type GetTaskArgs struct {
	WorkeeID int
}
type GetTaskReply struct {
	TaskType  TaskType
	TaskID    int
	FileNames []string
	NReduce   int
}

type TaskDoneArgs struct {
	WorkeeID        int
	ResultFileNames []string
}

type TaskDoneReply struct {
}

type RecvHeartBeatArgs struct {
	WorkeeID int
}

type RecvHeartBeatReply struct {
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
