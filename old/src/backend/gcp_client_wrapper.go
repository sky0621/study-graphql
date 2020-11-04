package backend

import (
	"context"
	"fmt"

	"google.golang.org/api/option"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2beta3"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"
)

type CloudTasksQueueKind int

const (
	CloudTasksQueueKindTodo CloudTasksQueueKind = iota + 1
	CloudTasksQueueKindUser
)

func NewGCPClientWrapper(projectID, locationID, credentialPath, taskExecURL string, queueIDMap map[CloudTasksQueueKind]string) *GCPClientWrapper {
	return &GCPClientWrapper{
		projectID:      projectID,
		locationID:     locationID,
		credentialPath: credentialPath,
		taskExecURL:    taskExecURL,
		queueIDMap:     queueIDMap,
	}
}

type GCPClientWrapper struct {
	projectID, locationID, credentialPath, taskExecURL string
	queueIDMap                                         map[CloudTasksQueueKind]string
}

func (w *GCPClientWrapper) CreateCloudTasksTask(ctx context.Context, kind CloudTasksQueueKind, message string) (*taskspb.Task, error) {
	// Create a new Cloud Tasks client instance.
	// See https://godoc.org/cloud.google.com/go/cloudtasks/apiv2
	client, err := cloudtasks.NewClient(ctx, option.WithCredentialsFile(w.credentialPath))
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}
	defer client.Close()

	// Build the Task queue path.
	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", w.projectID, w.locationID, w.queueIDMap[kind])

	// Build the Task payload.
	// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2beta3#CreateTaskRequest
	req := &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2beta3#HttpRequest
			PayloadType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        w.taskExecURL,
				},
			},
		},
	}

	// Add a payload message if one is present.
	req.Task.GetHttpRequest().Body = []byte(message)

	createdTask, err := client.CreateTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cloudtasks.CreateTask: %v", err)
	}

	return createdTask, nil
}
