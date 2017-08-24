package destroy

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/spf13/cobra"
)

//NewDestroy destroys a i2kit application
func NewDestroy(name string, awsConfig *aws.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: "Destroy a i2kit application",
		RunE: func(cmd *cobra.Command, args []string) error {
			svc := cloudformation.New(session.New(), awsConfig)
			response, err := svc.DescribeStacks(
				&cloudformation.DescribeStacksInput{
					StackName: aws.String(name),
				},
			)
			if len(response.Stacks) == 0 {
				fmt.Printf("Stack '%s' doesn't exist.\n", name)
				return nil
			}
			stackID := response.Stacks[0].StackId
			events, err := svc.DescribeStackEvents(
				&cloudformation.DescribeStackEventsInput{
					StackName: stackID,
				},
			)
			index := len(events.StackEvents)
			_, err = svc.DeleteStack(
				&cloudformation.DeleteStackInput{
					StackName: stackID,
				},
			)
			if err != nil {
				return err
			}
			errors := 0
			for {
				time.Sleep(10 * time.Second)
				response, err := svc.DescribeStacks(
					&cloudformation.DescribeStacksInput{
						StackName: stackID,
					},
				)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					if errors >= 3 {
						return err
					}
					continue
				}
				errors = 0
				events, err := svc.DescribeStackEvents(
					&cloudformation.DescribeStackEventsInput{
						StackName: stackID,
					},
				)
				for ; index < len(events.StackEvents); index++ {
					if events.StackEvents[index].ResourceStatusReason != nil {
						fmt.Printf("Index: %s\n", *events.StackEvents[index].ResourceStatusReason)
					}
				}
				status := *response.Stacks[0].StackStatus
				fmt.Printf("Status %s\n", status)
				if status != cloudformation.ResourceStatusDeleteInProgress {
					break
				}
			}
			return nil
		},
	}
	return cmd
}
