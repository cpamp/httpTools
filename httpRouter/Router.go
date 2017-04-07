package httpRouter

import (
	"httpTools/httpHelper"
	"net/http"
	"strings"
)

type HTTPVerb string

const (
	HTTPGet    HTTPVerb = "GET"
	HTTPPost            = "POST"
	HTTPPut             = "PUT"
	HTTPUpdate          = "UPDATE"
	HTTPDelete          = "DELETE"
)

type Router struct {
	tree          *node
	baseHandler   Handle
	contentType   httpHelper.ContentType
	safeResponses bool
}

func NewRouter(rootHandler Handle, contentType httpHelper.ContentType) *Router {
	node := node{component: "/", isParam: false, verbs: make(map[string]Handle)}
	return &Router{tree: &node, baseHandler: rootHandler, contentType: contentType, safeResponses: false}
}

func NewRouterSafe(rootHandler Handle, contentType httpHelper.ContentType) *Router {
	node := node{component: "/", isParam: false, verbs: make(map[string]Handle)}
	return &Router{tree: &node, baseHandler: rootHandler, contentType: contentType, safeResponses: true}
}

func (r *Router) Handle(verb HTTPVerb, path string, handler Handle) {
	if len(path) > 0 && path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.tree.addNode(string(verb), path, handler)
}

func (r *Router) RegisterRoutes(routes Routes) {
	for i := 0; i < len(routes); i++ {
		r.Handle(routes[i].Verb, routes[i].Path, routes[i].Handler)
	}
}

func (r *Router) RegisterRouteCollection(routeCollection RouteCollection) {
	for i := 0; i < len(routeCollection); i++ {
		r.RegisterRoutes(routeCollection[i])
	}
}

func (r *Router) Get(path string, handler Handle) {
	r.Handle(HTTPGet, path, handler)
}

func (r *Router) Post(path string, handler Handle) {
	r.Handle(HTTPPost, path, handler)
}

func (r *Router) Put(path string, handler Handle) {
	r.Handle(HTTPPut, path, handler)
}

func (r *Router) Update(path string, handler Handle) {
	r.Handle(HTTPUpdate, path, handler)
}

func (r *Router) Delete(path string, handler Handle) {
	r.Handle(HTTPDelete, path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	responder := httpHelper.NewWriter(w, r.contentType)
	request := httpHelper.NewRequest(req, params)
	if r.safeResponses {
		responder.SafeResponses()
	}
	if handler := node.verbs[req.Method]; handler != nil {
		handler(HandleHelper{responder, request, params})
	} else if r.baseHandler != nil {
		r.baseHandler(HandleHelper{responder, request, params})
	} else {
		responder.NotFound("", nil)
	}
}
