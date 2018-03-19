//
// Auto-generated with frappe-doctype-to-go
//
package bart


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

type Customer struct {
  NamingSeries string `json:"naming_series,omitempty"`
  CustomerName string `json:"customer_name"`
  CustomerType string `json:"customer_type"`
  TaxId string `json:"tax_id,omitempty"`
  Disabled bool `json:"disabled,omitempty"`
  Website string `json:"website,omitempty"`
  CreditLimit float32 `json:"credit_limit,omitempty"`
  BypassCreditLimitCheckAtSalesOrder bool `json:"bypass_credit_limit_check_at_sales_order,omitempty"`
  CustomerDetails string `json:"customer_details,omitempty"`
  IsFrozen bool `json:"is_frozen,omitempty"`
  DefaultCommissionRate float32 `json:"default_commission_rate,omitempty"`
  CustomerPosId string `json:"customer_pos_id,omitempty"`
  Links []Link `json:"links,omitempty"`
}