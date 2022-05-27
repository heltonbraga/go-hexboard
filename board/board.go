package board

import (
	"math"
	"sort"
)

type Board struct {
	Hex []Hexagono
}

type Hexagono struct {
	Id    int
	Coord Coordenada
	//Arestas [6]Aresta
}

type Coordenada struct {
	X int
	Y int
	Z int
}

type Aresta struct {
	Id      int
	Origem  int
	Destino int
}

func makeAresta(id, origem, destino int) *Aresta {
	aresta := new(Aresta)
	aresta.Id = id
	aresta.Origem = origem
	aresta.Destino = destino
	return aresta
}

func makeCoordenada(x, y, z int) *Coordenada {
	coord := new(Coordenada)
	coord.X = x
	coord.Y = y
	coord.Z = z
	return coord
}

func GetMatrix(board *Board) [][]int {
	tamBoard := math.Floor(float64(len(board.Hex) / 6))
	tamMatrix := tamBoard + 2
	offset := math.Floor(tamMatrix / 2)
	pos := 0
	var matrix = make([][]int, int(tamMatrix))
	sort.Slice(board.Hex, func(i, j int) bool { return board.Hex[i].Id < board.Hex[j].Id })

	for i := 0; float64(i) < tamMatrix; i++ {
		matrix[i] = make([]int, int(tamMatrix))
		for j := math.Max(0, offset); j < math.Min(tamMatrix, tamMatrix+offset); j++ {
			matrix[i][int(j)] = board.Hex[pos].Id
			pos++
		}
		offset--
	}

	return matrix
}

var idAresta = 0

func getIdAresta() int {
	idAresta++
	return idAresta
}

var idVertice = 0

func getIdVertice() int {
	idVertice++
	return idVertice
}

var idHexagono = 0

func getIdHexagono() int {
	idHexagono++
	return idHexagono
}

/*
func makeHexagono(id int, coord Coordenada, arestas [6]Aresta) *Hexagono {
	hex := new(Hexagono)
	hex.Id = id
	hex.Coord = coord
	hex.Arestas = arestas
	return hex
}
*/
func makeHexagono(id int, coord Coordenada) *Hexagono {
	hex := new(Hexagono)
	hex.Id = id
	hex.Coord = coord
	return hex
}

func MakeBoard(layers int) *Board {
	if layers < 1 {
		return nil
	}
	board := new(Board)
	addLayers(board, layers)
	return board
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addLayers(board *Board, layers int) {
	n := layers - 1
	for i := 0 - n; i <= n; i++ {
		for j := max(0-n, 0-i-n); j <= min(n, 0-i+n); j++ {
			board.Hex = append(board.Hex, *makeHexagono(getIdHexagono(), *makeCoordenada(i, j, 0-i-j)))
		}
	}
}
