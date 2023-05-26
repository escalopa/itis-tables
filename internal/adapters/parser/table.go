package parser

import (
	"context"
	"strings"
	"time"

	"github.com/escalopa/itis-tables/core"
	"github.com/escalopa/itis-tables/internal/application"
	"golang.org/x/net/html"
)

type TableParser struct {
	tr application.TableRepository
}

func NewTableParser(tr application.TableRepository) *TableParser {
	return &TableParser{tr: tr}
}

func (tp *TableParser) PraseTable(ctx context.Context, doc *html.Node) (map[time.Weekday][][]core.Subject, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Get the table body (<tbody> tag)
	doc = extracrtFirstTag(doc, "tbody")

	if doc == nil {
		return nil, core.ErrNotFound("table body tag not found")
	}

	table := make(map[time.Weekday][][]core.Subject, 0)
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "tr" {
			period := doTR(c)
			if len(period) == 0 {
				continue
			}
			if len(period) < len(core.DaysInOrder) {
				return nil, core.ErrInvalid("table row length is not equal to the number of days")
			}
			for i, day := range core.DaysInOrder {
				table[day] = append(table[day], period[i])
			}
		}
	}

	return table, nil
}

// Get the table rows (<tr> tags)
// Notice that tr tag represents a classes for alld days for a specific period of time
// For example a random tr tag will represent all the classes for all week days for the period betwen 8:30 and 10:00
func doTR(c *html.Node) [][]core.Subject {
	periodClasses := make([][]core.Subject, 0)
	cc := c.FirstChild.NextSibling // Skip the first <td> tag because it contains the period of time
	for ; cc != nil; cc = cc.NextSibling {
		if cc.Type == html.ElementNode && cc.Data == "td" {
			subjects := doTD(cc)
			// We append the subjects without checking if it is nil because we want to keep the order
			periodClasses = append(periodClasses, subjects)
		}
	}
	return periodClasses
}

// Get the classes for a specific day and period of time
func doTD(n *html.Node) []core.Subject {
	n = extracrtFirstTag(n, "div")
	if n == nil {
		return nil
	}

	var title []string
	var meta []core.SubjectMeta

	// Get the title and metadata from the <div> tag
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			switch c.Data {
			case "span":
				title = append(title, doSpan(c))
			case "font":
				meta = append(meta, doFont(c))
			}
		}
	}

	// Merge the title and metadata into a slice of subjSubject structs
	// The length of the title and metadata slices should be the same
	// Since for each title we have a metadata(font tags)
	subjects := make([]core.Subject, len(title))
	for i := range subjects {
		subjects[i] = core.Subject{Title: title[i], Meta: meta[i]}
	}
	return subjects
}

// Get title from <span> tag
func doSpan(n *html.Node) string {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			return c.Data
		}
	}
	return ""
}

// Get metadata from <font> tag
func doFont(n *html.Node) core.SubjectMeta {
	var values []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		var value string

		// Check if we have a <a> tag inside the <font> tag and get the value from it
		if c.Data == "a" {
			value = doA(c)
		} else if c.Type == html.TextNode { // If we don't have a <a> tag, get the value from the text node
			value = c.Data
		}
		if len(value) > 0 {
			values = append(values, value)
		}
		value = ""
	}

	values = strings.Split(strings.Join(values, ""), ",")

	if len(values) != 3 {
		return core.SubjectMeta{}
	}

	// Return the core.SubjectMeta struct
	return core.SubjectMeta{
		RoomNumber: strings.TrimSpace(values[0]),
		Location:   strings.TrimSpace(values[1]),
		Teacher:    formatTeacher(values[2]),
	}
}

// Get metadata from <a> tag
func doA(n *html.Node) string {
	if n == nil {
		return ""
	}
	var s string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			s += c.Data + doA(c.FirstChild)
		}
	}
	return s
}

// formatString formats string to remove new lines and commas and extra spaces
// Before:  чет. нед. Путинцев А. В. с 10.02.2023 по 26.05.2023
// After: Путинцев А. В.
func formatTeacher(teacher string) string {
	if strings.Contains(teacher, "нед.") {
		teacher = strings.Split(teacher, "нед")[1] // remove the prefix
	}
	teacher = strings.TrimSpace(strings.Split(teacher, " с ")[0]) // remove the suffix and trim spaces
	return teacher
}

// extracrtFirstTag returns the first node with the given tag name, starting the search at the root n.
func extracrtFirstTag(n *html.Node, tag string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := extracrtFirstTag(c, tag)
		if node != nil {
			return node
		}
	}
	return nil
}
