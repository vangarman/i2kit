package deploy

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/spf13/cobra"
)

var template = `
{
  "AWSTemplateFormatVersion" : "2010-09-09",

  "Description" : "AWS CloudFormation Sample Template S3_Website_Bucket_With_Retain_On_Delete: Sample template showing how to create a publicly accessible S3 bucket configured for website access with a deletion policy of retail on delete. **WARNING** This template creates an S3 bucket that will NOT be deleted when the stack is deleted. You will be billed for the AWS resources used if you create a stack from this template.",

  "Resources" : {
    "S3Bucket" : {
      "Type" : "AWS::S3::Bucket",
      "Properties" : {
        "AccessControl" : "PublicRead",
        "WebsiteConfiguration" : {
          "IndexDocument" : "index.html",
          "ErrorDocument" : "error.html"
         }
      },
      "DeletionPolicy" : "Retain"
    }
  },

  "Outputs" : {
    "WebsiteURL" : {
      "Value" : { "Fn::GetAtt" : [ "S3Bucket", "WebsiteURL" ] },
      "Description" : "URL for website hosted on S3"
    },
    "S3BucketSecureURL" : {
      "Value" : { "Fn::Join" : [ "", [ "https://", { "Fn::GetAtt" : [ "S3Bucket", "DomainName" ] } ] ] },
      "Description" : "Name of S3 bucket to hold website content"
    }
  }
}
`

//NewDeploy deploys a i2kit application
func NewDeploy(name, i2kitPath string, awsConfig *aws.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a i2kit application",
		RunE: func(cmd *cobra.Command, args []string) error {
			svc := cloudformation.New(session.New(), awsConfig)
			inStack := &cloudformation.CreateStackInput{
				Capabilities: []*string{aws.String("CAPABILITY_NAMED_IAM")},
				StackName:    &name,
				TemplateBody: &template,
				Tags: []*cloudformation.Tag{
					&cloudformation.Tag{
						Key:   aws.String("i2kit"),
						Value: aws.String("alpha"),
					},
				},
			}
			stack, err := svc.CreateStack(inStack)
			if err != nil {
				return err
			}

			errors := 0
			index := 0
			for {
				time.Sleep(10 * time.Second)
				response, err := svc.DescribeStacks(
					&cloudformation.DescribeStacksInput{
						StackName: stack.StackId,
					},
				)
				if err != nil {
					errors++
					fmt.Fprintln(os.Stderr, err)
					if errors >= 3 {
						return err
					}
					continue
				}
				errors = 0
				events, err := svc.DescribeStackEvents(
					&cloudformation.DescribeStackEventsInput{
						StackName: stack.StackId,
					},
				)
				for ; index < len(events.StackEvents); index++ {
					fmt.Println(events.StackEvents[index].ResourceStatusReason)
				}
				status := *response.Stacks[0].StackStatus
				fmt.Printf("Status %s\n", status)
				if status != cloudformation.ResourceStatusCreateInProgress {
					break
				}
			}
			return nil
		},
	}
	return cmd
}
