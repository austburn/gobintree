package main

import "testing"

func TestRootInit(t *testing.T) {
    x := search(4)

    if root != nil {
        t.Errorf("root should be undefined to start")
    }

    if x != -1 {
        t.Errorf("Empty tree search should return -1")
    }
}

func TestRootInitialAdd(t *testing.T) {
    add(8)

    if root.value != 8 {
        t.Errorf("Intial add does not seem to be working.")
    }

    x := search(8)
    if x != 8 {
        t.Errorf("should be able to find 8")
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

    x, y := search(5), search(1)
    if x != 5 || y != 1 {
        t.Errorf("Search failed")
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

    x, y := search(5), search(1)
    if x != 5 || y != 1 {
        t.Errorf("Search failed")
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

    preOrder := getPreOrderTraversal()
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

    postOrder := getPostOrderTraversal()
    expectedPostOrder := "43217865"

    if postOrder != expectedPostOrder {
        t.Errorf("Postorder traversal is incorrect, got %v, but expected %v", postOrder, expectedPostOrder)
    }

    random_search := search(6)
    if random_search != 6 {
        t.Errorf("Search failed")
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

    inOrder := getInOrderTraversal()
    expectedInOrder := "12345678"

    if inOrder != expectedInOrder {
        t.Errorf("Inorder traversal is incorrect, got %v, but expected %v", inOrder, expectedInOrder)
    }

    clearRoot()
}

func TestDepth(t *testing.T) {
    if getDepth() != 0 {
        t.Errorf("Initial depth should be 0")
    }

    add(5)
    //5
    if getDepth() != 1 {
        t.Errorf("Depth after add should be 1")
    }

    add(1)
    add(7)
    //1     5
    //2  1     7
    if getDepth() != 2 {
        t.Errorf("Depth after add should be 2")
    }

    add(2)
    //1     5
    //2  1     7
    //3   2
    if getDepth() != 3 {
        t.Errorf("Depth after add should be 3")
    }

    add(3)
    //1      5
    //2   1      7
    //3    2
    //4     3
    if getDepth() != 4 {
        t.Errorf("Depth after add should be 3")
    }

    add(8)
    add(6)
    add(11)
    add(10)
    //1      5
    //2   1      7
    //3    2    6   8
    //4     3         11
    //5              10
    if getDepth() != 5 {
        t.Errorf("Depth after add should be 5")
    }
}
