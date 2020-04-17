boolean true
number -273.15
string "She said \"Hello, BF\""
array ["gold", "silver", "bronze"]
object {"year": 1980,
"event": "archery",
"medals": ["gold", "silver", "bronze"]}


type Movie struct {
	Title string
	Year  int
	Color bool
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false
		Actors: []string{"Humprey Bogart", "Ingrid Bergman"}}
}


data, err := json.Marhsal(movies)
if err != nil {
	log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)



