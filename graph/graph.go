package main

import "fmt"

type Node struct {
	data int
}

type Graph struct {
	vertices []Node
	edges    []([]Node)
}

func (graph *Graph) adjacent() {

}

func (graph *Graph) neighbours() {

}

func hasVertex(graph *Graph, val int) int {
	for i, value := range graph.vertices {
		if value.data == val {
			return i
		}
	}
	return -1
}

func (graph *Graph) addVertex(val int) {
	if hasVertex(graph, val) == -1 {
		node := Node{data: val}
		graph.vertices = append(graph.vertices, node)
		graph.edges = append(graph.edges, node)
	} else {
		fmt.Println("Graph already has vertex:", val)
	}
}

func (graph *Graph) removeVertex(val int) {
	if index := hasVertex(graph, val); index != -1 {
		if len(graph.vertices) == 1 {
			graph.vertices = nil
		} else {
			graph.vertices = append(graph.vertices[:index], graph.vertices[index+1:]...)
		}
	} else {
		fmt.Printf("Vertex: %d not found", val)
	}

}

func (graph *Graph) addEdge(u int, v int) {
	if hasVertex(graph, u) && hasVertex(graph, v) {
		append(graph.edges[u], v)
		append(graph.edges[v], u)
	}
}

func (graph *Graph) removeEdge() {

}

func (graph *Graph) getVertexValue() {

}

func (graph *Graph) setVertexValue() {

}

func main() {
	graph := Graph{}
	graph.addVertex(3)
	graph.addVertex(5)
	graph.addVertex(7)
	graph.addVertex(7)
	graph.removeVertex(10)
}
