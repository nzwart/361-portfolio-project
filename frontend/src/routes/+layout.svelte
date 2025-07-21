<!-- src/routes/+layout.svelte -->

<slot />

<style>
    @import url("https://fonts.googleapis.com/css2?family=Jacquard+24:wght@400&display=swap");

    :global(body) {
        margin: 0;
        padding: 0;
    }

    :global(.container) {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        /* height: calc(100vh - 10px);
        width: calc(100vw - 10px);
        border: 5px solid white;
        border-radius: 10px;
        background-color: #bd4848;
        color: white;
        font-family: "serif";
        text-align: center; */
    }

    :global(h1) {
        font-size: 3rem;
        margin-bottom: 20px;
    }

    :global(.frontpage-title) {
        font-family: "Jacquard 24", serif;
        margin-bottom: 20px;
    }

    :global(.body) {
        margin: 0;
        max-width: 300px;
        width: 100%;
        font-size: 1rem;
        padding: 0;
        box-sizing: border-box;
    }

    :global(.button-container) {
        margin-top: 20px;
    }

    :global(button) {
        padding: 15px 30px;
        font-size: 1.2rem;
        background-color: #e0e0e0;
        border: none;
        border-radius: 8px;
        cursor: pointer;
        transition:
            background-color 0.2s,
            color 0.2s;
    }

    :global(button:hover:enabled) {
        background-color: #d4d4d4;
    }

    /* Instruction dialogue box css from encbuilder */
    :global(.instruction-dialog) {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        background-color: gray;
        padding: 20px;
        border-radius: 10px;
        box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
        text-align: center;
    }

    :global(.instruction-dialog h4) {
        margin-bottom: 10px;
    }

    :global(.instruction-dialog p) {
        margin: 10px 0;
    }

    /* search dialogue box css from encbuilder */
    :global(.search-dialog) {
        background-color: rgba(0, 0, 0, 0.8);
        color: white;
        position: fixed;
        transform: translate(-6%, 1%);
        z-index: 10000;
        width: 80%;
        height: 80%;
        padding: 20px;
        transform: translate(-50%, 1%);
        left: 50%;
        border-radius: 10px;
        display: flex;
        flex-direction: column;
        box-sizing: border-box;
    }

    /* Search input spanning the dropdown */
    :global(.search-input) {
        width: 100%; /* Ensure the input takes up the full width */
        padding: 10px;
        border: none;
        border-bottom: 1px solid black;
        font-size: 1rem;
        border-radius: 5px;
        box-sizing: border-box; /* Includes padding in width calculation */
    }

    /* Search results section */
    :global(.search-results) {
        flex: 1; /* Take up remaining space in the search dialog */
        overflow-y: auto; /* Enable vertical scrolling */
        max-height: 70vh; /* Limit the height to 70% of the viewport */
        margin-top: 10px;
        border: 1px solid #ccc; /* Optional: border around results */
        border-radius: 5px;
        padding: 10px;
        background-color: #f0f0f0; /* Light background */
    }

    /* Monster List Styling */
    :global(.monster-list) {
        list-style-type: none; /* Remove default bullet points */
        padding: 0;
        margin: 0;
    }

    /* Individual Monster Item */
    :global(.monster-item) {
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 10px 0;
        border-bottom: 1px solid #ccc;
        font-size: 1rem;
        color: black;
    }

    /* :global(.monster-name) {
        font-weight: bold;
        color: black;
        margin-right: 10px;
        font-size: 2rem;
        text-align: left;
    } */

    :global(.controls) {
        align-items: right;
    }

    :global(.search-dialog-close-btn) {
        margin-top: auto;
        padding: 10px 20px;
        background-color: #848484;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
    }

    /* Encbuilder page menu bar and tray and card styles */
    :global(.menu-bar) {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 60px;
        border: 5px solid white;
        border-radius: 10px;
        background-color: #bd4848;
        color: white;
        display: flex;
        justify-content: flex-start; /* Aligns buttons to the left */
        align-items: center;
        padding: 0 20px; /* Adds padding from the left */
        z-index: 1000;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        box-sizing: border-box;
    }

    :global(.menu-bar button) {
        margin-right: 10px;
        padding: 10px 20px;
        font-size: 16px;
        cursor: pointer;
        background-color: #e0e0e0;
        border: 1px solid black;
        border-radius: 8px;
        font-weight: bold;
        color: black;
    }

    /* Darker background when search is active */
    :global(.menu-bar button.active-search) {
        background-color: #333;
        color: white;
    }

    :global(.tray) {
        position: absolute;
        bottom: 0;
        left: 50%;
        transform: translateX(-50%);
        width: 80%;
        height: 80%;
        background-color: white;
        padding: 10px;
        border-radius: 10px;
        box-sizing: border-box;
        display: flex;
        flex-direction: row;
        overflow-x: auto;
        overflow-y: hidden;
        gap: 10px;
        box-shadow: inset 0 5px 10px rgba(0, 0, 0, 0.2);
    }

    :global(.card *) {
        max-width: 100%;
        overflow-wrap: break-word;
        word-wrap: break-word;
        hyphens: auto;
    }

    :global(.card) {
        background-color: #f8f8f8;
        color: black;
        padding: 10px;
        border: 1px solid #ddd;
        border-radius: 5px;
        min-width: 20%;
        max-width: 20%;
        flex-shrink: 0;
        display: flex;
        flex-direction: column;
        height: 100%;
        box-sizing: border-box;
        /* Add these properties */
        overflow-x: hidden;
        word-wrap: break-word;
        overflow-wrap: break-word;
    }

    :global(.monster-name) {
        font-weight: bold;
        color: black;
        font-size: 2rem;
        text-align: left;
        margin-bottom: 10px;
        /* Add text wrapping for long monster names */
        word-wrap: break-word;
        overflow-wrap: break-word;
        width: 100%;
    }

    :global(.card-body-text) {
        flex-grow: 1;
        overflow-y: auto;
        padding-right: 5px;
        height: calc(100% - 120px);
        display: flex;
        flex-direction: column;
    }

    :global(.attribute) {
        margin-bottom: 8px;
        word-wrap: break-word;
        overflow-wrap: break-word;
        width: 100%;
    }

    :global(.attribute strong) {
        margin-bottom: 2px;
    }

    :global(.remove-card-btn) {
        margin-top: auto;
        width: 100%;
    }

    :global(.expander-btn) {
        background: none;
        border: none;
        color: #666;
        cursor: pointer;
        padding: 5px;
        margin: 10px 0;
        width: 100%;
        text-align: center;
        font-size: 0.9rem;
    }

    :global(.expander-btn:hover) {
        background-color: #f0f0f0;
    }

    /* Card removal css */
    /* .confirmation-dialog {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
    } */

    :global(.confirmation-content) {
        background-color: white;
        padding: 20px;
        border-radius: 8px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        text-align: center;
    }

    :global(.confirmation-buttons) {
        display: flex;
        justify-content: center;
        gap: 10px;
        margin-top: 20px;
    }

    :global(.confirm-btn) {
        background-color: #dc3545;
        color: white;
        border: none;
        padding: 8px 16px;
        border-radius: 4px;
        cursor: pointer;
    }

    :global(.confirm-btn:hover) {
        background-color: #c82333;
    }

    :global(.cancel-btn) {
        background-color: #6c757d;
        color: white;
        border: none;
        padding: 8px 16px;
        border-radius: 4px;
        cursor: pointer;
    }

    :global(.cancel-btn:hover) {
        background-color: #5a6268;
    }

    :global(.empty-tray-msg) {
        color: #6c757d;
        z-index: 1000;
        font-size: 1.2rem;
    }

    :global(.empty-tray-container) {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
        width: 100%;
        position: absolute;
        top: 0;
        left: 0;
    }

    :global(.hp-tracker) {
        padding: 0.5rem;
        background-color: #f5f5f5;
        border-radius: 4px;
        margin: 0.5rem 0;
    }

    :global(.hp-controls) {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 0.5rem;
        margin-bottom: 0.5rem;
    }

    :global(.hp-btn) {
        padding: 0.25rem 0.5rem;
        font-size: 1.2rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        background-color: white;
        cursor: pointer;
    }

    :global(.hp-btn:hover) {
        background-color: #eee;
    }

    :global(.hp-display) {
        display: flex;
        align-items: center;
        gap: 0.25rem;
    }

    :global(.hp-input) {
        width: 4rem;
        padding: 0.25rem;
        text-align: center;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    :global(.hp-separator) {
        font-weight: bold;
    }

    :global(.hp-max) {
        min-width: 2rem;
        text-align: left;
    }

    :global(.hp-bar) {
        width: 100%;
        height: 0.5rem;
        background-color: #ddd;
        border-radius: 4px;
        overflow: hidden;
    }

    :global(.hp-bar-fill) {
        height: 100%;
        transition:
            width 0.3s ease,
            background-color 0.3s ease;
    }

    /* Name generator button styles */
    :global(.name-generator-wrapper) {
        position: fixed;
        top: 5rem;
        right: 1rem;
        z-index: 9999;
        width: 350px;
        max-height: calc(100vh - 7rem);
        overflow-y: auto;
        background-color: rgba(245, 245, 245, 0.95);
        border-radius: 0.5rem;
        box-shadow:
            0 4px 6px rgba(0, 0, 0, 0.1),
            0 2px 4px rgba(0, 0, 0, 0.06);
        backdrop-filter: blur(8px);
    }

    :global(.name-generator) {
        margin: 0;
        padding: 0.75rem;
    }

    :global(.name-generator-controls) {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    :global(.name-generator-btn) {
        padding: 0.4rem 0.75rem;
        font-size: 0.8rem;
        background-color: #4a5568;
        color: white;
        border: none;
        border-radius: 0.25rem;
        cursor: pointer;
        transition: all 0.2s;
        font-weight: 500;
        white-space: nowrap;
    }

    :global(.name-generator-btn:hover) {
        background-color: #2d3748;
        transform: translateY(-1px);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    :global(.name-generator-btn:disabled) {
        opacity: 0.5;
        cursor: not-allowed;
        transform: none;
        box-shadow: none;
    }

    :global(.name-generator-result) {
        font-size: 0.875rem;
        font-weight: 500;
        color: #2d3748;
        padding: 0.4rem;
        margin-top: 0.25rem;
        background-color: white;
        border-radius: 0.25rem;
        box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
        overflow-wrap: break-word;
        word-break: break-word;
        text-align: left;
        min-height: 2.5rem;
    }

    :global(.name-generator-error) {
        margin-top: 0.5rem;
        padding: 0.5rem;
        color: #e53e3e;
        background-color: #fff5f5;
        border: 1px solid #fc8181;
        border-radius: 0.25rem;
        font-size: 0.875rem;
        width: 100%;
    }

    :global(.name-generator-minimize) {
        position: absolute;
        top: 0.5rem;
        right: 0.5rem;
        padding: 0.25rem 0.5rem;
        font-size: 0.8rem;
        background: none;
        border: none;
        color: #4a5568;
        cursor: pointer;
    }

    /* Mobile responsive styles */
    @media (max-width: 640px) {
        :global(.name-generator-wrapper) {
            width: calc(100% - 2rem);
            top: auto;
            bottom: 1rem;
            right: 1rem;
            left: 1rem;
            max-height: 60vh;
        }

        :global(.name-generator-controls) {
            align-items: stretch;
        }

        :global(.name-generator-result) {
            text-align: center;
        }
    }
</style>
