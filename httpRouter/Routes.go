package httpRouter

type Route struct {
	Verb    HTTPVerb
	Path    string
	Handler Handle
}

type Routes []Route

type RouteCollection []Routes
