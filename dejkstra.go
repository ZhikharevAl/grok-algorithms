package main

import (
	"math"
)

// Граф представлен как карта смежности
type Graph map[string]map[string]int

func Dijkstra(graph Graph, start string) map[string]int {
	distances := make(map[string]int)
	visited := make(map[string]bool)

	// Инициализация расстояний
	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	for len(visited) < len(graph) {
		current := minDistanceNode(distances, visited)
		visited[current] = true

		for neighbor, weight := range graph[current] {
			if !visited[neighbor] {
				newDist := distances[current] + weight
				if newDist < distances[neighbor] {
					distances[neighbor] = newDist
				}
			}
		}
	}

	return distances
}

func minDistanceNode(distances map[string]int, visited map[string]bool) string {
	min := math.MaxInt32
	var minNode string

	for node, dist := range distances {
		if !visited[node] && dist < min {
			min = dist
			minNode = node
		}
	}

	return minNode
}
