// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v1

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/spdx"
)

func renderAnnotation2_1[T spdx.Annotation](ann T, w io.Writer) error {
	if ann.GetAnnotator() != "" && ann.GetAnnotatorType() != "" {
		fmt.Fprintf(w, "Annotator: %s: %s\n", ann.GetAnnotatorType(), ann.GetAnnotator())
	}
	if ann.GetAnnotationDate() != "" {
		fmt.Fprintf(w, "AnnotationDate: %s\n", ann.GetAnnotationDate())
	}
	if ann.GetAnnotationType() != "" {
		fmt.Fprintf(w, "AnnotationType: %s\n", ann.GetAnnotationType())
	}
	annIDStr := spdx.RenderDocElementID(ann.GetAnnotationSPDXIdentifier())
	if annIDStr != "SPDXRef-" {
		fmt.Fprintf(w, "SPDXREF: %s\n", annIDStr)
	}
	if ann.GetAnnotationComment() != "" {
		fmt.Fprintf(w, "AnnotationComment: %s\n", textify(ann.GetAnnotationComment()))
	}

	return nil
}
