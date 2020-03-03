package main

import (
    "bytes"
    "fmt"
    "math/rand"
    "os"
    "os/exec"
    "time"
)

func main(){
    //Setting the Seed of Rand to get different values each time
    rand.Seed(time.Now().UTC().UnixNano())
    //Grid Dimensions
    rows:= 40
    cols:= rows * 2
    // Create a grid of "cols" columns & "rows" rows
    grid:= CreatGrid(cols,rows)
    //Populate the grid randomly with 0 & 1
    PopulateGrid(grid)  
    //Print the grind on screen
    //number of iterations
    n:=500
    
    for i:=0; i<n; i++{
        //Clear console
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
        //Generate the new grid
        grid = GenerateNextGeneration(grid)
        PrintFancyGrid(grid)
        time.Sleep(time.Second / 30)
    }
}

func CreatGrid(cols, rows int) [][]int{
    grid:= make([][]int, rows)
    for i:= range grid{
        grid[i]= make([]int, cols)
    }
    return grid
}

func PopulateGrid(grid [][]int){
    for i:= range grid{
        for j:=range grid[i]{
            grid[i][j] = rand.Intn(2)
        }
    }
}

//Print the grid as * and spaces rather then 1 and 0
func PrintFancyGrid(grid [][]int) {
    var buf bytes.Buffer // it's a slice of strings used to store strings and then print them
    for i:= range grid{
        for j:= range grid[i]{
            b:= byte(' ') // initializing this byte as originally empty "a space"
            if grid[i][j] == 1 {
                b='*' // if a cell has a value of one change the byte to be * rather then a space
            }
            buf.WriteByte(b) // add the byte to the buffer

        }
        buf.WriteByte('\n') // once finished iterating throw every colons of one row, add a line break to the buffer
    }
    fmt.Printf(buf.String()) //Ones the buffer is ready print it
}

// Count the sum of the neighbours of a cell with coordinates (x,y)
func CountNeighboursOfCell(grid [][]int, x, y int) int{
    sum:= 0 //initialize the sum
    rows:= len(grid) //the numbers of rows in the grid
    cols:= len(grid[1]) //the numbers of colons in the grid (using the index 1 assuming all rows have the same length)
    //checking the 8 neighbours of a cell relative to it's position (x,y)
    for i:=-1; i<2; i++{
         for j:=-1; j<2; j++{
            //Count the neighbours of cell (x,y)
            sum += grid[(x+i+rows)%rows][(y+j+cols)%cols]
         }
    } 
    sum-= grid[x][y] // remove the cell itself's value from the sum so it's not counted
    
    return sum    
}

//Generates the next generation of cell based on the game of life rules
//takes a grid (old) and returns a grid (new)
func GenerateNextGeneration(oldGrid [][]int) [][]int{
    //Create a new grid with the size of the old grid
    newGrid:= CreatGrid(len(oldGrid[1]),len(oldGrid))
    for i:= range oldGrid{
         for j:=range oldGrid[i]{
            // count the neighbours of cell (i,j)
            neighbours:= CountNeighboursOfCell(oldGrid,i,j)
            currentCell:= oldGrid[i][j]

            //Game of life rules:

            //Any dead cell with exactly three live neighbours becomes a live cell.
            if((currentCell == 0 && neighbours == 3)){
                newGrid[i][j]=1
            //Any live cell with fewer than two or more than three live neighbours dies.
            }else if(currentCell == 1 && (neighbours>3 || neighbours <2)){
                newGrid[i][j]=0
                //println("i:",i,"j:",j)
            //Any live cell with two or three live neighbours lives on to the next generation.
            //Any dead cell with more or less then three live neighbours stays dead.
            }else{
                newGrid[i][j]=oldGrid[i][j]
            }

         }
    }   
    return newGrid      
}

// ----------------------- OBSOLETE ----------------------- //

//Print a grid with 1 & 0
func PrintGrid(grid [][]int){
    //println("\n")
    for i:= range grid{
        //print the grid on screen
        fmt.Printf("%v\n",grid[i])
    }
}