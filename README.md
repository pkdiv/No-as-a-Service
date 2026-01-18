# No-as-a-Service (NaaS)

A simple API built in Go that provides various ways to say "No". Perfect for those times when you need to decline a request but want to do it with a specific flair—be it corporate, sarcastic, or over-dramatic.

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.25.5 or later recommended)

### Running the Server

1.  **Clone the repository** (if you haven't already):
    ```bash
    git clone <repository-url>
    cd No-as-a-Service
    ```

2.  **Start the server**:
    ```bash
    go run main.go
    ```
    By default, the server runs on `http://localhost:8080`.

## Configuration

You can configure the port the server listens on using the `PORT` environment variable.

**Example: Running on port 3000**
```bash
PORT=3000 go run main.go
```

## API Usage

### 1. List All Categories
To see all available "No" categories, make a GET request to the root `/no/` endpoint.

**Request:**
```bash
curl http://localhost:8080/no/
```

**Response:**
```json
{
  "categories": [
    "random",
    "corporate_speak",
    "friendly_firm",
    "funny_light",
    "over_dramatic",
    "polite_professional",
    "passive_aggressive",
    "sarcastic",
    "tech_nerd"
  ]
}
```

### 2. Get a "No" Message
To get a random message from a specific category, append the category name to the `/no/` endpoint.

**Request (Corporate Speak):**
```bash
curl http://localhost:8080/no/corporate_speak
```

**Response:**
```json
{
  "message": "We appreciate the proposal, however, it doesn't align with our current bandwidth and strategic north star."
}
```

**Request (Sarcastic):**
```bash
curl http://localhost:8080/no/sarcastic
```

**Request (Random):**
If you want the API to pick a category for you:
```bash
curl http://localhost:8080/no/random
```

## Available Categories

| Category | Description |
| :--- | :--- |
| `corporate_speak` | Professional-sounding rejections for the workplace. |
| `friendly_firm` | Kind but unyielding. |
| `funny_light` | A humorous way to say no. |
| `over_dramatic` | For when a simple "no" isn't enough. |
| `polite_professional` | Standard professional etiquette. |
| `passive_aggressive` | For those subtle hints. |
| `sarcastic` | Sharp and witty rejections. |
| `tech_nerd` | Rejections in tech jargon. |
| `random` | A surprise "no" from any of the above. |
