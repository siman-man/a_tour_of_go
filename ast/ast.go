package main

import (
    "fmt"
    "go/parser"
    "go/ast"
    "github.com/fatih/color"
)

const t_line = "├────"
const l_line = "└────"

const SPACE = "        "
const SPACE_WITH_I = "│       "

func main() {
    doc := `
    1 + 2 / 3
    `
    expr, err := parser.ParseExpr(doc)
    if err != nil {
        panic(err)
    }

    traverse(expr, "", 0, true)
}

func traverse(n ast.Node, s string, d int, c bool) {
    blue := color.New(color.FgBlue).SprintFunc()

    switch n := n.(type) {
    case *ast.Ident:
        fmt.Printf(" %s\n", n.Name)
    case *ast.BinaryExpr:
        if (d == 0) {
            color.Magenta("BinaryExpr\n")
        } else {
            color.Magenta(" BinaryExpr\n")
        }

        fmt.Print(s + t_line)
        if (is_leaf_node(n.X)) {
            traverse(n.X, s, d+1, false)
        } else {
            traverse(n.X, s+SPACE_WITH_I, d+1, false)
        }

        fmt.Printf("%s :%s\n", s+t_line, n.Op)

        fmt.Print(s + l_line)
        if (is_leaf_node(n.Y)) {
            traverse(n.Y, s, d+1, true)
        } else {
            if (c) {
                traverse(n.Y, s+SPACE, d+1, true)
            } else {
                traverse(n.Y, s+SPACE_WITH_I, d+1, true)
            }
        }
    case *ast.UnaryExpr:
        traverse(n.X, "  ", d+1, true)
    case *ast.BasicLit:
        fmt.Printf(" :%s [%s]\n", n.Kind, blue(n.Value))
    case *ast.ParenExpr:
        color.Magenta(" ParenExpr\n")
        fmt.Print(s + l_line)
        traverse(n.X, s+SPACE, d+1, true)
    default:
        fmt.Println("Hello")
    }
}

func is_leaf_node(n ast.Node) bool {
    switch n.(type) {
    case *ast.Ident:
        return true
    case *ast.BinaryExpr:
        return false
    case *ast.UnaryExpr:
        return true
    case *ast.BasicLit:
        return true
    }

    return false
}
