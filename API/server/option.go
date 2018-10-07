/*
 * File: option.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 5:46:47 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

type Option struct {
	exe      string
	env      string
	debug    bool
	logpath  string
	datapath string
	address  string
	port     string
	ip       string
	index    bool
	reindex  bool
}

var (
	PROD = "production"
	DEV  = "développement"
	TEST = "test"
)

// Hydrate set datas for Option struc
func (o *Option) Hydrate(port, ip, env, dir, logpath, datapath string, debug bool, index bool, reindex bool) {
	o.port = port
	o.ip = ip
	o.datapath = datapath
	o.logpath = logpath
	o.env = env
	o.debug = debug
	o.exe = dir
	o.index = index
	o.reindex = reindex
}

// GetAddress concat ip and port and affect to address if needed
// else default address is define to 127.0.0.1:80
func (o *Option) GetAddress() string {
	return o.ip + ":" + o.port
}

/*
context.serverOption interface:

type IServerOption interface {
		GetExeFolder() string
		SetEnv(string)
		SetLogpath(string)
		SetDatapath(string)
		GetEnv() string
		GetLogpath() string
		GetDatapath() string
		GetAddress() string
		GetPort() string
		GetIp() string
	}
*/

// GetIP return ip
func (o *Option) GetIP() string {
	return o.ip
}

// GetPort return port
func (o *Option) GetPort() string {
	return o.port
}

// GetExeFolder return ./app.exe folder
func (o *Option) GetExeFolder() string {
	return o.exe
}

// GetLogpath return the path to logs folder
func (o *Option) GetLogpath() string {
	return o.logpath
}

// GetDatapath return the path to data folder
func (o *Option) GetDatapath() string {
	return o.datapath
}

// GetEnv return environnement description default on DEV: "développement"
func (o *Option) GetEnv() string {
	if o.env == "" {
		return DEV
	}
	return o.env
}

// GetIndex return if indexation is needed default false
func (o *Option) GetIndex() bool {
	return o.index
}

// GetReIndex return if reindexation is needed default false
func (o *Option) GetReindex() bool {
	return o.reindex
}

// SetEnv allow a IServerOption to change the environnement
// "prod" => "production"
// "test" => "test"
// (default) "dev" => développement
func (o *Option) SetEnv(env string) {
	if env == "prod" {
		o.env = PROD
	} else if env == "test" {
		o.env = TEST
	} else {
		o.env = DEV
	}
}

// SetLogpath allow IServerOption to change the path to logs
func (o *Option) SetLogpath(logpath string) {
	o.logpath = logpath
}

// SetDatapath allow IServerOption to change the path to data
func (o *Option) SetDatapath(datapath string) {
	o.datapath = datapath
}
