package modules

const TASK_STATE_NO_PROGRESS = "NO PROGRESS"
const TASK_STATE_PROGRESS = "PROGRESS"
const TASK_STATE_DONE = "DONE"

type Task struct {
	Client *SocketClient
	Data string
	Result string
	State string
}

func (this *Task) SetClient (client *SocketClient) {
	this.Client = client
}

func CreateTasks(config *Config) {
	clients := GetSocketClients()

	if (len(tasks) == 0 && 0 < len(clients)) {
		for range clients {
			task := &Task{
				Client: nil,
				Data: "",
				State: TASK_STATE_NO_PROGRESS,
			}
			tasks = append(tasks, task)
		}
	}
}

func CheckOldTasks() {
	if (0 < len(tasks)) {
		clients := GetSocketClients()
		freeClients := make([]*SocketClient, 0)

		for _, client := range clients {
			if (client.IsFree()) {
				freeClients = append(freeClients, client)
			}
		}

		noProgressTasks := make([]*Task, 0)

		for _, task := range tasks {
			if (task.State == TASK_STATE_NO_PROGRESS) {
				noProgressTasks = append(noProgressTasks, task)
			}
		}

		for 0 < len(freeClients) && 0 < len(noProgressTasks) {
			var freeClient *SocketClient
			var noProgressTask *Task

			freeClient, freeClients = freeClients[len(freeClients) - 1], freeClients[:len(freeClients) - 1]
			noProgressTask, noProgressTasks = noProgressTasks[len(noProgressTasks) - 1], noProgressTasks[:len(noProgressTasks) - 1]

			noProgressTask.Client = freeClient
			freeClient.SendTask(noProgressTask)
		}
	}
}

func ThatClientCompletedTask(client *SocketClient, data string) {
	for _, task := range tasks {
		if (task.Client == client) {
			task.State = TASK_STATE_DONE
			task.Result = data
		}
	}
}

func CheckCompletedTasks() {
	if (0 < len(tasks)) {

	}
}


// --- PRIVATE ---------------------------------------------------------------------------------------------------------
var tasks = make([]*Task, 0);

