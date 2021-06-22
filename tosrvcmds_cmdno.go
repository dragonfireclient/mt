// Code generated by cmdno.sh; DO NOT EDIT.

package mt

func (*ToSrvNil) toSrvCmdNo() uint16            { return 0 }
func (*ToSrvInit) toSrvCmdNo() uint16           { return 2 }
func (*ToSrvInit2) toSrvCmdNo() uint16          { return 17 }
func (*ToSrvModChanJoin) toSrvCmdNo() uint16    { return 23 }
func (*ToSrvModChanLeave) toSrvCmdNo() uint16   { return 24 }
func (*ToSrvModChanMsg) toSrvCmdNo() uint16     { return 25 }
func (*ToSrvPlayerPos) toSrvCmdNo() uint16      { return 35 }
func (*ToSrvGotBlks) toSrvCmdNo() uint16        { return 36 }
func (*ToSrvDeletedBlks) toSrvCmdNo() uint16    { return 37 }
func (*ToSrvInvAction) toSrvCmdNo() uint16      { return 49 }
func (*ToSrvChatMsg) toSrvCmdNo() uint16        { return 50 }
func (*ToSrvFallDmg) toSrvCmdNo() uint16        { return 53 }
func (*ToSrvSelectItem) toSrvCmdNo() uint16     { return 55 }
func (*ToSrvRespawn) toSrvCmdNo() uint16        { return 56 }
func (*ToSrvInteract) toSrvCmdNo() uint16       { return 57 }
func (*ToSrvRemovedSounds) toSrvCmdNo() uint16  { return 58 }
func (*ToSrvNodeMetaFields) toSrvCmdNo() uint16 { return 59 }
func (*ToSrvInvFields) toSrvCmdNo() uint16      { return 60 }
func (*ToSrvReqMedia) toSrvCmdNo() uint16       { return 64 }
func (*ToSrvCltReady) toSrvCmdNo() uint16       { return 67 }
func (*ToSrvFirstSRP) toSrvCmdNo() uint16       { return 80 }
func (*ToSrvSRPBytesA) toSrvCmdNo() uint16      { return 81 }
func (*ToSrvSRPBytesM) toSrvCmdNo() uint16      { return 82 }

var newToSrvCmd = map[uint16]func() Cmd{
	0:  func() Cmd { return new(ToSrvNil) },
	2:  func() Cmd { return new(ToSrvInit) },
	17: func() Cmd { return new(ToSrvInit2) },
	23: func() Cmd { return new(ToSrvModChanJoin) },
	24: func() Cmd { return new(ToSrvModChanLeave) },
	25: func() Cmd { return new(ToSrvModChanMsg) },
	35: func() Cmd { return new(ToSrvPlayerPos) },
	36: func() Cmd { return new(ToSrvGotBlks) },
	37: func() Cmd { return new(ToSrvDeletedBlks) },
	49: func() Cmd { return new(ToSrvInvAction) },
	50: func() Cmd { return new(ToSrvChatMsg) },
	53: func() Cmd { return new(ToSrvFallDmg) },
	55: func() Cmd { return new(ToSrvSelectItem) },
	56: func() Cmd { return new(ToSrvRespawn) },
	57: func() Cmd { return new(ToSrvInteract) },
	58: func() Cmd { return new(ToSrvRemovedSounds) },
	59: func() Cmd { return new(ToSrvNodeMetaFields) },
	60: func() Cmd { return new(ToSrvInvFields) },
	64: func() Cmd { return new(ToSrvReqMedia) },
	67: func() Cmd { return new(ToSrvCltReady) },
	80: func() Cmd { return new(ToSrvFirstSRP) },
	81: func() Cmd { return new(ToSrvSRPBytesA) },
	82: func() Cmd { return new(ToSrvSRPBytesM) },
}
