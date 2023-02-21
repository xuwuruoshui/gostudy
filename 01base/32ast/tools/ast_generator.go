package tools

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

// 添加检查程序检查时间
func TimeCalculateByCode(fileName, funcName, outFile string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f.Decls {
		switch node := decl.(type) {
		case *ast.GenDecl:
			node.Specs = append(node.Specs,
				&ast.ImportSpec{Path: &ast.BasicLit{Value: "\"log\"", Kind: token.STRING}},
				&ast.ImportSpec{Path: &ast.BasicLit{Value: "\"time\"", Kind: token.STRING}})
		}
	}
	ast.Inspect(f, func(node ast.Node) bool {
		switch tmp := node.(type) {
		case *ast.FuncDecl:
			if tmp.Name.Name == funcName {
				return true
			} else {
				fmt.Println("其他函数:", tmp.Name.Name)
				return false
			}
		case *ast.BlockStmt:
			//
			pre := &ast.AssignStmt{
				Lhs: []ast.Expr{&ast.Ident{Name: "start"}},
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						Sel: &ast.Ident{Name: "Now"},
						X:   &ast.Ident{Name: "time"},
					},
				}},
				Tok: token.DEFINE,
			}

			end := &ast.AssignStmt{
				Lhs: []ast.Expr{&ast.Ident{Name: "end"}},
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						Sel: &ast.Ident{Name: "Now"},
						X:   &ast.Ident{Name: "time"},
					},
				}},
				Tok: token.DEFINE,
			}

			end2 := &ast.ExprStmt{X: &ast.CallExpr{
				Args: []ast.Expr{&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						Sel: &ast.Ident{Name: "Milliseconds"},
						X: &ast.CallExpr{
							Args: []ast.Expr{&ast.Ident{Name: "start"}},
							Fun: &ast.SelectorExpr{
								Sel: &ast.Ident{
									Name: "Sub",
								},
								X: &ast.Ident{
									Name: "end",
								},
							},
						},
					},
				}},
				Fun: &ast.SelectorExpr{
					Sel: &ast.Ident{
						Name: "Println",
					},
					X: &ast.Ident{
						Name: "log",
					},
				},
			}}
			// 插入前面
			dataSlice := []ast.Stmt{pre}
			tmp.List = append(dataSlice, tmp.List...)
			tmp.List = append(tmp.List, end, end2)

			// 插入后面
		}
		return true
	})

	OutFile(f, fileName, outFile)
	//OutStdout(f)
}

func TimeCalculateTemplate(fileName, templateName, funcName, outFile string) {

	// 读取所需内容
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, templateName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	start := []ast.Stmt{}
	end := []ast.Stmt{}
	for _, decl := range f.Decls {
		switch node := decl.(type) {
		case *ast.FuncDecl:
			start = append(start, node.Body.List[0])
			end = append(end, node.Body.List[1:]...)
		}
	}

	// 解析所需
	fset2 := token.NewFileSet()
	f2, err := parser.ParseFile(fset2, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f2.Decls {
		switch node := decl.(type) {
		case *ast.GenDecl:
			spceMap := map[string]byte{}
			for _, spec := range node.Specs {
				tmp := spec.(*ast.ImportSpec)
				spceMap[tmp.Path.Value] = 0
			}
			for _, spec := range f.Imports {
				tmp := spec
				if _, ok := spceMap[spec.Path.Value]; !ok {
					node.Specs = append(node.Specs, tmp)
					spceMap[spec.Path.Value] = 0
				}
			}
		}
	}
	ast.Inspect(f2, func(node ast.Node) bool {
		switch tmp := node.(type) {
		case *ast.FuncDecl:
			if tmp.Name.Name == funcName {
				return true
			} else {
				fmt.Println("其他函数:", tmp.Name.Name)
				return false
			}
		case *ast.BlockStmt:
			tmp.List = append(start, tmp.List...)
			tmp.List = append(tmp.List, end...)
		}
		return true
	})
	OutFile(f2, fileName, outFile)
}

type EnumGenerator struct {
	Name        string
	Description string
	Data        []*EnumItem
}

type EnumItem struct {
	Key   string
	Value string
}

func (m *EnumGenerator) generate() *ast.GenDecl {
	var genDecl = &ast.GenDecl{
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names:  []*ast.Ident{&ast.Ident{}},
				Values: []ast.Expr{&ast.CompositeLit{}},
			},
		},
	}
	genDecl.Tok = token.VAR
	spec := genDecl.Specs[0].(*ast.ValueSpec)
	spec.Names[0].Name = m.Name
	lit := spec.Values[0].(*ast.CompositeLit)
	// 设置k v类型
	lit.Type = &ast.MapType{Key: &ast.Ident{Name: "int"}, Value: &ast.Ident{Name: "string"}}

	for _, data := range m.Data {
		tmp := data
		// 赋值
		lit.Elts = append(lit.Elts, &ast.KeyValueExpr{
			Key:   &ast.Ident{Name: tmp.Key},
			Value: &ast.BasicLit{Kind: token.STRING, Value: "\"" + tmp.Value + "\""}})
	}

	return genDecl
}

func (m *EnumGenerator) generate1() *dst.GenDecl {
	var genDecl = &dst.GenDecl{
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Names:  []*dst.Ident{&dst.Ident{}},
				Values: []dst.Expr{&dst.CompositeLit{}},
			},
		},
	}
	genDecl.Tok = token.VAR
	spec := genDecl.Specs[0].(*dst.ValueSpec)
	spec.Names[0].Name = m.Name
	lit := spec.Values[0].(*dst.CompositeLit)
	lit.Type = &dst.MapType{Key: &dst.Ident{Name: "int"}, Value: &dst.Ident{Name: "string"}}

	for _, data := range m.Data {
		tmp := data
		// 赋值
		lit.Elts = append(lit.Elts, &dst.KeyValueExpr{
			Key:   &dst.Ident{Name: tmp.Key},
			Value: &dst.BasicLit{Kind: token.STRING, Value: "\"" + tmp.Value + "\""},
			Decs: dst.KeyValueExprDecorations{
				NodeDecs: dst.NodeDecs{Before: dst.NewLine, After: dst.NewLine},
			},
		})
	}

	return genDecl
}
func ConstGenerateMap(fileName string) {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	mg := EnumGenerator{}
	ast.Inspect(f, func(node ast.Node) bool {
		switch tmp := node.(type) {
		case *ast.Comment:
			// 获取颜色参数
			str := strings.TrimSpace(tmp.Text)
			annotations := strings.Fields(str)
			switch annotations[1] {
			case "@Name":
				if len(annotations) != 3 {
					panic("参数必须为3个")
				}
				mg.Name = annotations[2]
			case "@Description":
				if len(annotations) != 3 {
					panic("参数必须为3个")
				}
				mg.Description = annotations[2]
			}
		case *ast.ValueSpec: // 获取数值
			key := tmp.Names[0].Obj.Name
			value := ""
			values := strings.Fields(tmp.Comment.List[0].Text)
			if len(values) == 2 {
				value = values[1]
			}
			mg.Data = append(mg.Data, &EnumItem{Key: key, Value: value})
		}
		return true
	})
	f.Decls = append(f.Decls, mg.generate())

	OutFile(f, fileName, "")
	OutStdout(f)
}

// const生成map
func ConstGenerateMap2(fileName string) {

	file, _ := os.ReadFile(fileName)
	f, err := decorator.Parse(string(file))
	if err != nil {
		panic(err)
	}

	mg := &EnumGenerator{}
	dst.Inspect(f, func(node dst.Node) bool {
		switch tmp := node.(type) {
		case *dst.GenDecl:
			// 获取颜色参数
			for _, data := range tmp.Decs.NodeDecs.Start {
				str := strings.TrimSpace(data)
				annotations := strings.Fields(str)
				if len(annotations) == 3 {
					switch annotations[1] {
					case "@Name":
						if len(annotations) != 3 {
							panic("参数必须为3个")
						}
						mg.Name = annotations[2]
					case "@Description":
						if len(annotations) != 3 {
							panic("参数必须为3个")
						}
						mg.Description = annotations[2]
					}
				}
			}
		case *dst.ValueSpec:
			fmt.Println()
			key := tmp.Names[0].Name
			value := ""
			values := strings.Fields(tmp.Decs.NodeDecs.End[0])
			if len(values) == 2 {
				value = values[1]
			}
			mg.Data = append(mg.Data, &EnumItem{Key: key, Value: value})
		}
		return true
	})

	f.Decls = append(f.Decls, mg.generate1())
	OutFile1(f, fileName, "")

}
