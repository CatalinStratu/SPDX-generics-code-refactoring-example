// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/spdx"
)

func renderRelationship2_2(rln *spdx.Relationship2_2, w io.Writer) error {
	rlnAStr := spdx.RenderDocElementID(rln.RefA)
	rlnBStr := spdx.RenderDocElementID(rln.RefB)
	if rlnAStr != "SPDXRef-" && rlnBStr != "SPDXRef-" && rln.Relationship != "" {
		fmt.Fprintf(w, "Relationship: %s %s %s\n", rlnAStr, rln.Relationship, rlnBStr)
	}
	if rln.RelationshipComment != "" {
		fmt.Fprintf(w, "RelationshipComment: %s\n", textify(rln.RelationshipComment))
	}

	return nil
}
