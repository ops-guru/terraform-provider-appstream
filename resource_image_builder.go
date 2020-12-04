package main

import (
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAppstreamImageBuilder() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppstreamImageBuilderCreate,
		Read:   resourceAppstreamImageBuilderRead,
		Update: resourceAppstreamImageBuilderUpdate,
		Delete: resourceAppstreamImageBuilderDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"appstream_agent_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"domain_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"directory_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"organizational_unit_distinguished_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"enable_default_internet_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"image_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"instance_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"iam_role_arn": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"image_arn": {
				Type:     schema.TypeString,
				Required: true,
			},

			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"vpc_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_group_ids": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet_ids": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAppstreamImageBuilderCreate(d *schema.ResourceData, meta interface{}) error {

	svc := meta.(*AwsClient).appstream

	CreateImageBuilderInputOpts := &appstream.CreateImageBuilderInput{}

	if v, ok := d.GetOk("name"); ok {
		CreateImageBuilderInputOpts.Name = aws.String(v.(string))
	}

	if v, ok := d.GetOk("appstream_agent_version"); ok {
		CreateImageBuilderInputOpts.AppstreamAgentVersion = aws.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		CreateImageBuilderInputOpts.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("display_name"); ok {
		CreateImageBuilderInputOpts.DisplayName = aws.String(v.(string))
	}

	DomainJoinInfoConfig := &appstream.DomainJoinInfo{}

	if dom, ok := d.GetOk("domain_info"); ok {
		DomainAttributes := dom.([]interface{})
		attr := DomainAttributes[0].(map[string]interface{})
		if v, ok := attr["directory_name"]; ok {
			DomainJoinInfoConfig.DirectoryName = aws.String(v.(string))
		}
		if v, ok := attr["organizational_unit_distinguished_name"]; ok {
			DomainJoinInfoConfig.OrganizationalUnitDistinguishedName = aws.String(v.(string))
		}
		CreateImageBuilderInputOpts.DomainJoinInfo = DomainJoinInfoConfig
	}

	if v, ok := d.GetOk("enable_default_internet_access"); ok {
		CreateImageBuilderInputOpts.EnableDefaultInternetAccess = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("image_name"); ok {
		CreateImageBuilderInputOpts.ImageName = aws.String(v.(string))
	}

	if v, ok := d.GetOk("instance_type"); ok {
		CreateImageBuilderInputOpts.InstanceType = aws.String(v.(string))
	}

	if v, ok := d.GetOk("iam_role_arn"); ok {
		CreateImageBuilderInputOpts.IamRoleArn = aws.String(v.(string))
	}

	if v, ok := d.GetOk("image_arn"); ok {
		CreateImageBuilderInputOpts.ImageArn = aws.String(v.(string))
	}

	VpcConfigConfig := &appstream.VpcConfig{}

	if vpc, ok := d.GetOk("vpc_config"); ok {
		VpcAttributes := vpc.([]interface{})
		attr := VpcAttributes[0].(map[string]interface{})
		if v, ok := attr["security_group_ids"]; ok {
			strSlice := strings.Split(v.(string), ",")
			for i, s := range strSlice {
				strSlice[i] = strings.TrimSpace(s)
			}
			VpcConfigConfig.SecurityGroupIds = aws.StringSlice(strSlice)
		}
		if v, ok := attr["subnet_ids"]; ok {
			strSlice := strings.Split(v.(string), ",")
			for i, s := range strSlice {
				strSlice[i] = strings.TrimSpace(s)
			}
			VpcConfigConfig.SubnetIds = aws.StringSlice(strSlice)
		}
		CreateImageBuilderInputOpts.VpcConfig = VpcConfigConfig
	}

	log.Printf("[DEBUG] Run configuration: %s", CreateImageBuilderInputOpts)

	resp, err := svc.CreateImageBuilder(CreateImageBuilderInputOpts)

	if err != nil {
		log.Printf("[ERROR] Error creating Appstream Image Builder: %s", err)
		return err
	}

	log.Printf("[DEBUG] Image builder created %s", resp)

	ImageBuilderName := aws.StringValue(CreateImageBuilderInputOpts.Name)
	for {

		resp, err := svc.DescribeImageBuilders(&appstream.DescribeImageBuildersInput{
			Names: aws.StringSlice([]string{ImageBuilderName}),
		})

		if err != nil {
			log.Printf("[ERROR] Error describing Appstream Image Builder: %s", err)
			return err
		}

		state := resp.ImageBuilders[0].State
		if aws.StringValue(state) == "RUNNING" {
			break
		}
		if aws.StringValue(state) != "RUNNING" {
			log.Printf("[DEBUG] Image Builder not running")
			time.Sleep(20 * time.Second)
			continue
		}

	}

	d.SetId(*CreateImageBuilderInputOpts.Name)

	return resourceAppstreamImageBuilderRead(d, meta)
}

func resourceAppstreamImageBuilderRead(d *schema.ResourceData, meta interface{}) error {

	svc := meta.(*AwsClient).appstream

	resp, err := svc.DescribeImageBuilders(&appstream.DescribeImageBuildersInput{})
	if err != nil {
		log.Printf("[ERROR] Error describing Appstream Image Builder: %s", err)
		return err
	}

	for _, v := range resp.ImageBuilders {

		if aws.StringValue(v.Name) == d.Get("name") {
			d.Set("name", v.Name)
			d.Set("description", v.Description)
			d.Set("display_name", v.DisplayName)
			d.Set("appstream_agent_version", v.AppstreamAgentVersion)
			d.Set("enable_default_internet_access", v.EnableDefaultInternetAccess)
			d.Set("instance_type", v.InstanceType)
			d.Set("image_name", d.Get("image_name"))
			d.Set("state", v.State)
			if v.VpcConfig != nil {
				vpcAttr := map[string]interface{}{}
				vpcConfigSg := aws.StringValueSlice(v.VpcConfig.SecurityGroupIds)
				vpcConfigSub := aws.StringValueSlice(v.VpcConfig.SubnetIds)
				vpcAttr["security_group_ids"] = aws.String(strings.Join(vpcConfigSg, ","))
				vpcAttr["subnet_ids"] = aws.String(strings.Join(vpcConfigSub, ","))
				d.Set("vpc_config", vpcAttr)
			}
			return nil
		}
	}

	d.SetId("")
	return nil

}

// Apstream2.0 doesn't support imageBuilder updates
func resourceAppstreamImageBuilderUpdate(d *schema.ResourceData, meta interface{}) error {

	svc := meta.(*AwsClient).appstream

	StartImageBuilderInputOptions := &appstream.StartImageBuilderInput{}
	StopImageBuilderInputOptions := &appstream.StopImageBuilderInput{}

	d.Partial(true)

	if v, ok := d.GetOk("name"); ok {
		StartImageBuilderInputOptions.Name = aws.String(v.(string))
		StopImageBuilderInputOptions.Name = aws.String(v.(string))
	}

	desiredState := d.Get("state")

	if d.HasChange("state") {
		d.SetPartial("state")
		if desiredState == "STOPPED" {
			svc.StopImageBuilder(StopImageBuilderInputOptions)
		} else if desiredState == "RUNNING" {
			svc.StartImageBuilder(StartImageBuilderInputOptions)
		}

		for {

			resp, err := svc.DescribeImageBuilders(&appstream.DescribeImageBuildersInput{
				Names: aws.StringSlice([]string{*StartImageBuilderInputOptions.Name}),
			})

			if err != nil {
				log.Printf("[ERROR] Error describing Appstream Image Builder: %s", err)
				return err
			}

			currState := resp.ImageBuilders[0].State
			if aws.StringValue(currState) == desiredState {
				break
			}
			if aws.StringValue(currState) != desiredState {
				time.Sleep(20 * time.Second)
				continue
			}

		}
	}

	d.Partial(false)
	return resourceAppstreamImageBuilderRead(d, meta)

}

func resourceAppstreamImageBuilderDelete(d *schema.ResourceData, meta interface{}) error {
	svc := meta.(*AwsClient).appstream

	ImageBuilderName := d.Id()

	resp, err := svc.DescribeImageBuilders(&appstream.DescribeImageBuildersInput{
		Names: aws.StringSlice([]string{ImageBuilderName}),
	})

	if err != nil {
		log.Printf("[ERROR] Error describing Appstream Image Builder: %s", err)
		return err
	}

	state := resp.ImageBuilders[0].State

	if aws.StringValue(state) == "RUNNING" {
		resp, err := svc.StopImageBuilder(&appstream.StopImageBuilderInput{
			Name: aws.String(d.Id()),
		})

		if err != nil {
			log.Printf("[ERROR] Error stopping Appstream Image Builder: %s", err)
			return err
		}

		log.Printf("[DEBUG] %s", resp)

		for {

			resp, err := svc.DescribeImageBuilders(&appstream.DescribeImageBuildersInput{
				Names: aws.StringSlice([]string{ImageBuilderName}),
			})

			if err != nil {
				log.Printf("[ERROR] Error describing Appstream Image Builder: %s", err)
				return err
			}

			state := resp.ImageBuilders[0].State
			if aws.StringValue(state) == "STOPPED" {
				break
			}
			if aws.StringValue(state) != "STOPPED" {
				log.Printf("[DEBUG] Image Builder not running")
				time.Sleep(20 * time.Second)
				continue
			}

		}
	}

	del, err := svc.DeleteImageBuilder(&appstream.DeleteImageBuilderInput{
		Name: aws.String(d.Id()),
	})
	if err != nil {
		log.Printf("[ERROR] Error deleting Appstream Image Builder: %s", err)
		return err
	}
	log.Printf("[DEBUG] %s", del)

	return nil
}
