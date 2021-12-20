package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		apprunner, err := createAppRunner(ctx)
		if err != nil {
			return err
		}

		ctx.Export("App Runner", apprunner.ID())
		return nil
	})
}
