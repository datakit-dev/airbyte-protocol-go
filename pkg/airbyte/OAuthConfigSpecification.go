package airbyte

type OAuthConfigSpecification struct {
	// OAuth specific blob. This is a Json Schema used to validate Json configurations produced                        
	// by the OAuth flows as they are                                                                                  
	// returned by the distant OAuth APIs.                                                                             
	// Must be a valid JSON describing the fields to merge back to                                                     
	// `ConnectorSpecification.connectionSpecification`.                                                               
	// For each field, a special annotation `path_in_connector_config` can be specified to                             
	// determine where to merge it,                                                                                    
	//                                                                                                                 
	// Examples:                                                                                                       
	//                                                                                                                 
	// complete_oauth_output_specification={                                                                           
	// refresh_token: {                                                                                                
	// type: string,                                                                                                   
	// path_in_connector_config: ['credentials', 'refresh_token']                                                      
	// }                                                                                                               
	// }                                                                                                               
	CompleteOauthOutputSpecification                                                            map[string]interface{} `json:"complete_oauth_output_specification,omitempty"`
	// OAuth specific blob. This is a Json Schema used to validate Json configurations persisted                       
	// as Airbyte Server configurations.                                                                               
	// Must be a valid non-nested JSON describing additional fields configured by the Airbyte                          
	// Instance or Workspace Admins to be used by the                                                                  
	// server when completing an OAuth flow (typically exchanging an auth code for refresh                             
	// token).                                                                                                         
	//                                                                                                                 
	// Examples:                                                                                                       
	//                                                                                                                 
	// complete_oauth_server_input_specification={                                                                     
	// client_id: {                                                                                                    
	// type: string                                                                                                    
	// },                                                                                                              
	// client_secret: {                                                                                                
	// type: string                                                                                                    
	// }                                                                                                               
	// }                                                                                                               
	CompleteOauthServerInputSpecification                                                       map[string]interface{} `json:"complete_oauth_server_input_specification,omitempty"`
	// OAuth specific blob. This is a Json Schema used to validate Json configurations persisted                       
	// as Airbyte Server configurations that                                                                           
	// also need to be merged back into the connector configuration at runtime.                                        
	// This is a subset configuration of `complete_oauth_server_input_specification` that                              
	// filters fields out to retain only the ones that                                                                 
	// are necessary for the connector to function with OAuth. (some fields could be used during                       
	// oauth flows but not needed afterwards, therefore                                                                
	// they would be listed in the `complete_oauth_server_input_specification` but not                                 
	// `complete_oauth_server_output_specification`)                                                                   
	// Must be a valid non-nested JSON describing additional fields configured by the Airbyte                          
	// Instance or Workspace Admins to be used by the                                                                  
	// connector when using OAuth flow APIs.                                                                           
	// These fields are to be merged back to `ConnectorSpecification.connectionSpecification`.                         
	// For each field, a special annotation `path_in_connector_config` can be specified to                             
	// determine where to merge it,                                                                                    
	//                                                                                                                 
	// Examples:                                                                                                       
	//                                                                                                                 
	// complete_oauth_server_output_specification={                                                                    
	// client_id: {                                                                                                    
	// type: string,                                                                                                   
	// path_in_connector_config: ['credentials', 'client_id']                                                          
	// },                                                                                                              
	// client_secret: {                                                                                                
	// type: string,                                                                                                   
	// path_in_connector_config: ['credentials', 'client_secret']                                                      
	// }                                                                                                               
	// }                                                                                                               
	CompleteOauthServerOutputSpecification                                                      map[string]interface{} `json:"complete_oauth_server_output_specification,omitempty"`
	// OAuth specific blob. This is a Json Schema used to validate Json configurations used as                         
	// input to OAuth.                                                                                                 
	// Must be a valid non-nested JSON that refers to properties from                                                  
	// ConnectorSpecification.connectionSpecification                                                                  
	// using special annotation 'path_in_connector_config'.                                                            
	// These are input values the user is entering through the UI to authenticate to the                               
	// connector, that might also shared                                                                               
	// as inputs for syncing data via the connector.                                                                   
	//                                                                                                                 
	// Examples:                                                                                                       
	//                                                                                                                 
	// if no connector values is shared during oauth flow,                                                             
	// oauth_user_input_from_connector_config_specification=[]                                                         
	// if connector values such as 'app_id' inside the top level are used to generate the API                          
	// url for the oauth flow,                                                                                         
	// oauth_user_input_from_connector_config_specification={                                                          
	// app_id: {                                                                                                       
	// type: string                                                                                                    
	// path_in_connector_config: ['app_id']                                                                            
	// }                                                                                                               
	// }                                                                                                               
	// if connector values such as 'info.app_id' nested inside another object are used to                              
	// generate the API url for the oauth flow,                                                                        
	// oauth_user_input_from_connector_config_specification={                                                          
	// app_id: {                                                                                                       
	// type: string                                                                                                    
	// path_in_connector_config: ['info', 'app_id']                                                                    
	// }                                                                                                               
	// }                                                                                                               
	OauthUserInputFromConnectorConfigSpecification                                              map[string]interface{} `json:"oauth_user_input_from_connector_config_specification,omitempty"`
}
