package urlshort

import (
	json2 "encoding/json"
	"gopkg.in/yaml.v2"
	"net/http"
)

type pathUrl struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if val, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, val, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func buildMap(parsedFile []pathUrl) map[string]string {
	pathMap := map[string]string{}
	for _, pathUrl := range parsedFile {
		pathMap[pathUrl.Path] = pathUrl.URL
	}
	return pathMap
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]pathUrl, error) {
	var pathURLs []pathUrl
	err := yaml.Unmarshal(yml, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
// [
//     {"path": "/value", "url": "https://..."},
//     {"path": "/value", "url": "https://..."},
//		...
// ]
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJson(json)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJson)
	return MapHandler(pathMap, fallback), nil
}

func parseJson(json []byte) ([]pathUrl, error) {
	var pathURLs []pathUrl
	err := json2.Unmarshal(json, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}
