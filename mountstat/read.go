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
    "strconv"
    "strings"
)

var NotNFS = errors.New("Not an NFS mount")

func (n *NFSStats) readDescription(b *bufio.Reader) (err error) {
    l,err := b.ReadString('\n')
    if err != nil {
        return
    }
    d := strings.Split(l, " ")
    if len(d) < 8 || d[7] != "nfs" {
        err = NotNFS
        return
    }
    n.Remote = d[1]
    n.Local = d[4]
    return
}

func readStatLine(b *bufio.Reader, labels []string) (key string, stats OpStats, err error) {
    l,err := b.ReadString('\n')
    if err != nil {
        return
    }
    kv := strings.Split(l, ":")
    key = strings.TrimSpace(kv[0])
    if len(key) < 2 {
        err = errors.New("Empty Line")
        return
    }
    stats = make(OpStats)
    vs := strings.Split(strings.TrimSpace(kv[1])," ")
    for i,v := range vs {
        stats[labels[i]], _ = strconv.ParseUint(v, 10, 64)
    }
    return
}

func (n * NFSStats) readOps(b *bufio.Reader) (err error) {
    var rpcKeys []string
    for k := range rpcLabels {
       rpcKeys = append(rpcKeys,k)
    }
    n.Ops = make(map[string]OpStats)
    for {
        op, stats, e := readStatLine(b, rpcKeys)
        if e == io.EOF {
            err = e
            return
        }
        if len(op) == 0 {
            continue
        }
        n.Ops[op] = stats
    }
    return
}

func ReadNFSStat(b *bufio.Reader) (n *NFSStats, err error){
    n = &NFSStats{}
    //Read Mount Description line
    err = n.readDescription(b)
    if err != nil {
        return
    }
    //Skip Ahead to Events
    for i := 0; i < 4; i++ {
        _,err = b.ReadString('\n')
        if err != nil {
            return
        }
    }
    //Read Event Stats
    _, n.Events, err = readStatLine(b,eventLabels)
    if err != nil {
        return
    }
    //Read Byte Stats
    _, n.Bytes, err = readStatLine(b,byteLabels)
    if err != nil {
        return
    }
    //skip line
    _,err = b.ReadString('\n')
    if err != nil {
        return
    }
    //Read XPRT Stats
    var xprtKeys []string
    for k := range xprtLabels {
       xprtKeys = append(xprtKeys,k)
    }
    _, n.XPRT, err = readStatLine(b,xprtKeys)
    if err != nil {
        return
    }
    //skip line
    _,err = b.ReadString('\n')
    if err != nil {
        return
    }
    //Read Op stats
    err = n.readOps(b)
    return
}