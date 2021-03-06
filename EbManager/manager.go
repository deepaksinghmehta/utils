package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	//"os"
	//"encoding/json"
	//"text/template"
	//"log"
)

var settings struct {
	ServerMode bool `json:"serverMode"`
	SourceDir  string `json:"sourceDir"`
	TargetDir  string `json:"targetDir"`
	Local_token 	   string `json:"local_token"`
}


// To configure option settings
//
// The following operation configures several options in the aws:elb:loadbalancer namespace:
func ExampleElasticBeanstalk_UpdateEnvironment_shared01() {
	//svc := elasticbeanstalk.New(session.New())
	svc := elasticbeanstalk.New(session.New(), aws.NewConfig().WithRegion("eu-west-1"))
	input := &elasticbeanstalk.UpdateEnvironmentInput{
		EnvironmentName: aws.String("docker-multicontainer-v2-dev"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			{
				Namespace:  aws.String("aws:autoscaling:scheduledaction:"),
				ResourceName: aws.String("ScheduledPeriodicScaleUp"),
				OptionName: aws.String("MinSize"),
				Value:      aws.String("2"),
			},
			{
				Namespace:  aws.String("aws:autoscaling:scheduledaction:"),
				ResourceName: aws.String("ScheduledPeriodicScaleUp"),
				OptionName: aws.String("MaxSize"),
				Value:      aws.String("4"),
			},
			{
				Namespace:  aws.String("aws:autoscaling:scheduledaction:"),
				ResourceName: aws.String("ScheduledPeriodicScaleUp"),
				OptionName: aws.String("DesiredCapacity"),
				Value:      aws.String("2"),
			},
			{
				Namespace:  aws.String("aws:autoscaling:scheduledaction:"),
				ResourceName: aws.String("ScheduledPeriodicScaleUp"),
				OptionName: aws.String("StartTime"),
				Value:      aws.String("2018-01-14T07:00:00Z"),
			},
			{
				Namespace:  aws.String("aws:autoscaling:scheduledaction:"),
				ResourceName: aws.String("ScheduledPeriodicScaleUp"),
				OptionName: aws.String("EndTime"),
				Value:      aws.String("2018-05-14T07:00:00Z"),
			},
			{
				Namespace:  aws.String("aws:autoscaling:scheduledaction:"),
				ResourceName: aws.String("ScheduledPeriodicScaleUp"),
				OptionName: aws.String("Recurrence"),
				Value:      aws.String("0 18 * * 5"),
			},
		},
	}

	result, err := svc.UpdateEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view a list of applications
//
// The following operation retrieves information about applications in the current region:
func ExampleElasticBeanstalk_DescribeApplications_shared00() {
	svc := elasticbeanstalk.New(session.New(), aws.NewConfig().WithRegion("eu-west-1"))
	input := &elasticbeanstalk.DescribeApplicationsInput{}

	result, err := svc.DescribeApplications(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	
	if result != nil {

		for _, element  := range result.Applications {
			fmt.Printf("Application name  is [%s]\n", element.ApplicationName)
		}
	}
	fmt.Println("printing result")
	fmt.Println(result)
}

//this method is used to create application
func creteebApplication(){
	//svc := elasticbeanstalk.New(session.New())
	// Create a ElasticBeanstalk client with additional configuration
	svc := elasticbeanstalk.New(session.New(), aws.NewConfig().WithRegion("eu-west-1"))
	input := &elasticbeanstalk.CreateApplicationInput{
		ApplicationName: aws.String("my-app-from-go"),
		Description:     aws.String("my application creatinhg using go"),
	}

	result, err := svc.CreateApplication(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyApplicationsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyApplicationsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println("printing result")
	fmt.Println(result)
}

// To configure option settings
//
// The following operation configures several options in the aws:elb:loadbalancer namespace:
//func ExampleElasticBeanstalk_UpdateEnvironment_shared01() {
//	svc := elasticbeanstalk.New(session.New())
//	input := &elasticbeanstalk.UpdateEnvironmentInput{
//		EnvironmentName: aws.String("my-env"),
//		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
//			{
//				Namespace:  aws.String("aws:elb:healthcheck"),
//				OptionName: aws.String("Interval"),
//				Value:      aws.String("15"),
//			},
//			{
//				Namespace:  aws.String("aws:elb:healthcheck"),
//				OptionName: aws.String("Timeout"),
//				Value:      aws.String("8"),
//			},
//			{
//				Namespace:  aws.String("aws:elb:healthcheck"),
//				OptionName: aws.String("HealthyThreshold"),
//				Value:      aws.String("2"),
//			},
//			{
//				Namespace:  aws.String("aws:elb:healthcheck"),
//				OptionName: aws.String("UnhealthyThreshold"),
//				Value:      aws.String("3"),
//			},
//		},
//	}
//
//	result, err := svc.UpdateEnvironment(input)
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//			switch aerr.Code() {
//			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
//				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
//			case elasticbeanstalk.ErrCodeTooManyBucketsException:
//				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
//			default:
//				fmt.Println(aerr.Error())
//			}
//		} else {
//			// Print the error, cast err to awserr.Error to get the Code and
//			// Message from an error.
//			fmt.Println(err.Error())
//		}
//		return
//	}
//
//	fmt.Println(result)
//}

func main() {
	//ExampleElasticBeanstalk_DescribeApplications_shared00()
	ExampleElasticBeanstalk_UpdateEnvironment_shared01()
	//creteebApplication()
	//t, err := template.ParseFiles( "envconfig.yaml.template")
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	//
	//f, err := os.Create("envconfig.yaml")
	//if err != nil {
	//	log.Println("create file: ", err)
	//	return
	//}
	//
	//// then config file settings
	//
	//configFile, err := os.Open("config.json")
	//if err != nil {
	//	fmt.Println("opening config file", err.Error())
	//}
	//defer configFile.Close()
	//jsonParser := json.NewDecoder(configFile)
	//if err = jsonParser.Decode(&settings); err != nil {
	//	fmt.Println("parsing config file", err.Error())
	//}

	//fmt.Printf("%v %s %s %s", settings.ServerMode, settings.SourceDir, settings.TargetDir, settings.Locale)

	//err = t.Execute(f, settings)
	//if err != nil {
	//	log.Print("execute: ", err)
	//	return
	//}
	return

}
