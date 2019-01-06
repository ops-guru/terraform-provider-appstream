package main

import (
	"log"
    	"github.com/hashicorp/terraform/helper/schema"
	"github.com/aws/aws-sdk-go/aws"
    	"github.com/aws/aws-sdk-go/aws/credentials"
    	"github.com/aws/aws-sdk-go/aws/session"
    	"github.com/aws/aws-sdk-go/service/appstream"
)

type Config struct {
	AccessKey       string
    	SecretKey       string
    	Token           string
    	Region          string
}

type AwsClient struct{
	appstream	*appstream.AppStream
}


func Provider() *schema.Provider {
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Access key id",
			},

			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Secret key id",
			},

			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Security token",
			},

            		"region": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"AWS_REGION",
					"AWS_DEFAULT_REGION",
				}, nil),
				Description:  "Region",
				InputDefault: "us-east-1",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
		    "appstream_stack":  resourceAppstreamStack(),
		    "appstream_image_builder":  resourceAppstreamImageBuilder(),
		    "appstream_fleet":  resourceAppstreamFleet(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := Config{
        	AccessKey:  d.Get("access_key").(string),
        	SecretKey:  d.Get("secret_key").(string),
		Token:      d.Get("token").(string),
		Region:     d.Get("region").(string),
	}

	var client AwsClient

	sess, err := session.NewSession(&aws.Config{
        	Region:      aws.String(c.Region),
        	Credentials: credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, c.Token),
    	})
	if err != nil {
		log.Printf("[ERROR] Error creating session: %s", err)
		return nil, err
	}
	client.appstream = appstream.New(sess)

	return &client, nil

}
