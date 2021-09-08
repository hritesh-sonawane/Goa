package main

import "fmt"

type Dep struct {
	Prereq int
	Job    int
}

type JobGraph struct {
	Nodes []*JobNode
	Graph map[int]*JobNode
}

type JobNode struct {
	Job          int
	Deps         []*JobNode
	NumOfPrereqs int
}

// O(j+d) time | O(j+d) space | j: jobs | d: dependencies
func TopologicalSort(jobs []int, deps []Dep) []int {
	jobGraph := createJobGraph(jobs, deps)
	return getOrderedJobs(jobGraph)
}

func createJobGraph(jobs []int, deps []Dep) *JobGraph {
	graph := NewJobGraph(jobs)
	for _, dep := range deps {
		graph.AddDep(dep.Prereq, dep.Job)
	}
	return graph
}

func getOrderedJobs(graph *JobGraph) []int {
	orderedJobs := []int{}
	nodesWithNoPrereqs := []*JobNode{}
	for _, node := range graph.Nodes {
		if node.NumOfPrereqs == 0 {
			nodesWithNoPrereqs = append(nodesWithNoPrereqs, node)
		}
	}
	for len(nodesWithNoPrereqs) > 0 {
		node := nodesWithNoPrereqs[len(nodesWithNoPrereqs)-1]
		nodesWithNoPrereqs = nodesWithNoPrereqs[:len(nodesWithNoPrereqs)-1]
		orderedJobs = append(orderedJobs, node.Job)
		removeDeps(node, &nodesWithNoPrereqs)
	}
	for _, node := range graph.Nodes {
		if node.NumOfPrereqs > 0 {
			return []int{}
		}
	}
	return orderedJobs
}

func removeDeps(node *JobNode, nodesWithNoPrereqs *[]*JobNode) {
	for len(node.Deps) > 0 {
		dep := node.Deps[len(node.Deps)-1]
		node.Deps = node.Deps[:len(node.Deps)-1]
		dep.NumOfPrereqs--
		if dep.NumOfPrereqs == 0 {
			*nodesWithNoPrereqs = append(*nodesWithNoPrereqs, dep)
		}
	}
}

func NewJobGraph(jobs []int) *JobGraph {
	g := &JobGraph{
		Graph: map[int]*JobNode{},
	}
	for _, job := range jobs {
		g.AddNode(job)
	}
	return g
}

func (g *JobGraph) AddDep(job, dep int) {
	JobNode, depNode := g.GetNode(job), g.GetNode(dep)
	JobNode.Deps = append(JobNode.Deps, depNode)
	depNode.NumOfPrereqs++
}

func (g *JobGraph) AddNode(job int) {
	g.Graph[job] = &JobNode{Job: job}
	g.Nodes = append(g.Nodes, g.Graph[job])
}

func (g *JobGraph) GetNode(job int) *JobNode {
	if _, found := g.Graph[job]; !found {
		g.AddNode(job)
	}
	return g.Graph[job]
}

func main() {
	jobs := []int{1, 2, 3, 4}
	deps := []Dep{{1, 2}, {1, 3}, {3, 2}, {4, 2}, {4, 3}}
	fmt.Println(TopologicalSort(jobs, deps))
}