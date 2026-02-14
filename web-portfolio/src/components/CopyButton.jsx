import { useState } from 'react'
import PropTypes from 'prop-types'
import './CopyButton.css'

/**
 * Copy-to-clipboard button with visual feedback.
 * 
 * @param {Object} props - Component props
 * @param {string} props.text - Text to copy to clipboard
 */
function CopyButton({ text }) {
    const [copied, setCopied] = useState(false)

    /**
     * Handles clipboard copy action.
     * Shows visual feedback for 2 seconds.
     */
    const handleCopy = async () => {
        try {
            await navigator.clipboard.writeText(text)
            setCopied(true)
            setTimeout(() => setCopied(false), 2000)
        } catch (err) {
            console.error('Failed to copy:', err)
        }
    }

    return (
        <button
            className={`copy-button ${copied ? 'copied' : ''}`}
            onClick={handleCopy}
            aria-label="Copy SSH command to clipboard"
        >
            <span className="button-icon">{copied ? '✓' : '📋'}</span>
            <span className="button-text">{copied ? 'Copied!' : 'Copy SSH Command'}</span>
        </button>
    )
}

CopyButton.propTypes = {
    text: PropTypes.string.isRequired
}

export default CopyButton
