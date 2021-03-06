// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"testing"

	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	api "github.com/kubeflow/pipelines/backend/api/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/apiserver/resource"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

var testWorkflow = util.NewWorkflow(&v1alpha1.Workflow{
	TypeMeta:   v1.TypeMeta{APIVersion: "argoproj.io/v1alpha1", Kind: "Workflow"},
	ObjectMeta: v1.ObjectMeta{Name: "workflow-name", UID: "workflow1"},
	Spec:       v1alpha1.WorkflowSpec{Arguments: v1alpha1.Arguments{Parameters: []v1alpha1.Parameter{{Name: "param1"}}}},
})

var testWorkflow2 = util.NewWorkflow(&v1alpha1.Workflow{
	TypeMeta:   v1.TypeMeta{APIVersion: "argoproj.io/v1alpha1", Kind: "Workflow"},
	ObjectMeta: v1.ObjectMeta{Name: "workflow-name", UID: "workflow2"},
	Spec:       v1alpha1.WorkflowSpec{Arguments: v1alpha1.Arguments{Parameters: []v1alpha1.Parameter{{Name: "param1"}}}},
})

var validReference = []*api.ResourceReference{
	{
		Key: &api.ResourceKey{
			Type: api.ResourceType_EXPERIMENT, Id: resource.DefaultFakeUUID},
		Relationship: api.Relationship_OWNER,
	},
}

func initWithExperiment(t *testing.T) (*resource.FakeClientManager, *resource.ResourceManager, *model.Experiment) {
	clientManager := resource.NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	resourceManager := resource.NewResourceManager(clientManager)
	experiment := &model.Experiment{Name: "123"}
	experiment, err := resourceManager.CreateExperiment(experiment)
	assert.Nil(t, err)
	return clientManager, resourceManager, experiment
}

func initWithOneTimeRun(t *testing.T) (*resource.FakeClientManager, *resource.ResourceManager, *model.RunDetail) {
	clientManager, manager, exp := initWithExperiment(t)
	apiRun := &api.Run{
		Name: "run1",
		PipelineSpec: &api.PipelineSpec{
			WorkflowManifest: testWorkflow.ToStringForStore(),
			Parameters: []*api.Parameter{
				{Name: "param1", Value: "world"},
			},
		},
		ResourceReferences: []*api.ResourceReference{
			{
				Key:          &api.ResourceKey{Type: api.ResourceType_EXPERIMENT, Id: exp.UUID},
				Relationship: api.Relationship_OWNER,
			},
		},
	}
	runDetail, err := manager.CreateRun(apiRun)
	assert.Nil(t, err)
	return clientManager, manager, runDetail
}

// Util function to create an initial state with pipeline uploaded
func initWithPipeline(t *testing.T) (*resource.FakeClientManager, *resource.ResourceManager, *model.Pipeline) {
	store := resource.NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	manager := resource.NewResourceManager(store)
	p, err := manager.CreatePipeline("p1", "", []byte(testWorkflow.ToStringForStore()))
	assert.Nil(t, err)
	return store, manager, p
}

func AssertUserError(t *testing.T, err error, expectedCode codes.Code) {
	userError, ok := err.(*util.UserError)
	assert.True(t, ok)
	assert.Equal(t, expectedCode, userError.ExternalStatusCode())
}
