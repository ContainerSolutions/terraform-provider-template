# Copyright 2015 Container Solutions
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

provider "awesome" {
  api_key     = "s3cur3t0k3n=="
  endpoint    = "https://api.example.org/v1"
  timeout     = 60
  max_retries = 5
}

resource "awesome_machine" "my-speedy-server" {
  name = "speedracer"
  cpus = 4
  ram  = 16384
}
