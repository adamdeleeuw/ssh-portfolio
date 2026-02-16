# SSH Portfolio

[![Deploy to Oracle VM](https://github.com/adamdeleeuw/ssh-portfolio/actions/workflows/deploy.yml/badge.svg)](https://github.com/adamdeleeuw/ssh-portfolio/actions/workflows/deploy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdeleeuw/ssh-portfolio)](https://goreportcard.com/report/github.com/adamdeleeuw/ssh-portfolio)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


A unique SSH-based portfolio with a beautiful TUI interface built with Go, Bubble Tea, and Lip Gloss.

> **NOTE:** I HIGHLY recommend using a dark terminal theme. If your terminal is light-themed, the dark TUI I made will be forced to a light theme that is very unpleasant. I will fix this in the future.

## Project Overview

Instead of a traditional website, this portfolio is accessed via SSH. Visitors connect to an interactive terminal interface with vim-inspired navigation, featuring:

- Tab-based navigation (Welcome, About, Projects, Future Plans)
- Beautiful TUI with Tokyo Night color scheme
- Secure, isolated environment
- Visitor counter and live statistics (coming soon)
- Interactive guest book (coming soon)
- Hidden easter eggs (coming soon)

## How to Connect

To access the portfolio, use the following SSH command:

```bash
ssh portfolio.adamdeleeuw.ca
```

If there are any issues connecting (handshake failed or any timeout behavior), please create an issue on the [GitHub repository](https://github.com/adamdeleeuw/ssh-portfolio).

## 🛠️ Built With

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea):** The fun, functional, and stateful terminal apps framework.
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss):** Style definitions for nice terminal layouts.
- **[Bubbles](https://github.com/charmbracelet/bubbles):** Some common Bubble Tea components.
- **[Glamour](https://github.com/charmbracelet/glamour):** Stylesheet-based Markdown renderer for your CLI apps.
- **[Glidessh](https://github.com/gliderlabs/ssh):** The SSH server library.

## Run Locally

If you want to view the portfolio TUI without SSH-ing into the server, you can run it locally:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/adamdeleeuw/ssh-portfolio.git
   cd ssh-portfolio
   ```

2. **Run the application:**
   ```bash
   go run cmd/server/main.go
   ```

3. **Connect locally:**
   ```bash
   ssh -p 2222 localhost
   ```

## Architecture & Infrastructure

To achieve a beautiful experience where anyone can connect without a port flag or password, this project utilizes a "Port-Swapped" infrastructure on an Oracle Cloud Ubuntu VM.

### The Network Stack

- **Public Entrance (Port 22):** The standard SSH port is bound to the Docker container running the Go application.
- **Administrative Entrance (Port 2022):** The host's native SSH daemon has been migrated to port 2022 to allow for remote management without interfering with the public portfolio.
- **Containerization:** The application is fully containerized, providing a layer of isolation between the guest session and the host OS.

### Deployment (CI/CD)

The project uses a deterministic GitHub Actions pipeline:

- **Build & Push:** Images are built and pushed to the GitHub Container Registry (GHCR) with the `latest` tag.
- **Deployment via SSH:** The workflow connects to the Oracle VM via SSH to execute a deployment script that:
  1. Updates the repo (`git fetch --all` & `git reset --hard origin/main`) to sync `docker-compose.yml`.
  2. Pulls the new image from GHCR (`docker-compose pull`).
  3. Restarts the containers (`docker-compose down` && `docker-compose up -d`) to apply changes and refresh network bindings.

## The "Debugging War Room"

Deploying a public SSH server surfaced several "invisible" networking bugs that provided a deep dive into TCP/IP and DNS layers.

### The "Ghost Bug" Theory: DNS vs. MTU

During development, the server suffered from intermittent "hangs" and "timeouts." I discovered I was actually fighting two distinct issues that mimicked each other:

1. **The DNS "Russian Roulette":** I originally had multiple A-Records (Vercel for web and Oracle for SSH) assigned to the root domain. Depending on the DNS lookup, users were often trying to SSH into a Vercel web server, leading to "Connection Timed Out" errors.
2. **The MTU "Ceiling":** Oracle Cloud Infrastructure (OCI) has a hard virtual network limit (MTU) of 1360 bytes. Standard Docker networks default to 1500 bytes. 

NOTE: This MTU configuration seemed to work for me, I don't know if this is a general issue. It's possible that the only bug was the DNS misconfiguration.

**The Result:** Handshakes (small packets) worked fine, but as soon as the "heavy" TUI data was sent, the packets exceeded the limit and were silently dropped by the cloud provider, causing the client to hang indefinitely.

### The Fix

- **MTU Clamping:** Configured the Docker network bridge to an MTU of 1350 and implemented iptables MSS clamping to force clients to negotiate smaller packet sizes during the initial handshake.
- **DNS Isolation:** Migrated the SSH server to its own dedicated subdomain (`portfolio.adamdeleeuw.ca`).

## Contributing & License

While this project is a showcase of my personal journey and technical skills, I welcome community interaction!

Please check out the [Contributing Guidelines](CONTRIBUTING.md) for more details.

**License:** This project is licensed under the [MIT License](LICENSE).

