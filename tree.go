package main

import "strconv"

type Node struct {
    value int
    leftNode *Node
    rightNode *Node
}

var root *Node

func clearRoot() {
    root = nil
}

func add(x int) {
    if root == nil {
        root = new(Node)
        root.value = x
    } else {
        node := new(Node)
        node.value = x
        addNode(root, node)
    }
}

func addNode(n *Node, leaf *Node) *Node {
    if n == nil {
        return leaf
    }

    if n.value > leaf.value {
        n.leftNode = addNode(n.leftNode, leaf)
    } else {
        n.rightNode = addNode(n.rightNode, leaf)
    }
    return n
}

func getPreOrderTraversal(n *Node) string {
    return preOrderTraversal(root, "")
}

func preOrderTraversal(n *Node, s string) string {
    if n == nil {
        return s
    }

    s = s + strconv.Itoa(n.value)

    if n.leftNode != nil {
        s = preOrderTraversal(n.leftNode, s)
    }

    if n.rightNode != nil {
        s = preOrderTraversal(n.rightNode, s)
    }

    return s
}

func getPostOrderTraversal(n *Node) string {
    return postOrderTraversal(root, "")
}

func postOrderTraversal(n *Node, s string) string {
    if n == nil {
        return s
    }

    if n.leftNode != nil {
        s = postOrderTraversal(n.leftNode, s)
    }

    if n.rightNode != nil {
        s = postOrderTraversal(n.rightNode, s)
    }

    s = s + strconv.Itoa(n.value)

    return s
}

func getInOrderTraversal(n *Node) string {
    return inOrderTraversal(root, "")
}

func inOrderTraversal(n *Node, s string) string {
    if n == nil {
        return s
    }

    if n.leftNode != nil {
        s = inOrderTraversal(n.leftNode, s)
    }

    s = s + strconv.Itoa(n.value)

    if n.rightNode != nil {
        s = inOrderTraversal(n.rightNode, s)
    }

    return s
}
