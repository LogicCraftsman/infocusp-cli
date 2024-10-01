package commands

import (
	"fmt"
	"log"
	"os"

	"infocusp-projects/constants"

	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func CloneRepo(repoURL string, folderName string) error {
	// Clone the repository using SSH authentication
	_, err := git.PlainClone(folderName, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	return err
}
func PromptRepositorySelection() (string, error) {
	// Create a list of repository names for the user to choose from
	repoNames := []string{
		"Ollama",
		"Go",
	}

	// Map repository names to URLs
	repoMap := map[string]string{
		"Ollama": constants.OllamaRepo,
		"Go":     constants.GoRepo,
	}

	// Create a promptui select prompt for the user to choose the repository
	prompt := promptui.Select{
		Label: "Select Repository to Clone",
		Items: repoNames,
	}

	// Prompt the user and get the selection index
	index, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// Get the selected repository URL from the map
	selectedRepoName := repoNames[index]
	return repoMap[selectedRepoName], nil
}

func CloneRepoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clone-repo",
		Short: "Clone a repository",
		Run: func(cmd *cobra.Command, args []string) {
			// Step 1: Prompt the user to select a repository
			repoURL, err := PromptRepositorySelection()
			if err != nil {
				log.Fatalf("Repository selection failed: %v", err)
			}

			// Step 2: Prompt the user for the folder name
			prompt := promptui.Prompt{
				Label: "Enter Folder Name",
			}

			// Prompt the user for input
			folderName, err := prompt.Run()
			if err != nil {
				log.Fatalf("Folder name input failed: %v", err)
			}

			// Step 3: Clone the repository
			err = CloneRepo(repoURL, folderName)
			if err != nil {
				log.Fatalf("Failed to clone repository: %v", err)
			}

			fmt.Printf("Successfully cloned repository to %s\n", folderName)
		},
	}
}
