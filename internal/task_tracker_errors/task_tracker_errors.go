package task_tracker_errors

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrInvalidStatus = errors.New("invalid status")
