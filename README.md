# SSH Portfolio

> A unique SSH-based portfolio with a beautiful TUI interface built with Go, Bubble Tea, and Lip Gloss.

NOTE: If your terminal is light themed, the dark TUI I made will be forced to a light theme that is very unpleasant. I will fix this in the future.

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
ssh -p 2222 portfolio@adamdeleeuw.ca
```

**Password:** `portfolio`

If there are any issues connecting (handshake failed or any timeout behaviour) please create an issue on the [GitHub repository](https://github.com/adamdeleeuw/ssh-portfolio).
