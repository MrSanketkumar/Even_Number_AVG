# 🧮 Even Number Average

<div align="center">

![Go Version](https://img.shields.io/github/go-mod/go-version/MrSanketKumar/Even_Number_AVG)
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-pr/MrSanketKumar/Even_Number_AVG)
![MIT License](https://img.shields.io/github/license/MrSanketKumar/Even_Number_AVG)


<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="Go" width="50" height="50"/>
<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original.svg" alt="Docker" width="50" height="50"/>
<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/redhat/redhat-original.svg" alt="Red Hat" width="50" height="50"/>
</div>

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 📚 Table of Contents

- [About](#-about)
- [Features](#-features)
- [Installation](#-installation)
- [Usage](#-usage)
- [API Documentation](#-api-documentation)
- [Testing](#-testing)
- [Contributing](#-contributing)
- [License](#-license)

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 🚀 About

The Even Number Average is a robust microservice designed to calculate the average of even numbers from a given set of inputs. This project showcases best practices in Go programming, including clean code structure, error handling, and comprehensive unit testing.

### Why I Made This

As a developer passionate about clean code and efficient algorithms, I created this project to:

1. **Demonstrate** how to build a simple yet powerful microservice in Go 1.22.
2. **Showcase** best practices in error handling and input validation.
3. **Provide** a real-world example of unit testing in Go.
4. **Offer** a template for developers looking to create similar microservices.

### Who Will This Help?

- 🎓 Go developers looking to improve their skills
- 💼 Teams seeking a template for building microservices
- 📚 Students learning about API development and testing

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## ✨ Features

- 📊 Calculate average of even numbers
- 🛡️ Robust error handling and input validation
- 🔄 RESTful API endpoint
- 📝 Detailed logging for requests and responses
- 🧪 Comprehensive unit tests
- 🎯 High performance and low resource usage
- 🐳 Docker support for easy deployment

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 🛠️ Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/MrSanketkumar/Even_Number_AVG.git
   ```

2. Navigate to the project directory:
   ```bash
   cd Even_Number_AVG
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 🖥️ Usage

1. Start the server:
   ```bash
   go run main.go
   ```

2. The server will start on `http://localhost:8080`.

3. Use the `/evenavg` endpoint to calculate averages:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"input_numbers": [1, 2, 3, 4, 5, 6]}' http://localhost:8080/evenavg
   ```

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 📘 API Documentation

### POST /average

Calculate the average of even numbers from the input array.

#### Request Body

```json
{
  "input_numbers": [float64]
}
```

#### Response

```json
{
  "message": string,
  "average": float64
}
```

#### Status Codes

- `200 OK`: Successful calculation
- `400 Bad Request`: Invalid input
- `405 Method Not Allowed`: Wrong HTTP method
- `500 Internal Server Error`: Server-side error

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 🧪 Testing

Run the test suite:

```bash
go test ./... -v
```

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

![----------------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

## 📄 License

This project is licensed under a sample license. See the [LICENSE](LICENSE) file for details.

---

<div align="center">
<p>Made with ❤️ by Sanket Chaudhary </p>
<p>
  <a href="https://github.com/MrSanketkumar" target="_blank" style="text-decoration:none;margin-right:20px;">
    <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/github/github-original.svg" width="40" height="40" alt="GitHub">
  </a>
  <a href="https://www.linkedin.com/in/sanketChaudhary/" target="_blank" style="text-decoration:none;">
    <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linkedin/linkedin-original.svg" width="40" height="40" alt="LinkedIn">
  </a>
</p>
</div>

### 🌟 A Fun Fact

Why did the even number break up with the odd number? Because it couldn't even!

[Back to top](#top)
