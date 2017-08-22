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
			_, err := svc.DeleteStack(
				&cloudformation.DeleteStackInput{
					StackName: aws.String(name),
				},
			)
			if err != nil {
				return err
			}

			errors := 0
			index := 0
			for {
				time.Sleep(10 * time.Second)
				response, err := svc.DescribeStacks(
					&cloudformation.DescribeStacksInput{
						StackName: aws.String(name),
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
				name = *response.Stacks[0].StackId
				errors = 0
				events, err := svc.DescribeStackEvents(
					&cloudformation.DescribeStackEventsInput{
						StackName: aws.String(name),
					},
				)
				for ; index < len(events.StackEvents); index++ {
					fmt.Println(events.StackEvents[index].ResourceStatusReason)
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
