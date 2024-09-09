package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// CreateReactAppCmd defines a Cobra command to generate a React application
// with options for Tailwind CSS, ESLint, TypeScript, and a testing framework.
//
// The command prompts the user for a project name and additional configurations
// such as whether to include Tailwind CSS, TypeScript, linting, and which
// testing framework to use.
//
// Returns:
//   *cobra.Command: A Cobra command object to run the React project generator.
func CreateReactAppCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "create-react-skeleton",
        Short: "Create a React app with custom options",
        Run: func(cmd *cobra.Command, args []string) {
            var projectName, useTailwind, useLinting, testingFramework, useTypeScript string

            // Prompt the user for the project name
            prompt := promptui.Prompt{Label: "Project Name"}
            projectName, _ = prompt.Run()

            // Prompt the user to decide if they want to include Tailwind CSS
            tailwindPrompt := promptui.Select{
                Label: "Do you want to include Tailwind CSS?",
                Items: []string{"Yes", "No"},
            }
            _, useTailwind, _ = tailwindPrompt.Run()

            // Prompt the user to decide if they want to include ESLint
            lintPrompt := promptui.Select{
                Label: "Do you want to include Linting (ESLint)?",
                Items: []string{"Yes", "No"},
            }
            _, useLinting, _ = lintPrompt.Run()

            // Prompt the user to select a testing framework
            testPrompt := promptui.Select{
                Label: "Choose a testing framework",
                Items: []string{"Jest", "Mocha", "None"},
            }
            _, testingFramework, _ = testPrompt.Run()

            // Prompt the user to decide if they want to use TypeScript
            tsPrompt := promptui.Select{
                Label: "Do you want to use TypeScript?",
                Items: []string{"Yes", "No"},
            }
            _, useTypeScript, _ = tsPrompt.Run()

            // Call the function to handle React project setup with the given user input
            CreateReactApp(projectName, useTailwind, useLinting, testingFramework, useTypeScript)
        },
    }
}

// CreateReactApp sets up a React project using the given configurations.
// It supports the following optional customizations:
// - Tailwind CSS integration
// - ESLint setup for linting
// - Testing frameworks (Jest or Mocha)
// - TypeScript support
//
// Parameters:
//   projectName (string): The name of the React project to be created.
//   useTailwind (string): "Yes" to include Tailwind CSS, "No" otherwise.
//   useLinting (string): "Yes" to include ESLint for linting, "No" otherwise.
//   testingFramework (string): The chosen testing framework ("Jest", "Mocha", or "None").
//   useTypeScript (string): "Yes" to use TypeScript, "No" otherwise.
//
// The function uses `npx create-react-app` to initialize the React project,
// and conditionally installs and configures Tailwind CSS, ESLint, and the
// selected testing framework based on the user's inputs.
func CreateReactApp(projectName, useTailwind, useLinting, testingFramework, useTypeScript string) {
    var createAppCmd *exec.Cmd

    // Determine whether to create the React app with or without TypeScript
    if strings.ToLower(useTypeScript) == "yes" {
        createAppCmd = exec.Command("npx", "create-react-app", projectName, "--template", "typescript")
    } else {
        createAppCmd = exec.Command("npx", "create-react-app", projectName)
    }

    // Run the command to create the React project
    createAppCmd.Stdout = os.Stdout
    createAppCmd.Stderr = os.Stderr
    createAppCmd.Run()

    // Change directory to the newly created project
    os.Chdir(projectName)

    // If the user selected Tailwind CSS, install and configure it
    if strings.ToLower(useTailwind) == "yes" {
        fmt.Println("Installing Tailwind CSS...")
        exec.Command("npm", "install", "-D", "tailwindcss", "postcss", "autoprefixer").Run()
        exec.Command("npx", "tailwindcss", "init").Run()
    }

    // If the user selected ESLint, set up linting
    if strings.ToLower(useLinting) == "yes" {
        fmt.Println("Setting up ESLint...")
        exec.Command("npm", "install", "-D", "eslint").Run()
        exec.Command("npx", "eslint", "--init").Run()
    }

    // Install the selected testing framework (Jest or Mocha)
    if testingFramework == "Jest" {
        fmt.Println("Setting up Jest...")
        exec.Command("npm", "install", "--save-dev", "jest").Run()
    } else if testingFramework == "Mocha" {
        fmt.Println("Setting up Mocha...")
        exec.Command("npm", "install", "--save-dev", "mocha").Run()
    }

    // Output a message indicating successful project creation
    fmt.Printf("Project '%s' created successfully!\n", projectName)
}
