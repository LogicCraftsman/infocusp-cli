package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// CreateFlaskSkeletonCmd defines a Cobra command to generate a Flask skeleton project.
// It prompts the user for a project name, testing framework, and generates models, routes, schemas, and tests.
func CreateFlaskSkeletonCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-flask-skeleton",
		Short: "Create a Flask project structure with dummy models, schemas, routes, and tests",
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

			// Create main.py with a basic Flask app
			mainPyContent := `from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/')
def index():
    return jsonify({"message": "Hello, World!"})

if __name__ == "__main__":
    app.run(debug=True)
`
			os.WriteFile(filepath.Join(appDir, "main.py"), []byte(mainPyContent), 0644)

			// Create models.py with a dummy model (this can later integrate with an ORM like SQLAlchemy)
			modelsContent := `class Item:
    def __init__(self, name, description, price, tax=None):
        self.name = name
        self.description = description
        self.price = price
        self.tax = tax
`
			os.WriteFile(filepath.Join(appDir, "models.py"), []byte(modelsContent), 0644)

			// Create schemas.py with simple schema logic
			schemasContent := `class ItemSchema:
    def __init__(self, name, description, price, tax=None):
        self.name = name
        self.description = description
        self.price = price
        self.tax = tax
`
			os.WriteFile(filepath.Join(appDir, "schemas.py"), []byte(schemasContent), 0644)

			// Create routes.py with example endpoints
			routesContent := `from flask import Blueprint, jsonify, request
from .schemas import ItemSchema

bp = Blueprint('routes', __name__)

@bp.route('/items', methods=['POST'])
def create_item():
    data = request.json
    item = ItemSchema(**data)
    return jsonify({"message": "Item created", "item": data})

@bp.route('/items/<int:item_id>', methods=['GET'])
def get_item(item_id):
    return jsonify({"message": "Get item", "item_id": item_id})
`
			os.WriteFile(filepath.Join(appDir, "routes.py"), []byte(routesContent), 0644)

			// Create requirements.txt
			requirementsContent := `flask
`
			os.WriteFile(filepath.Join(projectName, "requirements.txt"), []byte(requirementsContent), 0644)

			// Create Dockerfile
			dockerfileContent := `FROM python:3.8-slim

WORKDIR /app

COPY ./app /app

RUN pip install -r requirements.txt

CMD ["python", "main.py"]
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
				os.WriteFile(filepath.Join(testsDir, "__init__.py"), []byte(""), 0644)

				if testingFramework == "pytest" {
					// Write basic pytest test file
					pytestTestContent := `from app.main import app
import pytest

@pytest.fixture
def client():
    app.config['TESTING'] = True
    with app.test_client() as client:
        yield client

def test_index(client):
    rv = client.get('/')
    assert rv.status_code == 200
    assert rv.get_json() == {"message": "Hello, World!"}
`
					os.WriteFile(filepath.Join(testsDir, "test_main.py"), []byte(pytestTestContent), 0644)

					// Add pytest to requirements.txt
					requirementsContent += "\npytest\n"
					os.WriteFile(filepath.Join(projectName, "requirements.txt"), []byte(requirementsContent), 0644)

				} else if testingFramework == "unittest" {
					// Write basic unittest test file
					unittestTestContent := `import unittest
from app.main import app

class TestMain(unittest.TestCase):
    def setUp(self):
        app.config['TESTING'] = True
        self.client = app.test_client()

    def test_index(self):
        rv = self.client.get('/')
        self.assertEqual(rv.status_code, 200)
        self.assertEqual(rv.get_json(), {"message": "Hello, World!"})

if __name__ == '__main__':
    unittest.main()
`
					os.WriteFile(filepath.Join(testsDir, "test_main.py"), []byte(unittestTestContent), 0644)

					// unittest is part of the Python standard library, so no need to add it to requirements.txt
				}

				// Success message for tests
				fmt.Printf("Testing framework '%s' set up successfully in '%s/tests'.\n", testingFramework, projectName)
			}

			// Success message for the project
			fmt.Printf("Flask skeleton project '%s' created successfully!\n", projectName)
		},
	}
}
