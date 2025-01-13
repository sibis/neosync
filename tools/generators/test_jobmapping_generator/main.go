package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io"
	"os"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/nucleuscloud/go-antlrv4-parser/tsql"
	pg_query "github.com/pganalyze/pg_query_go/v5"
)

type Input struct {
	Folder  string `json:"folder"`
	SqlFile string `json:"sql_file"`
	Driver  string `json:"driver"`
}

type Column struct {
	Name        string
	TypeStr     string
	IsGenerated bool
}

type Table struct {
	Name    string
	Columns []*Column
}

type JobMapping struct {
	Table       string
	Column      string
	Transformer string
	Config      string
}

func parsePostegresStatements(sql string) ([]*Table, error) {
	tree, err := pg_query.Parse(sql)
	if err != nil {
		return nil, err
	}

	tables := []*Table{}
	for _, stmt := range tree.GetStmts() {
		s := stmt.GetStmt()
		switch s.Node.(type) { //nolint:gocritic
		case *pg_query.Node_CreateStmt:
			table := s.GetCreateStmt().GetRelation().GetRelname()
			columns := []*Column{}
			for _, col := range s.GetCreateStmt().GetTableElts() {
				if col.GetColumnDef() != nil {
					colDef := col.GetColumnDef()

					isGenerated := false
					// Check for GENERATED columns in constraints
					for _, constraint := range colDef.Constraints {
						if constraint.GetConstraint() != nil {
							// contype 5 is CONSTR_GENERATED
							if constraint.GetConstraint().Contype == 5 {
								isGenerated = true
								break
							}
						}
					}

					columns = append(columns, &Column{
						Name:        colDef.Colname,
						IsGenerated: isGenerated,
					})
				}
			}
			tables = append(tables, &Table{
				Name:    table,
				Columns: columns,
			})
		}
	}
	return tables, nil
}

var (
	reCreateTable         = regexp.MustCompile(`CREATE\s+TABLE\s+IF\s+NOT\s+EXISTS\s+(\w+)\s*\.\s*(\w+)\s*\(`)
	reCreateTableNoSchema = regexp.MustCompile(`CREATE\s+TABLE\s+IF\s+NOT\s+EXISTS\s+(\w+)\s*\(`)
	reColumn              = regexp.MustCompile(`^\s*(\w+)\s+[\w\(\)]+.*`)
)

// todo fix very brittle
func parseSQLStatements(sql string) []*Table {
	lines := strings.Split(sql, "\n")
	tableColumnsMap := make(map[string][]*Column)
	var currentTable string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if matches := reCreateTable.FindStringSubmatch(line); len(matches) > 2 {
			currentTable = matches[2]
		} else if matches := reCreateTableNoSchema.FindStringSubmatch(line); len(matches) > 1 {
			currentTable = matches[1]
		} else if currentTable != "" {
			if matches := reColumn.FindStringSubmatch(line); len(matches) > 1 {
				columnName := matches[1]
				if slices.Contains([]string{"primary key", "constraint", "key", "unique", "primary", "alter"}, strings.ToLower(matches[1])) {
					continue
				}
				isGenerated := false
				if strings.Contains(line, "GENERATED ALWAYS AS") || strings.Contains(line, "STORED") || strings.Contains(line, "VIRTUAL") {
					isGenerated = true
				}
				tableColumnsMap[currentTable] = append(tableColumnsMap[currentTable], &Column{
					Name:        columnName,
					IsGenerated: isGenerated,
				})
			} else if strings.HasPrefix(line, "PRIMARY KEY") || strings.HasPrefix(line, "CONSTRAINT") || strings.HasPrefix(line, "UNIQUE") || strings.HasPrefix(line, "KEY") || strings.HasPrefix(line, "ENGINE") || strings.HasPrefix(line, ")") {
				// Ignore key constraints and end of table definition
				if strings.HasPrefix(line, ")") {
					currentTable = ""
				}
			}
		}
	}
	res := []*Table{}
	for table, cols := range tableColumnsMap {
		res = append(res, &Table{
			Name:    table,
			Columns: cols,
		})
	}

	// Sort tables by name
	slices.SortFunc(res, func(a, b *Table) int {
		return strings.Compare(a.Name, b.Name)
	})

	return res
}

type TemplateData struct {
	SourceFile      string
	PackageName     string
	Tables          []*Table
	GenerateTypeMap bool
}

func formatJobMappings(pkgName, sqlFile string, tables []*Table, generateTypeMap bool) (string, error) {
	const tmpl = `
// Code generated by Neosync jobmapping_generator. DO NOT EDIT.
// source: {{ .SourceFile }}

package {{ .PackageName }}

import (
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
)

func GetDefaultSyncJobMappings(schema string)[]*mgmtv1alpha1.JobMapping {
  return []*mgmtv1alpha1.JobMapping{
		{{- range .Tables }}
		{{- $tableName := .Name }}
		{{- range .Columns }}
		{
			Schema: schema,
			Table:  "{{ $tableName }}",
			Column: "{{ .Name }}",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config: {{- if .IsGenerated }}
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_GenerateDefaultConfig{
							GenerateDefaultConfig: &mgmtv1alpha1.GenerateDefault{},
						},
					},
				{{- else }}
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
				{{- end }}
			},
		},
		{{- end }}
		{{- end }}
	}
}
{{ if .GenerateTypeMap }}


func GetTableColumnTypeMap() map[string]map[string]string {
	return map[string]map[string]string{
		{{- range .Tables }}
		"{{ .Name }}": {
		{{- range .Columns }}
			"{{ .Name }}": "{{ .TypeStr }}",
		{{- end }}
		},
		{{- end }}
	}
}
{{- end }}
`
	data := TemplateData{
		SourceFile:      sqlFile,
		PackageName:     pkgName,
		Tables:          tables,
		GenerateTypeMap: generateTypeMap,
	}
	t := template.Must(template.New("jobmappings").Parse(tmpl))
	var out bytes.Buffer
	err := t.Execute(&out, data)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("must provide necessary args")
	}

	configFile := args[1]
	gopackage := args[2]

	packageSplit := strings.Split(gopackage, "_")
	goPkg := packageSplit[len(packageSplit)-1]

	jsonFile, err := os.Open(configFile)
	if err != nil {
		fmt.Printf("failed to open file: %s\n", err) //nolint:forbidigo
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read file: %s\n", err) //nolint:forbidigo
		return
	}

	var inputs []Input
	if err := json.Unmarshal(byteValue, &inputs); err != nil {
		fmt.Printf("failed to unmarshal JSON: %s\n", err) //nolint:forbidigo
		return
	}
	for _, input := range inputs {
		folderSplit := strings.Split(input.Folder, "/")
		var goPkgName string
		if len(folderSplit) == 1 {
			goPkgName = strings.ReplaceAll(fmt.Sprintf("%s_%s", goPkg, input.Folder), "-", "")
		} else if len(folderSplit) > 1 {
			lastTwo := folderSplit[len(folderSplit)-2:]
			goPkgName = strings.ReplaceAll(strings.Join(lastTwo, "_"), "-", "")
		}
		sqlFile, err := os.Open(fmt.Sprintf("%s/%s", input.Folder, input.SqlFile))
		if err != nil {
			fmt.Printf("failed to open file: %s\n", err) //nolint:forbidigo
			return
		}

		byteValue, err := io.ReadAll(sqlFile)
		if err != nil {
			fmt.Printf("failed to read file: %s\n", err) //nolint:forbidigo
			return
		}

		sqlContent := string(byteValue)
		sqlFile.Close()

		var tables []*Table
		if input.Driver == "postgres" {
			t, err := parsePostegresStatements(sqlContent)
			if err != nil {
				fmt.Printf("Error parsing postgres SQL schema: %s\n", err) //nolint:forbidigo
				return
			}
			tables = t
		} else if input.Driver == "mysql" {
			t := parseSQLStatements(sqlContent)
			tables = t
		} else if input.Driver == "sqlserver" || input.Driver == "mssql" {
			t := parseTsql(sqlContent)
			tables = t
		}

		formattedJobMappings, err := formatJobMappings(goPkgName, input.SqlFile, tables, input.Driver == "sqlserver")
		if err != nil {
			fmt.Printf("Error formatting job mappings: %s\n", err) //nolint:forbidigo
			return
		}

		output := fmt.Sprintf("%s/job-mappings.go", input.Folder)
		outputFile, err := os.Create(output)
		if err != nil {
			fmt.Printf("Error creating jobmapping.go file: %s\n", err) //nolint:forbidigo
			return
		}

		_, err = outputFile.WriteString(formattedJobMappings)
		if err != nil {
			fmt.Printf("Error writing to jobmapping.go file: %s\n", err) //nolint:forbidigo
			return
		}
		outputFile.Close()
	}

	return
}

type tsqlListener struct {
	*parser.BaseTSqlParserListener
	inCreate     bool
	currentTable string
	currentCols  []*Column
	mappings     []*Table
}

func (l *tsqlListener) PushTable() {
	l.mappings = append(l.mappings, &Table{
		Name:    l.currentTable,
		Columns: l.currentCols,
	})
	l.currentTable = ""
	l.currentCols = []*Column{}
	l.inCreate = false
}

func (l *tsqlListener) PushColumn(name, typeStr string) {
	l.currentCols = append(l.currentCols, &Column{
		Name:    name,
		TypeStr: typeStr,
	})
}

func (l *tsqlListener) SetTable(schemaTable string) {
	split := strings.Split(schemaTable, ".")
	if len(split) == 1 {
		l.currentTable = split[0]
	} else if len(split) > 1 {
		l.currentTable = split[1]
	}
}

// EnterCreate_table is called when production create_table is entered.
func (l *tsqlListener) EnterCreate_table(ctx *parser.Create_tableContext) {
	l.inCreate = true
	table := ctx.Table_name().GetText()
	l.SetTable(table)
}

// ExitCreate_table is called when production create_table is exited.
func (l *tsqlListener) ExitCreate_table(ctx *parser.Create_tableContext) { //nolint:unparam
	l.PushTable()
}
func (l *tsqlListener) EnterColumn_definition(ctx *parser.Column_definitionContext) {
	l.PushColumn(ctx.Id_().GetText(), ctx.Data_type().GetText())
}

func parseTsql(sql string) []*Table {
	inputStream := antlr.NewInputStream(sql)

	// create the lexer
	lexer := parser.NewTSqlLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	p := parser.NewTSqlParser(tokens)

	listener := &tsqlListener{}
	tree := p.Tsql_file()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.mappings
}
