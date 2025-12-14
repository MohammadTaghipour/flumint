package task

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const dynamicDartTemplate = `/// AUTO-GENERATED FILE
/// DO NOT MODIFY

class AppConfig {
  static const config = %s;
}`

func toCamelCase(s string) string {
	parts := strings.Split(s, "-")

	caser := cases.Title(language.English)

	for i := 1; i < len(parts); i++ {
		parts[i] = caser.String(parts[i])
	}

	return strings.Join(parts, "")
}

var templateFuncs = template.FuncMap{
	"toCamelCase": toCamelCase,
	"isString":    func(i interface{}) bool { return reflect.TypeOf(i).Kind() == reflect.String },
	"isInt":       func(i interface{}) bool { return reflect.TypeOf(i).Kind() == reflect.Int },
	"isBool":      func(i interface{}) bool { return reflect.TypeOf(i).Kind() == reflect.Bool },
}

func GenerateDartFromMap(data map[string]interface{}, outputPath string) error {
	dartBody := ToDart(data, 1)

	dartCode := fmt.Sprintf(dynamicDartTemplate, dartBody)

	return os.WriteFile(outputPath, []byte(dartCode), 0644)
}

func ToDart(value interface{}, indent int) string {
	space := strings.Repeat("  ", indent)

	switch v := value.(type) {

	case map[string]interface{}:
		var b strings.Builder
		b.WriteString("{\n")
		for k, val := range v {
			b.WriteString(fmt.Sprintf(
				"%s  '%s': %s,\n",
				space,
				k,
				ToDart(val, indent+1),
			))
		}
		b.WriteString(space + "}")
		return b.String()

	case []interface{}:
		var b strings.Builder
		b.WriteString("[\n")
		for _, item := range v {
			b.WriteString(fmt.Sprintf(
				"%s  %s,\n",
				space,
				ToDart(item, indent+1),
			))
		}
		b.WriteString(space + "]")
		return b.String()

	case string:
		return fmt.Sprintf("'%s'", v)

	case bool, int, float64:
		return fmt.Sprintf("%v", v)

	default:
		return "null"
	}
}
