package service

import (
	"encoding/json"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/chaitin/veinmind-common-go/service/report/service/templates"
)

type Format string

const (
	CLI  Format = "cli"
	Json Format = "json"
	Html Format = "html"
)

type Table struct {
	eventType string
	title     string
	header    table.Row
	rows      []table.Row
	rowConfig table.RowConfig
	colConfig []table.ColumnConfig
}

func RenderTable(s *Service, f Format) map[string]*Table {
	tables := make(map[string]*Table, 0)
	for _, evt := range s.EventPool.Events {
		level := evt.BasicInfo.Level.String()
		switch f {
		case CLI:
			level = evt.BasicInfo.Level.Color()
		}

		if _, ok := tables[evt.AlertType.String()]; ok {
			tables[evt.AlertType.String()].rows = append(tables[evt.AlertType.String()].rows, evt.AlertDetail.RenderTable(evt.BasicInfo.ID, level)...)
		} else {
			tables[evt.AlertType.String()] = &Table{
				eventType: evt.EventType.String(),
				title:     evt.AlertDetail.RenderTableTitle(),
				header:    evt.AlertDetail.RenderTableHeader(),
				rows:      evt.AlertDetail.RenderTable(evt.BasicInfo.ID, level),
				rowConfig: evt.AlertDetail.RenderRowConfig(),
				colConfig: evt.AlertDetail.RenderColumnConfig(),
			}
		}
	}
	return tables
}

func WriteTable(s *Service) error {
	tables := RenderTable(s, CLI)
	// render table
	for _, t := range tables {
		if t.eventType == "info" && !s.Options.verbose {
			continue
		}
		tw := table.NewWriter()
		tw.SetOutputMirror(os.Stdout)
		tw.SetStyle(table.StyleLight)
		tw.Style().Options.SeparateRows = true
		tw.SetTitle(color.CyanString(t.title))
		tw.SetColumnConfigs(t.colConfig)
		tw.AppendHeader(func() table.Row {
			for i, h := range t.header {
				t.header[i] = color.MagentaString(h.(string))
			}
			return t.header
		}(), t.rowConfig)
		for _, r := range t.rows {
			tw.AppendRow(r, t.rowConfig)
		}
		tw.SetCaption(color.GreenString(" For More Info, Try Using Json Format"))
		tw.Render()
	}

	return nil
}

func WriteJson(s *Service) error {
	outputDir := s.Options.output
	if outputDir == "" {
		outputDir = "."
	}
	file, err := os.Create(path.Join(outputDir, "result.json"))
	if err != nil {
		return err
	}
	defer file.Close()
	eventsBytes, err := json.MarshalIndent(s.EventPool.Events, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(eventsBytes)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte("\n"))
	return err
}

func WriteHtml(s *Service) error {
	outputDir := s.Options.output
	if outputDir == "" {
		outputDir = "."
	}
	file, err := os.Create(path.Join(outputDir, "result.html"))
	if err != nil {
		return err
	}
	defer file.Close()
	tmpl, err := template.ParseFS(templates.TmplFS, "html.tpl")
	if err != nil {
		return err
	}
	data := map[string]template.HTML{}
	tables := RenderTable(s, Html)
	for key, value := range tables {
		tw := table.NewWriter()
		tw.SetStyle(table.StyleLight)
		tw.Style().Options.SeparateRows = true
		tw.Style().HTML = table.HTMLOptions{
			CSSClass: "ui celled structured unstackable table",
		}

		if len(value.rows) > 0 {
			tw.SetColumnConfigs(value.colConfig)
			tw.AppendHeader(value.header, value.rowConfig)
			for _, r := range value.rows {
				tw.AppendRow(r, value.rowConfig)
			}
		} else {
			tw.SetCaption("No Info")
		}
		data[key] = (template.HTML)(ColorHtml(tw.RenderHTML()))
	}
	return tmpl.Execute(file, data)
}

func ColorHtml(data string) string {
	data = strings.Replace(data, "<td>Low</td>", "<td><a class=\"ui blue label\">Low</a></td>", -1)
	data = strings.Replace(data, "<td>Medium</td>", "<td><a class=\"ui yellow label\">Medium</a></td>", -1)
	data = strings.Replace(data, "<td>High</td>", "<td><a class=\"ui red label\">High</a></td>", -1)
	data = strings.Replace(data, "<td>Critical</td>", "<td><a class=\"ui red label\">Critical</a></td>", -1)
	data = strings.Replace(data, "<td>None</td>", "<td><a class=\"ui teal label\">None</a></td>", -1)
	return data
}
