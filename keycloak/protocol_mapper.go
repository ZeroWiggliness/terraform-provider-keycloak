package keycloak

import (
	"context"
	"fmt"
)

// https://www.keycloak.org/docs-api/4.2/rest-api/index.html#_protocolmapperrepresentation
type protocolMapper struct {
	Id             string            `json:"id,omitempty"`
	Name           string            `json:"name"`
	Protocol       string            `json:"protocol"`
	ProtocolMapper string            `json:"protocolMapper"`
	Config         map[string]string `json:"config"`
}

var (
	addToAccessTokenField                = "access.token.claim"
	addToIdTokenField                    = "id.token.claim"
	addToUserInfoField                   = "userinfo.token.claim"
	addToIntrospectionTokenField         = "introspection.token.claim"
	addToLightweightClaimField           = "lightweight.claim"
	attributeNameField                   = "attribute.name"
	attributeNameFormatField             = "attribute.nameformat"
	claimNameField                       = "claim.name"
	claimValueField                      = "claim.value"
	claimValueTypeField                  = "jsonType.label"
	jsonTypeField                        = "jsonType.label"
	friendlyNameField                    = "friendly.name"
	fullPathField                        = "full.path"
	includedClientAudienceField          = "included.client.audience"
	includedCustomAudienceField          = "included.custom.audience"
	multivaluedField                     = "multivalued"
	samlScriptField                      = "Script" // needs to start with uppercase S for SAML script mapper
	scriptField                          = "script"
	singleValueAttributeField            = "single"
	userAttributeField                   = "user.attribute"
	userPropertyField                    = "user.attribute"
	userRealmRoleMappingRolePrefixField  = "usermodel.realmRoleMapping.rolePrefix"
	userClientRoleMappingClientIdField   = "usermodel.clientRoleMapping.clientId"
	userClientRoleMappingRolePrefixField = "usermodel.clientRoleMapping.rolePrefix"
	userSessionNoteField                 = "user.session.note"
	aggregateAttributeValuesField        = "aggregate.attrs"

	protocolMapperIgnore = []string{addToIntrospectionTokenField, addToUserInfoField, addToIdTokenField, addToAccessTokenField, addToLightweightClaimField, claimNameField, jsonTypeField}
)

func protocolMapperPath(realmId, clientId, clientScopeId string) string {
	parentResourceId := clientId
	parentResourcePath := "clients"

	if clientScopeId != "" {
		parentResourceId = clientScopeId
		parentResourcePath = "client-scopes"
	}

	return fmt.Sprintf("/realms/%s/%s/%s/protocol-mappers/models", realmId, parentResourcePath, parentResourceId)
}

func individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId string) string {
	return fmt.Sprintf("%s/%s", protocolMapperPath(realmId, clientId, clientScopeId), mapperId)
}

func (keycloakClient *KeycloakClient) listGenericProtocolMappers(ctx context.Context, realmId, clientId, clientScopeId string) ([]*protocolMapper, error) {
	var protocolMappers []*protocolMapper

	err := keycloakClient.get(ctx, protocolMapperPath(realmId, clientId, clientScopeId), &protocolMappers, nil)
	if err != nil {
		return nil, err
	}

	return protocolMappers, nil
}
