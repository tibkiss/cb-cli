package cli

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/sequenceiq/hdc-cli/client/blueprints"
	"github.com/urfave/cli"
)

func ListBlueprints(c *cli.Context) error {
	client := NewOAuth2HTTPClient(c.String(FlCBServer.Name), c.String(FlCBUsername.Name), c.String(FlCBPassword.Name)).Cloudbreak

	// make the request to get all items
	resp, err := client.Blueprints.GetPublics(&blueprints.GetPublicsParams{})
	if err != nil {
		log.Error(err)
	}

	for _, v := range resp.Payload {
		fmt.Printf("%s\n", v.Name)
	}
	return nil

}
