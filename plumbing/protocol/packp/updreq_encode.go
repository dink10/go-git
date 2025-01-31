package packp

import (
	"fmt"
	"io"

	"github.com/dink10/go-git.v4/plumbing"
	"github.com/dink10/go-git.v4/plumbing/format/pktline"
	"github.com/dink10/go-git.v4/plumbing/protocol/packp/capability"
)

var (
	zeroHashString = plumbing.ZeroHash.String()
)

// Encode writes the ReferenceUpdateRequest encoding to the stream.
func (r *ReferenceUpdateRequest) Encode(w io.Writer) error {
	if err := r.validate(); err != nil {
		return err
	}

	e := pktline.NewEncoder(w)

	if err := r.encodeShallow(e, r.Shallow); err != nil {
		return err
	}

	if err := r.encodeCommands(e, r.Commands, r.Capabilities); err != nil {
		return err
	}

	if r.Packfile != nil {
		if _, err := io.Copy(w, r.Packfile); err != nil {
			return err
		}

		return r.Packfile.Close()
	}

	return nil
}

func (r *ReferenceUpdateRequest) encodeShallow(e *pktline.Encoder,
	h *plumbing.Hash) error {

	if h == nil {
		return nil
	}

	objId := []byte(h.String())
	return e.Encodef("%s%s", shallow, objId)
}

func (r *ReferenceUpdateRequest) encodeCommands(e *pktline.Encoder,
	cmds []*Command, cap *capability.List) error {

	if err := e.Encodef("%s\x00%s",
		formatCommand(cmds[0]), cap.String()); err != nil {
		return err
	}

	for _, cmd := range cmds[1:] {
		if err := e.Encodef(formatCommand(cmd)); err != nil {
			return err
		}
	}

	return e.Flush()
}

func formatCommand(cmd *Command) string {
	o := cmd.Old.String()
	n := cmd.New.String()
	return fmt.Sprintf("%s %s %s", o, n, cmd.Name)
}
