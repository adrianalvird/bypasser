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
  ## 🌐 Connect With Me

<div align="center">
  <a href="https://github.com/adrianalvird">
    <img src="https://img.shields.io/badge/GitHub-@adrianalvird-181717?style=for-the-badge&logo=github" alt="GitHub" />
  </a>
  <a href="https://linkedin.com/in/adrianalvird">
    <img src="https://img.shields.io/badge/LinkedIn-@adrianalvird-0077B5?style=for-the-badge&logo=linkedin" alt="LinkedIn" />
  </a>
  <a href="https://instagram.com/adrian.alvird">
    <img src="https://img.shields.io/badge/Instagram-@adrian.alvird-E4405F?style=for-the-badge&logo=instagram" alt="Instagram" />
  </a>
  <a href="https://x.com/adrianalvird">
    <img src="https://img.shields.io/badge/X-@adrianalvird-000000?style=for-the-badge&logo=x" alt="X (Twitter)" />
  </a>
</div>

---

## 🤝 Support Me

<div align="center">

[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-ffdd00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black)](https://www.buymeacoffee.com/adrianalvird)
[![PayPal](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://paypal.me/adrianalvird)
[![GitHub Sponsors](https://img.shields.io/badge/GitHub%20Sponsors-30363D?style=for-the-badge&logo=github&logoColor=white)](https://github.com/sponsors/adrianalvird)

</div>

<div align="center">
  <p><strong>☕ Every coffee keeps me coding through the night!</strong></p>
  <p><code>https://www.buymeacoffee.com/adrianalvird</code></p>
</div>
---

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  <p>Made with ❤️ by <a href="https://github.com/adrianalvird">@adrianalvird</a></p>
  <p>Version 1.0.9 | © 2024 Bypasser</p>
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
