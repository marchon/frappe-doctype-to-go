//
// Auto-generated with frappe-doctype-to-go
//
package erpnext_api

// Link define a link to any DocType
type Link struct {
  Creation    string `json:"creation,omitempty"`
  Docstatus   int    `json:"docstatus,omitempty"`
  Doctype     string `json:"doctype,omitempty"`
  Idx         int    `json:"idx,omitempty"`
  LinkDoctype string `json:"link_doctype"`
  LinkName    string `json:"link_name"`
  Modified    string `json:"modified,omitempty"`
  ModifiedBy  string `json:"modified_by,omitempty"`
  Name        string `json:"name,omitempty"`
  Owner       string `json:"owner,omitempty"`
  Parent      string `json:"parent,omitempty"`
  Parentfield string `json:"parentfield,omitempty"`
  Parenttype  string `json:"parenttype,omitempty`
}

// Validate run validation for mandatory fields
func (d *Link) Validate() (isValid bool, m map[string]string) {
  isValid = true

  if len(d.LinkDoctype) == 0 {
    m["LinkDoctype"] = "Field 'LinkDoctype' cannot be empty"
    isValid = false
  }

  if len(d.LinkName) == 0 {
    m["LinkName"] = "Field 'LinkName' cannot be empty"
    isValid = false
  }
  return
}