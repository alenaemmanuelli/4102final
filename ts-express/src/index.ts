import express, { Request, Response } from "express"; // import express and request and response

const app = express(); // create express app, similar to Python FastAPI

app.use(express.json()); // allows express to read JSON

// equivalent to movie class in Python
interface Movie {
    id: string;
    title: string;
    director: string;
    year: number;
    genre: string;
    rating?: number;
}

const moviesDB: Movie[] = []; // equivalent to movies_db in Python

// helper function to generate unique ID
function generateID(): string {
    return Math.random().toString(36).substring(2,10);
}

// gets all movies
app.get("/movies", (req: Request, res: Response) => {
    res.json(moviesDB);
});

// get single movie based on ID
app.get("/movies/:id", (req: Request, res: Response) => {
  const movie = moviesDB.find((m) => m.id === req.params.id);
  if (!movie) { // manual error handling
    res.status(404).json({ error: "movie not found" });
    return;
  }
  res.json(movie);
});

// create movie
app.post("/movies", (req: Request, res: Response) => {
    const {
        title,
        director,
        year,
        genre,
        rating,
    } = req.body;

    // manual validation
    if (!title || !director || !year || !genre) {
    res.status(400).json({ error: "title, director, year, and genre are required" });
    return;
  }

  const newMovie: Movie = {
    id: generateID(),
    title,
    director,
    year,
    genre,
    rating,
  };
 
  moviesDB.push(newMovie);
  res.status(201).json(newMovie);
});

// replace movie details
app.put("/movies/:id", (req: Request, res: Response) => {
  const id = req.params.id;

  if (Array.isArray(id)) {
    res.status(400).json({ error: "invalid movie id" });
    return;
  }

  const index = moviesDB.findIndex((m) => m.id === id);

  if (index === -1) {
    res.status(404).json({ error: "movie not found" });
    return;
  }

  const {
    title,
    director,
    year,
    genre,
    rating,
  } = req.body;

  // manual validation
  if (!title || !director || !year || !genre) {
    res.status(400).json({ error: "title, director, year, and genre are required" });
    return;
  }

  moviesDB[index] = {
    id,
    title,
    director,
    year,
    genre,
    rating,
  };

  res.json(moviesDB[index]);
});

// delete movie
app.delete("/movies/:id", (req: Request, res: Response) => {
  const index = moviesDB.findIndex((m) => m.id === req.params.id);

  if (index === -1){
    res.status(404).json({ error: "movie not found" });
    return;
  }

  // remove movie at given index
  moviesDB.splice(index, 1);
  res.status(204).send();
});

// forces server to start on port 3000
app.listen(3000, () => {
  console.log("Server running on: http://localhost:3000");
});