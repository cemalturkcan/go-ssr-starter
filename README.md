# Go Fiber SSR Starter

This is a **starter project** for rendering **Server-Side Rendered (SSR)** pages using **Go Fiber** and **HTML templating**. It is designed to streamline development and improve productivity with reusable components and live reloading during development.

---

## Project Purpose

The main purpose of this project is to:

- Provide a boilerplate for SSR applications using Go and Fiber.
- Enable modular and reusable HTML components.
- Simplify the development process with **live reload** integration via [Live.js](https://livejs.com/).

---

## Project Structure

The project follows a clean and organized structure:

```bash
go-ssr-starter/
│
├── cmd/
│   └── main.go         # Application entry point
├── src/
│   ├── general/
│   │   ├── config/     # Configurations
│   │   ├── server/     # Server setup
│   │   ├── templ/      # Templating utilities
│   │   └── util/       # Additional utilities
│   ├── routes/         # Routing logic
│   │   └── blog/       # example route
│   ├── tmp/            # Temporary files
│   └── views/          # View templates
│       ├── layouts/    # Base layout templates
│       ├── partials/   # Components
│       └── styles/     # Style templates
│
├── Dockerfile          # Docker configuration
├── Makefile            # Build and run commands
├── docker-compose.yml  # Docker Compose for local development
```

---

## Features

### 1. Live Reload in Development

When running in development mode, **Live.js** is injected into the HTML header, automatically reloading the browser when any file changes.

- No manual refresh needed.
- Improves development speed and workflow.

To enable live reload:
1. Ensure `https://livejs.com/live.js` is included in your `<head>` during development.
2. Changes in your HTML, CSS, or JS files will trigger an auto-refresh.

---

### 2. Reusable Components

The project allows you to define and reuse **HTML components** using Go Fiber's templating engine.

#### Defining a Component

Example of a `button` component:

```html
{{ define "button" }}
<button>
    {{ .label }}
    {{ .context.Name }}
</button>
{{ end }}
```

#### Passing Props to Components

You can pass props to a component using the custom `putValues` method:

```html
{{ template "button" putValues "context" . "label" "Hello" }}
```

The `putValues` function dynamically maps keys and values to props.

---

### 3. Custom `putValues` Method

The custom `putValues` method is implemented as follows:

```go
engine.AddFunc(
    "putValues", func(values ...interface{}) fiber.Map {
        if len(values)%2 != 0 {
            return fiber.Map{}
        }
        result := make(fiber.Map, len(values)/2)
        for i := 0; i < len(values); i += 2 {
            key, ok := values[i].(string)
            if !ok {
                log.Error("putValues: key is not a string")
                continue
            }
            result[key] = values[i+1]
        }
        return result
    },
)
```

- **Purpose**: Allows you to pass dynamic key-value pairs to components in the HTML templates.
- **Usage**: Simplifies the process of injecting props into reusable components.

---

---
## HTML Templating Documentation
For more details on Go Fiber's HTML templating engine, check out the [official documentation](https://docs.gofiber.io/template/html/).

---

## License

This project is licensed under the MIT License.

Build the application
```bash
make build
```
Run the application
```bash
make run
```
Clean up binary from the last build:
```bash
make clean
```

Live reload the application:
```bash
make watch
```
