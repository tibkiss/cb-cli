package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/cloud"
	"github.com/hortonworks/cb-cli/cli/types"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1credentials"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

func CreateAwsCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	createCredential(c)
}

func CreateAzureCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	createCredential(c)
}

func CreateGcpCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.GCP)
	createCredential(c)
}

func CreateOpenstackCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.OPENSTACK)
	createCredential(c)
}

func CreateCredentialFromFile(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "create credential")

	req := assembleCredentialRequest(c)
	cbClient := NewCloudbreakHTTPClientFromContext(c)
	postCredential(cbClient.Cloudbreak.V1credentials, c.Bool(FlPublicOptional.Name), req)
}

func assembleCredentialRequest(c *cli.Context) *models_cloudbreak.CredentialRequest {
	path := c.String(FlInputJson.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		utils.LogErrorAndExit(err)
	}

	log.Infof("[assembleCredentialRequest] read credential json from file: %s", path)
	content := utils.ReadFile(path)

	var req models_cloudbreak.CredentialRequest
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid json format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}

	credName := c.String(FlNameOptional.Name)
	if len(credName) != 0 {
		req.Name = &credName
	}
	if req.Name == nil || len(*req.Name) == 0 {
		utils.LogErrorMessageAndExit("Name of the credential must be set either in the template or with the --name command line option.")
	}

	return &req
}

func createCredential(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "create credential")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	createCredentialImpl(c.String, c.Bool, cbClient.Cloudbreak.V1credentials)
}

type createCredentialClient interface {
	PostPublicCredential(*v1credentials.PostPublicCredentialParams) (*v1credentials.PostPublicCredentialOK, error)
	PostPrivateCredential(*v1credentials.PostPrivateCredentialParams) (*v1credentials.PostPrivateCredentialOK, error)
}

func createCredentialImpl(stringFinder func(string) string, boolFinder func(string) bool, client createCredentialClient) *models_cloudbreak.CredentialResponse {
	provider := cloud.GetProvider()
	credentialMap, err := provider.GetCredentialParameters(stringFinder, boolFinder)
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	name := stringFinder(FlName.Name)
	credReq := &models_cloudbreak.CredentialRequest{
		Name:          &name,
		Description:   &(&types.S{S: stringFinder(FlDescriptionOptional.Name)}).S,
		CloudPlatform: provider.GetName(),
		Parameters:    credentialMap,
	}
	public := boolFinder(FlPublicOptional.Name)
	return postCredential(client, public, credReq)
}

func postCredential(client createCredentialClient, public bool, credReq *models_cloudbreak.CredentialRequest) *models_cloudbreak.CredentialResponse {
	var credential *models_cloudbreak.CredentialResponse
	if public {
		log.Infof("[postCredential] sending create public credential request")
		resp, err := client.PostPublicCredential(v1credentials.NewPostPublicCredentialParams().WithBody(credReq))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		credential = resp.Payload
	} else {
		log.Infof("[postCredential] sending create private credential request")
		resp, err := client.PostPrivateCredential(v1credentials.NewPostPrivateCredentialParams().WithBody(credReq))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		credential = resp.Payload
	}
	log.Infof("[postCredential] credential created: %s (id: %d)", *credential.Name, credential.ID)
	return credential
}

func DescribeCredential(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "describe credential")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	resp, err := cbClient.Cloudbreak.V1credentials.GetPublicCredential(v1credentials.NewGetPublicCredentialParams().WithName(c.String(FlName.Name)))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	cred := resp.Payload
	output.Write(cloudResourceHeader, &cloudResourceOut{*cred.Name, *cred.Description, *cred.CloudPlatform})
}

func DeleteCredential(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "delete credential")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	name := c.String(FlName.Name)
	log.Infof("[DeleteCredential] sending delete credential request with name: %s", name)
	if err := cbClient.Cloudbreak.V1credentials.DeletePublicCredential(v1credentials.NewDeletePublicCredentialParams().WithName(name)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteCredential] credential deleted, name: %s", name)
}

func ListCredentials(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list credentials")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	listCredentialsImpl(cbClient.Cloudbreak.V1credentials, output.WriteList)
}

type getPublicsCredentialClient interface {
	GetPublicsCredential(*v1credentials.GetPublicsCredentialParams) (*v1credentials.GetPublicsCredentialOK, error)
}

func listCredentialsImpl(client getPublicsCredentialClient, writer func([]string, []utils.Row)) {
	log.Infof("[listCredentialsImpl] sending credential list request")
	credResp, err := client.GetPublicsCredential(v1credentials.NewGetPublicsCredentialParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, cred := range credResp.Payload {
		tableRows = append(tableRows, &cloudResourceOut{*cred.Name, *cred.Description, *cred.CloudPlatform})
	}

	writer(cloudResourceHeader, tableRows)
}
