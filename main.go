package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	Name     string
	IsDir    bool
	Children []*Node
}

func (n *Node) CreateDirectory(child *Node) {
	n.Children = append(n.Children, child)
	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].Name < n.Children[j].Name
	})
}

func (n *Node) DeleteDirectory(name string) bool {
	for i, child := range n.Children {
		if child.Name == name {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			return true
		}
	}
	return false
}

func (n *Node) FindDirectory(name string) *Node {
	if n.Name == name {
		return n
	}
	for _, child := range n.Children {
		if result := child.FindDirectory(name); result != nil {
			return result
		}
	}
	return nil
}

func (n *Node) GetChildDirectory(name string) *Node {
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

func (n *Node) MoveDirectory(srcPath, destPath string) bool {
	srcParts := strings.Split(srcPath, "/")
	destParts := strings.Split(destPath, "/")

	srcParent := n
	for i := 0; i < len(srcParts)-1; i++ {
		srcParent = srcParent.FindDirectory(srcParts[i])
		if srcParent == nil {
			fmt.Println("Source parent not found:", srcParts[i])
			return false
		}
	}
	nodeToMove := srcParent.FindDirectory(srcParts[len(srcParts)-1])
	if nodeToMove == nil {
		fmt.Println("Node to move not found:", srcParts[len(srcParts)-1])
		return false
	}

	destParent := n
	for _, part := range destParts {
		child := destParent.FindDirectory(part)
		if child == nil {
			child = &Node{Name: part, IsDir: true}
			destParent.CreateDirectory(child)
		}
		destParent = child
	}

	if destParent.FindDirectory(srcParts[len(srcParts)-1]) != nil {
		fmt.Println("Cannot move a node to its descendant:", srcParts[len(srcParts)-1])
		return false
	}

	if !srcParent.DeleteDirectory(srcParts[len(srcParts)-1]) {
		fmt.Println("Failed to remove node from current parent:", srcParts[len(srcParts)-1])
		return false
	}

	destParent.CreateDirectory(nodeToMove)
	return true
}

func (n *Node) PrintDirectoryTree(indent string) {
	for _, child := range n.Children {
		fmt.Println(indent + child.Name)
		child.PrintDirectoryTree(indent + "  ")
	}
}

var Actions = map[string]struct{}{
	"CREATE": {},
	"LIST":   {},
	"DELETE": {},
	"MOVE":   {},
}

func isValidAction(action string) bool {
	_, exists := Actions[action]
	return exists
}

func createPath(root *Node, path string) {
	parts := strings.Split(path, "/")
	currentNode := root

	for _, part := range parts {
		childNode := currentNode.FindDirectory(part)
		if childNode == nil {
			childNode = &Node{Name: part, IsDir: true}
			currentNode.CreateDirectory(childNode)
		}
		currentNode = childNode
	}
}

func deletePath(root *Node, path string) bool {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return false
	}

	currentNode := root
	for i := 0; i < len(parts)-1; i++ {
		currentNode = currentNode.GetChildDirectory(parts[i])
		if currentNode == nil {
			return false
		}
	}

	return currentNode.DeleteDirectory(parts[len(parts)-1])
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := &Node{Name: "/", IsDir: true}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) > 0 && isValidAction(parts[0]) {
			switch parts[0] {
			case "CREATE":
				if len(parts) > 1 {
					createPath(root, parts[1])
					fmt.Println("CREATE", parts[1])
				}
			case "LIST":
				fmt.Println("LIST")
				root.PrintDirectoryTree("")
			case "DELETE":
				if len(parts) > 1 {
					fmt.Println("DELETE", parts[1])
					if !deletePath(root, parts[1]) {
						topLevelDirectory := strings.Split(parts[1], "/")
						fmt.Println("Cannot delete", parts[1], "-", topLevelDirectory[0], "does not exist")
					}
				}
			case "MOVE":
				if len(parts) > 2 {
					if root.MoveDirectory(parts[1], parts[2]) {
						fmt.Println("MOVE", parts[1], parts[2])
					} else {
						fmt.Println("Cannot move:", parts[1], "to", parts[2])
					}
				}
			}
		} else {
			fmt.Println("Invalid action or empty line")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
