package main

import (
  "strings"
  "io/ioutil"
)

// This is a helper function that loads a text file into a string it returns
// and panics if an error occurs.
func loadTextFilePanicOnError(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// loadHtmlFile loads the file given to it in the parameter, adds the layout file
// around it, parses and executes all @include statments in the file and in included
// files and returns the resulting html string.
func loadHtmlFile(file string) string {
	// load layout file
	layoutHtml := loadTextFilePanicOnError("layout.html")

	// load html files
	html := loadTextFilePanicOnError(file)

	// add layout around the html view
	html = strings.Replace(layoutHtml, "@yield", html, 1)

	// execute @include statements in file
	html = includeFiles(html)

	return html
}

// includeFiles parses and executes @include statements. It recurses until no more
// @include's are left in the file.
func includeFiles(htmlBeforeIncludes string) string {
	var output string

	lines := strings.Split(htmlBeforeIncludes, "\n")

	// this is to check if any new files were included in this pass
	filesIncluded := false

	for _, line := range lines {
		matches := includeRegexp.FindStringSubmatch(line)
		if len(matches) > 0 {
			filesIncluded = true
			output += loadTextFilePanicOnError(matches[1])
		} else {
			output += line + "\n"
		}
	}

	// call this method recursively if there were files included
	// to also handle @include statements in the included files
	if filesIncluded {
		return includeFiles(output)
	}
	return output
}