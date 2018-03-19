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

// Contact type
type Contact struct {
  FirstName string `json:"first_name"`
  LastName string `json:"last_name,omitempty"`
  EmailId string `json:"email_id,omitempty"`
  Status string `json:"status,omitempty"`
  Phone string `json:"phone,omitempty"`
  MobileNo string `json:"mobile_no,omitempty"`
  IsPrimaryContact bool `json:"is_primary_contact,omitempty"`
  Department string `json:"department,omitempty"`
  Designation string `json:"designation,omitempty"`
  Unsubscribed bool `json:"unsubscribed,omitempty"`
  Links []Link `json:"links,omitempty"`
}