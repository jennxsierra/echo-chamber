# Echo Chamber

A TCP echo server written in Go that reflects client messages with a twist. It not only echoes input but also supports a set of interactive commands for an engaging user experience. This project served as an educational journey into TCP networking, concurrency with goroutines, and robust command handling.

> [!TIP]
>
> Watch a demo of echo-chamber in action on [YouTube](https://youtu.be/DemoVideoID).

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Makefile Uses](#makefile-uses)
- [Educational Highlights](#educational-highlights)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- **Concurrency:** Utilizes goroutines to handle multiple client connections concurrently.
- **Command Handling:** Supports interactive commands like `/time`, `/quit`, `/echo`, and `/help`.
- **Graceful Shutdown:** Implements proper resource cleanup including TCP linger strategy before closing connections.
- **Custom Logging:** Logs server and client activities with timestamps for clear traceability.
- **Inactivity Timeout:** Closes client connections after a period of inactivity to maintain server health.

## Installation

> [!NOTE]
>
> Ensure you have [Go](https://golang.org/dl/) (version 1.17 or later) installed before proceeding.

Follow these steps to install and build the project:

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/jennxsierra/echo-chamber.git
    cd echo-chamber
    ```

2. **Build the Application:**

    - **Using Make (recommended):**

      ```bash
      make build
      ```

      The binary will be located in the root directory or specified output directory.

    - **Without Make:**

      ```bash
      go fmt ./...
      go vet ./...
      go build -o echo-chamber cmd/server/main.go
      ```

3. **Run the Server:**

> [!NOTE]
>
> Replace `./bin/echo-chamber` with simply `./echo-chamber` if you built without the `Makefile`.

- To start the server on the default port (configured in `/internal/config`):

     ```bash
     ./bin/echo-chamber
     ```

- To run on a custom port (e.g., 5000):

     ```bash
     ./bin/echo-chamber -port=5000
     ```

> [!WARNING]
>
> Ensure the port you choose is not already in use by another application. You can check this using [netstat](https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/netstat) or similar tools.

## Usage

Clients can connect to the server using any TCP client (like [netcat](https://netcat.sourceforge.net/) or [telnet](https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/telnet)).
To connect using netcat, run the following command in a separate terminal:

```bash
nc localhost 4000
```

> [!NOTE]
>
> Replace `4000` with the port number you specified when starting the server.

Once connected, the server will greet you with a welcome message, and you can start sending messages. Try using the available commands prefixed by a slash (`/`).

## Commands

The following commands are supported:

- **/time**: Displays the current server time.

- **/echo**: Echoes back the provided message.

- **/help**: Lists all available commands with usage details.

- **/quit**: Closes your connection to the server.

These commands are implemented in [`internal/server/commands.go`](https://github.com/jennxsierra/echo-chamber/blob/main/internal/server/commands.go).

## Makefile Uses

> [!NOTE]
>
> You may need to install `make` if it's not already available on your system.

The provided `Makefile` streamlines several tasks involved in the development workflow:

- **fmt:** Formats the source code.
- **vet:** Runs static analysis to catch potential issues.
- **build:** Builds the application binary.
- **test:** Runs unit tests with verbose output.
- **clean:** Cleans up build artifacts and log files.
- **check:** Runs format, vet, and tests sequentially.

Execute any target with the `make <target>` command. For example:

```bash
make build
```

## Educational Highlights

- **Most Educationally Enriching:** The functionality of the personality mode and command handling was the most educationally enriching aspect of this project. It provided a deep dive into how to manage user interactions and implement maps, which are crucial for building interactive applications. I had a great time learning more about [Go maps](https://go.dev/blog/maps) and how to use them effectively for mapping strings to functions. This experience has significantly improved my understanding of Go's capabilities and how to leverage them for building robust applications.

- **Most Research Required:** The most part that required me to do the most research was implementing the TCP linger strategy for graceful shutdowns. It was a bit challenging and required a deep dive into Go's [net](https://pkg.go.dev/net) package in order to gain a solid understanding of TCP connections and how to manage them effectively. I also learned a lot about the importance of resource management and ensuring that connections are closed properly to avoid data loss or corruption, even though I had to go through a lot of trial and error to get it right. This experience has taught me the importance of thorough testing and debugging, especially when dealing with network programming.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- This project was developed as part of an assignment for the **[CMPS2242] Systems Programming & Computer Organization** course  under the Associate of Information Technology program at the [University of Belize](https://www.ub.edu.bz/).
- Special thanks to Mr. Dalwin Lewis for his guidance and support.
