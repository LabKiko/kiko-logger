/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/5
* Time: 16:44
* Software: GoLand
 */

package kiko_logger

type Hook func() *HookStruct

type HookStruct struct {
	LoggerName string
}
