
# Getting Started

Here are two example on how to use the library.

## Hello World Example

Assuming this simple 'Hello World' BPMN example should just print "hello world". \
![hello_world.png](./examples/hello_world/simple_task.png)

Then a simple (and verbose) code to execute this looks like this

<!-- MARKDOWN-AUTO-DOCS:START (CODE:src=./examples/hello_world/hello_world.go) -->
<!-- The below code snippet is automatically added from ./examples/hello_world/hello_world.go -->
```go
package main

import (
	"fmt"
	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine"
)

func main() {
	// create a new named engine
	bpmnEngine := bpmn_engine.New("a name")
	// basic example loading a BPMN from file,
	process, err := bpmnEngine.LoadFromFile("simple_task.bpmn")
	if err != nil {
		panic("file \"simple_task.bpmn\" can't be read.")
	}
	// register a handler for a service task by defined task type
	bpmnEngine.AddTaskHandler("hello-world", printContextHandler)
	// setup some variables
	variables := map[string]interface{}{}
	variables["foo"] = "bar"
	// and execute the process
	bpmnEngine.CreateAndRunInstance(process.ProcessKey, variables)
}

func printContextHandler(job bpmn_engine.ActivatedJob) {
	println("< Hello World >")
	println(fmt.Sprintf("ElementId                = %s", job.ElementId))
	println(fmt.Sprintf("BpmnProcessId            = %s", job.BpmnProcessId))
	println(fmt.Sprintf("ProcessDefinitionKey     = %d", job.ProcessDefinitionKey))
	println(fmt.Sprintf("ProcessDefinitionVersion = %d", job.ProcessDefinitionVersion))
	println(fmt.Sprintf("CreatedAt                = %s", job.CreatedAt))
	println(fmt.Sprintf("Variable 'foo'           = %s", job.GetVariable("foo")))
	job.Complete() // don't forget this one, or job.Fail("foobar")
}
```
<!-- MARKDOWN-AUTO-DOCS:END -->
