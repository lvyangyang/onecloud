// Copyright 2019 Yunion
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

package architecture

/*
                                                   +-----------+                                                   +-----------+
                                                   | RecordSet |                  +---------------+                | RecordSet |                                         +--------+
                                                 __| (A)       |_________         | TrafficPolicy |                | (TXT)     |__                                      _|  Vpc   |
                                               _/  +-----------+         \________| (Google)      |___             +-----------+  \___                                _/ +--------+
                                            __/                                   +---------------+   \____                           \__                            /
                  +-----------------------+/       +-----------+                                           \____   +-----------+         \___  +-----------------+ _/    +--------+
    API           |  DnsZone  example.com |        | RecordSet |                                                \__| RecordSet |             \_| DnsZone abc.app |-------|  Vpc   |
                  |  (Public)             |--------| (AAAA)    |                  +---------------+                | (CAA)     |---------------| (Private)       |       +--------+
                  +-----------------------+        +-----------+                  | TrafficPolicy |                +-----------+             __+-----------------+_
                          ^                \_                                _____| (Telecom)     |____                                  ___/          ^           \_
                          │                  \_    +-----------+     _______/     +---------------+    \_______    +-----------+      __/              |             \_  +--------+
                          │                    \_  | RecordSet |____/                                          \___| RecordSet |  ___/                 |               \_|  Vpc   |
                          │                      \_| (NS)      |                                                   | (PTR)     |_/                     |                 +--------+
                          │                        +-----------+                                                   +-----------+                       |
                          │                                                                                                                            |
                          │                                                                                                                            |
               ___________│____________________________________________________________________________________________________________________________|__________________________________
                          │                                                                                                                            |
                          v                                                                                                                            |
                  +-----------------+                                                                                                                  |
                  |                 |                                                                                                                  v
                  |                 |            +----------+
                  |  example.com <──|──────────> | Account1 |                                                          +----------+           +---------------+
                  |                 |            | (阿里云) |                                                          | Account3 | <-------> |     abc.app   |
                  |                 |            +----------+                              +------------+              | (Aws)    |           +---------------+
                  |                 |                                                      | Account2   |              +----------+
                  |  example.com <──|────────────────────────────────────────────────────> | (腾讯云)   |
   Cache          |                 |                                                      +------------+
                  |                 |
                  |                 |            +----------+
                  |  example.com <──|──────────> | Account4 |
                  |                 |            | (阿里云) |
                  |                 |            +----------+
                  +-----------------+
               ___________________________________________________________________________________________________________________________________________________________________________




                                                *************                           ***************                  *************
                                            ****             ****                   ****               ****          ****             ****
   公有云                                    **     阿里云    **                     **      腾讯云     **            **	  Aws      **
                                            ****             ****                   ****               ****          ****             ****
                                                *************                           ***************                  *************








*/
