from collections import defaultdict
import sys

type Weight = int
type Vertex = int
type Edge = tuple[Vertex, Vertex, Weight]
type Graph = dict[int, dict[int, int]]


def read_data() -> tuple[list[Edge], tuple[Vertex, Vertex]]:
    read_int_line = lambda: tuple(map(int, sys.stdin.readline().strip().split(" ")))
    _, edge_amount = read_int_line()
    edges = [read_int_line() for _ in range(edge_amount)]
    vertexes = read_int_line()
    return edges, vertexes


def build_graph(edges: list[Edge]) -> Graph:
    graph = defaultdict(dict)
    for first, second, weight in edges:
        graph[first][second] = weight
    return graph


def find_min_distance(graph: Graph, from_vertex: Vertex, to_vertex: Vertex) -> int:
    vertices = set(graph.keys())
    for nbrs in graph.values():
        vertices.update(nbrs.keys())
    vertices.add(from_vertex)
    vertices.add(to_vertex)

    unvisited = set(vertices)
    distances = {vertex: float("inf") for vertex in vertices}
    distances[from_vertex] = 0

    while unvisited:
        current_vertex = min(unvisited, key=lambda v: distances[v])
        unvisited.remove(current_vertex)

        if distances[current_vertex] == float("inf"):
            break

        for neighbor, weight in graph.get(current_vertex, {}).items():
            alternative_distance = distances[current_vertex] + weight
            if alternative_distance < distances[neighbor]:
                distances[neighbor] = alternative_distance

    return distances[to_vertex] if distances[to_vertex] != float("inf") else -1


if __name__ == "__main__":
    edges, vertexes = read_data()
    graph = build_graph(edges)
    result = find_min_distance(graph, *vertexes)
    print(result)
