provider "ibm"{
}
// provision apigateway resource instance
resource "ibm_resource_instance" "apigateway"{
    name     = "testname"
    location = "global"
    service  = "api-gateway"
    plan     = "lite"
 }

//provision endpoint resource for one api document
variable "file_path"{
}
resource "ibm_api_gateway_endpoint" "endpoint"{
    service_instance_crn = ibm_resource_instance.apigateway.id
    name="test-endpoint"
    managed="true"
    open_api_doc_name = var.file_path
    type="share" //required only when updating action
}

// provision endpoint resources for a directory of api douments..
variable "dir_path"{
}
resource "ibm_api_gateway_endpoint" "endpoint"{
    for_each=fileset(var.dir_path, "*.json")
    service_instance_crn = ibm_resource_instance.apigateway.id
    managed="true"
    name=replace("endpoint-${each.key}",".json","")
    open_api_doc_name=format("%s%s",var.dir_path,each.key)
    type="share" //required only when updating action
}

//provisioning subscription resource for one endpoint..
data "ibm_api_gateway" "endpoint"{
    service_instance_crn =ibm_resource_instance.apigateway.id
}
resource "ibm_api_gateway_endpoint_subscription" "subs" {
  artifact_id = data.ibm_api_gateway.endpoint.endpoints[0].endpoint_id
  client_id   = "testclientID"
  name        = "testname"
  type        = "bluemix"
  client_secret="testsecret"
}

// example datasource for API gateway endpoint..
data "ibm_api_gateway" "endpoint"{
    service_instance_crn =ibm_resource_instance.apigateway.id
}






