// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appregistryiface provides an interface to enable mocking the AWS Service Catalog App Registry service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appregistryiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appregistry"
)

// AppRegistryAPI provides an interface to enable mocking the
// appregistry.AppRegistry service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Service Catalog App Registry.
//    func myFunc(svc appregistryiface.AppRegistryAPI) bool {
//        // Make svc.AssociateAttributeGroup request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := appregistry.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAppRegistryClient struct {
//        appregistryiface.AppRegistryAPI
//    }
//    func (m *mockAppRegistryClient) AssociateAttributeGroup(input *appregistry.AssociateAttributeGroupInput) (*appregistry.AssociateAttributeGroupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAppRegistryClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AppRegistryAPI interface {
	AssociateAttributeGroup(*appregistry.AssociateAttributeGroupInput) (*appregistry.AssociateAttributeGroupOutput, error)
	AssociateAttributeGroupWithContext(aws.Context, *appregistry.AssociateAttributeGroupInput, ...request.Option) (*appregistry.AssociateAttributeGroupOutput, error)
	AssociateAttributeGroupRequest(*appregistry.AssociateAttributeGroupInput) (*request.Request, *appregistry.AssociateAttributeGroupOutput)

	AssociateResource(*appregistry.AssociateResourceInput) (*appregistry.AssociateResourceOutput, error)
	AssociateResourceWithContext(aws.Context, *appregistry.AssociateResourceInput, ...request.Option) (*appregistry.AssociateResourceOutput, error)
	AssociateResourceRequest(*appregistry.AssociateResourceInput) (*request.Request, *appregistry.AssociateResourceOutput)

	CreateApplication(*appregistry.CreateApplicationInput) (*appregistry.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *appregistry.CreateApplicationInput, ...request.Option) (*appregistry.CreateApplicationOutput, error)
	CreateApplicationRequest(*appregistry.CreateApplicationInput) (*request.Request, *appregistry.CreateApplicationOutput)

	CreateAttributeGroup(*appregistry.CreateAttributeGroupInput) (*appregistry.CreateAttributeGroupOutput, error)
	CreateAttributeGroupWithContext(aws.Context, *appregistry.CreateAttributeGroupInput, ...request.Option) (*appregistry.CreateAttributeGroupOutput, error)
	CreateAttributeGroupRequest(*appregistry.CreateAttributeGroupInput) (*request.Request, *appregistry.CreateAttributeGroupOutput)

	DeleteApplication(*appregistry.DeleteApplicationInput) (*appregistry.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *appregistry.DeleteApplicationInput, ...request.Option) (*appregistry.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*appregistry.DeleteApplicationInput) (*request.Request, *appregistry.DeleteApplicationOutput)

	DeleteAttributeGroup(*appregistry.DeleteAttributeGroupInput) (*appregistry.DeleteAttributeGroupOutput, error)
	DeleteAttributeGroupWithContext(aws.Context, *appregistry.DeleteAttributeGroupInput, ...request.Option) (*appregistry.DeleteAttributeGroupOutput, error)
	DeleteAttributeGroupRequest(*appregistry.DeleteAttributeGroupInput) (*request.Request, *appregistry.DeleteAttributeGroupOutput)

	DisassociateAttributeGroup(*appregistry.DisassociateAttributeGroupInput) (*appregistry.DisassociateAttributeGroupOutput, error)
	DisassociateAttributeGroupWithContext(aws.Context, *appregistry.DisassociateAttributeGroupInput, ...request.Option) (*appregistry.DisassociateAttributeGroupOutput, error)
	DisassociateAttributeGroupRequest(*appregistry.DisassociateAttributeGroupInput) (*request.Request, *appregistry.DisassociateAttributeGroupOutput)

	DisassociateResource(*appregistry.DisassociateResourceInput) (*appregistry.DisassociateResourceOutput, error)
	DisassociateResourceWithContext(aws.Context, *appregistry.DisassociateResourceInput, ...request.Option) (*appregistry.DisassociateResourceOutput, error)
	DisassociateResourceRequest(*appregistry.DisassociateResourceInput) (*request.Request, *appregistry.DisassociateResourceOutput)

	GetApplication(*appregistry.GetApplicationInput) (*appregistry.GetApplicationOutput, error)
	GetApplicationWithContext(aws.Context, *appregistry.GetApplicationInput, ...request.Option) (*appregistry.GetApplicationOutput, error)
	GetApplicationRequest(*appregistry.GetApplicationInput) (*request.Request, *appregistry.GetApplicationOutput)

	GetAttributeGroup(*appregistry.GetAttributeGroupInput) (*appregistry.GetAttributeGroupOutput, error)
	GetAttributeGroupWithContext(aws.Context, *appregistry.GetAttributeGroupInput, ...request.Option) (*appregistry.GetAttributeGroupOutput, error)
	GetAttributeGroupRequest(*appregistry.GetAttributeGroupInput) (*request.Request, *appregistry.GetAttributeGroupOutput)

	ListApplications(*appregistry.ListApplicationsInput) (*appregistry.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *appregistry.ListApplicationsInput, ...request.Option) (*appregistry.ListApplicationsOutput, error)
	ListApplicationsRequest(*appregistry.ListApplicationsInput) (*request.Request, *appregistry.ListApplicationsOutput)

	ListApplicationsPages(*appregistry.ListApplicationsInput, func(*appregistry.ListApplicationsOutput, bool) bool) error
	ListApplicationsPagesWithContext(aws.Context, *appregistry.ListApplicationsInput, func(*appregistry.ListApplicationsOutput, bool) bool, ...request.Option) error

	ListAssociatedAttributeGroups(*appregistry.ListAssociatedAttributeGroupsInput) (*appregistry.ListAssociatedAttributeGroupsOutput, error)
	ListAssociatedAttributeGroupsWithContext(aws.Context, *appregistry.ListAssociatedAttributeGroupsInput, ...request.Option) (*appregistry.ListAssociatedAttributeGroupsOutput, error)
	ListAssociatedAttributeGroupsRequest(*appregistry.ListAssociatedAttributeGroupsInput) (*request.Request, *appregistry.ListAssociatedAttributeGroupsOutput)

	ListAssociatedAttributeGroupsPages(*appregistry.ListAssociatedAttributeGroupsInput, func(*appregistry.ListAssociatedAttributeGroupsOutput, bool) bool) error
	ListAssociatedAttributeGroupsPagesWithContext(aws.Context, *appregistry.ListAssociatedAttributeGroupsInput, func(*appregistry.ListAssociatedAttributeGroupsOutput, bool) bool, ...request.Option) error

	ListAssociatedResources(*appregistry.ListAssociatedResourcesInput) (*appregistry.ListAssociatedResourcesOutput, error)
	ListAssociatedResourcesWithContext(aws.Context, *appregistry.ListAssociatedResourcesInput, ...request.Option) (*appregistry.ListAssociatedResourcesOutput, error)
	ListAssociatedResourcesRequest(*appregistry.ListAssociatedResourcesInput) (*request.Request, *appregistry.ListAssociatedResourcesOutput)

	ListAssociatedResourcesPages(*appregistry.ListAssociatedResourcesInput, func(*appregistry.ListAssociatedResourcesOutput, bool) bool) error
	ListAssociatedResourcesPagesWithContext(aws.Context, *appregistry.ListAssociatedResourcesInput, func(*appregistry.ListAssociatedResourcesOutput, bool) bool, ...request.Option) error

	ListAttributeGroups(*appregistry.ListAttributeGroupsInput) (*appregistry.ListAttributeGroupsOutput, error)
	ListAttributeGroupsWithContext(aws.Context, *appregistry.ListAttributeGroupsInput, ...request.Option) (*appregistry.ListAttributeGroupsOutput, error)
	ListAttributeGroupsRequest(*appregistry.ListAttributeGroupsInput) (*request.Request, *appregistry.ListAttributeGroupsOutput)

	ListAttributeGroupsPages(*appregistry.ListAttributeGroupsInput, func(*appregistry.ListAttributeGroupsOutput, bool) bool) error
	ListAttributeGroupsPagesWithContext(aws.Context, *appregistry.ListAttributeGroupsInput, func(*appregistry.ListAttributeGroupsOutput, bool) bool, ...request.Option) error

	UpdateApplication(*appregistry.UpdateApplicationInput) (*appregistry.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *appregistry.UpdateApplicationInput, ...request.Option) (*appregistry.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*appregistry.UpdateApplicationInput) (*request.Request, *appregistry.UpdateApplicationOutput)

	UpdateAttributeGroup(*appregistry.UpdateAttributeGroupInput) (*appregistry.UpdateAttributeGroupOutput, error)
	UpdateAttributeGroupWithContext(aws.Context, *appregistry.UpdateAttributeGroupInput, ...request.Option) (*appregistry.UpdateAttributeGroupOutput, error)
	UpdateAttributeGroupRequest(*appregistry.UpdateAttributeGroupInput) (*request.Request, *appregistry.UpdateAttributeGroupOutput)
}

var _ AppRegistryAPI = (*appregistry.AppRegistry)(nil)
