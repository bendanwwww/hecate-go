package impl

import (
	builderApi "github.com/bendanwwww/hecate-go/pkg/core/driver/builder/api"
	"github.com/bendanwwww/hecate-go/pkg/core/driver/scheduler/api"
	"github.com/bendanwwww/hecate-go/pkg/core/entities"
	"github.com/bendanwwww/hecate-go/pkg/framework/common/log"
	"github.com/bendanwwww/hecate-go/pkg/framework/constant"
	driverContext "github.com/bendanwwww/hecate-go/pkg/framework/context"
	"github.com/bendanwwww/hecate-go/pkg/framework/tools"
)

type SimpleZagSchedulerImpl[T any] struct {
}

func NewSimpleZagScheduler[T any]() *SimpleZagSchedulerImpl[T] {
	return &SimpleZagSchedulerImpl[T]{}
}

func (d *SimpleZagSchedulerImpl[T]) Run(ctx *driverContext.RuntimeContext[T], builder builderApi.Builder[T]) chan string {
	// 已执行节点列表
	runNodeArray := make([]int, builder.GetNodeNumber())
	// 实时入度
	inDegree := tools.BuildNewBigBinary(builder.GetEdgeNumber())
	// 模拟栈
	var nodeStack [][]int
	// 遍历头结点
	headNodes := builder.GetHeads()
	for _, nodeIndex := range headNodes {
		node := builder.GetNodeByIndex(nodeIndex)
		logContext := driverContext.NewLogContextWithFirstNodeInfo(ctx.Context,
			ctx.RequestId, ctx.ScenesName, ctx.ExecutionStartTime,
			nodeIndex, node.GetNodeName(), ctx.NeedPrintLog, ctx.NeedWriteSimpleLog, ctx.MustWriteTime, &ctx.LogBuffer)
		log.DefaultNodeLog.Infof(logContext, "access node %s", node.GetNodeName())
		log.DefaultNodeLog.Infof(logContext, "node %s start run", node.GetNodeName())
		node.GetNodeOperator().Execute(logContext, ctx.ServingContext)
		// 随机模拟结果
		logContext.NodeSchedulerState = randomSchedulerResult()
		logContext.NodeActuatorState = randomExecuteResult()
		runNodeArray[nodeIndex] = 1
		log.DefaultNodeLog.Infof(logContext, "node %s end run", node.GetNodeName())
		// 下游节点入栈
		if builder.HasNextNode(nodeIndex) {
			for _, nextNode := range builder.GetNextNodeByIndex(nodeIndex) {
				// 节点下标
				nextNodeIndex := nextNode[0]
				// 边下标
				edgeIndex := nextNode[1]
				// 修改入度
				inDegree.ChangeBitToTrue(edgeIndex)
				// 创建调度上下文
				stackContext := []int{nodeIndex, nextNodeIndex}
				nodeStack = append(nodeStack, stackContext)
			}
		}
	}
	// 模拟递归
	for {
		if len(nodeStack) <= 0 {
			break
		}
		stackContext := nodeStack[len(nodeStack)-1]
		fromNodeIndex := stackContext[0]
		nodeIndex := stackContext[1]
		nodeStack = nodeStack[0 : len(nodeStack)-1]
		fromNode := builder.GetNodeByIndex(fromNodeIndex)
		node := builder.GetNodeByIndex(nodeIndex)
		logContext := driverContext.NewLogContext(ctx.Context,
			ctx.RequestId, ctx.ScenesName, ctx.ExecutionStartTime,
			nodeIndex, node.GetNodeName(), fromNodeIndex, fromNode.GetNodeName(),
			ctx.NeedPrintLog, ctx.NeedWriteSimpleLog, ctx.MustWriteTime, &ctx.LogBuffer)
		log.DefaultNodeLog.Infof(logContext, "access node %s", node.GetNodeName())
		// 节点实时入度是否满足调度条件
		if !builder.NodeCanRun(nodeIndex, inDegree) {
			log.DefaultNodeLog.Infof(logContext, "node %s is not ready", node.GetNodeName())
			continue
		}
		// 调度
		if runNodeArray[nodeIndex] != 1 {
			if node.GetNodeName() == constant.EndKey {
				log.DefaultNodeLog.Infof(logContext, "node %s run", constant.EndKey)
			} else {
				log.DefaultNodeLog.Infof(logContext, "node %s start run", node.GetNodeName())
				node.GetNodeOperator().Execute(logContext, ctx.ServingContext)
				// 随机模拟结果
				logContext.NodeSchedulerState = randomSchedulerResult()
				logContext.NodeActuatorState = randomExecuteResult()
				log.DefaultNodeLog.Infof(logContext, "node %s end run", node.GetNodeName())
			}
			runNodeArray[nodeIndex] = 1
		} else {
			log.DefaultNodeLog.Infof(logContext, "node %s is already run", node.GetNodeName())
			continue
		}
		// 下游节点入栈
		if builder.HasNextNode(nodeIndex) {
			for _, nextNode := range builder.GetNextNodeByIndex(nodeIndex) {
				// 节点下标
				nextNodeIndex := nextNode[0]
				// 边下标
				edgeIndex := nextNode[1]
				// 修改入度
				inDegree.ChangeBitToTrue(edgeIndex)
				// 创建调度上下文
				nextStackContext := []int{nodeIndex, nextNodeIndex}
				nodeStack = append(nodeStack, nextStackContext)
			}
		}
	}
	return nil
}

func randomSchedulerResult() string {
	rand := tools.RandIntN(100)
	if rand%8 == 0 {
		return "Skip"
	}
	return "Done"
}

func randomExecuteResult() string {
	rand := tools.RandIntN(100)
	if rand%8 == 0 {
		return string(entities.NodeTimeout)
	}
	if rand%8 == 0 {
		return string(entities.NodeSkip)
	}
	if rand%8 == 0 {
		return string(entities.NodeDefaultError)
	}
	return string(entities.NodeSuccess)
}

var _ api.Scheduler[any] = (*SimpleZagSchedulerImpl[any])(nil)
