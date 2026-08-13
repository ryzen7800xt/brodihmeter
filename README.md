# Brodihmeter

A fun and simple **Meanness Counter** web application to track instances of meanness in real-time. Built with Go backend and an interactive web frontend with heatmap visualization. It's purpose is to ragebait my friend, Brodie.

## ✨ Features

- **Real-time Counter**: Track meanness instances with a live counter display
- **Increment & Reset**: Easy-to-use buttons to add to the counter or reset to zero
- **History Tracking**: Maintains a complete timestamp-based history of all count changes
- **Heatmap Visualization**: Visual representation of counter activity intensity
- **REST API**: Full-featured JSON API for programmatic access
- **Thread-Safe**: Uses mutex locks to ensure data consistency under concurrent access
- **Static File Server**: Serves HTML/CSS/JS files automatically from the application directory

##  Quick Start

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher
- Any modern web browser

### Installation & Running

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/brodihmeter.git
   cd brodihmeter/src
   ```

2. **Build the application** (optional, or download the pre-built executable):
   ```bash
   go build -o brodihmeter.exe main.go
   ```

3. **Run the server**:
   ```bash
   ./brodihmeter.exe
   # or
   go run main.go
   ```

4. **Open in browser**:
   Navigate to `http://localhost:8080` in your web browser

The server will start on **port 8080** and display:
```
Brodih Program
Starting server on port 8080...
```

## Usage

### Web Interface

1. **View Counter**: The current "meanness" count displays prominently
2. **Increment**: Click the "+ Increment" button to add 1 to the counter
3. **Reset**: Click the "Reset" button to clear the counter and history back to zero
4. **Heatmap**: Visual heatmap shows activity patterns over time

### API Endpoints

All endpoints return JSON responses.

#### Get Current Count
```
GET /api/count
```
**Response**:
```json
{
  "meanness": 42
}
```

#### Increment Counter
```
POST /api/increment
```
**Response**:
```json
{
  "meanness": 43
}
```

#### Reset Counter
```
POST /api/reset
```
**Response**:
```json
{
  "meanness": 0
}
```

#### Get History
```
GET /api/history
```
**Response**:
```json
{
  "entries": [
    {
      "count": 1,
      "timestamp": "2024-01-15T10:30:45.123456Z"
    },
    {
      "count": 2,
      "timestamp": "2024-01-15T10:31:20.654321Z"
    }
  ]
}
```

## 🏗️ Project Structure

```
brodihmeter/
├── README.md          # This file
├── LICENSE            # Apache 2.0 License
└── src/
    ├── main.go        # Go backend server
    ├── index.html     # Web interface
    ├── test.html      # Testing page
    └── brodihmeter.exe # Compiled executable
```

## 🛠️ Building from Source

```bash
cd src
go build -o brodihmeter.exe main.go
```

The resulting executable will be ready to run.

##  Example API Usage

### Using cURL

```bash
# Increment the counter
curl -X POST http://localhost:8080/api/increment

# Get current count
curl http://localhost:8080/api/count

# Get history
curl http://localhost:8080/api/history

# Reset counter
curl -X POST http://localhost:8080/api/reset
```

### Using JavaScript/Fetch

```javascript
// Increment counter
fetch('/api/increment', { method: 'POST' })
  .then(res => res.json())
  .then(data => console.log('Current count:', data.meanness));

// Get current count
fetch('/api/count')
  .then(res => res.json())
  .then(data => console.log('Meanness level:', data.meanness));
```

##  Thread Safety

The application uses mutex locks to ensure thread-safe operations on the counter and history. Multiple concurrent requests are handled safely without data corruption.

##  License

This project is licensed under the **Apache License 2.0** - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to submit issues and enhancement requests!

---

**Built with Go** 🐹
