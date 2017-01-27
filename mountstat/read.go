//
// Copyright © 2017 Bryan T. Meyers <bmeyers@datadrake.com>
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

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

var line string

// ErrNotStatLine is an internal error for reporting the end of the RPC stats section
var ErrNotStatLine = errors.New("Not a Stat Line")

func readLine(b *bufio.Reader) (err error) {
	line, err = b.ReadString('\n')
	return
}

func parseDescription() (n *NFSStats) {
	d := strings.Split(line, " ")
	if len(d) < 8 || d[7] != "nfs" {
		return
	}
	n = &NFSStats{
		Remote: d[1],
		Local:  d[4],
	}
	return
}

func parseStatLine(b *bufio.Reader, labels []string) (key string, stats OpStats, err error) {
	err = readLine(b)
	if err != nil {
		return
	}
	kv := strings.Split(line, ":")
	if len(kv) != 2 {
		err = ErrNotStatLine
		return
	}
	key = strings.TrimSpace(kv[0])
	stats = make(OpStats)
	vs := strings.Split(strings.TrimSpace(kv[1]), " ")
	for i, v := range vs {
		stats[labels[i]], _ = strconv.ParseUint(v, 10, 64)
	}
	return
}

func (n *NFSStats) readOps(b *bufio.Reader) (err error) {
	n.Ops = make(map[string]OpStats)
	for {
		op, stats, e := parseStatLine(b, rpcKeys)
		if e != nil {
			if e != ErrNotStatLine {
				err = e
			}
			return
		}
		if len(op) == 0 {
			continue
		}
		n.Ops[op] = stats
	}
}

func parseNFSStat(b *bufio.Reader) (n *NFSStats, err error) {
	if len(line) == 0 {
		//Read Mount Description line
		err = readLine(b)
		if err != nil {
			return
		}
	}
	n = parseDescription()
	if n == nil {
		return
	}
	//Skip Ahead to Events
	for i := 0; i < 4; i++ {
		err = readLine(b)
		if err != nil {
			return
		}
	}
	//Read Event Stats
	_, n.Events, err = parseStatLine(b, eventLabels)
	if err != nil {
		return
	}
	//Read Byte Stats
	_, n.Bytes, err = parseStatLine(b, byteLabels)
	if err != nil {
		return
	}
	//skip line
	err = readLine(b)
	if err != nil {
		return
	}
	//Read XPRT Stats
	_, n.XPRT, err = parseStatLine(b, xprtKeys)
	if err != nil {
		return
	}
	//skip line
	err = readLine(b)
	if err != nil {
		return
	}
	//Read Op stats
	err = n.readOps(b)
	return
}

// ReadMountStats gets all of the NFS entries in /proc/self/mountstats
func ReadMountStats() (stats []*NFSStats, err error) {
	f, err := os.Open("/proc/self/mountstats")
	if err != nil {
		return
	}
	defer f.Close()
	b := bufio.NewReader(f)
	stats = make([]*NFSStats, 0)
	for {
		line = ""
		stat, e := parseNFSStat(b)
		if e != nil {
			if e != io.EOF {
				err = e
			}
			return
		}
		if stat == nil {
			continue
		}
		stats = append(stats, stat)
	}
}
