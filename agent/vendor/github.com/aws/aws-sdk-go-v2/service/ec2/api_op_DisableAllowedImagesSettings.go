// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables Allowed AMIs for your account in the specified Amazon Web Services
// Region. When set to disabled , the image criteria in your Allowed AMIs settings
// do not apply, and no restrictions are placed on AMI discoverability or usage.
// Users in your account can launch instances using any public AMI or AMI shared
// with your account.
//
// The Allowed AMIs feature does not restrict the AMIs owned by your account.
// Regardless of the criteria you set, the AMIs created by your account will always
// be discoverable and usable by users in your account.
//
// For more information, see [Control the discovery and use of AMIs in Amazon EC2 with Allowed AMIs] in Amazon EC2 User Guide.
//
// [Control the discovery and use of AMIs in Amazon EC2 with Allowed AMIs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-allowed-amis.html
func (c *Client) DisableAllowedImagesSettings(ctx context.Context, params *DisableAllowedImagesSettingsInput, optFns ...func(*Options)) (*DisableAllowedImagesSettingsOutput, error) {
	if params == nil {
		params = &DisableAllowedImagesSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableAllowedImagesSettings", params, optFns, c.addOperationDisableAllowedImagesSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableAllowedImagesSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableAllowedImagesSettingsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DisableAllowedImagesSettingsOutput struct {

	// Returns disabled if the request succeeds; otherwise, it returns an error.
	AllowedImagesSettingsState types.AllowedImagesSettingsDisabledState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableAllowedImagesSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisableAllowedImagesSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisableAllowedImagesSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisableAllowedImagesSettings"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableAllowedImagesSettings(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDisableAllowedImagesSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisableAllowedImagesSettings",
	}
}