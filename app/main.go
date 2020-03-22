package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

type briefInfo struct {
	nonTag map[string]struct{}
	level  float64
}

var (
	briefMap  = make(map[string]map[string]*briefInfo)
	nonTagMap = make(map[string]struct{})
	tags      []string
	nodes     []app.Node
	akhr      *hello
)

func init() {

	akhr = &hello{}

	tags = []string{
		"治疗",
		"支援",
		"输出",
		"群攻",
		"减速",

		"生存",
		"防护",
		"削弱",
		"位移",
		"控场",

		"爆发",
		"召唤",
		"快速复活",
		// "费用回复",
		// "支援机械",
	}

	var m interface{}

	jsonFile, err := http.Get("https://127.0.0.1:8888/web/akhr.json")
	if err != nil {
		log.Error(err.Error())
		return
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}

	json.Unmarshal(jsonBytes, &m)

	nonTagMap["近战位"] = struct{}{}
	nonTagMap["远程位"] = struct{}{}

	for _, v := range m.([]interface{}) {

		if vv := v.(map[string]interface{}); vv["level"].(float64) < 3 || vv["hidden"].(bool) == true {
			continue
		} else {

			name := vv["name"].(string)

			tagList, ok := vv["tags"].([]interface{})
			if ok && tagList != nil {
				for _, vvv := range tagList {

					switch vvv.(string) {
					case "近战位", "远程位",
						"资深干员", "高级资深干员":
						continue
					}

					if briefMap[vvv.(string)] == nil {
						briefMap[vvv.(string)] = make(map[string]*briefInfo)

					}

					if briefMap[vvv.(string)][name] == nil {
						briefMap[vvv.(string)][name] = &briefInfo{
							nonTag: make(map[string]struct{}),
						}
						briefMap[vvv.(string)][name].nonTag[func() string {
							for _, pos := range tagList {
								switch pos.(string) {
								case "近战位", "远程位":
									return pos.(string)
								}
							}
							return ""
						}()] = struct{}{}
					}

					briefMap[vvv.(string)][name].nonTag[vv["type"].(string)] = struct{}{}
					nonTagMap[vv["type"].(string)] = struct{}{}

					briefMap[vvv.(string)][name].level = vv["level"].(float64)
				}
			}

		}

	}

}

type hello struct {
	app.Compo
	names []string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.Range(tags).Slice(func(i int) app.UI {
				return app.Button().
					OnClick(func(src app.Value, e app.Event) {
						h.names = []string{}
					out:
						for k := range nonTagMap {
							result := make(map[string]map[string]struct{})
							for kk, vv := range briefMap[tags[i]] {
								if _, ok := vv.nonTag[k]; !ok {
									continue
								}
								if vv.level <= 3 {
									continue out
								}
								if result[k] == nil {
									result[k] = make(map[string]struct{})
								}
								result[k][kk] = struct{}{}
							}

							if len(result) != 0 {
								akhr.names = append(akhr.names, fmt.Sprint(result))
								akhr.Update()
							}
						}
					}).
					Body(
						app.Text(tags[i]),
					)
			}),
		),
		app.Range(h.names).Slice(func(i int) app.UI {
			return app.P().Body(
				app.Text(h.names[i]),
			)
		}),
	)
}

func main() {
	app.Route("/", akhr)
	app.Run()
}
