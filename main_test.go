package main

import (
	"strings"
	"testing"
)

// Helper function to set up a root node and apply a series of commands
func setupTree(commands []string) *Node {
	root := &Node{Name: "root", IsDir: true}
	for _, command := range commands {
		parts := strings.Split(command, " ")
		if len(parts) > 0 && isValidAction(parts[0]) {
			switch parts[0] {
			case "CREATE":
				if len(parts) > 1 {
					createPath(root, parts[1])
				}
			case "DELETE":
				if len(parts) > 1 {
					deletePath(root, parts[1])
				}
			case "MOVE":
				if len(parts) > 2 {
					root.MoveDirectory(parts[1], parts[2])
				}
			}
		}
	}
	return root
}

func TestCreatePath(t *testing.T) {
	commands := []string{
		"CREATE fruits",
		"CREATE fruits/apples",
	}
	root := setupTree(commands)

	if root.GetChildDirectory("fruits") == nil {
		t.Errorf("Expected 'fruits' directory to be created")
	}
	if root.GetChildDirectory("fruits").GetChildDirectory("apples") == nil {
		t.Errorf("Expected 'apples' directory to be created under 'fruits'")
	}
}

func TestDeletePath(t *testing.T) {
	commands := []string{
		"CREATE fruits",
		"CREATE fruits/apples",
		"DELETE fruits/apples",
	}
	root := setupTree(commands)

	if root.GetChildDirectory("fruits").GetChildDirectory("apples") != nil {
		t.Errorf("Expected 'apples' directory to be deleted")
	}
}

func TestMovePath(t *testing.T) {
	commands := []string{
		"CREATE fruits",
		"CREATE fruits/apples",
		"CREATE vegetables",
		"MOVE fruits/apples vegetables",
	}
	root := setupTree(commands)

	if root.GetChildDirectory("fruits").GetChildDirectory("apples") != nil {
		t.Errorf("Expected 'apples' directory to be moved from 'fruits'")
	}
	if root.GetChildDirectory("vegetables").GetChildDirectory("apples") == nil {
		t.Errorf("Expected 'apples' directory to be moved to 'vegetables'")
	}
}

func TestListSubdirectories(t *testing.T) {
	commands := []string{
		"CREATE fruits",
		"CREATE vegetables",
	}
	root := setupTree(commands)

	var sb strings.Builder
	root.PrintTree(&sb, "")

	output := sb.String()
	expected := "fruits\nvegetables\n"

	if output != expected {
		t.Errorf("Expected output to be %s, but got %s", expected, output)
	}
}

// Modify the PrintTree method to accept a strings.Builder
func (n *Node) PrintTree(sb *strings.Builder, indent string) {
	for _, child := range n.Children {
		sb.WriteString(indent + child.Name + "\n")
		child.PrintTree(sb, indent+"  ")
	}
}
