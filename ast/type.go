package main

import (
    "fmt"
    "go/ast"
    "go/importer"
    "go/parser"
    "go/token"
    "go/types"
    "log"
)

func main() {
    fs := token.NewFileSet()
    f, err := parser.ParseFile(fs, "/Users/siman/Programming/golang/a_tour_of_go/ast/gopher.go", nil, 0)
    //f, err := parser.ParseFile(fs, "main.go", src, 0)
    if err != nil {
        log.Fatal(err)
    }

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

    for idnt, o := range info.Defs {
        if o == nil {
            continue
        }

        if _, ok := o.Type().Underlying().(*types.Struct); ok {
            fmt.Println(fs.Position(idnt.Pos()), idnt)
        }
    }

    for idnt, o := range info.Uses {
        if o == nil {
            continue
        }

        if _, ok := o.Type().Underlying().(*types.Struct); ok {
            fmt.Println(fs.Position(idnt.Pos()), idnt)
        }
    }

    /*
    for idnt, o := range info.Defs {
        if o != nil && types.Identical(o.Type(), it) {
            fmt.Printf("%T %T\n", idnt, o)
            fmt.Println(fs.Position(idnt.Pos()), idnt)
        }
    }
    */
}