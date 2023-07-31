
# Weather App in Go

A simple command-line weather app written in Go that fetches current weather data for a given city using the OpenWeatherMap API.

## Features

- Fetches current weather data (temperature and weather description) for a specified city.
- Uses the OpenWeatherMap API to get real-time weather information.
- Provides an easy-to-use command-line interface for quick weather checks.

## Installation

1. Make sure you have Go installed on your system. If not, download and install it from the official website: [https://golang.org/](https://golang.org/)

2. Clone this repository or download the `main.go` file.

3. Sign up for a free API key at [https://openweathermap.org/api](https://openweathermap.org/api).

4. Open the `main.go` file in a text editor and replace `"..."` with your actual API key obtained from OpenWeatherMap.

## Getting an API Key

To use this weather app, you'll need to obtain an API key from OpenWeatherMap. The API key is required to make requests and fetch weather data from their servers.

Follow the steps below to get your API key:

1. **Sign Up on OpenWeatherMap**: Go to the OpenWeatherMap website at [https://openweathermap.org/](https://openweathermap.org/) and click on the "Sign Up" button. Fill in the required details to create an account.

2. **Verify Your Email**: After signing up, you will receive an email from OpenWeatherMap to verify your email address. Click on the verification link in the email to complete the registration.

3. **Log In to OpenWeatherMap**: Once your email is verified, log in to your OpenWeatherMap account.

4. **Subscribe to the API**: After logging in, go to the "API Keys" section in your account dashboard. Select the type of API you want to use; for our weather app, choose the "Current Weather Data" API.

5. **Choose a Plan**: Select a plan that suits your needs. OpenWeatherMap offers both free and paid plans with varying levels of access and features. The free plan usually has some limitations on the number of requests you can make per minute or per day.

6. **Generate an API Key**: Once you've chosen a plan, click on the "Generate" button to create your API key. The API key will be displayed on the screen, and you will also receive an email with your API key.

7. **Copy Your API Key**: Copy the API key and keep it safe. You will need to use this key in your weather app to make API requests.

## Usage

In your terminal or command prompt, navigate to the directory where the `main.go` file is located.

To run the weather app and check the weather for a specific city, use the following command:

```bash
go run main.go <city_name>
```

Replace `<city_name>` with the name of the city you want to check the weather for.

Example:

```bash
go run main.go Mostar
```

The app will display the weather information for the specified city, including the temperature and a brief description of the weather conditions.
