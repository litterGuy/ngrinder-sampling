package utils

import (
	"encoding/json"
	"reflect"
)

func JSONMerger(dst, src []byte) map[string]interface{} {
	var m1, m2 map[string]interface{}

	json.Unmarshal((dst), &m1)
	json.Unmarshal((src), &m2)

	return JsonMerge(m1, m2)
}

func JsonMerge(dst, src map[string]interface{}) map[string]interface{} {
	return jsMerge(dst, src, 0)
}

func jsMerge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	for key, srcVal := range src {

		if dstVal, ok := dst[key]; ok {

			srcMap, srcMapOk := jsMapify(srcVal)
			dstMap, dstMapOk := jsMapify(dstVal)

			if srcMapOk && dstMapOk {
				srcVal = jsMerge(dstMap, srcMap, depth+1)
			}
		}

		dst[key] = srcVal
	}

	return dst
}

func jsMapify(i interface{}) (map[string]interface{}, bool) {

	value := reflect.ValueOf(i)

	if value.Kind() == reflect.Map {

		m := map[string]interface{}{}

		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}

		return m, true
	}

	return map[string]interface{}{}, false
}
