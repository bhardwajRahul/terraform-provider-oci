// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DocumentLanguageClassificationFeature Detecting the language of the document.
type DocumentLanguageClassificationFeature struct {

	// The maximum number of results to return.
	MaxResults *int `mandatory:"false" json:"maxResults"`
}

func (m DocumentLanguageClassificationFeature) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DocumentLanguageClassificationFeature) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DocumentLanguageClassificationFeature) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDocumentLanguageClassificationFeature DocumentLanguageClassificationFeature
	s := struct {
		DiscriminatorParam string `json:"featureType"`
		MarshalTypeDocumentLanguageClassificationFeature
	}{
		"LANGUAGE_CLASSIFICATION",
		(MarshalTypeDocumentLanguageClassificationFeature)(m),
	}

	return json.Marshal(&s)
}
