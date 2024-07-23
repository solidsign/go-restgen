package codegen

import (
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	filePath    string
	indentCount int
	body        string
}

func New(module, filename string) *Generator {
	return &Generator{
		filePath:    filename + ".go",
		indentCount: 0,
		body:        "package " + module + "\n\n",
	}
}

func (g *Generator) FuncVoidStart(name string, args ...string) *Generator {
	g.body += g.indent() + "func " + name + "(" + strings.Join(args, ", ") + ") {\n"
	g.indentCount++
	return g
}

func (g *Generator) FuncStart(name string, returns string, args ...string) *Generator {
	g.body += g.indent() + "func " + name + "(" + strings.Join(args, ", ") + ") " + returns + " {\n"
	g.indentCount++
	return g
}

func (g *Generator) FuncEnd() *Generator {
	g.indentCount--
	g.body += g.indent() + "}\n" + "\n"
	return g
}

func (g *Generator) AppendLine(line string) *Generator {
	g.body += g.indent() + line + "\n"
	return g
}

func (g *Generator) Append(str string) *Generator {
	g.body += g.indent() + str
	return g
}

func (g *Generator) Import(imports ...string) *Generator {
	g.body += g.indent() + "import (\n"
	g.indentCount++
	for _, imp := range imports {
		g.body += g.indent() + "\"" + imp + "\"" + "\n"
	}
	g.indentCount--
	g.body += g.indent() + ")" + "\n\n"
	return g
}

func (g *Generator) Struct(name string, fields ...string) *Generator {
	g.body += g.indent() + "type " + name + " struct {\n"
	g.indentCount++
	for i, field := range fields {
		if i < len(fields)-1 {
			g.body += g.indent() + field + "\n"
		} else {
			g.body += g.indent() + field + "," + "\n"
		}
	}
	g.indentCount--
	g.body += g.indent() + "}" + "\n"
	return g
}

func (g *Generator) Interface(name string, methods ...string) *Generator {
	g.body += g.indent() + "type " + name + " interface {\n"
	g.indentCount++
	for i, field := range methods {
		if i < len(methods)-1 {
			g.body += g.indent() + field + "\n"
		} else {
			g.body += g.indent() + field + "," + "\n"
		}
	}
	g.indentCount--
	g.body += g.indent() + "}" + "\n"
	return g
}

func (g *Generator) indent() string {
	return strings.Repeat("\t", g.indentCount)
}

func (g *Generator) Write() error {
	err := os.MkdirAll(filepath.Dir(g.filePath), 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(g.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(g.body)

	return err
}
