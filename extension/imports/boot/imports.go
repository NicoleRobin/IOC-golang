//go:build !disableAOP

/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package boot

import (
	_ "github.com/alibaba/ioc-golang/extension/aop/call"
	_ "github.com/alibaba/ioc-golang/extension/aop/dynamic_plugin"
	_ "github.com/alibaba/ioc-golang/extension/aop/list"
	_ "github.com/alibaba/ioc-golang/extension/aop/monitor"
	_ "github.com/alibaba/ioc-golang/extension/aop/trace"
	_ "github.com/alibaba/ioc-golang/extension/aop/transaction"
	_ "github.com/alibaba/ioc-golang/extension/aop/watch"
)
