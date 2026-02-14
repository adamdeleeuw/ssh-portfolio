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
    const typedRef = useRef(null)
    const el = useRef(null)

    // SSH connection command
    const sshCommand = 'ssh -p 2222 adam@yourserver.com'

    useEffect(() => {
        // Typing animation configuration
        const options = {
            strings: [
                'Connecting to portfolio...^500',
                'Connecting to portfolio...^500\nTraditional websites are so 2020.^800',
                'Connecting to portfolio...^500\nTraditional websites are so 2020.^800\n \nTo explore my portfolio, SSH instead:^500',
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

                            <div className="password-hint">
                                Password: <span className="highlight">[displayed on connection]</span>
                            </div>

                            <div className="button-container">
                                <CopyButton text={sshCommand} />
                            </div>

                            <div className="help-text">
                                <p>Need help? Check the <a href="https://github.com/adamdeleeuw/ssh-portfolio" target="_blank" rel="noopener noreferrer">README</a></p>
                            </div>
                        </div>
                    )}

                    {/* Cursor */}
                    <span className={`cursor ${typingComplete ? 'blink' : ''}`}>_</span>
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
