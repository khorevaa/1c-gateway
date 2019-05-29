package native

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
)

var ScalarTypes = map[string]bool{
	"String":   true,
	"Int":      true,
	"Int16":    true,
	"Int64":    true,
	"Double":   true,
	"Float":    true,
	"DateTime": true,
	"Binary":   true,
	"Stream":   true,
	"Boolean":  true,
	"Guid":     true,
}

func (g *Generator) GenFilters(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenFilter(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenFilter(source shared.OneCType) string {
	result := ""
	result += g.GenFilterStruct(source)
	result += "\n"
	result += g.GenFilterToStringFunc(source)
	result += "\n"
	result += g.GenFilterScalarToStringFunc(source)
	return result
}

func (g *Generator) GenFilterStruct(source shared.OneCType) string {
	result := ""
	t := g.TranslateType(source.Name)
	result += fmt.Sprintf("type %sFilter struct {\n", t)
	result += fmt.Sprintf("	AND *[]%sFilter\n", t)
	result += fmt.Sprintf("	OR *[]%sFilter\n", t)
	for _, prop := range source.Properties {
		propType := g.TranslateType(prop.Type)
		propName := g.TranslateName(prop.Name)
		fmt.Println(propName)
		if _, ok := ScalarTypes[propType]; ok {
			result += fmt.Sprintf("	%sEq *%s\n", propName, propType)
			result += fmt.Sprintf("	%sNe *%s\n", propName, propType)
		} else {
			if !strings.HasPrefix(propType, "[") {
				result += fmt.Sprintf("	%s: %sFilter\n", propName, propType)
			}
		}
	}
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenFilterToStringFunc(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf(
		`func (t %sFilter) ToString() string {
	if t.AND != nil && len(*t.AND) > 0 {
		tmp := %sFilter{AND: t.AND}
		t.AND = nil
		*tmp.AND = append(*tmp.AND, t)
		result := "true"
		for _, f := range *tmp.AND {
			result += " and " + f.ToString()
		}
		return "("+result+")"
	}
	if t.OR != nil && len(*t.OR) > 0 {
		tmp := %sFilter{OR: t.OR}
		t.OR = nil
		*tmp.OR = append(*tmp.OR, t)
		result := "false"
		for _, f := range *tmp.OR {
			result += " or " + f.ToString()
		}
		return "("+result+")"
	}
	return t.ScalarToString()
}`, t, t, t)
	return result
}

func (g *Generator) GenFilterScalarToStringFunc(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf(
		`func (f *%sFilter) ScalarToString() string {
	result := "true"
`, t)
	for _, prop := range source.Properties {
		propType := g.TranslateType(prop.Type)
		propName := g.TranslateName(prop.Name)
		if _, ok := ScalarTypes[propType]; ok {
			result += fmt.Sprintf("	if f.%sNe != nil {\n", propName)
			result += fmt.Sprintf(`		result += " and %s ne " + f.%sNe.AsParameter()`+"\n", prop.Name, propName)
			result += fmt.Sprintf("	}\n")
			result += fmt.Sprintf(`	if f.%sEq != nil {`+"\n", propName)
			result += fmt.Sprintf(`		result += " and %s eq " + f.%sEq.AsParameter()`+"\n", prop.Name, propName)
			result += fmt.Sprintf("	}\n")
		} else {
			if !strings.HasPrefix(propType, "[") {
				result += fmt.Sprintf("	%s: %sFilter\n", propName, propType)
			}
		}
	}
	result += fmt.Sprintf("	return `(` + result + `)`\n}")
	return result
}