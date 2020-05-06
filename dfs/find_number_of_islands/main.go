package main

import (
	"fmt"
)

var islandMapSize int
var wave int = 0
const (
	water = 0
	land = 1
)

type islandMap struct {
	area int
	visited bool
}

func main() {
	myMap, err := makeMap()
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println("Islands discovered : " ,discoverIslands(myMap))
	return
}

func makeMap() ([][]islandMap, error) {
	
	fmt.Print("Enter thse map size: ")
	_, err := fmt.Scanf("%d\n", &islandMapSize)
	if err != nil {
		return nil, err
	}

	myMap := make([][] islandMap , islandMapSize )
	for i := range myMap {
		myMap[i] = make([] islandMap, islandMapSize)
	}

	for i := 0 ; i < islandMapSize; i++ {
		for j:=0; j < islandMapSize; j++ {
			fmt.Printf("Enter value of [%v][%v] : ",i,j)
			fmt.Scanf("%d\n",&myMap[i][j].area)
		}
	}
	return myMap, nil
}

func discoverIslands(myMap [][]islandMap) (int){
	discovered := 0
	for i := 0; i < islandMapSize ; i++ {
		for j := 0; j < islandMapSize ; j++ {
			if !myMap[i][j].visited {
				size := explore(myMap, i, j)
				if size != 0 {
					fmt.Printf("Discovered island at : %d-%d, size: %d\n",i,j, size)
					discovered += 1
				}
			}
		}
	}
	return discovered
}

func explore (myMap [][]islandMap, i, j int) (int) {
	if i < 0 ||  j < 0 || i >= islandMapSize || j >= islandMapSize || myMap[i][j].visited {
		return 0
	}
	areaExpanded := 0
	if myMap[i][j].area == land {
		myMap[i][j].visited = true
		areaExpanded = 1
		areaExpanded += explore(myMap,i+1,j)
		areaExpanded += explore(myMap,i,j+1)
		areaExpanded += explore(myMap,i-1,j)
		areaExpanded += explore(myMap,i,j-1)
		areaExpanded += explore(myMap,i+1,j-1)
		areaExpanded += explore(myMap,i-1,j+1)
		areaExpanded += explore(myMap,i-1,j-1)
		areaExpanded += explore(myMap,i+1,j+1)
	}
	return areaExpanded
} 