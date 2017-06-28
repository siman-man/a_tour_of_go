package main

import (
    "go/token"
    "go/parser"
    "go/ast"
    "strings"
    "fmt"
    "strconv"
    "log"
    "go/types"
    "go/importer"
)

func main() {
    fs := token.NewFileSet()
    f, err := parser.ParseFile(fs, "/Users/siman/Programming/golang/a_tour_of_go/ast/gopher.go", nil, 0)
    count := 0

    config := &types.Config{
        Importer: importer.Default(),
    }

    info := &types.Info{
        Defs: map[*ast.Ident]types.Object{},
        Uses: map[*ast.Ident]types.Object{},
    }

    _, err = config.Check("main", fs, []*ast.File{f}, info)
    if err != nil {
        log.Fatal(err)
    }

    ast.Inspect(f, func(n ast.Node) bool {
        switch t := n.(type) {
        case *ast.Ident:
            if (info.Defs[t] != nil) {
                if _, ok := info.Defs[t].Type().Underlying().(*types.Struct); ok {
                    count += 1
                    fmt.Println(fs.Position(t.Pos()))
                }
            }
            if (info.Uses[t] != nil) {
                if _, ok := info.Uses[t].Type().Underlying().(*types.Struct); ok {
                    count += 1
                    fmt.Println(fs.Position(t.Pos()))
                }
            }
            name := strings.ToLower(t.Name)

            if strings.Contains(name, "gopher") {
                //fmt.Println(fs.Position(t.Pos()))
            }
        case *ast.BasicLit:
            val := t.Value
            v, err := strconv.Unquote(val)
            if err != nil {
                log.Fatal(err)
            }
            name := strings.ToLower(v)
            if strings.Contains(name, "gopher") {
                //fmt.Println(fs.Position(t.Pos()))
            }
        }
        return true
    })

    fmt.Printf("Gopher Count = %d\n", count)

    ast.Print(fs, f)
}
