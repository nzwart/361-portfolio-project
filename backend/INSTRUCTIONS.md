# Running with all microservices

-   Start each microservice in its own process, by calling the appropriate commands from terminal instances within each project dir.
    -   For my microservices: (example) `./builds/zwartn-microserviceD-mac-arm64`
    -   For the classmate microservice:
        ```bash
        source .venv/bin/activate
        python name_gen_server.py
        ```
-   Start this webapp as normal.

# Setup

### Go setup

-   Make sure you have Go installed by running:

    ```
    go version
    ```

### Svelte setup

-   Make sure you have Node.js installed by running:

    ```
    node -v
    npm -v
    ```

-   Install the Svelte template and make it multi-page:

    ```
    npx degit sveltejs/template frontend
    cd frontend
    npm install
    npm install svelte-spa-router
    ```

-   Inside the `frontend` directory, build the Svelte app:

    ```
    npm run build
    ```

    This command will generate the static files inside the public directory, which will be served by the Go server.

### Running the app with hot reloading

-   Will need two processes. Inside process #1, from the root directory, run:

    ```
    cd frontend
    npm run dev -- --open

    ```

-   Inside process #2, from the root directory, run:

    ```
    go run main.go monster.go

    ```

    Use `http://localhost:5173` to access the app.

    Why? `5173` is the default development server port used by Vite, the build tool used by Svelte’s default template. Vite powers the development environment to provide hot-reloading.
