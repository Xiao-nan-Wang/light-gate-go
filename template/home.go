package template

import (
	"strconv"
	"strings"
)

func GetHome(service map[string][]string) string {
	part1 := "<html lang=\"zh\"><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/></head><style>    table {        margin: auto;        text-align: center;    }    table tr {        height: 60px    }    #my-table {        font-family: \"Trebuchet MS\", Arial, Helvetica, sans-serif; width: 100%;border-collapse: collapse;} #my-table td, #my-table th {font-size: 1em;height: 50px; padding: 3px 7px 2px 7px;} #my-table th {font-size: 1.1em; text-align: left; padding-top: 5px; height: 50px; padding-bottom: 4px; background-color: #979797;color: #ffffff;} #my-table tr.alt td {background-color: #e8e8e8;}</style><table id=\"my-table\"><tbody><tr><th style=\"width: 10%;text-align: center;\">服务名</th><th style=\"width: 10%;text-align: center;\">实例数</th><th style=\"width: 80%;text-align: center;\">实例地址</th></tr>"
	part2 := ""
	part3 := "</tbody></table></html>"
	var i = 1
	for key, value := range service {
		if i%2 == 0 {
			part2 += "<tr class=\"alt\">"
		} else {
			part2 += "<tr>"
		}
		part2 += "<td>" + key + "</td><td>" + strconv.Itoa(len(value)) + "</td><td>" + strings.Join(value, ", ") + "</td></tr>"
		i++
	}
	return part1 + part2 + part3
}
