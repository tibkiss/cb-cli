package models_autoscale

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*AlertRuleDefinitionEntry alert rule definition entry

swagger:model AlertRuleDefinitionEntry
*/
type AlertRuleDefinitionEntry struct {

	/* label
	 */
	Label *string `json:"label,omitempty"`

	/* name
	 */
	Name *string `json:"name,omitempty"`
}

// Validate validates this alert rule definition entry
func (m *AlertRuleDefinitionEntry) Validate(formats strfmt.Registry) error {
	return nil
}