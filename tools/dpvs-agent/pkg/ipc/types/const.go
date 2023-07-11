// Copyright 2023 IQiYi Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"strings"
)

type lcoreid_t uint8
type SockoptType uint32
type DpvsErrType int32
type DpvsFwdMode uint32

const (
	DPVS_FWD_MASQ DpvsFwdMode = iota
	DPVS_FWD_LOCALNODE
	DPVS_FWD_MODE_TUNNEL
	DPVS_FWD_MODE_DR
	DPVS_FWD_MODE_BYPASS
	DPVS_FWD_MODE_FNAT
	DPVS_FWD_MODE_SNAT
)
const DPVS_FWD_MODE_NAT DpvsFwdMode = DPVS_FWD_MASQ

const (
	DPVS_SVC_F_PERSISTENT = 0x1 << iota
	DPVS_SVC_F_HASHED
	DPVS_SVC_F_ONEPACKET
	DPVS_SVC_F_SCHED1
	DPVS_SVC_F_SCHED2
	DPVS_SVC_F_SCHED3
)

const (
	DPVS_SVC_F_SIP_HASH uint32 = 0x100 << iota
	DPVS_SVC_F_QID_HASH
	DPVS_SVC_F_MATCH
)

const (
	DPVS_DEST_HC_PASSIVE = 0x1 << iota
	DPVS_DEST_HC_TCP
	DPVS_DEST_HC_UDP
	DPVS_DEST_HC_PING
)

const (
	DPVS_DEST_F_AVAILABLE = 0x1 << iota
	DPVS_DEST_F_OVERLOAD
	DPVS_DEST_F_INHIBITED
)

const (
	DPVS_CONN_F_MASQ = iota
	DPVS_CONN_F_LOCALNODE
	DPVS_CONN_F_TUNNEL
	DPVS_CONN_F_DROUTE
	DPVS_CONN_F_BYPASS
	DPVS_CONN_F_FULLNAT
	DPVS_CONN_F_SNAT
	DPVS_CONN_F_FWD_MASK
)

const (
	DPVS_CONN_F_SYNPROXY = 0x10 << iota
	DPVS_CONN_F_EXPIRE_QUIESCENT
	_
	_
	DPVS_CONN_F_HASHED
	DPVS_CONN_F_INACTIVE
	DPVS_CONN_F_TEMPLATE
	DPVS_CONN_F_ONE_PACKET
	DPVS_CONN_F_IN_TIMER
	DPVS_CONN_F_REDIRECT_HASHED
	DPVS_CONN_F_NOFASTXMIT
)

func (e *DpvsFwdMode) String() string {
	switch *e {
	case DPVS_FWD_MASQ: // DPVS_FWD_MASQ == DPVS_FWD_MODE_NAT
		return "MASQ"
	case DPVS_FWD_LOCALNODE:
		return "LOCALNODE"
	case DPVS_FWD_MODE_DR:
		return "DR"
	case DPVS_FWD_MODE_TUNNEL:
		return "TUNNLE"
	case DPVS_FWD_MODE_BYPASS:
		return "BYPASS"
	case DPVS_FWD_MODE_SNAT:
		return "SNAT"
	case DPVS_FWD_MODE_FNAT:
		return "FNAT"
	}
	return "UNKNOW"
}

func (e *DpvsFwdMode) FromString(name string) {
	switch strings.ToUpper(name) {
	case "MASQ":
		*e = DPVS_FWD_MASQ
	case "LOCALNODE":
		*e = DPVS_FWD_LOCALNODE
	case "DR":
		*e = DPVS_FWD_MODE_DR
	case "TUNNLE":
		*e = DPVS_FWD_MODE_TUNNEL
	case "BYPASS":
		*e = DPVS_FWD_MODE_BYPASS
	case "SNAT":
		*e = DPVS_FWD_MODE_SNAT
	case "NAT":
		*e = DPVS_FWD_MODE_NAT
	case "FNAT":
		*e = DPVS_FWD_MODE_FNAT
	default:
		*e = DPVS_FWD_MODE_FNAT
	}
}

const (
	EDPVS_INPROGRESS DpvsErrType = 2 - iota
	EDPVS_KNICONTINUE
	EDPVS_OK
	EDPVS_INVAL
	EDPVS_NOMEM
	EDPVS_EXIST
	EDPVS_NOTEXIST
	EDPVS_INVPKT
	EDPVS_DROP
	EDPVS_NOPROT
	EDPVS_NOROUTE
	EDPVS_DEFRAG
	EDPVS_FRAG
	EDPVS_DPDKAPIFAIL
	EDPVS_IDLE
	EDPVS_BUSY
	EDPVS_NOTSUPP
	EDPVS_RESOURCE
	EDPVS_OVERLOAD
	EDPVS_NOSERV
	EDPVS_DISABLED
	EDPVS_NOROOM
	EDPVS_NONEALCORE
	EDPVS_CALLBACKFAIL
	EDPVS_IO
	EDPVS_MSG_FAIL
	EDPVS_MSG_DROP
	EDPVS_PKTSTOLEN
	EDPVS_SYSCALL
	EDPVS_NODEV
)

var dpvsErrNames = [...]string{"EDPVS_INPROGRESS", "EDPVS_KNICONTINUE", "EDPVS_OK", "EDPVS_INVAL", "EDPVS_NOMEM", "EDPVS_EXIST", "EDPVS_NOTEXIST", "EDPVS_INVPKT", "EDPVS_DROP", "EDPVS_NOPROT", "EDPVS_NOROUTE", "EDPVS_DEFRAG", "EDPVS_FRAG", "EDPVS_DPDKAPIFAIL", "EDPVS_IDLE", "EDPVS_BUSY", "EDPVS_NOTSUPP", "EDPVS_RESOURCE", "EDPVS_OVERLOAD", "EDPVS_NOSERV", "EDPVS_DISABLED", "EDPVS_NOROOM", "EDPVS_NONEALCORE", "EDPVS_CALLBACKFAIL", "EDPVS_IO", "EDPVS_MSG_FAIL", "EDPVS_MSG_DROP", "EDPVS_PKTSTOLEN", "EDPVS_SYSCALL", "EDPVS_NODEV"}

func (e *DpvsErrType) String() string {
	return dpvsErrNames[EDPVS_INPROGRESS-*e]
}

const (
	SOCKOPT_VERSION = 0x10000
)

const (
	SOCKOPT_GET SockoptType = iota
	SOCKOPT_SET
	SOCKOPT_TYPE_MAX
)

const (
	SOCKOPT_SET_LADDR_ADD = iota
	SOCKOPT_SET_LADDR_DEL
	SOCKOPT_SET_LADDR_FLUSH
	SOCKOPT_GET_LADDR_GETALL
	SOCKOPT_GET_LADDR_MAX

	DPVSAGENT_VS_GET_LADDR
	DPVSAGENT_VS_ADD_LADDR
	DPVSAGENT_VS_DEL_LADDR

	DPVS_SO_SET_FLUSH
	DPVS_SO_SET_ZERO
	DPVS_SO_SET_ADD
	DPVS_SO_SET_EDIT
	DPVS_SO_SET_DEL
	DPVS_SO_SET_ADDDEST
	DPVS_SO_SET_EDITDEST
	DPVS_SO_SET_DELDEST
	DPVS_SO_SET_GRATARP
	DPVS_SO_GET_VERSION
	DPVS_SO_GET_INFO
	DPVS_SO_GET_SERVICES
	DPVS_SO_GET_SERVICE
	DPVS_SO_GET_DESTS
	DPVSAGENT_SO_GET_SERVICES
	SOCKOPT_SVC_MAX

	DPVSAGENT_VS_GET_DESTS
	DPVSAGENT_VS_ADD_DESTS
	DPVSAGENT_VS_EDIT_DESTS
	DPVSAGENT_VS_DEL_DESTS

	SOCKOPT_SET_ROUTE_ADD
	SOCKOPT_SET_ROUTE_DEL
	SOCKOPT_SET_ROUTE_SET
	SOCKOPT_SET_ROUTE_FLUSH
	SOCKOPT_GET_ROUTE_SHOW

	SOCKOPT_SET_ROUTE6_ADD_DEL
	SOCKOPT_SET_ROUTE6_FLUSH
	SOCKOPT_GET_ROUTE6_SHOW

	DPVSAGENT_ROUTE_GET
	DPVSAGENT_ROUTE_ADD
	DPVSAGENT_ROUTE_DEL
	DPVSAGENT_ROUTE6_GET
	DPVSAGENT_ROUTE6_ADD
	DPVSAGENT_ROUTE6_DEL

	DPVSAGENT_IFADDR_GET_BASE
	DPVSAGENT_IFADDR_GET_STATS
	DPVSAGENT_IFADDR_GET_VERBOSE
	DPVSAGENT_IFADDR_ADD
	DPVSAGENT_IFADDR_DEL

	SOCKOPT_SET_IFADDR_ADD
	SOCKOPT_SET_IFADDR_DEL
	SOCKOPT_SET_IFADDR_SET
	SOCKOPT_SET_IFADDR_FLUSH
	SOCKOPT_GET_IFADDR_SHOW

	SOCKOPT_NETIF_SET_LCORE
	SOCKOPT_NETIF_SET_PORT
	SOCKOPT_NETIF_SET_BOND
	SOCKOPT_NETIF_SET_MAX
	SOCKOPT_NETIF_GET_LCORE_MASK
	SOCKOPT_NETIF_GET_LCORE_BASIC
	SOCKOPT_NETIF_GET_LCORE_STATS
	SOCKOPT_NETIF_GET_PORT_LIST
	SOCKOPT_NETIF_GET_PORT_BASIC
	SOCKOPT_NETIF_GET_PORT_STATS
	SOCKOPT_NETIF_GET_PORT_XSTATS
	SOCKOPT_NETIF_GET_PORT_EXT_INFO
	SOCKOPT_NETIF_GET_BOND_STATUS
	SOCKOPT_NETIF_GET_MAX

	SOCKOPT_SET_NEIGH_ADD
	SOCKOPT_SET_NEIGH_DEL
	SOCKOPT_GET_NEIGH_SHOW

	SOCKOPT_SET_BLKLST_ADD
	SOCKOPT_SET_BLKLST_DEL
	SOCKOPT_SET_BLKLST_FLUSH
	SOCKOPT_GET_BLKLST_GETALL

	SOCKOPT_SET_WHTLST_ADD
	SOCKOPT_SET_WHTLST_DEL
	SOCKOPT_SET_WHTLST_FLUSH
	SOCKOPT_GET_WHTLST_GETALL

	SOCKOPT_SET_VLAN_ADD
	SOCKOPT_SET_VLAN_DEL
	SOCKOPT_GET_VLAN_SHOW

	SOCKOPT_TC_ADD
	SOCKOPT_TC_DEL
	SOCKOPT_TC_CHANGE
	SOCKOPT_TC_REPLACE
	SOCKOPT_TC_SHOW

	SOCKOPT_SET_CONN
	SOCKOPT_GET_CONN_ALL
	SOCKOPT_GET_CONN_SPECIFIED

	SOCKOPT_IP6_SET
	SOCKOPT_IP6_STATS

	SOCKOPT_TUNNEL_ADD
	SOCKOPT_TUNNEL_DEL
	SOCKOPT_TUNNEL_CHANGE
	SOCKOPT_TUNNEL_REPLACE
	SOCKOPT_TUNNEL_SHOW

	SOCKOPT_SET_KNI_ADD
	SOCKOPT_SET_KNI_DEL
	SOCKOPT_SET_KNI_FLUSH
	SOCKOPT_GET_KNI_LIST

	SOCKOPT_SET_IPSET
	SOCKOPT_GET_IPSET_TEST
	SOCKOPT_GET_IPSET_LIST

	SOCKOPT_SET_IFTRAF_ADD
	SOCKOPT_SET_IFTRAF_DEL
	SOCKOPT_GET_IFTRAF_SHOW
)

const (
	NETIF_NIC_PROMISC_ON = 1 << iota
	NETIF_NIC_PROMISC_OFF
	NETIF_NIC_LINK_UP
	NETIF_NIC_LINK_DOWN
	NETIF_NIC_FWD2KNI_ON
	NETIF_NIC_FWD2KNI_OFF
	NETIF_NIC_TC_EGRESS_ON
	NETIF_NIC_TC_EGRESS_OFF
	NETIF_NIC_TC_INGRESS_ON
	NETIF_NIC_TC_INGRESS_OFF
)

const (
	RTF_UP = 1 << iota
	RTF_GATEWAY
	RTF_HOST
	RTF_REINSTATE
	RTF_DYNAMIC
	RTF_MODIFIED
	RTF_MTU
	RTF_WINDOW
	RTF_IRTT
	RTF_REJECT

	RTF_FORWARD
	RTF_LOCALIN
	RTF_DEFAULT
	RTF_KNI
	RTF_OUTWALL
)

const (
	IFA_SCOPE_GLOBAL uint8 = iota
	IFA_SCOPE_SITE
	IFA_SCOPE_LINK
	IFA_SCOPE_HOST
	IFA_SCOPE_NONE = 255
)
