package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// blueprint for movie
type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Director string `json:"director"`
	Year int `json:"year"`
	Genre string `json:"genre"`
	Rating float64 `json:"rating,omitempty"` // omitempty - omit empty, optional field
}

var (
	movies = map[string]Movie{} // mock database using a map, equivalent to py dict and ts array
	mu = sync.Mutex{} // allows multiple requests to be made at the same time
)

func main() {
	router := gin.Default() // equivalent to app = FastAPI (py) and express (ts)

	// register routes
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovie)
	router.POST("/movies", createMovie)
	router.PUT("/movies/:id", updateMovie)
	router.DELETE("/movies/:id", deleteMovie)

	router.Run() // uses default port 8080
}

// gets all movies
func getMovies(c *gin.Context){

	// "locks" resources while function is active, then unlocks when done
	mu.Lock()
	defer mu.Unlock()
	
	list := []Movie{} // equivalent to a list or array
	for _, m := range movies {
		list = append(list, m) // builds list from the map
	}

	c.JSON(http.StatusOK, list)
}

// get single movie based on ID
func getMovie(c *gin.Context){
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	movie, ok := movies[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}
	
	c.JSON(http.StatusOK, movie)
}

// create movie
func createMovie(c *gin.Context) {
	var input Movie
	if err := c.ShouldBindJSON(&input); err != nil { // automatic binding + validatio
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	input.ID = uuid.New().String() // generate unique id
	movies[input.ID] = input
	c.JSON(http.StatusCreated, input) // 201 Created
}

// replace movie details
func updateMovie(c *gin.Context) {
	id := c.Param("id")

	var input Movie
	if err := c.ShouldBindJSON(&input); err != nil { // automatic binding + validation
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := movies[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	input.ID = id // preserve the original id
	movies[id] = input // overwrite the old entry
	c.JSON(http.StatusOK, input)
}


// delete movie
func deleteMovie(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	if _, exists := movies[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	delete(movies, id) // removes key from map based on id
	c.Status(http.StatusNoContent)
}