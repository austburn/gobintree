package main

import (
    "math"
    "strconv"
)

type Node struct {
    value     int
    leftNode  *Node
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

func search(x int) int {
    return searchPrivate(x, root)
}

func searchPrivate(x int, n *Node) int {
    if n == nil {
        return -1
    }

    if n.value == x {
        return n.value
    }

    if n.value > x {
        return searchPrivate(x, n.leftNode)
    } else {
        return searchPrivate(x, n.rightNode)
    }
}

func getDepth() float64 {
    return depth(root, 0)
}

func depth(n *Node, d float64) float64 {
    if n == nil {
        return d
    } else {
        d = d + 1
    }

    if n.leftNode != nil && n.rightNode != nil {
        return math.Max(depth(n.leftNode, d), depth(n.rightNode, d))
    } else if n.leftNode != nil {
        return depth(n.leftNode, d)
    } else {
        return depth(n.rightNode, d)
    }
}

func getPreOrderTraversal() string {
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

func getPostOrderTraversal() string {
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

func getInOrderTraversal() string {
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
