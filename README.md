<div align="center">
  <h1> Bypasser</h1>
  <p><strong>The Most Powerful Open Source WAF Bypass Tool</strong></p>
  <p>Bypass 4XX Status Codes with Advanced Techniques & Payloads</p>

  <!-- Version Badge -->
  <img src="https://img.shields.io/badge/version-1.0.9-blue.svg" alt="Version 1.0.9">
  <img src="https://img.shields.io/badge/go-1.24.2-green.svg" alt="Go 1.24.2">
  <img src="https://img.shields.io/badge/license-MIT-orange.svg" alt="MIT License">
  <img src="https://img.shields.io/badge/status-stable-brightgreen.svg" alt="Stable">

  <!-- Typing Effect -->
  <p>
    <img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=500&size=22&duration=3000&pause=1000&color=00FF00&center=true&vCenter=true&width=600&lines=Bypass+403%2C+401%2C+404+Like+a+Pro;Advanced+Payload+Collection;Rate+Limiting+%2B+Concurrent+Requests;HTTP%2F3+%2F+HTTP%2F2+%2F+HTTP%2F1.1+Support;Path+%2B+Header+%2B+Protocol+Manipulation" alt="Typing SVG" />
  </p>
</div>

---

## ✨ Features

| Category | Techniques |
|----------|------------|
| 🔀 **Path Manipulation** | Path traversal, encoding tricks, special characters |
| 🧠 **Header Injection** | X-Forwarded-For, X-Original-URL, X-Rewrite-URL, Client-IP |
| 🔄 **Method Manipulation** | GET, POST, PUT, TRACE, OPTIONS, custom methods |
| 📁 **Extension Injection** | .php, .js, .html, .asp, .aspx, .json |
| 📡 **Protocol Downgrade** | HTTP/3, HTTP/2, HTTP/1.1, HTTP/1.0, HTTP/0.9 |
| 🚦 **Rate Limiting** | Dynamic rate adjustment, user-agent rotation |
| ⚡ **Concurrency** | Multi-worker concurrent requests |
| 📝 **Request File Support** | Parse captured requests from files |
| 🔌 **Proxy Support** | HTTP/HTTPS proxy integration |

---

## 🚀 Installation

### Via Go Install (Recommended)
```bash
go install github.com/adrianalvird/bypasser/bypasser@v1.0.9
```

## Usage :

```bash
echo "http://example.com" | bypasser -verbose
```

```bash
echo "http://example.com" | bypasser -proxy http://127.0.0.1:8080
```

```bash
bypasser -request request.txt
```

```bash
bypasser -request request.txt -proxy http://127.0.0.1:8080
```
