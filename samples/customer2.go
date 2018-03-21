/dev/stdout
//
// Auto-generated with frappe-doctype-to-go
//
package erpnext_api

// Customer type
type Customer struct {
  Creation    string `json:"creation,omitempty"`
  Docstatus   int    `json:"docstatus,omitempty"`
  Doctype     string `json:"doctype,omitempty"`
  Modified    string `json:"modified,omitempty"`
  ModifiedBy  string `json:"modified_by,omitempty"`
  Name        string `json:"name,omitempty"`
  Owner       string `json:"owner,omitempty"`
  NamingSeries string `json:"naming_series,omitempty"`
  CustomerName string `json:"customer_name"`
  CustomerType string `json:"customer_type"`
  TaxId string `json:"tax_id,omitempty"`
  Disabled int `json:"disabled,omitempty"`
  Website string `json:"website,omitempty"`
  CreditLimit float32 `json:"credit_limit,omitempty"`
  BypassCreditLimitCheckAtSalesOrder int `json:"bypass_credit_limit_check_at_sales_order,omitempty"`
  CustomerDetails string `json:"customer_details,omitempty"`
  IsFrozen int `json:"is_frozen,omitempty"`
  DefaultCommissionRate float32 `json:"default_commission_rate,omitempty"`
  CustomerPosId string `json:"customer_pos_id,omitempty"`
  Links []Link `json:"links,omitempty"`
}

// Validate run validation for mandatory fields
func (d *Customer) Validate() (bool, string) {
  if len(d.CustomerName) == 0 {
    return false,  "Field 'CustomerName' cannot be empty"
  }
  if len(d.CustomerType) == 0 {
    return false,  "Field 'CustomerType' cannot be empty"
  }
  return true, ""
}

// CustomerOps Define operations for Customer
type CustomerOps struct {
  Session *Session
}

// $opsType Define methods for Customer
func NewCustomerOps(Session *s) *CustomerOps {
  return &CustomerOps{
    Session: s
  }
}

// Create creates a new Customer in ERPNext
func (CustomerOps) Create(doc Customer) (err error) {
  return
}

// Delete an existing Customer from ERPNext
func (CustomerOps) Delete(doc Customer) (err error) {
  return
}

// Update an existing Customer in ERPNext
func (CustomerOps) Update(doc Customer) (err error) {
  return
}

// Get an existing Customer from ERPNext
func (CustomerOps) Get(key string) (doc Customer, err error) {
  return
}

// Get existing Customer based on a filter from ERPNext
func (CustomerOps) Filter(filter Filter) (docs []Customer, err error)  {
  return
}
