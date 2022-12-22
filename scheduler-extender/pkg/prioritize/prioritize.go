package prioritize

import (
	"math/rand"
	"time"

	extenderv1 "k8s.io/kube-scheduler/extender/v1"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Prioritize(args extenderv1.ExtenderArgs) extenderv1.HostPriorityList {
	hostPriority := make(extenderv1.HostPriorityList, 0)

	for _, node := range args.Nodes.Items {
		hostPriority = append(hostPriority, extenderv1.HostPriority{
			Host:  node.Name,
			Score: rand.Int63n(extenderv1.MaxExtenderPriority + 1),
		})
	}

	return hostPriority
}
