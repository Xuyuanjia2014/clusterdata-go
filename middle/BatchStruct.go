package middle

type Batch struct {
	AllJobs map[string]Job `yaml:"Jobs,flow"`
}

type Job struct {
	FatalTag string
	SelfTasks map[string]Task `yaml:"Tasks,flow"`
}

type Task struct {
	InstanceNumber int
	Type string
	TaskStatus string
	StartTime int
	EndTime int
	Cpu int
	Mem int
	NextTaskName string
	SelfInstances map[string]Instance `yaml:"Instances,flow"`
}

type Instance struct {
	Type string
	InstanceStatus string
	StartTime int
	EndTime int
	MachineId string
	SeqNo int
	TotalSeqNo int
	CpuAvg float64
	CpuMax float64
	MemAvg float64
	MemMax float64
}



