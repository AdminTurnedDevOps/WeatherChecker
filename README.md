# WeatherChecker
WeatherChecker is a command-line tool written using [Cobra](https://github.com/spf13/cobra) that you can use to check the weather forecast in your area using a zip code.

## Prerequisites
The only thing you need to make this application work is a FREE account from weatherstack. If you aren't familiar with Weatherstack, it's an API that you can use to get real-time and historical weather data. You can sign up for a free account [here](https://weatherstack.com/)

There are paid options, but you don't need them. The only caveat is you can't use HTTPS for the API call.

## Using weathercheck

Install the tool by running `go install`

To use the weathercheck command, run the following:
```go
weathercheck checkweather --apikey your_api_key --location your_zip_code
```