package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v2"
	"log"
	"encoding/json"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request){
		
		path := r.URL.Path;
		val, ok := pathsToUrls[path]
		if ok{
			http.Redirect(w, r, val, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
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


type yamlFormat struct{
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}

type jsonFormat struct{
	Path string `json:"path`
	Url string  `json:"url"`
}

func parseJson(jsonBlob []byte)([]jsonFormat,error){
	var entries []jsonFormat
	err := json.Unmarshal(jsonBlob,&entries)
	if (err != nil){
		return nil, err
	}
	return entries,nil
}

func parseYAML(yml []byte)([]yamlFormat, error){
	//so we need to take the []byte array make a new array of unmarshalled Json From the 
	var entries []yamlFormat
	err := yaml.Unmarshal(yml,&entries)
	if (err != nil){
		return nil , err
	}
	return entries, nil
	
}

func makeMap(jsonEntries []jsonFormat)(map[string]string){
	constructedUrls := make(map[string]string)
	for _,val := range jsonEntries{
		constructedUrls[val.Path] = val.Url
	}
	return constructedUrls
}

func JSONHandler(jsonBlob []byte, fallback http.Handler)(http.HandlerFunc,error){
	parsedJson, err := parseJson(jsonBlob)
	if err != nil{
		return nil, err
	}
	pathMap := makeMap(parsedJson)
	return MapHandler(pathMap,fallback),nil
}

func constructMap(yamlEntries []yamlFormat)(map[string]string){
	constructedUrls := make(map[string]string)
	for _,val := range yamlEntries{
		constructedUrls[val.Path] = val.Url
	}
	return constructedUrls
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := constructMap(parsedYaml)
	log.Printf("constructed pathMap has %d entries", len(pathMap))
	return MapHandler(pathMap, fallback), nil
}


