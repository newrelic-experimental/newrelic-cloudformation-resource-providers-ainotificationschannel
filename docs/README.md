# NewRelic::Observability::AINotificationsChannel

CRUD operations for New Relic AI Notifications Channels via the NerdGraph API

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "NewRelic::Observability::AINotificationsChannel",
    "Properties" : {
        "<a href="#channel" title="Channel">Channel</a>" : <i>String</i>,
        "<a href="#listqueryfilter" title="ListQueryFilter">ListQueryFilter</a>" : <i>String</i>,
        "<a href="#listquerysorter" title="ListQuerySorter">ListQuerySorter</a>" : <i>String</i>,
        "<a href="#variables" title="Variables">Variables</a>" : <i><a href="variables.md">Variables</a></i>
    }
}
</pre>

### YAML

<pre>
Type: NewRelic::Observability::AINotificationsChannel
Properties:
    <a href="#channel" title="Channel">Channel</a>: <i>String</i>
    <a href="#listqueryfilter" title="ListQueryFilter">ListQueryFilter</a>: <i>String</i>
    <a href="#listquerysorter" title="ListQuerySorter">ListQuerySorter</a>: <i>String</i>
    <a href="#variables" title="Variables">Variables</a>: <i><a href="variables.md">Variables</a></i>
</pre>

## Properties

#### Channel

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ListQueryFilter

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ListQuerySorter

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Variables

_Required_: No

_Type_: <a href="variables.md">Variables</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Guid.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Guid

Returns the <code>Guid</code> value.

