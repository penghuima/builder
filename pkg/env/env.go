// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package env specifies environment variables used to configure buildpack behavior.
package env

import (
	"fmt"
	"os"
	"strconv"
)

const (

	// Runtime is an env var used constrain autodetection in runtime buildpacks or to set runtime name in App Engine buildpacks.
	// Runtime must be respected by each runtime buildpack.
	// Example: `nodejs` will cause the nodejs/runtime buildpack to opt-in.
	Runtime = "FUNC_RUNTIME"

	// RuntimeVersion is an env var used to specify which runtime version to install.
	// RuntimeVersion must be respected by each runtime buildpack.
	// Example: `13.7.0` for Node.js, `1.14.1` for Go.
	RuntimeVersion = "FUNC_RUNTIME_VERSION"

	// DebugMode enables more verbose logging. The value is unused; only the presence of the env var is required to enable.
	DebugMode = "FUNC_DEBUG"

	// DevMode is an env var used to enable development mode in buildpacks.
	// DevMode should be respected by all buildpacks that are not product-specific.
	// Example: `true`, `True`, `1` will enable development mode.
	DevMode = "FUNC_DEVMODE"

	// Entrypoint is an env var used to override the default entrypoint.
	// Entrypoint should be respected by at least one buildpack in builders that are not product-specific.
	// Example: `gunicorn -p :8080 main:app` for Python.
	Entrypoint = "FUNC_ENTRYPOINT"

	// ClearSource is an env var used to clear source files from the final image.
	// Buildpacks for Go and Java support clearing the source.
	ClearSource = "FUNC_CLEAR_SOURCE"

	// Buildable is an env var used to specify the buildable unit to build.
	// Buildable should be respected by buildpacks that build source.
	// Example: `./maindir` for Go will build the package rooted at maindir.
	Buildable = "FUNC_BUILDABLE"

	// BuildArgs is an env var used to append arguments to the build command.
	// Example: `-Pprod` for Maven apps run "mvn clear package ... -Pprod" command.
	BuildArgs = "FUNC_BUILD_ARGS"

	// GAEMain is an env var used to specify path or fully qualified package name of the main package in App Engine buildpacks.
	// Behavior: In Go, the value is cleaned up and passed on to subsequent buildpacks as GOOGLE_BUILDABLE.
	GAEMain = "GAE_YAML_MAIN"

	// AppEngineAPIs is an env var that enables access to App Engine APIs. Set to TRUE to enable.
	// Example: `true`, `True`, `1` will enable API access.
	AppEngineAPIs = "GAE_APP_ENGINE_APIS"

	// FunctionTarget is an env var used to specify function name.
	// FunctionTarget must be respected by all functions-framework buildpacks.
	// Example: `helloWorld` or any exported function name.
	FunctionTarget = "FUNC_NAME"
	// FunctionTargetLaunch is a launch time version of FunctionTarget.
	FunctionTargetLaunch = "FUNCTION_TARGET"

	// FunctionSource is an env var used to specify function source location.
	// FunctionSource must be respected by all functions-framework buildpacks.
	// Example: `./path/to/source` will build the function at the specfied path.
	FunctionSource = "FUNC_SRC"
	// FunctionSourceLaunch is a launch time version of FunctionSource.
	FunctionSourceLaunch = "FUNCTION_SOURCE"

	// FunctionSignatureType is an env var used to specify function signature type.
	// FunctionSignatureType must be respected by all functions-framework buildpacks.
	// Example: `http` for HTTP-triggered functions or `event` for event-triggered functions.
	FunctionSignatureType = "FUNC_TYPE"
	// FunctionSignatureTypeLaunch is a launch time version of FunctionSignatureType.
	FunctionSignatureTypeLaunch = "FUNCTION_SIGNATURE_TYPE"

	// GoGCFlags is an env var used to pass through compilation flags to the Go compiler.
	// Example: `-N -l` is used during debugging to disable optimizations and inlining.
	GoGCFlags = "FUNC_GOGCFLAGS"
	// GoLDFlags is an env var used to pass through linker flags to the Go linker.
	// Example: `-s -w` is sometimes used to strip and reduce binary size.
	GoLDFlags = "FUNC_GOLDFLAGS"
	// GoProxy is an env var used to proxy go mod
	GoProxy = "FUNC_GOPROXY"

	// UseNativeImage is used to enable the GraalVM Java buildpack for native image compilation.
	// Example: `true`, `True`, `1` will enable development mode.
	UseNativeImage = "FUNC_JAVA_USE_NATIVE_IMAGE"

	// LabelPrefix is a prefix for values that will be added to the final
	// built user container. The prefix is stripped and the remainder forms the
	// label key. For example, "GOOGLE_LABEL_ABC=Some-Value" will result in a
	// label on the final container of "abc=Some-Value". The label key itself is
	// lowercased, underscores changed to dashes, and is prefixed with "google.".
	LabelPrefix = "FUNC_LABEL_"
)

// IsDebugMode returns true if the buildpack debug mode is enabled.
func IsDebugMode() (bool, error) {
	return isPresentAndTrue(DebugMode)
}

// IsDevMode indicates that the builder is running in Development mode.
func IsDevMode() (bool, error) {
	return isPresentAndTrue(DevMode)
}

// IsUsingNativeImage returns true if the Java application should be built as a native image.
func IsUsingNativeImage() (bool, error) {
	return isPresentAndTrue(UseNativeImage)
}

// Returns true if the environment variable evaluates to True.
func isPresentAndTrue(varName string) (bool, error) {
	//LookupEnv檢索由鍵命名的環境變量的值。如果環境中存在變量，則返回值(可能為空)，並且布爾值為true。否則，返回值將為空，布爾值將為false
	varValue, present := os.LookupEnv(varName)
	if !present {
		return false, nil
	}
	//strconv是一个类型转换包  其中ParseBool 是将 string--> bool 类型   一般如果环境变量中存在该值，这一步parsebool一般转化失败，返回false
	parsed, err := strconv.ParseBool(varValue)
	if err != nil {
		return false, fmt.Errorf("parsing %s: %v", varName, err) //%v普通占位符  fmt.Errorf根据于格式说明符进行格式化，并将字符串作为满足 error 的值返回，其返回类型是error
	}

	return parsed, nil
}
