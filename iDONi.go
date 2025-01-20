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

	// Формируем HTML-ответ с использованием CSS
	response := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Weather</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 0;
				padding: 0;
				background: linear-gradient(to bottom, #87CEFA, #4682B4);
				color: #fff;
				text-align: center;
			}
			.container {
				padding: 20px;
			}
			.card {
				background-color: rgba(255, 255, 255, 0.1);
				border-radius: 10px;
				padding: 20px;
				box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
				display: inline-block;
				max-width: 90%;
				width: 300px;
			}
			h1 {
				font-size: 2rem;
				margin-bottom: 20px;
			}
			p {
				font-size: 1.2rem;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="card">
				<h1>Weather in %s</h1>
				<p>The temperature is %.2f°C</p>
			</div>
		</div>
	</body>
	</html>
	`, weather.Name, weather.Main.Temp)

	// Устанавливаем Content-Type для HTML-ответа
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
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
				padding: 15px;
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
