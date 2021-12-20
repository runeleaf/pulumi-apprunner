package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apprunner"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createAppRunner(ctx *pulumi.Context) (*apprunner.Service, error) {
	repository := pulumi.String("https://github.com/runeleaf/pulumi-app-example")
	arn := os.Getenv("GITHUB_CONNECTION_ARN")

	res, err := apprunner.NewService(ctx, "pulumi-app-example", &apprunner.ServiceArgs{
		ServiceName: pulumi.String("pulumi-app-example"),
		SourceConfiguration: &apprunner.ServiceSourceConfigurationArgs{
			AuthenticationConfiguration: &apprunner.ServiceSourceConfigurationAuthenticationConfigurationArgs{
				ConnectionArn: pulumi.String(arn),
			},
			CodeRepository: &apprunner.ServiceSourceConfigurationCodeRepositoryArgs{
				CodeConfiguration: &apprunner.ServiceSourceConfigurationCodeRepositoryCodeConfigurationArgs{
					CodeConfigurationValues: &apprunner.ServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesArgs{
						BuildCommand: pulumi.String("npm install && npm run build"),
						Port:         pulumi.String("3000"),
						Runtime:      pulumi.String("NODEJS_12"),
						StartCommand: pulumi.String("npm start"),
					},
					ConfigurationSource: pulumi.String("API"),
				},
				RepositoryUrl: repository,
				SourceCodeVersion: &apprunner.ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs{
					Type:  pulumi.String("BRANCH"),
					Value: pulumi.String("main"),
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
