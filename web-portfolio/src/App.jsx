import { useState, useEffect, useRef } from 'react'
import Typed from 'typed.js'
import './App.css'
import Terminal from './components/Terminal'
import CopyButton from './components/CopyButton'

/**
 * Main application component for SSH portfolio landing page.
 * Displays terminal-style interface with typing animation.
 */
function App() {
    const [typingComplete, setTypingComplete] = useState(false)
    const [terminalClosed, setTerminalClosed] = useState(false)
    const typedRef = useRef(null)
    const el = useRef(null)

    // SSH connection command
    const sshCommand = 'ssh portfolio.adamdeleeuw.ca'

    useEffect(() => {
        // Typing animation configuration
        const options = {
            strings: [
                'Connecting to portfolio...^500\nHTTP is for losers.^800\n \nTo explore my portfolio, SSH instead:^500',
            ],
            typeSpeed: 40,
            backSpeed: 0,
            fadeOut: false,
            loop: false,
            showCursor: false,
            onComplete: () => {
                setTypingComplete(true)
            }
        }

        // Initialize typed.js
        typedRef.current = new Typed(el.current, options)

        // Cleanup
        return () => {
            typedRef.current.destroy()
        }
    }, [])

    // Handle Ctrl+C keypress
    useEffect(() => {
        const handleKeyDown = (e) => {
            // Check for Ctrl+C (Windows/Linux) or Cmd+C (Mac) when terminal is open
            if ((e.ctrlKey || e.metaKey) && e.key === 'c' && !terminalClosed) {
                // Only close if nothing is selected (allow copy)
                if (window.getSelection().toString() === '') {
                    e.preventDefault()
                    setTerminalClosed(true)
                }
            }
        }

        window.addEventListener('keydown', handleKeyDown)
        return () => window.removeEventListener('keydown', handleKeyDown)
    }, [terminalClosed])

    // Reopen terminal
    const reopenTerminal = () => {
        setTerminalClosed(false)
    }

    // If terminal is closed, show reopen button
    if (terminalClosed) {
        return (
            <div className="app">
                <div className="terminal-closed">
                    <p className="closed-message">Terminal closed</p>
                    <button className="reopen-button" onClick={reopenTerminal}>
                        Reopen Terminal
                    </button>
                </div>
            </div>
        )
    }

    return (
        <div className="app">
            <Terminal>
                <div className="terminal-content">
                    {/* Typing animation */}
                    <div className="typed-lines">
                        <span ref={el}></span>
                    </div>

                    {/* SSH command display (shown after typing) */}
                    {typingComplete && (
                        <div className="ssh-command-section">
                            <div className="command-line">
                                <span className="prompt">$</span>
                                <span className="command">{sshCommand}</span>
                            </div>

                            <div className="button-container">
                                <CopyButton text={sshCommand} />
                            </div>

                            <div className="help-text">
                                <p>Need help? Check the <a href="https://github.com/adamdeleeuw/ssh-portfolio" target="_blank" rel="noopener noreferrer">README</a></p>
                            </div>
                        </div>
                    )}
                </div>

                {/* Terminal footer hint */}
                <div className="terminal-footer">
                    <span className="footer-hint">Press Ctrl+C to exit</span>
                </div>
            </Terminal>
        </div>
    )
}

export default App
