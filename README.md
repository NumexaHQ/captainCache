 <img src="https://github.com/NumexaHQ/captainCache/assets/28846178/4f786b2b-dd2c-44b6-b6cd-7ac46a933461" alt="image" width="250" height="250">


## captainCache

Proxy for the Generative AI application using OpenAI, Use it for cahing the prompts and monitoring the number of API calls. 

## Prerequisites

- Go installed on your machine
- Redis installed and running locally or accessible via a Redis server

## Installation

1. Clone this repository.

2. Change into the Project directory.
3. Install the project dependencies.
   ```
   go mod download
   ```
4. Set up the Redis connection details in the main.go file
   ```
   redisClient = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // Enter your Redis password if applicable
    DB:       0,  // Use default Redis database
    })
   ```
## Usage
- start the Proxy server using
  ```
  go run main.go
  ```
- Make HTTP requests via Proxy
  ```
  GET http://localhost:8080/?api_key=<YOUR_API_KEY>&prompt=Hello%20World
  ```
