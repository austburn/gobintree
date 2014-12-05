package main

import "testing"

func TestRootInit(t *testing.T) {
    if root != nil {
        t.Errorf("root should be undefined to start")
    }
}

func TestRootInitialAdd(t *testing.T) {
    add(8)

    if root.value != 8 {
        t.Errorf("Intial add does not seem to be working.")
    }

    clearRoot()
}

func TestAddToLeftTree(t *testing.T) {
    add(5)
    add(1)

    if root.value != 5 {
        t.Errorf("Add seemed to not work for the root node")
    }

    if root.rightNode != nil {
        t.Errorf("rightNode should have been left nil")
    }

    if root.leftNode == nil {
        t.Errorf("Looks as though we did not add the leftNode")
    }

    if root.leftNode != nil && root.leftNode.value != 1 {
        t.Errorf("leftNode was created but does not contain the right value")
    }

    clearRoot()
}

func TestAddToRightTree(t *testing.T) {
    add(1)
    add(5)

    if root.value != 1 {
        t.Errorf("Add seemed to not work for the root node")
    }

    if root.leftNode != nil {
        t.Errorf("leftNode should have been left nil")
    }

    if root.rightNode == nil {
        t.Errorf("Looks as though we did not add the rightNode")
    }

    if root.rightNode != nil && root.rightNode.value != 5 {
        t.Errorf("rightNode was created but does not contain the right value")
    }

    clearRoot()
}

func TestAddToBothTrees(t *testing.T) {
    add(5)
    add(1)
    add(9)

    if root.value != 5 {
        t.Errorf("Add seemed to not work for the root node")
    }

    if root.leftNode == nil {
        t.Errorf("Looks as though we did not add the leftNode")
    }

    if root.rightNode == nil {
        t.Errorf("Looks as though we did not add the rightNode")
    }

    if root.leftNode != nil && root.leftNode.value != 1 {
        t.Errorf("leftNode was created but does not contain the right value")
    }

    if root.rightNode != nil && root.rightNode.value != 9 {
        t.Errorf("rightNode was created but does not contain the right value")
    }

    clearRoot()
}

func TestPreOrderTraversal(t *testing.T) {
    add(5)
    add(1)
    add(2)
    add(3)
    add(6)
    add(8)
    add(7)
    add(4)

    preOrder := getPreOrderTraversal(root)
    expectedPreOrder := "51234687"

    if preOrder != expectedPreOrder {
        t.Errorf("Preorder traversal is incorrect, got %v, but expected %v", preOrder, expectedPreOrder)
    }

    clearRoot()
}

func TestPostOrderTraversal(t *testing.T) {
    add(5)
    add(1)
    add(2)
    add(3)
    add(6)
    add(8)
    add(7)
    add(4)

    postOrder := getPostOrderTraversal(root)
    expectedPostOrder := "43217865"

    if postOrder != expectedPostOrder {
        t.Errorf("Postorder traversal is incorrect, got %v, but expected %v", postOrder, expectedPostOrder)
    }

    clearRoot()
}

func TestInOrderTraversal(t *testing.T) {
    add(5)
    add(1)
    add(2)
    add(3)
    add(6)
    add(8)
    add(7)
    add(4)

    inOrder := getInOrderTraversal(root)
    expectedInOrder := "12345678"

    if inOrder != expectedInOrder {
        t.Errorf("Inorder traversal is incorrect, got %v, but expected %v", inOrder, expectedInOrder)
    }

    clearRoot()
}
