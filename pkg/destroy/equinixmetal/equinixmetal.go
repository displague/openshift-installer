package equinixmetal

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/openshift/installer/pkg/asset/installconfig/equinixmetal"
	"github.com/openshift/installer/pkg/destroy/providers"
	"github.com/openshift/installer/pkg/types"
)

// ClusterUninstaller holds the various options for the cluster we want to delete.
type ClusterUninstaller struct {
	Metadata types.ClusterMetadata
	Logger   logrus.FieldLogger
}

// Run is the entrypoint to start the uninstall process.
func (uninstaller *ClusterUninstaller) Run() error {
	_, err := equinixmetal.NewConnection()
	if err != nil {
		return fmt.Errorf("failed to initialize connection to equinixmetal-engine's %s", err)
	}
	// @TODO(displague) delete each thing
	//if err := uninstaller.deleteThing(con); err != nil {
	//		uninstaller.Logger.Errorf("Failed to remove Thing: %s", err)
	//	}

	return nil
}

// New returns Equinix Metal Uninstaller from ClusterMetadata.
func New(logger logrus.FieldLogger, metadata *types.ClusterMetadata) (providers.Destroyer, error) {
	return &ClusterUninstaller{
		Metadata: *metadata,
		Logger:   logger,
	}, nil
}
