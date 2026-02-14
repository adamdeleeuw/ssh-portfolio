# SSH Portfolio - Landing Page

Terminal-themed landing page directing visitors to the SSH portfolio.

## Development

```bash
# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## Structure

```
src/
├── App.jsx              # Main app with typing animation
├── App.css              # Main app styling
├── index.css            # Global styles and theme
├── main.jsx             # React entry point
└── components/
    ├── Terminal.jsx     # Terminal window UI
    ├── Terminal.css     # Terminal styling
    ├── CopyButton.jsx   # Copy to clipboard button
    └── CopyButton.css   # Button styling
```

## Features

- Animated typing effect using typed.js
- One-click SSH command copy
- Tokyo Night color theme matching TUI
- Fully responsive design
- Fast load times with Vite

