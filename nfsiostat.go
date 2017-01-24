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

package main

import (
    "bufio"
    "fmt"
	"github.com/DataDrake/go-nfsiostat/mountstat"
    "io"
	"os"
)

func main() {
    mountstat.PrintStatUsage(os.Stdout)
    b := bufio.NewReader(os.Stdin)
    for {
        n,e := mountstat.ReadNFSStat(b)
        if e == mountstat.NotNFS {
            continue
        }
        if n == nil {
            os.Exit(2)
        } else {
            fmt.Printf("%v\n", n.Ops)
        }
        if e == io.EOF {
            return
        }
    }
}
