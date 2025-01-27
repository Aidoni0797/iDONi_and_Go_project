Trial link ended, why?

Link my first project with Go: https://idoni-and-go-project.fly.dev/

Project Description: 

Weather App using Go
This is my first project using the Go programming language. It is a simple weather application that allows users to check the current weather for a specific city by using the OpenWeatherMap API. The project demonstrates how to interact with an external API, process user input from an HTML form, and dynamically display the results on a web page.

Features:
Users can enter the name of a city into an input field.
The application fetches weather data from the OpenWeatherMap API.
Weather information is displayed on the same page without reloading.

Technologies Used:
Programming Language: Go (Golang)
API: OpenWeatherMap
Frontend: HTML, CSS
Server Framework: net/http (built-in Go package)

How It Works:
The user opens the HTML page and enters a city name in the input field.
The form sends a request to the Go backend server.
The Go server processes the request, makes an API call to OpenWeatherMap, and retrieves weather data.
The retrieved weather information (e.g., temperature, weather condition) is displayed dynamically below the input field.

How to Run the Project:
Clone the repository:
git clone (https://github.com/Aidoni0797/iDONi_and_Go_project)
cd your-repository

Set up your OpenWeatherMap API key:
Sign up at OpenWeatherMap.
Copy your API key and add it to the Go code where indicated.

Run the Go server:
go run main.go
Open the application in your browser at http://localhost:8080.

Future Improvements:
Add support for more detailed weather data, such as forecasts or wind speed.
Implement error handling for invalid city names or API errors.
Make the design more user-friendly with advanced styling.

Feel free to adapt or modify this explanation as needed! 😊
