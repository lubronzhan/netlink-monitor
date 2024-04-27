package syscall_lib

import "syscall"

var syscallTypeNameMap = map[uint16]string{
	syscall.RTM_DELACTION:    "RTM_DELACTION",
	syscall.RTM_DELADDR:      "RTM_DELADDR",
	syscall.RTM_DELADDRLABEL: "RTM_DELADDRLABEL",
	syscall.RTM_DELLINK:      "RTM_DELLINK",
	syscall.RTM_DELNEIGH:     "RTM_DELNEIGH",
	syscall.RTM_DELQDISC:     "RTM_DELQDISC",
	syscall.RTM_DELROUTE:     "RTM_DELROUTE",
	syscall.RTM_DELRULE:      "RTM_DELRULE",
	syscall.RTM_DELTCLASS:    "RTM_DELTCLASS",
	syscall.RTM_DELTFILTER:   "RTM_DELTFILTER",
	syscall.RTM_F_CLONED:     "RTM_F_CLONED",
	syscall.RTM_F_EQUALIZE:   "RTM_F_EQUALIZE",
	syscall.RTM_F_NOTIFY:     "RTM_F_NOTIFY",
	syscall.RTM_F_PREFIX:     "RTM_F_PREFIX",
	syscall.RTM_GETACTION:    "RTM_GETACTION",
	syscall.RTM_GETADDR:      "RTM_GETADDR",
	syscall.RTM_GETADDRLABEL: "RTM_GETADDRLABEL",
	syscall.RTM_GETANYCAST:   "RTM_GETANYCAST",
	syscall.RTM_GETDCB:       "RTM_GETDCB",
	syscall.RTM_GETLINK:      "RTM_GETLINK",
	syscall.RTM_GETMULTICAST: "RTM_GETMULTICAST",
	syscall.RTM_GETNEIGH:     "RTM_GETNEIGH",
	syscall.RTM_GETNEIGHTBL:  "RTM_GETNEIGHTBL",
	syscall.RTM_GETQDISC:     "RTM_GETQDISC",
	syscall.RTM_GETROUTE:     "RTM_GETROUTE",
	syscall.RTM_GETRULE:      "RTM_GETRULE",
	syscall.RTM_GETTCLASS:    "RTM_GETTCLASS",
	syscall.RTM_GETTFILTER:   "RTM_GETTFILTER",
	syscall.RTM_MAX:          "RTM_MAX",
	syscall.RTM_NEWACTION:    "RTM_NEWACTION",
	syscall.RTM_NEWADDR:      "RTM_NEWADDR",
	syscall.RTM_NEWADDRLABEL: "RTM_NEWADDRLABEL",
	syscall.RTM_NEWLINK:      "RTM_NEWLINK",
	syscall.RTM_NEWNDUSEROPT: "RTM_NEWNDUSEROPT",
	syscall.RTM_NEWNEIGH:     "RTM_NEWNEIGH",
	syscall.RTM_NEWNEIGHTBL:  "RTM_NEWNEIGHTBL",
	syscall.RTM_NEWPREFIX:    "RTM_NEWPREFIX",
	syscall.RTM_NEWQDISC:     "RTM_NEWQDISC",
	syscall.RTM_NEWROUTE:     "RTM_NEWROUTE",
	syscall.RTM_NEWRULE:      "RTM_NEWRULE",
	syscall.RTM_NEWTCLASS:    "RTM_NEWTCLASS",
	syscall.RTM_NEWTFILTER:   "RTM_NEWTFILTER",
	syscall.RTM_SETLINK:      "RTM_SETLINK",
	syscall.RTM_SETNEIGHTBL:  "RTM_SETNEIGHTBL",
}

func GetTypeName(val uint16) string {
	return syscallTypeNameMap[val]
}
