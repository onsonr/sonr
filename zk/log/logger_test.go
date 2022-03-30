/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidLogLevel(t *testing.T) {
	_, err := NewStdoutLogger("test", "Invalid log level", FORMAT_SHORT)
	assert.NotNil(t, err, "should produce an error due to invalid log level")
}

func TestInvalidFormat(t *testing.T) {
	_, err := NewStdoutLogger("test", INFO, "Invalid format")
	assert.NotNil(t, err, "should produce an error due to invalid log format")
}

func TestWithLogFileThatShouldNotBeCreated(t *testing.T) {
	_, err := NewFileLogger("test", "/shouldnotbecreated.txt", INFO, FORMAT_SHORT)
	assert.NotNil(t, err, "should produce an error because of invalid path")
}
