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
    <img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=500&size=22&duration=3000&pause=1000&color=00FF00&center=true&vCenter=true&width=600&lines=Bypass+403%2C+401%2C+404+Like+a+Pro;Advanced+Payload+Collection;Rate+Limiting+%2B+Concurrent+Requests;Burp+Suite+Integration;HTTP%2F3+%2F+HTTP%2F2+%2F+HTTP%2F1.1+Support;Path+%2B+Header+%2B+Protocol+Manipulation;Unlimited+Bypass+with+-r+Flag" alt="Typing SVG" />
  </p>
  
  <!-- Social Badge -->
  <p>
    <a href="https://github.com/adrianalvird">
      <img src="https://img.shields.io/badge/Author-@adrianalvird-blue?style=flat&logo=github" alt="Author" />
    </a>
  </p>
</div>

---

## ✨ Features

| Category | Techniques & Capabilities |
|----------|--------------------------|
| 🔀 **Path Manipulation** | Path traversal, encoding tricks, special characters |
| 🧠 **Header Injection** | X-Forwarded-For, X-Original-URL, X-Rewrite-URL, Client-IP |
| 🔄 **Method Manipulation** | GET, POST, PUT, TRACE, OPTIONS, custom methods |
| 📁 **Extension Injection** | .php, .js, .html, .asp, .aspx, .json |
| 📡 **Protocol Downgrade** | HTTP/3, HTTP/2, HTTP/1.1, HTTP/1.0, HTTP/0.9 |
| 🚦 **Rate Limiting** | Dynamic rate adjustment to avoid blocking |
| ⚡ **Concurrent Requests** | Multiple requests sent simultaneously with different patterns |
| 🔌 **Burp Suite Integration** | Route all traffic through Burp for inspection |
| ♾️ **Continuous Mode** | `-r` flag for unlimited bypass attempts |
| 📝 **Request File Support** | Parse captured requests from files |
| 🔍 **Verbose Mode** | See all requests and responses in real-time |

---

## 🚀 Installation

### Via Go Install (Recommended)
```bash
go install github.com/adrianalvird/bypasser/bypasser@v1.0.9
```

## Usage :

```bash
echo "http://notsecure.in" | bypasser -verbose
```

```bash
echo "http://notsecure.in" | bypasser -proxy http://127.0.0.1:8080
```

```bash
bypasser -request request.txt
```

```bash
bypasser -request request.txt -proxy http://127.0.0.1:8080
```
