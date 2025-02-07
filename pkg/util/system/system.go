// Copyright 2024 The frp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !android

package system

// EnableCompatibilityMode enables compatibility mode for different system.
// For example, on Android, the inability to obtain the correct time zone will result in incorrect log time output.
func EnableCompatibilityMode() {
	// 非 Android 平台不需要做任何兼容性处理
}
