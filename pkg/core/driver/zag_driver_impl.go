package driver

import (
	"context"
	"fmt"
	config "github.com/bendanwwww/hecate-go/pkg/core/config"
	"github.com/bendanwwww/hecate-go/pkg/core/constant"
	actuatorImpl "github.com/bendanwwww/hecate-go/pkg/core/driver/actuator"
	actuatorApi "github.com/bendanwwww/hecate-go/pkg/core/driver/actuator/api"
	builderImpl "github.com/bendanwwww/hecate-go/pkg/core/driver/builder"
	builderApi "github.com/bendanwwww/hecate-go/pkg/core/driver/builder/api"
	schedulerImpl "github.com/bendanwwww/hecate-go/pkg/core/driver/scheduler"
	schedulerApi "github.com/bendanwwww/hecate-go/pkg/core/driver/scheduler/api"
	"github.com/bendanwwww/hecate-go/pkg/core/entities"
	"github.com/bendanwwww/hecate-go/pkg/framework/common/errors"
	frameWorkConstant "github.com/bendanwwww/hecate-go/pkg/framework/constant"
	driverContext "github.com/bendanwwww/hecate-go/pkg/framework/context"
	"sync"
)

type ZagDriverImpl[T any] struct {
	/** driver config */
	driverConfig config.DriverConfig
	/** all nodes map (key: string value: []zagEntities.NodeWrapper[T]) */
	wrappers sync.Map
	/** builder type */
	builderType constant.BuilderType
	/** scheduler type */
	schedulerType constant.SchedulerType
	/** actuator type */
	actuatorType constant.ActuatorType
	/** map collection (key: string value: zagMapApi.ZagMap[T]) */
	scenesZagMap sync.Map
	/** scheduler */
	zagScheduler schedulerApi.Scheduler[T]
	/** actuator */
	zagActuator actuatorApi.Actuator[T]
}

func NewDefaultZagDriver[T any](driverConfig config.DriverConfig) *ZagDriverImpl[T] {
	return NewZagDriver[T](driverConfig, constant.DefaultBuilder, constant.DefaultScheduler, constant.DefaultActuator)
}

func NewZagDriver[T any](driverConfig config.DriverConfig,
	builderType constant.BuilderType, schedulerType constant.SchedulerType, actuator constant.ActuatorType) *ZagDriverImpl[T] {
	res := &ZagDriverImpl[T]{
		builderType:  builderType,
		driverConfig: driverConfig,
	}
	// choose actuator
	switch actuator {
	case constant.DefaultActuator:
		res.zagActuator = actuatorImpl.NewDefaultActuator[T](driverConfig.MaxConcurrency(), driverConfig.UsePool())
	}
	// choose scheduler
	switch schedulerType {
	case constant.DefaultScheduler:
		res.zagScheduler = schedulerImpl.NewDefaultScheduler[T](res.zagActuator)
	case constant.SimpleScheduler:
		res.zagScheduler = schedulerImpl.NewSimpleZagScheduler[T]()
	}
	return res
}

func (zagDriver *ZagDriverImpl[T]) AddNodes(scenes string, node ...entities.NodeWrapper[T]) {
	for _, item := range node {
		addNode(zagDriver, scenes, frameWorkConstant.DefaultKey, -1, item)
	}
}

func (zagDriver *ZagDriverImpl[T]) AddGroupNodes(scenes string, groupName string, node ...entities.NodeWrapper[T]) {
	zagDriver.AddGroupNodesWithTimeout(scenes, groupName, -1, node...)
}

func (zagDriver *ZagDriverImpl[T]) AddGroupNodesWithTimeout(scenes string, groupName string, groupTimeout int64, node ...entities.NodeWrapper[T]) {
	var nodeList []entities.NodeWrapper[T]
	data, isIn := zagDriver.wrappers.Load(scenes)
	if isIn {
		nodeList = data.([]entities.NodeWrapper[T])
	}
	for _, item := range node {
		needAdd := true
		for index := range nodeList {
			// check the node is added
			if nodeList[index].GetZagNode().GetNodeName() == item.GetZagNode().GetNodeName() {
				nodeList[index].AddGroupNameWithTimeout(groupName, groupTimeout)
				needAdd = false
				break
			}
		}
		if needAdd {
			addNode(zagDriver, scenes, groupName, groupTimeout, item)
		}
	}
}

func (zagDriver *ZagDriverImpl[T]) BuildAll() {
	zagDriver.wrappers.Range(func(key, value any) bool {
		scenes := key.(string)
		nodes := value.([]entities.NodeWrapper[T])
		switch zagDriver.builderType {
		case constant.DefaultBuilder:
			zagDriver.scenesZagMap.Store(scenes, builderImpl.NewDefaultBuilderImpl[T](scenes, nodes))
		}
		return true
	})
}

func (zagDriver *ZagDriverImpl[T]) Build(scenes string) {
	data, isIn := zagDriver.wrappers.Load(scenes)
	if !isIn {
		return
	}
	nodes := data.([]entities.NodeWrapper[T])
	switch zagDriver.builderType {
	case constant.DefaultBuilder:
		zagDriver.scenesZagMap.Store(scenes, builderImpl.NewDefaultBuilderImpl[T](scenes, nodes))
	}
}

func (zagDriver *ZagDriverImpl[T]) Run(ctx context.Context, servingContext *T, scenes string, timeout int64) (chan string, error) {
	data, isIn := zagDriver.scenesZagMap.Load(scenes)
	if !isIn {
		err := errors.NewHecateExceptionWithCode(errors.ScenesNotFound)
		return nil, err
	}
	zagMap := data.(builderApi.Builder[T])
	requestContext := driverContext.NewRuntimeContext[T](ctx, scenes, zagDriver.driverConfig.LogPrint(),
		zagDriver.driverConfig.MustWriteTime(), zagDriver.driverConfig.LogUpload(), zagDriver.driverConfig.SamplingRate(),
		servingContext, zagMap.GetNodeNumber(), timeout)
	return zagDriver.zagScheduler.Run(requestContext, zagMap), nil
}

func (zagDriver *ZagDriverImpl[T]) Clear(scenes string) {
	// delete nodes
	zagDriver.wrappers.Delete(scenes)
	// delete map
	zagDriver.scenesZagMap.Delete(scenes)
}

func (zagDriver *ZagDriverImpl[T]) ToString(scenes string) string {
	data, isIn := zagDriver.scenesZagMap.Load(scenes)
	if !isIn {
		return ""
	}
	zagMap := data.(builderApi.Builder[T])
	return fmt.Sprint(zagMap)
}

func addNode[T any](zagDriver *ZagDriverImpl[T], scenes string, groupName string, groupTimeout int64, node entities.NodeWrapper[T]) {
	node.AddGroupNameWithTimeout(groupName, groupTimeout)
	var nodeList []entities.NodeWrapper[T]
	data, isIn := zagDriver.wrappers.Load(scenes)
	if !isIn {
		nodeList = []entities.NodeWrapper[T]{}
	} else {
		nodeList = data.([]entities.NodeWrapper[T])
	}
	nodeList = append(nodeList, node)
	zagDriver.wrappers.Store(scenes, nodeList)
}

var _ ZagDriver[any] = (*ZagDriverImpl[any])(nil)
