# ResLog

A lightweight Go CLI tool for parsing logs and redacting sensitive data on the fly.

## Installation

```bash
go install [github.com/dlcellucci/reslog/cmd/resolog@latest](https://github.com/dlcellucci/reslog/cmd/resolog@latest)
```

## Usage

```bash
echo "User test@example.com logged in from 192.168.1.1" | reslog
# Output: User [REDACTED EMAIL] logged in from [REDACTED IP]
```

### Command Line Options

- `-emails=false`: Disable email redaction
- `-ips=false`: Disable IP redaction
- `-cards=false`: Disable credit card redaction

## License

Licensed under the MIT License.
