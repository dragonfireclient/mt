package mt

import (
	"fmt"
	"io"
	"net"

	"github.com/dragonfireclient/mt/rudp"
)

// A Pkt is a deserialized rudp.Pkt.
type Pkt struct {
	Cmd
	rudp.PktInfo
}

// Peer wraps rudp.Conn, adding (de)serialization.
type Peer struct {
	*rudp.Conn
}

func SerializePkt(pkt Cmd, w io.WriteCloser, toSrv bool) bool {
	var cmdNo uint16
	if toSrv {
		cmdNo = pkt.(ToSrvCmd).toSrvCmdNo()
	} else {
		cmdNo = pkt.(ToCltCmd).toCltCmdNo()
	}

	if cmdNo == 0xffff {
		return false
	}

	go func() (err error) {
		// defer w.CloseWithError(err)
		defer w.Close()

		buf := make([]byte, 2)
		be.PutUint16(buf, cmdNo)
		if _, err := w.Write(buf); err != nil {
			return err
		}
		return serialize(w, pkt)
	}()

	return true
}

func (p Peer) Send(pkt Pkt) (ack <-chan struct{}, err error) {
	r, w := io.Pipe()
	if !SerializePkt(pkt.Cmd, w, p.IsSrv()) {
		return nil, p.Close()
	}

	return p.Conn.Send(rudp.Pkt{r, pkt.PktInfo})
}

// SendCmd is equivalent to Send(Pkt{cmd, cmd.DefaultPktInfo()}).
func (p Peer) SendCmd(cmd Cmd) (ack <-chan struct{}, err error) {
	return p.Send(Pkt{cmd, cmd.DefaultPktInfo()})
}

func DeserializePkt(pkt io.Reader, fromSrv bool) (*Cmd, error) {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(pkt, buf); err != nil {
		return nil, err
	}
	cmdNo := be.Uint16(buf)

	var newCmd func() Cmd
	if fromSrv {
		newCmd = newToCltCmd[cmdNo]
	} else {
		newCmd = newToSrvCmd[cmdNo]
	}

	if newCmd == nil {
		return nil, fmt.Errorf("unknown cmd: %d", cmdNo)
	}
	cmd := newCmd()

	if err := deserialize(pkt, cmd); err != nil {
		return nil, fmt.Errorf("%T: %w", cmd, err)
	}

	extra, err := io.ReadAll(pkt)
	if len(extra) > 0 {
		err = fmt.Errorf("%T: %w", cmd, rudp.TrailingDataError(extra))
	}

	return &cmd, err
}

func (p Peer) Recv() (_ Pkt, rerr error) {
	pkt, err := p.Conn.Recv()
	if err != nil {
		return Pkt{}, err
	}

	cmd, err := DeserializePkt(pkt, p.IsSrv())

	if cmd == nil {
		return Pkt{}, err
	} else {
		return Pkt{*cmd, pkt.PktInfo}, err
	}
}

func Connect(conn net.Conn) Peer {
	return Peer{rudp.Connect(conn)}
}

type Listener struct {
	*rudp.Listener
}

func Listen(conn net.PacketConn) Listener {
	return Listener{rudp.Listen(conn)}
}

func (l Listener) Accept() (Peer, error) {
	rpeer, err := l.Listener.Accept()
	return Peer{rpeer}, err
}
