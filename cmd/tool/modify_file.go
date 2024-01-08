package tool

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func modifyMethodInFile(filename, receiverTypeName, methodName, newContent string) error {
	fset := token.NewFileSet()

	// Parse the Go source file
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// Iterate through the declarations in the file
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
			// Not a method declaration
			continue
		}

		recvType, ok := funcDecl.Recv.List[0].Type.(*ast.StarExpr) // Check if the receiver is of pointer type
		if !ok {
			continue
		}

		ident, ok := recvType.X.(*ast.Ident)
		if !ok || ident.Name != receiverTypeName || funcDecl.Name.Name != methodName {
			// The receiver type or method name does not match
			continue
		}

		// Clear the existing body and add new content as a comment or basic statement
		funcDecl.Body = &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.BasicLit{
						Kind:  token.STRING,
						Value: newContent,
					},
				},
			},
		}
	}

	// Create a new file to write the modified contents
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write the modified AST back to the file
	if err := printer.Fprint(outputFile, fset, node); err != nil {
		return err
	}

	return nil
}
