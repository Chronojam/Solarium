# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/solarium.proto](#proto/solarium.proto)
    - [DoActionRequest](#solarium.DoActionRequest)
    - [DoActionResponse](#solarium.DoActionResponse)
    - [GameEvent](#solarium.GameEvent)
    - [GameStatusRequest](#solarium.GameStatusRequest)
    - [GameStatusResponse](#solarium.GameStatusResponse)
    - [GameUpdateRequest](#solarium.GameUpdateRequest)
    - [GameUpdateResponse](#solarium.GameUpdateResponse)
    - [GlobalEvent](#solarium.GlobalEvent)
    - [GlobalUpdateRequest](#solarium.GlobalUpdateRequest)
    - [GlobalUpdateResponse](#solarium.GlobalUpdateResponse)
    - [JoinGameRequest](#solarium.JoinGameRequest)
    - [JoinGameResponse](#solarium.JoinGameResponse)
    - [NewGameRequest](#solarium.NewGameRequest)
    - [NewGameResponse](#solarium.NewGameResponse)
    - [Player](#solarium.Player)
  
  
  
    - [Solarium](#solarium.Solarium)
  

- [pkg/gamemodes/desert-planet/proto/desert.proto](#pkg/gamemodes/desert-planet/proto/desert.proto)
    - [DesertPlanetAction](#.DesertPlanetAction)
    - [DesertPlanetEvent](#.DesertPlanetEvent)
    - [DesertPlanetFailed](#.DesertPlanetFailed)
    - [DesertPlanetGatherComponents](#.DesertPlanetGatherComponents)
    - [DesertPlanetGatherFood](#.DesertPlanetGatherFood)
    - [DesertPlanetGatherWater](#.DesertPlanetGatherWater)
    - [DesertPlanetGatheredComponents](#.DesertPlanetGatheredComponents)
    - [DesertPlanetGatheredFood](#.DesertPlanetGatheredFood)
    - [DesertPlanetGatheredWater](#.DesertPlanetGatheredWater)
    - [DesertPlanetPlayerStatus](#.DesertPlanetPlayerStatus)
    - [DesertPlanetStatus](#.DesertPlanetStatus)
    - [DesertPlanetSucceeded](#.DesertPlanetSucceeded)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/solarium.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/solarium.proto



<a name="solarium.DoActionRequest"></a>

### DoActionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PlayerID | [string](#string) |  |  |
| GameID | [string](#string) |  |  |
| DesertPlanet | [DesertPlanetAction](#DesertPlanetAction) |  | Desert Planet Scenario |






<a name="solarium.DoActionResponse"></a>

### DoActionResponse







<a name="solarium.GameEvent"></a>

### GameEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  | The name of the event. |
| Desc | [string](#string) |  | The Description of the event. |
| InitatingPlayers | [Player](#solarium.Player) | repeated | A list of players who initiated this event |
| AffectedPlayers | [Player](#solarium.Player) | repeated | A list of players who are affected by this event |
| DesertPlanet | [DesertPlanetEvent](#DesertPlanetEvent) |  | Desert Planet Scenario |






<a name="solarium.GameStatusRequest"></a>

### GameStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  |  |






<a name="solarium.GameStatusResponse"></a>

### GameStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| DesertPlanet | [DesertPlanetStatus](#DesertPlanetStatus) |  |  |






<a name="solarium.GameUpdateRequest"></a>

### GameUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  |  |






<a name="solarium.GameUpdateResponse"></a>

### GameUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Events | [GameEvent](#solarium.GameEvent) | repeated |  |






<a name="solarium.GlobalEvent"></a>

### GlobalEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Notification | [string](#string) |  |  |






<a name="solarium.GlobalUpdateRequest"></a>

### GlobalUpdateRequest







<a name="solarium.GlobalUpdateResponse"></a>

### GlobalUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Events | [GlobalEvent](#solarium.GlobalEvent) | repeated |  |






<a name="solarium.JoinGameRequest"></a>

### JoinGameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="solarium.JoinGameResponse"></a>

### JoinGameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SecretKey | [string](#string) |  |  |






<a name="solarium.NewGameRequest"></a>

### NewGameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Gamemode | [string](#string) |  |  |
| Difficulty | [int64](#int64) |  |  |






<a name="solarium.NewGameResponse"></a>

### NewGameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="solarium.Player"></a>

### Player



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |





 

 

 


<a name="solarium.Solarium"></a>

### Solarium


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| NewGame | [NewGameRequest](#solarium.NewGameRequest) | [NewGameResponse](#solarium.NewGameResponse) |  |
| JoinGame | [JoinGameRequest](#solarium.JoinGameRequest) | [JoinGameResponse](#solarium.JoinGameResponse) |  |
| GameUpdate | [GameUpdateRequest](#solarium.GameUpdateRequest) | [GameUpdateResponse](#solarium.GameUpdateResponse) stream |  |
| GlobalUpdate | [GlobalUpdateRequest](#solarium.GlobalUpdateRequest) | [GlobalUpdateResponse](#solarium.GlobalUpdateResponse) stream |  |
| DoAction | [DoActionRequest](#solarium.DoActionRequest) | [DoActionResponse](#solarium.DoActionResponse) |  |
| GameStatus | [GameStatusRequest](#solarium.GameStatusRequest) | [GameStatusResponse](#solarium.GameStatusResponse) |  |

 



<a name="pkg/gamemodes/desert-planet/proto/desert.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pkg/gamemodes/desert-planet/proto/desert.proto



<a name=".DesertPlanetAction"></a>

### DesertPlanetAction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GatherWater | [DesertPlanetGatherWater](#DesertPlanetGatherWater) |  |  |
| GatherFood | [DesertPlanetGatherFood](#DesertPlanetGatherFood) |  |  |
| GatherComponents | [DesertPlanetGatherComponents](#DesertPlanetGatherComponents) |  |  |






<a name=".DesertPlanetEvent"></a>

### DesertPlanetEvent
## NOTIFICATIONS 
NotificationInterface


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| DesertPlanetGatheredWater | [DesertPlanetGatheredWater](#DesertPlanetGatheredWater) |  |  |
| DesertPlanetGatheredFood | [DesertPlanetGatheredFood](#DesertPlanetGatheredFood) |  |  |
| DesertPlanetGatheredComponents | [DesertPlanetGatheredComponents](#DesertPlanetGatheredComponents) |  |  |
| DesertPlanetFailed | [DesertPlanetFailed](#DesertPlanetFailed) |  |  |
| DesertPlanetSucceeded | [DesertPlanetSucceeded](#DesertPlanetSucceeded) |  |  |






<a name=".DesertPlanetFailed"></a>

### DesertPlanetFailed
## FAILURE AND SUCCESS CONDITIONS


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Score | [int32](#int32) |  |  |






<a name=".DesertPlanetGatherComponents"></a>

### DesertPlanetGatherComponents
Instruct the player to gather some components.






<a name=".DesertPlanetGatherFood"></a>

### DesertPlanetGatherFood
Instruct the player to gather some food.






<a name=".DesertPlanetGatherWater"></a>

### DesertPlanetGatherWater
Instruct the player to gather some water.






<a name=".DesertPlanetGatheredComponents"></a>

### DesertPlanetGatheredComponents
A player gathered some components


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Quantity | [int32](#int32) |  | How much water did they gather? |






<a name=".DesertPlanetGatheredFood"></a>

### DesertPlanetGatheredFood
A player gathered some food


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Quantity | [int32](#int32) |  | How much water did they gather? |






<a name=".DesertPlanetGatheredWater"></a>

### DesertPlanetGatheredWater
A player gathered some water


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Quantity | [int32](#int32) |  | How much water did they gather? |






<a name=".DesertPlanetPlayerStatus"></a>

### DesertPlanetPlayerStatus
## STATUS MESSAGES


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PlayerID | [string](#string) |  |  |
| PlayerName | [string](#string) |  |  |
| Thirst | [int32](#int32) |  |  |
| Hunger | [int32](#int32) |  |  |
| Incapaciated | [bool](#bool) |  |  |
| Status | [string](#string) | repeated |  |






<a name=".DesertPlanetStatus"></a>

### DesertPlanetStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Water | [int32](#int32) |  |  |
| Food | [int32](#int32) |  |  |
| Fuel | [int32](#int32) |  |  |
| Components | [int32](#int32) |  |  |
| TargetComponents | [int32](#int32) |  |  |
| PlayerStatus | [DesertPlanetPlayerStatus](#DesertPlanetPlayerStatus) | repeated |  |






<a name=".DesertPlanetSucceeded"></a>

### DesertPlanetSucceeded



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Score | [int32](#int32) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

