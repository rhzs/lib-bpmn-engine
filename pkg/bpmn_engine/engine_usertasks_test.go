package bpmn_engine

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/rhzs/lib-bpmn-engine/pkg/spec/BPMN20/process_instance"
)

func Test_user_tasks_can_be_handled(t *testing.T) {
	// setup
	bpmnEngine := New("name")
	process, err := bpmnEngine.LoadFromFile("../../test-cases/simple-user-task.bpmn")
	then.AssertThat(t, err, is.Nil())
	cp := CallPath{}
	bpmnEngine.AddTaskHandler("user-task", cp.CallPathHandler)

	instance, _ := bpmnEngine.CreateAndRunInstance(process.ProcessKey, nil)

	then.AssertThat(t, instance.state, is.EqualTo(process_instance.COMPLETED))
	then.AssertThat(t, cp.CallPath, is.EqualTo("user-task"))
}
