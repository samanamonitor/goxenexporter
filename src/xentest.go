package main

import (
	"flag"
	"os"

	"xenapi"
	"log"
	"encoding/json"
)

var HOST_FLAG = flag.String("host", "127.0.0.1", "the Host of the form ip[:port] pointing at the server")
var USERNAME_FLAG = flag.String("username", "", "the username of the host (e.g. root)")
var PASSWORD_FLAG = flag.String("password", "", "the password of the host")
var VERIFYSSL_FLAG = flag.Bool("verifyssl", true, "verify SSL certificate from the server")
/*
var CA_CERT_PATH_FLAG = flag.String("ca_cert_path", "", "the CA certificate file path for the host")
var NFS_SERVER_FLAG = flag.String("nfs_server", "", "the ip address pointing at the nfs server")
var NFS_PATH_FLAG = flag.String("nfs_path", "", "the nfs server path")
var IP1_FLAG = flag.String("ip1", "", "the URL of the form https://ip[:port] pointing at another host1")
var USERNAME1_FLAG = flag.String("username1", "", "the username of the host1 (e.g. root)")
var PASSWORD1_FLAG = flag.String("password1", "", "the password of the host1")
*/


type SammXen struct {
	verifySsl bool
	Session *xenapi.Session
	sessionRef xenapi.SessionRef
	SessionRec xenapi.SessionRecord
}

func NewSammXen(host string, user string, password string, verifySsl bool) (*SammXen, error) {
	x := &SammXen{
		verifySsl: verifySsl,
		Session: xenapi.NewSession(&xenapi.ClientOpts{
			URL: "https://" + host,
			Headers: map[string]string{
				"User-Agent": "SAMM exporter v2.0",
			},
		}),
	}
	var err error
	x.sessionRef, err = x.Session.LoginWithPassword(user, password, "1.0", "Samm exporter v2.0")
	if err != nil {
		return nil, err
	}
	x.SessionRec, err = x.Session.GetRecord(x.sessionRef)
	if err != nil {
		return nil, err
	}
	return x, nil
}

func (self SammXen) SessionId() (string) {
	return self.SessionRec.UUID
}

func (self SammXen) GetUpdatesRrd(hostUuid string) ([]byte, error) {
	//var host xenapi.HostRef
	//host, err = xenapi.Host.GetByUUID(self.sessionRef, hostUuid)
	return json.Marshal([]string{})
}

var session *xenapi.Session

func main() {
	flag.Parse()
	x, err := NewSammXen(*HOST_FLAG, *USERNAME_FLAG, *PASSWORD_FLAG, false)
	if err != nil {
		panic(err)
		return
	}
	/*
	session = xenapi.NewSession(&xenapi.ClientOpts{
		URL: "http://" + *HOST_FLAG,
		Headers: map[string]string{
			"User-Agent": "XS SDK for Go - Examples v1.0",
		},
	})
	session_id , err := session.LoginWithPassword(*USERNAME_FLAG, *PASSWORD_FLAG, "1.0", "Go sdk samples")
	if err != nil {
		panic(err)
		return
	}
	log.Print("api version: ", session.APIVersion)
	log.Print("xapi rpm version: ", session.XAPIVersion)
	rec, _ := session.GetRecord(session_id)
	*/
	log.Printf("%s\n", x.SessionId())
	allHosts, nerr := xenapi.Host.GetAll(x.Session)
	if nerr != nil {
		panic(nerr)
		return
	}
	for _, hostRef := range(allHosts) {
		host, err := xenapi.Host.GetRecord(x.Session, hostRef)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("%s\n", host.UUID)
		}
	}

	if err := x.Session.Logout(); err != nil {
		log.Print(err)
	}
	os.Exit(0)
}
