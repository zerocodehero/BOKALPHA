/*
 @author: lynn
 @date: 2023/5/18
 @time: 23:22
*/

package models

type RequestData struct {
	Url     string
	Params  map[string]string
	Method  string
	Body    map[string]interface{}
	Headers map[string]string
}
