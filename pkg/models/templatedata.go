package models

// TemplateData will hold whatever kind of dataset we'll pass from handler to template(s). We may pass string, int, float, date and whatever other type exist.Also, we may pass 1,2,3 or n number of strings, floats, dates, int or someother data type. For this, we'll create maps of common data-types which we know will be passed commonly such as string, int, floats but to send any other type such as struct which'll hold other data in it, we'll create special type called Data and it'll map[string] but actual content of it can be anything and in Go, when we're not sure of data-type, we call it interface and because we're calling it datatype, we'll have to put curly braces {} after it.
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
