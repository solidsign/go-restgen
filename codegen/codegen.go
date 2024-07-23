package codegen

import "strings"

type Generator struct {
	filePath    string
	indentCount int
	body        string
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) FuncStart(name string, args ...string) *Generator {
	g.body += g.indent() + "func " + name + "(" + strings.Join(args, ", ") + ") {\n"
	g.indentCount++
	return g
}
func (g *Generator) FuncEnd() *Generator {
	g.indentCount--
	g.body += g.indent() + "}\n"
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

func (g *Generator) Module(name string) *Generator {
	g.body += "module " + name
	return g
}

func (g *Generator) Import(imports ...string) *Generator {
	g.body += g.indent() + "import (\n"
	g.indentCount++
	for _, imp := range imports {
		g.body += g.indent() + imp + "\n"
	}
	g.indentCount--
	g.body += g.indent() + ")"
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
	g.body += g.indent() + "}"
	return g
}

func (g *Generator) indent() string {
	return strings.Repeat("\t", g.indentCount)
}
