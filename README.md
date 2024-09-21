# Fastest CEP API Response using Multithreading in Golang
============================================================

## Project Overview
This project aims to implement a Golang program that uses multithreading to fetch the fastest response from two different APIs, BrasilAPI and ViaCEP, for a given CEP (Brazilian postal code). The program will send simultaneous requests to both APIs, accept the response from the fastest API, and discard the slower response.

## API Endpoints
* **BrasilAPI**: `https://brasilapi.com.br/api/cep/v1/{cep}`
* **ViaCEP**: `http://viacep.com.br/ws/{cep}/json/`

## Requirements
* Send simultaneous requests to both APIs using multithreading
* Accept the response from the fastest API and discard the slower response
* Display the response data, including the API that sent the response, in the command line
* Limit the response time to 1 second; if exceeded, display a timeout error

## Implementation
The program will use Golang's built-in `net/http` package to send HTTP requests to both APIs. The `sync` package will be used to implement multithreading and synchronize the responses. A timeout mechanism will be implemented using Golang's `context` package to limit the response time to 1 second.

## Code Structure
* **main.go**: The main program that sends requests to both APIs and handles the responses
* **api.go**: A package that defines the API endpoints and sends HTTP requests
* **utils.go**: A package that provides utility functions for handling responses and errors

## Getting Started
To run the program, simply execute the following command in your terminal:
