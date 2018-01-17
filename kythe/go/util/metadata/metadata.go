/*
 * Copyright 2017 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package metadata provides support code for processing Kythe metadata
// records, of the type generated by instrumented code generators for
// cross-language linkage.
package metadata

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"kythe.io/kythe/go/util/schema/edges"

	protopb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	spb "kythe.io/kythe/proto/storage_go_proto"
)

// TODO(fromberger): Add a link to the format documentation here.

// Rules are a collection of metadata rules.
type Rules []Rule

// MarshalJSON encodes the specified rule set as a JSON file.
func (rs Rules) MarshalJSON() ([]byte, error) {
	f := file{
		Type: fileType,
		Meta: make([]rule, len(rs)),
	}
	for i, r := range rs {
		kind := r.EdgeOut
		if r.Reverse {
			kind = edges.Mirror(kind)
		}
		rtype := "nop"
		if r.EdgeIn == edges.DefinesBinding {
			rtype = "anchor_defines"
		}
		f.Meta[i] = rule{
			Type:  rtype,
			Begin: r.Begin,
			End:   r.End,
			VName: r.VName,
			Edge:  kind,
		}
	}
	return json.Marshal(f)
}

// A Rule denotes a single metadata rule, associating type linkage information
// for an anchor spanning a given range of text.
type Rule struct {
	// The Begin and End fields represent a half-closed interval of byte
	// positions to match. Begin is inclusive, End is exclusive.
	Begin, End int

	EdgeIn  string     // edge kind to match over the anchor spanned
	EdgeOut string     // outbound edge kind to emit
	VName   *spb.VName // the vname to create an edge to or from
	Reverse bool       // whether to draw to vname (false) or from it (true)
}

// The types below are intermediate structures used for JSON marshaling.

const fileType = "kythe0" // protocol marker

// A file represents an encoded set of rules in JSON notation.
type file struct {
	Type string `json:"type"` // required: must equal fileType
	Meta []rule `json:"meta,omitempty"`
}

// A rule is the encoded format of a single rule.
type rule struct {
	Type  string     `json:"type"`
	Begin int        `json:"begin"`
	End   int        `json:"end"`
	Edge  string     `json:"edge,omitempty"`
	VName *spb.VName `json:"vname,omitempty"`
}

// Parse parses a single JSON metadata object from r and returns the
// corresponding rules. It is an error if there are extra data after the
// metadata object, or if the type tag of the object does not match the current
// format code.
func Parse(r io.Reader) (Rules, error) {
	dec := json.NewDecoder(r)
	var f file
	if err := dec.Decode(&f); err != nil {
		return nil, fmt.Errorf("metadata: invalid file: %v", err)
	} else if _, err := dec.Token(); err != io.EOF {
		return nil, errors.New("metadata: extra junk at end of input")
	} else if f.Type != fileType {
		return nil, fmt.Errorf("metadata: wrong type tag: %q", f.Type)
	}

	rs := make(Rules, len(f.Meta))
	for i, meta := range f.Meta {
		rs[i] = Rule{
			Begin:   meta.Begin,
			End:     meta.End,
			EdgeOut: edges.Canonical(meta.Edge),
			Reverse: edges.IsReverse(meta.Edge),
			VName:   meta.VName,
		}
		switch t := meta.Type; t {
		case "nop":
			// ok, no special behaviour
		case "anchor_defines":
			rs[i].EdgeIn = edges.DefinesBinding
		default:
			return nil, fmt.Errorf("metadata: unknown rule type: %q", t)
		}
	}
	return rs, nil
}

// FromGeneratedCodeInfo constructs a set of rules from the corresponding
// protobuf descriptor message and the vname of the metadata file from which
// the generated descriptor was loaded.
func FromGeneratedCodeInfo(msg *protopb.GeneratedCodeInfo, vname *spb.VName) Rules {
	rs := make(Rules, len(msg.Annotation))
	for i, anno := range msg.Annotation {
		// Convert the path to a dot-separated string, e.g., 1.0.3.2,
		// for use in the vname signature.
		sig := make([]string, len(anno.Path))
		for i, elt := range anno.Path {
			sig[i] = strconv.Itoa(int(elt))
		}

		// TODO(fromberger): Work out how to derive the correct corpus and root
		// labels. When the protobuf source file is in the same corpus as its
		// metadata, this will work as-is.
		//
		// If the protobuf inputs live in a different corpus, it will be
		// necessary to make the extractor map the metadata file to the correct
		// corpus and root at build time. Since the metadata file does not get
		// pointed to directly, this ensures we get the right contact.
		//
		// This does NOT solve how to deal with generated .proto files, but
		// that is a much less common case, and we can address it separately.
		vname := &spb.VName{
			Corpus:    vname.GetCorpus(),
			Root:      vname.GetRoot(),
			Path:      anno.GetSourceFile(),
			Language:  "protobuf",
			Signature: strings.Join(sig, "."),
		}
		rs[i] = Rule{
			EdgeIn:  edges.DefinesBinding,
			EdgeOut: edges.Generates,
			Reverse: true,
			Begin:   int(anno.GetBegin()),
			End:     int(anno.GetEnd()),
			VName:   vname,
		}
	}
	return rs
}
