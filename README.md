# 🌦️ Weather API (GraphQL + Go + Redis)

A GraphQL-based weather API built with Go that fetches real-time and historical weather data from [Visual Crossing](https://www.visualcrossing.com/resources/documentation/weather-api/timeline-weather-api/). The API supports querying by city and optional date, caches responses using Redis, and applies rate limiting using API keys to prevent abuse.

---

## 🚀 Features

- 🔍 **GraphQL API** using `gqlgen`
- 🌦️ Fetch weather data by **city** and optional **date** (defaults to today)
- ⚡ **Caching with Redis** to avoid redundant external API calls
- 🛡️ **Rate Limiting** based on API keys to prevent abuse
- 📆 **Query by City & Date** (defaults to current date if not provided)
- 🐳 Docker-based Redis setup
- 🛠️ Developer-friendly with VSCode debug support

---

## 📦 Prerequisites

- ✅ Go installed ([Download Go](https://go.dev/dl/))
- ✅ Docker installed ([Install Docker](https://docs.docker.com/get-docker/))
- ✅ Redis running locally (or via Docker)
- ✅ Visual Crossing API Key from: [https://www.visualcrossing.com](https://www.visualcrossing.com)

---

## 🔧 Technologies Used

- [Golang](https://go.dev/)
- [gqlgen](https://gqlgen.com/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/)
- [Visual Crossing Weather API](https://www.visualcrossing.com/weather-api)
- GraphQL

---

## 🖼️ Architecture Diagram

![Architecture Diagram](architecture.png)

---

## 📦 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Agarwalsahil/weather-api.git
cd weather-api
```

### 2. Install Go dependencies

```bash
go mod tidy
```

### 3. Generate GraphQL code (from schema)

```bash
go run github.com/99designs/gqlgen generate
```

### 4. Set the Environment Variable

#### 💻 For local run (Linux/macOS):
```bash
# 💻 On Linux/macOS:
export WEATHER_API_KEY=your_key_here

# 🪟 On Windows (cmd):
set WEATHER_API_KEY=your_key_here
```

---

## ▶️ Running the Project

```bash
go run server.go
```

> Server starts on: [http://localhost:8080](http://localhost:8080)

---

#### 🧠 For Debugging in VSCode (`launch.json`):

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "weather-api",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/server.go",
      "env": {
        "WEATHER_API_KEY": "Your Api Key"
      },
      "args": []
    }
  ]
}
```

---

## 🧪 Sample GraphQL Query

```graphql
query {
  getWeatherData(city: "New Delhi") {
    latitude
    longitude
    resolvedAddress
    days {
      temp
      tempmax
      windgust
      windspeed
    }
    isProcessed
  }
}
```

---

## 🗃️ Redis Caching Strategy

- **Key**: `${city}:${date}`
- Checks Redis cache first.
- If not found → Fetch from external API → Store result in cache → Return response.

---

## 🔐 Rate Limiting

- API key must be passed in request headers:
  
  ```json
  {
    "X-API-Key": "your-api-key"
  }
  ```

- Requests without valid API key are denied.
- Rate limits configurable (e.g., max 5 req/min per key).

---

## 🐳 Docker (for Redis)

```bash
docker run --name weather-api-redis -p 6379:6379 -d redis
```

---

## 🧠 Contributing

Pull requests are welcome! For major changes, open an issue first to discuss what you'd like to change.

---

## 👨‍💻 Author

Built with ❤️ using Go and GraphQL  
By [Sahil Agarwal]