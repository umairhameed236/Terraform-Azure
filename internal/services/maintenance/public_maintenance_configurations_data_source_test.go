package maintenance_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type PublicMaintenanceConfigurationsDataSource struct{}

func TestAccDataSourcePublicMaintenanceConfigurations_noFilters(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_public_maintenance_configurations", "test")
	r := PublicMaintenanceConfigurationsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.noFilters(),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("public_maintenance_configurations.30.name").Exists(),
			),
		},
	})
}

func TestAccDataSourcePublicMaintenanceConfigurations_allFilters(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_public_maintenance_configurations", "test")
	r := PublicMaintenanceConfigurationsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.allFilters(),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("public_maintenance_configurations.#").HasValue("1"),
				check.That(data.ResourceName).Key("public_maintenance_configurations.0.maintenance_scope").HasValue("SQLManagedInstance"),
				check.That(data.ResourceName).Key("public_maintenance_configurations.0.name").HasValue("SQL_WestEurope_MI_1"),
				check.That(data.ResourceName).Key("public_maintenance_configurations.0.recur_every").HasValue("week Monday, Tuesday, Wednesday, Thursday"),
			),
		},
	})
}

func TestAccDataSourcePublicMaintenanceConfigurations_recurEvery(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_public_maintenance_configurations", "test")
	r := PublicMaintenanceConfigurationsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.recurEvery(),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("public_maintenance_configurations.0.maintenance_scope").HasValue("SQLManagedInstance"),
				check.That(data.ResourceName).Key("public_maintenance_configurations.0.recur_every").HasValue("week Friday, Saturday, Sunday"),
			),
		},
	})
}

func (PublicMaintenanceConfigurationsDataSource) allFilters() string {
	return fmt.Sprintf(`
data "azurerm_public_maintenance_configurations" "test" {
  location_filter    = "westeurope"
  scope_filter       = "SQLManagedInstance"
  recur_every_filter = "weekMondayToThursday"
}
`)
}

func (PublicMaintenanceConfigurationsDataSource) noFilters() string {
	return fmt.Sprintf(`
data "azurerm_public_maintenance_configurations" "test" {

}
`)
}

func (PublicMaintenanceConfigurationsDataSource) recurEvery() string {
	return fmt.Sprintf(`
data "azurerm_public_maintenance_configurations" "test" {
  scope_filter       = "SQLManagedInstance"
  recur_every_filter = "weekFridayToSunday"
}
`)
}
