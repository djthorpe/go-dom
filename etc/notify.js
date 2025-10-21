// Server-Sent Events for live reload and error notifications
(function () {
    'use strict';

    // Initialize when DOM is fully loaded
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        // DOM is already loaded
        init();
    }

    function init() {
        // Create error overlay DOM elements
        createErrorOverlay();

        // Start SSE connection
        startSSEConnection();
    }

    function createErrorOverlay() {
        // Create the error overlay div
        const errorDiv = document.createElement('div');
        errorDiv.id = 'compile-error';

        // Create close button
        const closeButton = document.createElement('button');
        closeButton.id = 'compile-error-close';
        closeButton.textContent = '✕ Close';
        closeButton.onclick = hideCompileError;

        // Create title
        const title = document.createElement('div');
        title.id = 'compile-error-title';
        title.textContent = '⚠️ Compilation Error';

        // Create content div
        const content = document.createElement('div');
        content.id = 'compile-error-content';

        // Assemble the error overlay
        errorDiv.appendChild(closeButton);
        errorDiv.appendChild(title);
        errorDiv.appendChild(content);

        // Append to body
        document.body.appendChild(errorDiv);
    }

    function startSSEConnection() {
        const eventSource = new EventSource('/_notify');

        eventSource.onmessage = function (event) {
            if (event.data === 'reload') {
                console.log('wasmbuild: reloading page...');
                // Hide any error overlay before reload
                hideCompileError();
                window.location.reload();
            }
        };

        eventSource.addEventListener('compileerror', function (event) {
            console.error('wasmbuild compilation error:', event.data);
            showCompileError(event.data);
        });

        eventSource.onerror = function (error) {
            console.log('wasmbuild: connection error, will retry...');
        };
    }

    function showCompileError(errorMsg) {
        const errorDiv = document.getElementById('compile-error');
        const errorContent = document.getElementById('compile-error-content');
        if (errorContent) {
            errorContent.textContent = errorMsg;
        }
        if (errorDiv) {
            errorDiv.style.display = 'block';
        }
    }

    function hideCompileError() {
        const errorDiv = document.getElementById('compile-error');
        if (errorDiv) {
            errorDiv.style.display = 'none';
        }
    }

    // Expose hideCompileError globally for programmatic access
    window.hideCompileError = hideCompileError;
})();

