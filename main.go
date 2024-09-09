package main

import (
	"infocusp-projects/commands"

	"github.com/spf13/cobra"
)

// main is the entry point for the Infocusp Projects CLI.
// This CLI tool allows users to bootstrap skeletons for various tech stacks such as React, FastAPI, and Flask.
// It also provides customization options like Tailwind CSS, TypeScript, linting, testing frameworks, and Docker.
//
// Usage:
//   Run `infocusp` followed by a command to create a project skeleton.
//   Example commands:
//   - infocusp create-react-app: Create a React app with optional features like Tailwind CSS and TypeScript.
//   - infocusp create-flask-skeleton: Create a basic Flask project structure.
//   - infocusp create-fastapi-skeleton: Create a basic FastAPI project structure with optional testing and Google project settings.
//
// Features:
//   - Generate skeletons for popular frameworks.
//   - Customize projects with various options.
//   - Automatically install dependencies and create Dockerfiles where applicable.
func main() {
	// rootCmd defines the root command for the Infocusp Projects CLI.
	// It provides a brief description of the CLI and the available features.
	var rootCmd = &cobra.Command{
		Use:   "infocusp",
		Short: `Welcome to Infocusp Projects CLI!
This powerful and intuitive tool helps you quickly bootstrap skeletons for different tech stacks, 
create demo projects, and streamline development workflows.

With Infocusp Projects CLI, you can:
- 🚀 Scaffold full-stack applications in minutes.
- 🔥 Generate skeletons for popular frameworks like React, FastAPI, and more.
- 🧰 Customize your projects with options like Tailwind CSS, TypeScript, linting, testing frameworks, and Docker.
- 🛠️ Explore ready-to-use demo projects for rapid prototyping.

More commands and features are on the way!

Get started by running one of the commands and let Infocusp Projects CLI handle the heavy lifting so you can focus on building something awesome!

Happy Coding! 😎
		`,
	}

	// Add command for creating a React app.
	rootCmd.AddCommand(commands.CreateReactAppCmd())

	// Add command for creating a Flask skeleton project.
	rootCmd.AddCommand(commands.CreateFlaskSkeletonCmd())

	// Add command for creating a FastAPI skeleton project.
	rootCmd.AddCommand(commands.CreateFastAPISkeletonCmd())

	// Execute the root command to start the CLI.
	// This will listen for any subcommands (such as 'create-react-app', 'create-flask-skeleton', etc.)
	// and delegate the processing to the respective functions.
	rootCmd.Execute()
}
