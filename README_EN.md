## 🇷🇺  [Русская версия](./README.md)

**Go-Watcher** is a lightweight security tool written in Go that monitors suspicious activity and blocks it.

_This repository is a test task for an internship at CyberOK_

## Key features

- **Real-time Nginx log monitoring**: Go-Watcher continuously monitors Nginx access logs for suspicious requests.
- **Attack Detection**: Using customizable rules and regular expressions, Go-Watcher can identify various types of attacks, such as:
- Port Scans
- SQL Injections
- Cross-Site Scripting (XSS)
- Brute Force
- **IP Blocking**: When suspicious activity is detected, Go-Watcher can automatically block the attacker's IP address by adding it to the Nginx blacklist.
- **Flexible Configuration**: You can configure Go-Watcher to:
- Monitor specific paths
- Ignore specific IP addresses
- Detection sensitivity settings
- **Event Logging**: Go-Watcher keeps a detailed log of all detected events, including:
- Date and time
- Attacker's IP address
- Blocked URL
- Attack type
- **Ease of Use**: Go-Watcher is easy to configure and use, making it accessible even for novice users.

## Installation

### Step 1: Download the latest version of Go-Watcher

```bash
git clone https://github.com/trottling/Go-Watcher.git
```

### Step 2: Configure Go-Watcher

1. Edit the `config.json` configuration file:
- Set the detection rules.
- Configure other settings.

Example `config.json` configuration:
```json
{ 
"Proxy_Server": { 
"Port": 8080, 
"Show_Connections_STDOUT": true 
}, 
"Activity_Handler": { 
"Non_legit_Ports_RPM": 10, 
"Legit_Ports_RPM": 50, 
"Legit_Paths_Ignore_Regex": [".*/favicon.ico$"], 
"Legit_Ports": [0, 80, 443], 
"Legit_Paths_Brute_Regex": [".*/login$"], 
"Legit_Paths_Brute_RPM": 10, 
"Block_IPs": true, 
"Block_IPs_time": 600, 
"Dump_Requests": true,
"Requests_Dump_Ignore_Regex": [".*.ico$", ".*.css$", ".*.png$", ".*.js$"]
}
}

```

### Step 3: Run Go-Watcher

Run the Go-Watcher executable:

```bash
go run main.go
```