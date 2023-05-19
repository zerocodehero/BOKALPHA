/*
 @author: lynn
 @date: 2023/5/19
 @time: 09:00
*/

package api

type RequestData struct {
	Url     string
	Params  map[string]string
	Method  string
	Body    map[string]interface{}
	Headers map[string]string
}
