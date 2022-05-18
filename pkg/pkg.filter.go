package pkg

import (
	"log"
	"strconv"
	"strings"
)

// filter=role_access|building:tenant,role_name|building manager
func ParseFilter(filter string) map[string]interface{} {
	filters := map[string][]string{}
	if filter != "" {
		for _, value := range strings.Split(filter, ",") {
			values := strings.Split(value, "|")
			filters[values[0]] = strings.Split(values[1], ",")
		}
	}
	filter_return := make(map[string]interface{})
	for k, v := range filters {
		for _, vv := range v {
			if vv == "true" || vv == "false" {
				pb, _ := strconv.ParseBool(vv)
				filter_return[k] = pb
			} else if number, err := strconv.Atoi(vv); err == nil {
				log.Printf("%q looks like a number.\n", v)
				filter_return[k] = number
			} else {
				filter_return[k] = vv
			}
		}
	}
	return filter_return
}
