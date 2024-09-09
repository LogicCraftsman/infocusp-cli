# Infocusp Projects CLI

Welcome to **Infocusp Projects CLI**, a powerful tool to help you rapidly scaffold project skeletons, generate demo applications, and streamline your development workflows. Whether you're building with React, FastAPI, Flask, or other popular frameworks, this CLI makes it easy to bootstrap your projects with the right configurations and dependencies.

## 🚀 Features

- **Scaffold Full-Stack Applications**: Quickly generate skeletons for popular frameworks such as React, FastAPI, and Flask.
- **Customizable Options**: Choose from options like Tailwind CSS, TypeScript, linting, testing frameworks, and Docker.
- **Demo Projects**: Explore ready-to-use demo projects for rapid prototyping.
- **Google Project Support**: Generate dependency files (`requirements.in`, `requirements.txt`) for projects using Google Cloud environments.

## 📦 Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/infocusp-projects-cli.git
   ```

2. Navigate to the project directory:

   ```bash
   cd infocusp-projects-cli
   ```

3. Build the CLI tool:

   ```bash
   go build -o infocusp
   ```

4. Add the binary to your system's PATH (optional but recommended):
   ```bash
   sudo mv infocusp /usr/local/bin/
   ```

## 🛠️ Usage

Once the CLI is installed, you can start creating project skeletons by running the following commands:

### Create a React App

```bash
infocusp create-react-app
```

This command scaffolds a new React app. It will prompt for:

- Project name
- Whether to include **Tailwind CSS**
- Whether to include **ESLint** for linting
- Choice of testing frameworks (Jest or Mocha)
- TypeScript support

### Create a FastAPI Skeleton

```bash
infocusp create-fastapi-skeleton
```

This command creates a basic FastAPI project with:

- A directory structure (`app/`, `main.py`, `__init__.py`)
- `requirements.txt` for dependencies
- Option to create a **Google Project** (adds `requirements.in` and runs `pip-compile`)

### Create a Flask Skeleton

```bash
infocusp create-flask-skeleton
```

This command creates a basic Flask project with:

- A directory structure (`app/`, `main.py`, `__init__.py`)
- `requirements.txt` for dependencies
- Dummy models, routes, and tests.

## 🧰 Available Commands

| Command                            | Description                                               |
| ---------------------------------- | --------------------------------------------------------- |
| `infocusp create-react-app`        | Create a new React project with custom configurations.    |
| `infocusp create-fastapi-skeleton` | Generate a basic FastAPI project with Docker and testing. |
| `infocusp create-flask-skeleton`   | Generate a Flask project with dummy models and routes.    |

## 📝 Examples

### Example: Creating a React App with TypeScript

```bash
infocusp create-react-app
```

- Enter the project name: `my-react-app`
- Include Tailwind CSS: `Yes`
- Include Linting: `Yes`
- Choose testing framework: `Jest`
- Use TypeScript: `Yes`

### Example: Creating a FastAPI Project for Google Cloud

```bash
infocusp create-fastapi-skeleton
```

- Enter the project name: `my-fastapi-app`
- Is this a Google project? `Yes`

This will create a FastAPI skeleton with a `requirements.in` file and run `pip-compile` to generate a `requirements.txt` file with dependency hashes.

## 🤝 Contributing

We welcome contributions to improve this CLI! Feel free to submit issues and pull requests to enhance the tool’s features or fix bugs.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Feel free to adjust the details like the repository URL or licensing information based on your project specifics.
