import PropTypes from 'prop-types'
import './Terminal.css'

/**
 * Terminal window component with macOS-style window controls.
 * Provides the aesthetic container for terminal content.
 * 
 * @param {Object} props - Component props
 * @param {React.ReactNode} props.children - Terminal content to display
 */
function Terminal({ children }) {
    return (
        <div className="terminal-window">
            {/* Terminal header with window controls */}
            <div className="terminal-header">
                <div className="window-controls">
                    <span className="control close"></span>
                    <span className="control minimize"></span>
                    <span className="control maximize"></span>
                </div>
                <div className="terminal-title">portfolio@adamdeleeuw.ca</div>
            </div>

            {/* Terminal body */}
            <div className="terminal-body">
                {children}
            </div>
        </div>
    )
}

Terminal.propTypes = {
    children: PropTypes.node.isRequired
}

export default Terminal
