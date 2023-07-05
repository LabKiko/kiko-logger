/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2023/7/5
* Time: 18:14
* Software: GoLand
 */

package zapcore

import (
	"errors"
)

var (
	// ErrLogPathNotSet is an error that indicates the log path is not set.
	ErrLogPathNotSet = errors.New("log path must be set")
)
