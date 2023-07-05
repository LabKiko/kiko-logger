/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2023/7/5
* Time: 16:44
* Software: GoLand
 */

package kiko_logger

type Hook func() *HookStruct

type HookStruct struct {
	LoggerName string
}
