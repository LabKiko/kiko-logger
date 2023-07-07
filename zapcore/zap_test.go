/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/7
* Time: 14:47
* Software: GoLand
 */

package zapcore

import (
	"context"
	"testing"
)

func TestNewRun(t *testing.T) {
	log := New(WithFilename("test"), WithBasePath("./"), WithConsole(true), WithDisableDisk(false))
	ctx := context.Background()
	log.WithContext(ctx).WithFields(map[string]interface{}{"datda": "value"}).Info("data")
	log.WithContext(ctx).Error("failed to call")
	log.WithContext(ctx).Errorf("failed to call %s", "aaaaaaaaaaa")
	defer log.Sync()
}
