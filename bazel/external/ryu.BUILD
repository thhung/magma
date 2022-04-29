# Copyright 2022 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@python_deps//:requirements.bzl", "requirement")
load("@rules_python//python:defs.bzl", "py_library")

py_library(
    name = "ryu_patched",
    srcs = glob(
        ["ryu/**/*.py"],
        exclude = ["**/tests/**"],
    ),
    visibility = ["//visibility:public"],
    deps = [
        requirement("oslo-config"),
        requirement("routes"),
        requirement("tinyrpc"),
        requirement("webob"),
        requirement("ovs"),
        requirement("eventlet"),
    ],
)
