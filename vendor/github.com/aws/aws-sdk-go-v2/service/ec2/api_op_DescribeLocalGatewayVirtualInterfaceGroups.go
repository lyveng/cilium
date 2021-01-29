// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified local gateway virtual interface groups.
func (c *Client) DescribeLocalGatewayVirtualInterfaceGroups(ctx context.Context, params *DescribeLocalGatewayVirtualInterfaceGroupsInput, optFns ...func(*Options)) (*DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	if params == nil {
		params = &DescribeLocalGatewayVirtualInterfaceGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocalGatewayVirtualInterfaceGroups", params, optFns, addOperationDescribeLocalGatewayVirtualInterfaceGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocalGatewayVirtualInterfaceGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewayVirtualInterfaceGroupsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * local-gateway-id - The ID of a local gateway.
	//
	// *
	// local-gateway-virtual-interface-id - The ID of the virtual interface.
	//
	// *
	// local-gateway-virtual-interface-group-id - The ID of the virtual interface
	// group.
	Filters []types.Filter

	// The IDs of the virtual interface groups.
	LocalGatewayVirtualInterfaceGroupIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeLocalGatewayVirtualInterfaceGroupsOutput struct {

	// The virtual interface groups.
	LocalGatewayVirtualInterfaceGroups []types.LocalGatewayVirtualInterfaceGroup

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLocalGatewayVirtualInterfaceGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGatewayVirtualInterfaceGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGatewayVirtualInterfaceGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGatewayVirtualInterfaceGroups(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeLocalGatewayVirtualInterfaceGroupsAPIClient is a client that implements
// the DescribeLocalGatewayVirtualInterfaceGroups operation.
type DescribeLocalGatewayVirtualInterfaceGroupsAPIClient interface {
	DescribeLocalGatewayVirtualInterfaceGroups(context.Context, *DescribeLocalGatewayVirtualInterfaceGroupsInput, ...func(*Options)) (*DescribeLocalGatewayVirtualInterfaceGroupsOutput, error)
}

var _ DescribeLocalGatewayVirtualInterfaceGroupsAPIClient = (*Client)(nil)

// DescribeLocalGatewayVirtualInterfaceGroupsPaginatorOptions is the paginator
// options for DescribeLocalGatewayVirtualInterfaceGroups
type DescribeLocalGatewayVirtualInterfaceGroupsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLocalGatewayVirtualInterfaceGroupsPaginator is a paginator for
// DescribeLocalGatewayVirtualInterfaceGroups
type DescribeLocalGatewayVirtualInterfaceGroupsPaginator struct {
	options   DescribeLocalGatewayVirtualInterfaceGroupsPaginatorOptions
	client    DescribeLocalGatewayVirtualInterfaceGroupsAPIClient
	params    *DescribeLocalGatewayVirtualInterfaceGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLocalGatewayVirtualInterfaceGroupsPaginator returns a new
// DescribeLocalGatewayVirtualInterfaceGroupsPaginator
func NewDescribeLocalGatewayVirtualInterfaceGroupsPaginator(client DescribeLocalGatewayVirtualInterfaceGroupsAPIClient, params *DescribeLocalGatewayVirtualInterfaceGroupsInput, optFns ...func(*DescribeLocalGatewayVirtualInterfaceGroupsPaginatorOptions)) *DescribeLocalGatewayVirtualInterfaceGroupsPaginator {
	options := DescribeLocalGatewayVirtualInterfaceGroupsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeLocalGatewayVirtualInterfaceGroupsInput{}
	}

	return &DescribeLocalGatewayVirtualInterfaceGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLocalGatewayVirtualInterfaceGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeLocalGatewayVirtualInterfaceGroups page.
func (p *DescribeLocalGatewayVirtualInterfaceGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeLocalGatewayVirtualInterfaceGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeLocalGatewayVirtualInterfaceGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLocalGatewayVirtualInterfaceGroups",
	}
}
