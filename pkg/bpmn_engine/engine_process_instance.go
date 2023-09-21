package bpmn_engine

import (
	"time"

	"github.com/rhzs/lib-bpmn-engine/pkg/spec/BPMN20/process_instance"
)

type ProcessInstanceInfo struct {
	processInfo     *ProcessInfo
	instanceKey     int64
	variableContext map[string]interface{}
	createdAt       time.Time
	state           process_instance.State
	caughtEvents    []CatchEvent
}

type ProcessInstance interface {
	GetProcessInfo() *ProcessInfo
	GetInstanceKey() int64
	GetVariable(key string) interface{}
	SetVariable(key string, value interface{})
	GetCreatedAt() time.Time
	GetState() process_instance.State
}

func (pii *ProcessInstanceInfo) GetProcessInfo() *ProcessInfo {
	return pii.processInfo
}

func (pii *ProcessInstanceInfo) GetInstanceKey() int64 {
	return pii.instanceKey
}

func (pii *ProcessInstanceInfo) GetVariable(key string) interface{} {
	pii.initVariableContext()

	return pii.variableContext[key]
}

func (pii *ProcessInstanceInfo) SetVariable(key string, value interface{}) {
	pii.initVariableContext()

	delete(pii.variableContext, key)
	pii.variableContext[key] = value
}

func (pii *ProcessInstanceInfo) initVariableContext() {
	if pii.variableContext == nil {
		pii.variableContext = make(map[string]interface{})
	}
}

func (pii *ProcessInstanceInfo) GetCreatedAt() time.Time {
	return pii.createdAt
}

func (pii *ProcessInstanceInfo) Purge() {
	for k := range pii.variableContext {
		delete(pii.variableContext, k)
	}

	pii.caughtEvents = nil
	pii.variableContext = nil
}

// GetState returns one of [ProcessInstanceReady,ProcessInstanceActive,ProcessInstanceCompleted]
//
//	┌─────┐
//	│Ready│
//	└──┬──┘
//
// ┌───▽──┐
// │Active│
// └───┬──┘
// ┌────▽────┐
// │Completed│
// └─────────┘
func (pii *ProcessInstanceInfo) GetState() process_instance.State {
	return pii.state
}
