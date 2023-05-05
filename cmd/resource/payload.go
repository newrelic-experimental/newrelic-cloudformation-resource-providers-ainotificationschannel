package resource

import (
   "fmt"
   "github.com/newrelic-experimental/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

//
// Generic, should be able to leave these as-is
//

type Payload struct {
   model  *Model
   models []interface{}
}

func NewPayload(m *Model) *Payload {
   return &Payload{
      model:  m,
      models: make([]interface{}, 0),
   }
}

func (p *Payload) GetResourceModel() interface{} {
   return p.model
}

func (p *Payload) GetResourceModels() []interface{} {
   log.Debugf("GetResourceModels: returning %+v", p.models)
   return p.models
}

func (p *Payload) AppendToResourceModels(m model.Model) {
   p.models = append(p.models, m.GetResourceModel())
}

func (p *Payload) GetTags() map[string]string {
   // return p.model.Tags
   return nil
}

func (p *Payload) HasTags() bool {
   // return p.model.Tags != nil
   return false
}

//
// These are usually API specific, MAY BE configured per API
//

var typeName = "NewRelic::Observability::AINotificationsChannel"

func (p *Payload) NewModelFromGuid(g interface{}) (m model.Model) {
   s := fmt.Sprintf("%s", g)
   return NewPayload(&Model{Guid: &s})
}

func (p *Payload) GetGraphQLFragment() *string {
   return p.model.Channel
}

func (p *Payload) SetGuid(g *string) {
   p.model.Guid = g
   log.Debugf("SetGuid: %s", *p.model.Guid)
}

func (p *Payload) GetGuid() *string {
   return p.model.Guid
}

func (p *Payload) GetGuidKey() string {
   return "id"
}

func (p *Payload) GetVariables() map[string]string {
   vars := make(map[string]string)
   if p.model.Variables != nil {
      for k, v := range p.model.Variables {
         vars[k] = v
      }
   }

   if p.model.Guid != nil {
      vars["GUID"] = *p.model.Guid
   }

   if p.model.Channel != nil {
      vars["FRAGMENT"] = *p.model.Channel
   }

   lqf := ""
   if p.model.ListQueryFilter != nil {
      lqf = *p.model.ListQueryFilter
   }
   vars["LISTQUERYFILTER"] = lqf

   return vars
}

func (p *Payload) GetErrorKey() string {
   return "type"
}

func (p *Payload) GetResultKey(a model.Action) string {
   switch a {
   case model.Delete:
      return "ids"
   }
   return p.GetGuidKey()
}

func (p *Payload) NeedsPropagationDelay(a model.Action) bool {
   return true
}
func (p *Payload) GetCreateMutation() string {
   return `
mutation {
    aiNotificationsCreateChannel(accountId: {{{ACCOUNTID}}}, {{{FRAGMENT}}} ) {
        error {
            ... on AiNotificationsConstraintsError {
                constraints {
                    dependencies
                    name
                }
            }
            ... on AiNotificationsDataValidationError {
                details
                fields {
                    field
                    message
                }
            }
            ... on AiNotificationsResponseError {
                description
                details
                type
            }
            ... on AiNotificationsSuggestionError {
                description
                details
                type
            }
        }
        channel {
            id
        }
    }
}
`
}

func (p *Payload) GetDeleteMutation() string {
   return `
mutation {
    aiNotificationsDeleteChannel(accountId: {{{ACCOUNTID}}}, channelId: "{{{GUID}}}") {
        error {
            description
            details
            type
        }
        ids
    }
}
`
}

func (p *Payload) GetUpdateMutation() string {
   return `
mutation {
    aiNotificationsUpdateChannel(accountId: {{{ACCOUNTID}}},  {{{FRAGMENT}}} , channelId: "{{{GUID}}}") {
        error {
            ... on AiNotificationsConstraintsError {
                constraints {
                    dependencies
                    name
                }
            }
            ... on AiNotificationsDataValidationError {
                details
                fields {
                    field
                    message
                }
            }
            ... on AiNotificationsResponseError {
                description
                details
                type
            }
            ... on AiNotificationsSuggestionError {
                description
                details
                type
            }
        }
        channel {
            id
         }
    }
}
`
}

func (p *Payload) GetReadQuery() string {
   return `
{
    actor {
        account(id: {{{ACCOUNTID}}}) {
            aiNotifications {
                channels(filters: {id: "{{{GUID}}}"}) {
                    entities {
                        id
                        type
                    }
                    error {
                        description
                        details
                        type
                    }
                    nextCursor
                    totalCount
                }
            }
        }
    }
}
`
}

func (p *Payload) GetListQuery() string {
   return `
{
    actor {
        account(id: {{{ACCOUNTID}}}) {
            aiNotifications {
                channels (cursor: "{{{NEXTCURSOR}}}"){
                    entities {
                        id
                        type
                    }
                    error {
                        description
                        details
                        type
                    }
                    nextCursor
                    totalCount
                }
            }
        }
    }
}
`
}

func (p *Payload) GetListQueryNextCursor() string {
   return ""
}
