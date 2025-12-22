package abstractions

import (
	"testing"
)

func TestNode_NewNode(t *testing.T) {
	tests := map[string]struct {
		name         string
		expectedName string
		expectedNext int
	}{
		"simple_node": {
			name:         "node1",
			expectedName: "node1",
			expectedNext: 0,
		},
		"empty_name": {
			name:         "",
			expectedName: "",
			expectedNext: 0,
		},
		"node_with_special_chars": {
			name:         "node-123_abc",
			expectedName: "node-123_abc",
			expectedNext: 0,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			node := NewNode(tc.name, false)

			if node == nil {
				t.Fatal("Expected node, got nil")
			}

			if node.name != tc.expectedName {
				t.Errorf("Expected name %q, got %q", tc.expectedName, node.name)
			}

			if node.next == nil {
				t.Error("Expected next to be initialized, got nil")
			}

			if len(node.next) != tc.expectedNext {
				t.Errorf("Expected %d next nodes, got %d", tc.expectedNext, len(node.next))
			}
		})
	}
}

func TestNode_AddNext(t *testing.T) {
	tests := map[string]struct {
		nodeName     string
		nextNodes    []string
		expectedSize int
	}{
		"add_single_node": {
			nodeName:     "root",
			nextNodes:    []string{"child1"},
			expectedSize: 1,
		},
		"add_multiple_nodes": {
			nodeName:     "root",
			nextNodes:    []string{"child1", "child2", "child3"},
			expectedSize: 3,
		},
		"add_no_nodes": {
			nodeName:     "root",
			nextNodes:    []string{},
			expectedSize: 0,
		},
		"add_same_name_nodes": {
			nodeName:     "root",
			nextNodes:    []string{"child", "child", "child"},
			expectedSize: 3,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			node := NewNode(tc.nodeName, false)

			for _, nextName := range tc.nextNodes {
				nextNode := NewNode(nextName, false)
				node.AddNext(nextNode)
			}

			if len(node.next) != tc.expectedSize {
				t.Errorf("Expected %d next nodes, got %d", tc.expectedSize, len(node.next))
			}

			// Verify all nodes were added in order
			for i, expectedName := range tc.nextNodes {
				if node.next[i].name != expectedName {
					t.Errorf("Next node %d: expected name %q, got %q", i, expectedName, node.next[i].name)
				}
			}
		})
	}
}

func TestNode_FindNodeByName(t *testing.T) {
	tests := map[string]struct {
		setupTree    func() *Node
		searchName   string
		shouldFind   bool
		expectedName string
	}{
		"find_root_node": {
			setupTree: func() *Node {
				return NewNode("root", false)
			},
			searchName:   "root",
			shouldFind:   true,
			expectedName: "root",
		},
		"find_direct_child": {
			setupTree: func() *Node {
				root := NewNode("root", false)
				child := NewNode("child1", false)
				root.AddNext(child)
				return root
			},
			searchName:   "child1",
			shouldFind:   true,
			expectedName: "child1",
		},
		"find_grandchild": {
			setupTree: func() *Node {
				root := NewNode("root", false)
				child := NewNode("child1", false)
				grandchild := NewNode("grandchild1", false)
				child.AddNext(grandchild)
				root.AddNext(child)
				return root
			},
			searchName:   "grandchild1",
			shouldFind:   true,
			expectedName: "grandchild1",
		},
		"find_in_multiple_branches": {
			setupTree: func() *Node {
				root := NewNode("root", false)
				child1 := NewNode("child1", false)
				child2 := NewNode("child2", false)
				grandchild := NewNode("target", false)
				child2.AddNext(grandchild)
				root.AddNext(child1)
				root.AddNext(child2)
				return root
			},
			searchName:   "target",
			shouldFind:   true,
			expectedName: "target",
		},
		"find_deep_nested_node": {
			setupTree: func() *Node {
				root := NewNode("root", false)
				level1 := NewNode("level1", false)
				level2 := NewNode("level2", false)
				level3 := NewNode("level3", false)
				level4 := NewNode("deep_target", false)
				level3.AddNext(level4)
				level2.AddNext(level3)
				level1.AddNext(level2)
				root.AddNext(level1)
				return root
			},
			searchName:   "deep_target",
			shouldFind:   true,
			expectedName: "deep_target",
		},
		"node_not_found": {
			setupTree: func() *Node {
				root := NewNode("root", false)
				child := NewNode("child1", false)
				root.AddNext(child)
				return root
			},
			searchName: "nonexistent",
			shouldFind: false,
		},
		"search_in_empty_tree": {
			setupTree: func() *Node {
				return NewNode("root", false)
			},
			searchName: "missing",
			shouldFind: false,
		},
		"find_first_occurrence_in_tree": {
			setupTree: func() *Node {
				root := NewNode("root", false)
				child1 := NewNode("duplicate", false)
				child2 := NewNode("child2", false)
				child3 := NewNode("duplicate", false)
				root.AddNext(child1)
				root.AddNext(child2)
				root.AddNext(child3)
				return root
			},
			searchName:   "duplicate",
			shouldFind:   true,
			expectedName: "duplicate",
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			tree := tc.setupTree()
			result := tree.FindNodeByName(tc.searchName)

			if tc.shouldFind {
				if result == nil {
					t.Fatalf("Expected to find node %q, got nil", tc.searchName)
				}
				if result.name != tc.expectedName {
					t.Errorf("Expected node name %q, got %q", tc.expectedName, result.name)
				}
			} else {
				if result != nil {
					t.Errorf("Expected nil for node %q, got node with name %q", tc.searchName, result.name)
				}
			}
		})
	}
}

func TestNode_FindNodeByName_ComplexTree(t *testing.T) {
	// Build a more complex tree structure
	//        root
	//       /    \
	//     A        B
	//    / \      / \
	//   C   D    E   F
	//  /         \
	// G           H

	root := NewNode("root", false)
	nodeA := NewNode("A", false)
	nodeB := NewNode("B", false)
	nodeC := NewNode("C", false)
	nodeD := NewNode("D", false)
	nodeE := NewNode("E", false)
	nodeF := NewNode("F", false)
	nodeG := NewNode("G", false)
	nodeH := NewNode("H", false)

	root.AddNext(nodeA)
	root.AddNext(nodeB)
	nodeA.AddNext(nodeC)
	nodeA.AddNext(nodeD)
	nodeB.AddNext(nodeE)
	nodeB.AddNext(nodeF)
	nodeC.AddNext(nodeG)
	nodeE.AddNext(nodeH)

	tests := map[string]struct {
		searchName string
		shouldFind bool
	}{
		"find_root":        {"root", true},
		"find_A":           {"A", true},
		"find_B":           {"B", true},
		"find_C":           {"C", true},
		"find_D":           {"D", true},
		"find_E":           {"E", true},
		"find_F":           {"F", true},
		"find_G":           {"G", true},
		"find_H":           {"H", true},
		"find_nonexistent": {"Z", false},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			result := root.FindNodeByName(tc.searchName)

			if tc.shouldFind {
				if result == nil {
					t.Errorf("Expected to find node %q, got nil", tc.searchName)
				} else if result.name != tc.searchName {
					t.Errorf("Expected node name %q, got %q", tc.searchName, result.name)
				}
			} else {
				if result != nil {
					t.Errorf("Expected nil for node %q, got node with name %q", tc.searchName, result.name)
				}
			}
		})
	}
}

func TestNode_AddNext_ChainedOperations(t *testing.T) {
	root := NewNode("root", false)
	child1 := NewNode("child1", false)
	child2 := NewNode("child2", false)
	child3 := NewNode("child3", false)

	// Add nodes one by one
	root.AddNext(child1)
	if len(root.next) != 1 {
		t.Errorf("After adding 1 node, expected 1, got %d", len(root.next))
	}

	root.AddNext(child2)
	if len(root.next) != 2 {
		t.Errorf("After adding 2 nodes, expected 2, got %d", len(root.next))
	}

	root.AddNext(child3)
	if len(root.next) != 3 {
		t.Errorf("After adding 3 nodes, expected 3, got %d", len(root.next))
	}

	// Verify order is preserved
	if root.next[0].name != "child1" {
		t.Errorf("First child should be child1, got %s", root.next[0].name)
	}
	if root.next[1].name != "child2" {
		t.Errorf("Second child should be child2, got %s", root.next[1].name)
	}
	if root.next[2].name != "child3" {
		t.Errorf("Third child should be child3, got %s", root.next[2].name)
	}
}
