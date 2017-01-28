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
	"fmt"
	"io"
	"sort"
)

// OpStats is a map of RPC statistics for each RPC method
type OpStats map[string]int64

// NFSStats is an struct representation of an NFS entry in /proc/self/mountstats
type NFSStats struct {
	Remote string
	Local  string
	Events OpStats
	Bytes  OpStats
	XPRT   OpStats
	Ops    map[string]OpStats
}

var eventLabels = []string{
	"inoderevalidate",
	"dentryrevalidate",
	"datainvalidate",
	"attrinvalidate",
	"vfsopen",
	"vfslookup",
	"vfsaccess",
	"vfsupdatepage",
	"vfsreadpage",
	"vfsreadpages",
	"vfswritepage",
	"vfswritepages",
	"vfsgetdents",
	"vfssetattr",
	"vfsflush",
	"vfsfsync",
	"vfslock",
	"vfsrelease",
	"congestionwait",
	"setattrtrunc",
	"extendwrite",
	"sillyrename",
	"shortread",
	"shortwrite",
	"delay",
	"pnfs_read",
	"pnfs_write",
}

func printEventUsage(w io.Writer) {
	fmt.Fprintf(w, "\033[1mEVENT STATISTICS\033[0m\n\n")
	fmt.Fprint(w, "Configuration Example\n")
	fmt.Fprint(w, "    events:\n")
	fmt.Fprint(w, "        -shortread\n")
	fmt.Fprint(w, "        -shortwrite\n")
	fmt.Fprint(w, "\nAvailable Stats:\n")
	sort.Strings(eventLabels)
	for _, l := range eventLabels {
		fmt.Fprintf(w, "    %s\n", l)
	}
}

var byteLabels = []string{
	"normalreadbytes",
	"normalwritebytes",
	"directreadbytes",
	"directwritebytes",
	"serverreadbytes",
	"serverwritebytes",
	"readpages",
	"writepages",
}

func printByteUsage(w io.Writer) {
	fmt.Fprint(w, "\033[1mBYTE STATISTICS\033[0m\n\n")
	fmt.Fprint(w, "Configuration Example\n")
	fmt.Fprint(w, "    byte:\n")
	fmt.Fprint(w, "        -readpages\n")
	fmt.Fprint(w, "        -writepages\n")
	fmt.Fprint(w, "\nAvailable Stats:\n")
	sort.Strings(byteLabels)
	for _, l := range byteLabels {
		fmt.Fprintf(w, "    %s\n", l)
	}
}

var xprtKeys = []string{
	"protocol",
	"srcport",
	"bind_count",
	"connect_count",
	"connect_time",
	"idle_time",
	"rpcsends",
	"rpcrecvs",
	"badxids",
	"req_u",
	"bklog_u",
	"max_slots",
	"sending_u",
	"pending_u",
}
var xprtLabels = map[string]string{
	"protocol":      "Connection protocol",
	"srcport":       "Ephemeral port",
	"bind_count":    "How many rpcbind operations",
	"connect_count": "How many TCP connects",
	"connect_time":  "How long connects have taken",
	"idle_time":     "How long transport has been idle",
	"rpcsends":      "How many socket sends",
	"rpcrecvs":      "How many socket receives",
	"badxids":       "How many unmatchable XIDs have been received",
	"req_u":         "Average requests on the wire (slot table utilization)",
	"bklog_u":       "backlog queue utilization (average length of baklog queue)",
	"max_slots":     "max rpc_slots used",
	"sending_u":     "send q utilization",
	"pending_u":     "pend q utilization",
}

func printXprtUsage(w io.Writer) {
	fmt.Fprint(w, "\033[1mXPRT STATISTICS\033[0m\n\n")
	fmt.Fprint(w, "Configuration Example\n")
	fmt.Fprint(w, "    xprt:\n")
	fmt.Fprint(w, "        -req_u\n")
	fmt.Fprint(w, "        -bklog_u\n")
	fmt.Fprint(w, "\nAvailable Stats:\n")
	sort.Strings(xprtKeys)
	for _, l := range xprtKeys {
		fmt.Fprintf(w, "    %14s - %s\n", l, xprtLabels[l])
	}
}

var rpcKeys = []string{
	"ops",
	"trans",
	"timeouts",
	"bytes_sent",
	"bytes_recv",
	"queue",
	"rtt",
	"execute",
}

var rpcLabels = map[string]string{
	"ops":        "How many ops of this type have been requested",
	"trans":      "How many transmissions of this op type have been sent",
	"timeouts":   "How many timeouts of this op type have occurred",
	"bytes_sent": "How many bytes have been sent for this op type",
	"bytes_recv": "How many bytes have been received for this op type",
	"queue":      "How long ops of this type have waited in queue before being transmitted (microsecond)",
	"rtt":        "How long the client waited to receive replies of this op type from the server (microsecond)",
	"execute":    "How long ops of this type take to execute (from rpc_init_task to rpc_exit_task) (microsecond)",
}

func printRPCUsage(w io.Writer) {
	fmt.Fprint(w, "\033[1mRPC STATISTICS\033[0m\n\n")
	fmt.Fprint(w, "Configuration Example\n")
	fmt.Fprint(w, "    rpc:\n")
	fmt.Fprint(w, "        READ:\n")
	fmt.Fprint(w, "            -ops\n")
	fmt.Fprint(w, "            -trans\n")
	fmt.Fprint(w, "            -timeouts\n")
	fmt.Fprint(w, "        WRITE:\n")
	fmt.Fprint(w, "            -ops\n")
	fmt.Fprint(w, "            -trans\n")
	fmt.Fprint(w, "\nAvailable Stats:\n")
	sort.Strings(rpcKeys)
	for _, l := range rpcKeys {
		fmt.Fprintf(w, "    %10s - %s\n", l, rpcLabels[l])
	}
}

// PrintStatUsage prints out the configuration documentation for the NFSStats
func PrintStatUsage(w io.Writer) {
	printEventUsage(w)
	fmt.Fprintln(w, "")
	printByteUsage(w)
	fmt.Fprintln(w, "")
	printXprtUsage(w)
	fmt.Fprintln(w, "")
	printRPCUsage(w)
}
