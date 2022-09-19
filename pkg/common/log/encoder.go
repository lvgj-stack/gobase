/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package log

import (
	"time"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

func zapFields(args ...interface{}) []zap.Field {
	if len(args) == 1 {
		fs, ok := args[0].([]interface{})
		if ok {
			args = fs
		}
	}
	fields := make([]zap.Field, 0, len(args)/2+1)
	for i := 0; i < len(args)-1; i += 2 {
		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			fields = append(fields, zap.Any("exceeds", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			fields = append(fields, zap.Any("invalidKey", val))
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
	}
	return fields
}
