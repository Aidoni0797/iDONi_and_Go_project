package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const apiKey = "31621498deb5472106162e10ff88dc44"

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func getWeather(city string) (WeatherResponse, error) {
	var weather WeatherResponse
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return weather, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return weather, err
	}

	err = json.Unmarshal(body, &weather)
	return weather, err
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	weather, err := getWeather(city)
	if err != nil {
		http.Error(w, "Unable to fetch weather data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "The temperature in %s is %.2f°C\n", weather.Name, weather.Main.Temp)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
    <html>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f0f8ff;
				margin: 0;
				padding: 0;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
			}
			.container {
				background-color: white;
				padding: 20px;
				border-radius: 10px;
				box-shadow: 0px 0px 15px rgba(0, 0, 0, 0.1);
				width: 400px;
				text-align: center;
			}
			h2 {
				color: #1E90FF;
			}
			input[type="text"] {
				padding: 10px;
				width: 80%;
				margin: 10px 0;
				border: 1px solid #ccc;
				border-radius: 5px;
				font-size: 16px;
			}
			input[type="submit"] {
				padding: 10px 20px;
				background-color: #1E90FF;
				border: none;
				color: white;
				font-size: 16px;
				border-radius: 5px;
				cursor: pointer;
			}
			input[type="submit"]:hover {
				background-color: #45a049;
			}
			#weatherResult {
				margin-top: 20px;
				padding: 10px;
				background-color: #f1f1f1;
				border-radius: 5px;
				display: inline-block;
				min-width: 100%;
				word-wrap: break-word;
			}
		</style>
      <body>
        <form action="/weather" method="get">
		  <h2>iDONi</h2>
		  is creating their first project using the Go programming language.<br><hr>
          <label for="city">Please, enter your city:</label>
          <input type="text" id="city" name="city" required>
          <input type="submit" value="Get Weather">
        </form>
      </body>
    </html>
    `
	fmt.Fprint(w, html)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/weather", weatherHandler).Methods("GET")
	http.Handle("/", r)

	port := "8080"
	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
