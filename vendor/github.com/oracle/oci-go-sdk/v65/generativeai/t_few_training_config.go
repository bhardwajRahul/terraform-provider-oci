// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TFewTrainingConfig The TFEW training method hyperparameters.
type TFewTrainingConfig struct {

	// The maximum number of training epochs to run for.
	TotalTrainingEpochs *int `mandatory:"false" json:"totalTrainingEpochs"`

	// The initial learning rate to be used during training
	LearningRate *float64 `mandatory:"false" json:"learningRate"`

	// The batch size used during training.
	TrainingBatchSize *int `mandatory:"false" json:"trainingBatchSize"`

	// Stop training if the loss metric does not improve beyond 'early_stopping_threshold' for this many times of evaluation.
	EarlyStoppingPatience *int `mandatory:"false" json:"earlyStoppingPatience"`

	// How much the loss must improve to prevent early stopping.
	EarlyStoppingThreshold *float64 `mandatory:"false" json:"earlyStoppingThreshold"`

	// Determines how frequently to log model metrics.
	// Every step is logged for the first 20 steps and then follows this parameter for log frequency. Set to 0 to disable logging the model metrics.
	LogModelMetricsIntervalInSteps *int `mandatory:"false" json:"logModelMetricsIntervalInSteps"`
}

// GetTotalTrainingEpochs returns TotalTrainingEpochs
func (m TFewTrainingConfig) GetTotalTrainingEpochs() *int {
	return m.TotalTrainingEpochs
}

// GetLearningRate returns LearningRate
func (m TFewTrainingConfig) GetLearningRate() *float64 {
	return m.LearningRate
}

// GetTrainingBatchSize returns TrainingBatchSize
func (m TFewTrainingConfig) GetTrainingBatchSize() *int {
	return m.TrainingBatchSize
}

// GetEarlyStoppingPatience returns EarlyStoppingPatience
func (m TFewTrainingConfig) GetEarlyStoppingPatience() *int {
	return m.EarlyStoppingPatience
}

// GetEarlyStoppingThreshold returns EarlyStoppingThreshold
func (m TFewTrainingConfig) GetEarlyStoppingThreshold() *float64 {
	return m.EarlyStoppingThreshold
}

// GetLogModelMetricsIntervalInSteps returns LogModelMetricsIntervalInSteps
func (m TFewTrainingConfig) GetLogModelMetricsIntervalInSteps() *int {
	return m.LogModelMetricsIntervalInSteps
}

func (m TFewTrainingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TFewTrainingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TFewTrainingConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTFewTrainingConfig TFewTrainingConfig
	s := struct {
		DiscriminatorParam string `json:"trainingConfigType"`
		MarshalTypeTFewTrainingConfig
	}{
		"TFEW_TRAINING_CONFIG",
		(MarshalTypeTFewTrainingConfig)(m),
	}

	return json.Marshal(&s)
}
