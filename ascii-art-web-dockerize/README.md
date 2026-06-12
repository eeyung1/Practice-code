# Docker Setup

## Create the Dockerfile

Create a file named `Dockerfile` in the project root and add:

```dockerfile
FROM golang

WORKDIR /app

COPY . .

RUN go build

CMD ["./project"]
```

### Explanation

* `FROM golang`

  * Uses the official Go image as the base environment.

* `WORKDIR /app`

  * Creates and switches to the `/app` directory inside the container.

* `COPY . .`

  * Copies the entire project into the container.

* `RUN go build`

  * Builds the application and creates an executable named `project`.

* `CMD ["./project"]`

  * Starts the application when the container runs.

---

# Build the Image

From the project root, run:

```bash
docker build -t ascii-art .
```

### Explanation

* `docker build`

  * Creates an image from the Dockerfile.

* `-t ascii-art`

  * Assigns the name `ascii-art` to the image.

* `.`

  * Uses the current directory as the build context.

---

# Verify the Image

List available images:

```bash
docker images
```

You should see:

```text
REPOSITORY   TAG       IMAGE ID
ascii-art    latest    ...
```

---

# Run the Container

Start the application:

```bash
docker run -p 8080:8080 ascii-art
```

### Explanation

* `docker run`

  * Creates and starts a container from the image.

* `-p 8080:8080`

  * Maps host port 8080 to container port 8080.

* `ascii-art`

  * Specifies the image to run.

If successful, the server should output:

```text
Server started on :8080
```

---

# View Running Containers

```bash
docker ps
```

Example:

```text
CONTAINER ID   IMAGE       PORTS
7217092406c3   ascii-art   0.0.0.0:8080->8080/tcp
```

---

# Stop a Running Container

Using the container ID:

```bash
docker stop 7217092406c3
```

Or using the container name:

```bash
docker stop magical_sammet
```

---

# Container Lifecycle

```text
Dockerfile
    ↓
docker build
    ↓
Image
    ↓
docker run
    ↓
Container
    ↓
Running Server
```

* **Image** → Blueprint of the application.
* **Container** → Running instance of the image.

Stopping a container does not remove the image. A new container can always be created from the same image.
