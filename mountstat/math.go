//
// Copyright 2017 Bryan T. Meyers <bmeyers@datadrake.com>
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
//

package mountstat

func diffStats(curr, prev OpStats) (n OpStats) {
	n = make(OpStats)
	for k, v := range curr {
		n[k] = v - prev[k]
	}
	return
}

func diffOps(curr, prev map[string]OpStats) (n map[string]OpStats) {
	n = make(map[string]OpStats)
	for k, v := range curr {
		n[k] = diffStats(v, prev[k])
	}
	return
}

// Diff calculates the change between NFSStats
func Diff(curr, prev *NFSStats) (n *NFSStats) {
	n = &NFSStats{
		Remote: curr.Remote,
		Local:  curr.Local,
		Events: diffStats(curr.Events, prev.Events),
		Bytes:  diffStats(curr.Bytes, prev.Bytes),
		XPRT:   diffStats(curr.XPRT, prev.XPRT),
		Ops:    diffOps(curr.Ops, prev.Ops),
	}
	return
}
