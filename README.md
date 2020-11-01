# WeatherChecker
WeatherChecker is a command-line tool written using [Cobra](https://github.com/spf13/cobra) that you can use to check the weather forecast in your area using a zip code.

## Prerequisites
The only thing you need to make this application work is a FREE account from weatherstack. If you aren't familiar with Weatherstack, it's an API that you can use to get real-time and historical weather data. You can sign up for a free account [here](https://weatherstack.com/)

There are paid options, but you don't need them. The only caveat is you can't use HTTPS for the API call.

## Using weathercheck

### Command
The command used in the weathercheck tool is `checkweather`

```shell
weathercheck is a command-line tool you can use to check real-time weather in your area

Usage:
  weathercheck [command]

Available Commands:
  checkweather A brief description of your command
  help         Help about any command

Flags:
  -a, --apikey          Used to pass in your Weatherstack API key
  -l, --location        Used to pass in a zip code that you want to check the weather for
```

### Install and Run
Install the tool by running `go install`

To use the weathercheck command, run the following:
```go
weathercheck checkweather --apikey your_api_key --location your_zip_code
```