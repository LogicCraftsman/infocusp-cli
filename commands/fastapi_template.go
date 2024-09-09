package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// CreateFastAPISkeletonCmd defines a Cobra command to generate a FastAPI skeleton project.
// It prompts the user for a project name, then creates the directory structure and files
// necessary for a basic FastAPI application, including models, schemas, routes, and optional tests.
func CreateFastAPISkeletonCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-fastapi-skeleton",
		Short: "Create a FastAPI project structure with dummy models, schemas, routes, and tests",
		Run: func(cmd *cobra.Command, args []string) {
			// Prompt the user to input the project name
			prompt := promptui.Prompt{
				Label: "Project Name",
			}
			projectName, _ := prompt.Run()

			// Prompt for testing framework
			testPrompt := promptui.Select{
				Label: "Choose a testing framework",
				Items: []string{"unittest", "pytest", "None"},
			}
			_, testingFramework, _ := testPrompt.Run()

			// Create the root project directory
			err := os.Mkdir(projectName, 0755)
			if err != nil {
				fmt.Println("Error creating project directory:", err)
				return
			}

			// Create app directory
			appDir := filepath.Join(projectName, "app")
			os.Mkdir(appDir, 0755)

			// Create __init__.py
			os.WriteFile(filepath.Join(appDir, "__init__.py"), []byte(""), 0644)

			// Create main.py with router imports
			mainPyContent := `from fastapi import FastAPI
from .routes import router

app = FastAPI()

app.include_router(router)

@app.get("/")
def read_root():
    return {"message": "Hello, World!"}
`
			os.WriteFile(filepath.Join(appDir, "main.py"), []byte(mainPyContent), 0644)

			// Create models.py with a dummy model
			modelsContent := `from pydantic import BaseModel

class Item(BaseModel):
    name: str
    description: str = None
    price: float
    tax: float = None
`
			os.WriteFile(filepath.Join(appDir, "models.py"), []byte(modelsContent), 0644)

			// Create schemas.py with Pydantic schema
			schemasContent := `from pydantic import BaseModel

class ItemSchema(BaseModel):
    name: str
    description: str = None
    price: float
    tax: float = None
`
			os.WriteFile(filepath.Join(appDir, "schemas.py"), []byte(schemasContent), 0644)

			// Create routes.py with example endpoints
			routesContent := `from fastapi import APIRouter
from .schemas import ItemSchema
from .models import Item

router = APIRouter()

@router.post("/items/")
def create_item(item: ItemSchema):
    return {"message": "Item created", "item": item}

@router.get("/items/{item_id}")
def get_item(item_id: int):
    return {"message": "Get item", "item_id": item_id}
`
			os.WriteFile(filepath.Join(appDir, "routes.py"), []byte(routesContent), 0644)

			// Create requirements.txt
			requirementsContent := `fastapi
fastapi[standard]
uvicorn[standard]
`
			os.WriteFile(filepath.Join(projectName, "requirements.txt"), []byte(requirementsContent), 0644)

			// Create Dockerfile
			dockerfileContent := `FROM tiangolo/uvicorn-gunicorn-fastapi:python3.8

COPY ./app /app
`
			os.WriteFile(filepath.Join(projectName, "Dockerfile"), []byte(dockerfileContent), 0644)

			// Create .gitignore
			gitignoreContent := `.venv/
__pycache__/
*.pyc
`
			os.WriteFile(filepath.Join(projectName, ".gitignore"), []byte(gitignoreContent), 0644)

			// Set up testing framework based on user choice
			if testingFramework != "None" {
				// Create tests directory
				testsDir := filepath.Join(projectName, "tests")
				os.Mkdir(testsDir, 0755)
				// Create __init__.py file
				os.WriteFile(filepath.Join(testsDir, "__init__.py"), []byte(""), 0644)

				if testingFramework == "pytest" {
					// Write basic pytest test file
					pytestTestContent := `from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)

def test_read_root():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"message": "Hello, World!"}
`
					os.WriteFile(filepath.Join(testsDir, "test_main.py"), []byte(pytestTestContent), 0644)

					// Add pytest to requirements.txt
					requirementsContent += "\npytest\n"
					os.WriteFile(filepath.Join(projectName, "requirements.txt"), []byte(requirementsContent), 0644)

				} else if testingFramework == "unittest" {
					// Write basic unittest test file
					unittestTestContent := `import unittest
from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)

class TestMain(unittest.TestCase):
    def test_read_root(self):
        response = client.get("/")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json(), {"message": "Hello, World!"})

if __name__ == '__main__':
    unittest.main()
`
					os.WriteFile(filepath.Join(testsDir, "test_main.py"), []byte(unittestTestContent), 0644)

					// Add unittest to requirements.txt (though it's part of the standard library)
					// Optionally add pytest to run unittest cases if needed
				}

				// Success message for tests
				fmt.Printf("Testing framework '%s' set up successfully in '%s/tests'.\n", testingFramework, projectName)
			}

			// Success message for the project
			fmt.Printf("FastAPI skeleton project '%s' created successfully!\n", projectName)
		},
	}
}
