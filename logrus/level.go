/**
* Author: Jason
* Email:  jason_w96@163.com
* Date: 2023/7/3
* Time: 18:48
* Software: GoLand
 */

package logrus

import (
	kLogger "github.com/LabKiko.kiko-logger"
	"github.com/sirupsen/logrus"
)

type Level int8

var (
	LevelMap = map[kLogger.Level]logrus.Level{
		kLogger.DebugLevel: logrus.InfoLevel,
		kLogger.ErrorLevel: logrus.ErrorLevel,
		kLogger.WarnLevel:  logrus.WarnLevel,
		kLogger.FatalLevel: logrus.FatalLevel,
		kLogger.InfoLevel:  logrus.InfoLevel,
	}
)
