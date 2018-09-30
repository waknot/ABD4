/*
 * File: context.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 5:43:00 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/database/manager"
	"ABD4/API/iserial"
	"ABD4/API/logger"
	"ABD4/API/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	appName = "abd4"
	SECRET  = "==+VDMEG@ABD4"
)

// IServerOption :
// SetEnv setter for environnement ("prod", "dev", "test")
// GetEnv getter for environnement
// GetExeFolder getter for .exe folder
// GetLogpath getter for log folder
// GetDataPath getter for bolt data folder
// GetAddress return full host name (IP:Port)
// GetPort return port
// GetIp return ip
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
	GetIP() string
}

// IResponseWriter interface define the required methods to
// use the AppContext.Rw variable into the API
type IResponseWriter interface {
	Send(*AppContext, http.ResponseWriter, int, iserial.Serializable, string, string)
	SendError(*AppContext, http.ResponseWriter, int, string, string)
	SendItSelf(*AppContext, http.ResponseWriter)
	NewResponse(int, string, string, string) IResponseWriter
}

// ISessionUser abstract the user from model
type ISessionUser interface {
	GetID() string
}

// AppContext define globals tools and variable usefull in the API
// It embed the dao's objects (XxxManager *manager.XxxManager),
// a ResponseWriter which offer shorthand to send uniformised Response
type AppContext struct {
	Rw          IResponseWriter
	SessionUser ISessionUser
	Opts        IServerOption
	UserManager *manager.UserManager
	Log         *logger.AppLogger
	Exe         string
	Logpath     string
	DataPath    string
}

// Instanciate the global ctx variable
// AppContext know: manager, iserial, utils and logger packages
// AppContext don't need to know: server, model and handler packages
// AppContext rely on three interfaces implementing usefull function for API
// ISerial to serializable model object (model abstraction)
// ISessionUser to typical api user (model abstraction)
// IResponseWriter working with ISerial, it should provide basic shorthand
// to write responses logging and managing errors in handling process
// The context allow us to embed usefull data as datapath or ip and port (Opts)
// in handlers implementing CustomHandler and HandlerWrapper.
// It also centralise common functionnalities as database access, logging process
// and response/error formatting
// the app must define how routing will work and must implement some interfaces:
// - IResponseWriter: Response process where serialisation apply
// - ISerial: On each model entity we want to return
// - ISessionUser: On the User model entity
// - IServerOption: On the structure embedding configuration from flags and harcoded values
func (ctx *AppContext) Instanciate(opts IServerOption) *AppContext {
	if opts.GetExeFolder() == "" {
		log.Fatalf("No exe folder defined, %s unable to provide defaults values", appName)
	}
	// check opts content and define default values if required
	if opts.GetLogpath() == "" {
		opts.SetLogpath(filepath.Join(opts.GetExeFolder(), "log/"))
	}
	if opts.GetDatapath() == "" {
		opts.SetDatapath(filepath.Join(opts.GetExeFolder(), "data/"))
	}
	// create path if don't exist
	err := os.MkdirAll(opts.GetLogpath(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(opts.GetDatapath(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// open a file for log, for now, just one file is defined
	// but we can move it easy to three log files (debug, info, error)
	// for now, debug mode is useless
	logFile, err := os.OpenFile(
		filepath.Join(opts.GetLogpath(), appName+".log"),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// we define an output able to log on file and standart output
	output := io.MultiWriter(logFile, os.Stdout)
	loggers := logger.Instanciate(output, output, output)
	// define dao access (database/manager package)
	um, err := manager.NewUserManager("users.dat", opts.GetDatapath(), SECRET)
	if err != nil {
		loggers.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}
	// instanciate the ctx to return
	ctx.Opts = opts
	ctx.UserManager = um
	//SessionUser: &model.User{},
	ctx.Log = loggers
	ctx.Exe = opts.GetExeFolder()
	ctx.Logpath = opts.GetLogpath()
	ctx.DataPath = opts.GetDatapath()
	ctx.Log.Info.Printf("%s RootDir: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Exe)
	ctx.Log.Info.Printf("%s LogPath: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Logpath)
	ctx.Log.Info.Printf("%s Datapath: %s", utils.Use().GetStack(ctx.Instanciate), ctx.DataPath)
	ctx.Log.Info.Printf("%s IP: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetIP())
	ctx.Log.Info.Printf("%s Port: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetPort())
	ctx.Log.Info.Printf("%s Address: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetAddress())
	return ctx
}
