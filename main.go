package main

import (
	"fmt"
	"log"
	"os"

	packer "github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2021-04-30/client/packer_service"

	"github.com/hashicorp/hcp-sdk-go/httpclient"
)

func main() {

	// Initialize SDK http client
	cl, err := httpclient.New(httpclient.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Import versioned client for each service.
	packerClient := packer.New(cl, nil)

	// These IDs can be obtained from the portal URL
	orgID := os.Getenv("HCP_ORGANIZATION_ID")
	projID := os.Getenv("HCP_PROJECT_ID")
	bucketSlug := os.Getenv("HCP_BUCKET_SLUG")
	listParamsA := packer.NewPackerServiceGetBucketParams()
	listParamsA.LocationOrganizationID = orgID
	listParamsA.LocationProjectID = projID
	listParamsA.BucketSlug = bucketSlug
	respA, errA := packerClient.PackerServiceGetBucket(listParamsA, nil)
	latestIterationID := respA.Payload.Bucket.LatestIteration.ID

	if errA != nil {
		log.Fatal(errA)
	}

	if len(respA.Payload.Bucket.ID) > 0 {
		fmt.Println("### HCP Bucket Slug ###")
		fmt.Printf(respA.Payload.Bucket.Slug)
		fmt.Println()
		fmt.Println("### HCP Bucket ID ###")
		fmt.Printf(respA.Payload.Bucket.ID)
		fmt.Println()
		fmt.Println("### HCP Bucket Description ###")
		fmt.Printf(respA.Payload.Bucket.Description)
		fmt.Println("############################")
		fmt.Println("### HCP Latest Iteration ###")
		fmt.Printf(respA.Payload.Bucket.LatestIteration.ID)
		fmt.Println()

	} else {
		fmt.Printf("Response: %#v\n\n", respA.Payload.Bucket.ID)
	}

	listParamsB := packer.NewPackerServiceGetIterationParams()
	listParamsB.LocationOrganizationID = orgID
	listParamsB.LocationProjectID = projID
	listParamsB.BucketSlug = bucketSlug
	listParamsB.IterationID = &latestIterationID
	respB, errB := packerClient.PackerServiceGetIteration(listParamsB, nil)

	if errB != nil {
		log.Fatal(errB)
	}

	if len(respB.Payload.Iteration.BucketSlug) > 0 {
		fmt.Println("### HCP Iteration AuthorID ###")
		fmt.Printf(respB.Payload.Iteration.AuthorID)
		fmt.Println()
		fmt.Println("### HCP Iteration Ancestor ID ###")
		fmt.Printf(respB.Payload.Iteration.IterationAncestorID)
		fmt.Println()
		fmt.Println("### HCP Iteration Fingerprint  ###")
		fmt.Printf(respB.Payload.Iteration.Fingerprint)
		fmt.Println()
		fmt.Println("### Image ID ###")
		fmt.Printf(respB.Payload.Iteration.Builds[len(respB.Payload.Iteration.Builds)-1].Images[len(respB.Payload.Iteration.Builds[len(respB.Payload.Iteration.Builds)-1].Images)-1].ImageID)
		fmt.Println()
		//fmt.Println("### Image Creation Date ###")
		//fmt.Printf(respB.Payload.Iteration.Builds[len(respB.Payload.Iteration.Builds)-1].CreatedAt.String())
		//fmt.Println()
		fmt.Println("### Image Region ###")
		fmt.Printf(respB.Payload.Iteration.Builds[len(respB.Payload.Iteration.Builds)-1].Images[len(respB.Payload.Iteration.Builds[len(respB.Payload.Iteration.Builds)-1].Images)-1].Region)
		fmt.Println()

	} else {
		fmt.Printf("Response: %#v\n\n", respB.Payload.Iteration)
	}

	listParamsC := packer.NewPackerServiceGetRegistryTFCRunTaskAPIParams()
	listParamsC.LocationOrganizationID = orgID
	listParamsC.LocationProjectID = projID
	taskID := "validation"
	listParamsC.TaskType = taskID
	respC, errC := packerClient.PackerServiceGetRegistryTFCRunTaskAPI(listParamsC, nil)


	if errC != nil {
		log.Fatal(errC)
	}

	if len(respC.Payload.APIURL) > 0 {
		fmt.Println("### HCP API URL ###")
		fmt.Printf(respC.Payload.APIURL)
		fmt.Println()
		fmt.Println("### HCP HMAC ###")
		fmt.Printf(respC.Payload.HmacKey)
		fmt.Println()
	} else {
		fmt.Printf("Response: %#v\n\n", respC.Payload)
	}
}
