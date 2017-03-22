# go-nfsiostat
A tools for reporting NFS mounstats to logging services

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/go-nfsiostat)](https://goreportcard.com/report/github.com/DataDrake/go-nfsiostat) [![license](https://img.shields.io/github/license/DataDrake/go-nfsiostat.svg)]() 

## Motivation

Lots of companies use NFS for accessing data, but it can be difficult to capture the statistics of a set of mount-points for logging and analysis. This tool seeks to easily log select statistics to a variety of logging services.

## Goals

 * Read from mounstats directly
 * Single file configuration
 * Support for logging endpoints:
    * syslog
    * zabbix_sender
 * A+ Rating on [Report Card](https://goreportcard.com/report/github.com/DataDrake/go-nfsiostat)
 
## License
 
Copyright 2017 Bryan T. Meyers <bmeyers@datadrake.com>
 
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 
http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
 
