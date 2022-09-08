package bpmn_engine

import (
	"errors"
	"fmt"
	"time"

	"github.com/rhzs/lib-bpmn-engine/pkg/spec/BPMN20"
	"github.com/senseyeio/duration"
)

type Timer struct {
	ElementId          string
	ElementInstanceKey int64
	ProcessInstanceKey int64
	State              TimerState
	CreatedAt          time.Time
	DueDate            time.Time
}

type TimerState string

const (
	TimerCreated   TimerState = "CREATED"
	TimerTriggered TimerState = "TRIGGERED"
	TimerCancelled TimerState = "CANCELLED"
)

func (state *BpmnEngineState) handleIntermediateTimerCatchEvent(process *ProcessInfo, instance *ProcessInstanceInfo, ice BPMN20.TIntermediateCatchEvent) bool {
	timer := findExistingTimerNotYetTriggered(state, ice.Id, instance)
	if timer == nil {
		newTimer, err := createNewTimer(process, instance, ice, state.generateKey)
		if err != nil {
			return false
		}
		timer = newTimer
		state.timers = append(state.timers, timer)
	}
	if time.Now().After(timer.DueDate) {
		timer.State = TimerTriggered
		return true
	}
	return false
}

func createNewTimer(process *ProcessInfo, instance *ProcessInstanceInfo, ice BPMN20.TIntermediateCatchEvent,
	generateKey func() int64,
) (*Timer, error) {
	durationVal, err := findDurationValue(ice, process)
	if err != nil {
		return nil, &BpmnEngineError{Msg: fmt.Sprintf("Error parsing 'timeDuration' value "+
			"from element with ID=%s. Error:%s", ice.Id, err.Error())}
	}
	now := time.Now()
	return &Timer{
		ElementId:          ice.Id,
		ElementInstanceKey: generateKey(),
		ProcessInstanceKey: instance.instanceKey,
		State:              TimerCreated,
		CreatedAt:          now,
		DueDate:            durationVal.Shift(now),
	}, nil
}

func findExistingTimerNotYetTriggered(state *BpmnEngineState, id string, instance *ProcessInstanceInfo) *Timer {
	var t *Timer
	for _, timer := range state.timers {
		if timer.ElementId == id && timer.ProcessInstanceKey == instance.GetInstanceKey() && timer.State == TimerCreated {
			t = timer
			break
		}
	}
	return t
}

func findDurationValue(ice BPMN20.TIntermediateCatchEvent, process *ProcessInfo) (duration.Duration, error) {
	durationStr := &ice.TimerEventDefinition.TimeDuration.XMLText
	if durationStr == nil {
		return duration.Duration{}, errors.New(fmt.Sprintf("Can't find 'timeDuration' value for INTERMEDIATE_CATCH_EVENT with id=%s", ice.Id))
	}
	d, err := duration.ParseISO8601(*durationStr)
	return d, err
}

func checkDueTimersAndFindIntermediateCatchEvent(timers []*Timer, intermediateCatchEvents []BPMN20.TIntermediateCatchEvent, instance *ProcessInstanceInfo) *BPMN20.TIntermediateCatchEvent {
	for _, timer := range timers {
		if timer.ProcessInstanceKey == instance.GetInstanceKey() && timer.State == TimerCreated {
			if time.Now().After(timer.DueDate) {
				for _, ice := range intermediateCatchEvents {
					if ice.Id == timer.ElementId {
						return &ice
					}
				}
			}
		}
	}
	return nil
}
