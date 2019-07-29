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
  
    - [NewGameRequest.DifficultyLevel](#solarium.NewGameRequest.DifficultyLevel)
    - [NewGameRequest.GameMode](#solarium.NewGameRequest.GameMode)
  
  
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
  
  
  
  

- [pkg/gamemodes/thewolfgame/proto/wolf.proto](#pkg/gamemodes/thewolfgame/proto/wolf.proto)
    - [TheWolfGameAction](#.TheWolfGameAction)
    - [TheWolfGameAction.VillagerVictory](#.TheWolfGameAction.VillagerVictory)
    - [TheWolfGameAction.VoteMurder](#.TheWolfGameAction.VoteMurder)
    - [TheWolfGameAction.VoteStart](#.TheWolfGameAction.VoteStart)
    - [TheWolfGameAction.WerewolfVictory](#.TheWolfGameAction.WerewolfVictory)
    - [TheWolfGameEvent](#.TheWolfGameEvent)
    - [TheWolfGameStatus](#.TheWolfGameStatus)
    - [TheWolfGameStatusPlayer](#.TheWolfGameStatusPlayer)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/solarium.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/solarium.proto



<a name="solarium.DoActionRequest"></a>

### DoActionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PlayerID | [string](#string) |  | What playerID is performing the action? this is the same as the SecretKey from a JoinGameResponse |
| GameID | [string](#string) |  | What gameId is this player in? |
| DesertPlanet | [DesertPlanetAction](#DesertPlanetAction) |  | DesertPlanet Specific actions, Populate me in order to perform actions in this scenario |
| TheWolfGame | [TheWolfGameAction](#TheWolfGameAction) |  |  |






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
| IsGameOver | [bool](#bool) |  | If you recieve this, then the game is over. |
| DesertPlanet | [DesertPlanetEvent](#DesertPlanetEvent) |  | DesertPlanet specific events. |
| TheWolfGame | [TheWolfGameEvent](#TheWolfGameEvent) |  |  |






<a name="solarium.GameStatusRequest"></a>

### GameStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  | The Id of the game you wish to get the status from |






<a name="solarium.GameStatusResponse"></a>

### GameStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| DesertPlanet | [DesertPlanetStatus](#DesertPlanetStatus) |  | DesertPlanet specific statuses |
| TheWolfGame | [TheWolfGameStatus](#TheWolfGameStatus) |  |  |






<a name="solarium.GameUpdateRequest"></a>

### GameUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  | the Id of the game you wish to subscribe to events from |






<a name="solarium.GameUpdateResponse"></a>

### GameUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Events | [GameEvent](#solarium.GameEvent) | repeated | A bulk list of recent events from the game. |






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
| Events | [GlobalEvent](#solarium.GlobalEvent) | repeated | A list of Global events |






<a name="solarium.JoinGameRequest"></a>

### JoinGameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  | The Id of the game you wish to join |
| Name | [string](#string) |  | The name of your player/character |






<a name="solarium.JoinGameResponse"></a>

### JoinGameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SecretKey | [string](#string) |  | A special secret key used for playeractions, this should be kept hidden from other players. |






<a name="solarium.NewGameRequest"></a>

### NewGameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Gamemode | [NewGameRequest.GameMode](#solarium.NewGameRequest.GameMode) |  | Which gamemode you would like to request. |
| Difficulty | [NewGameRequest.DifficultyLevel](#solarium.NewGameRequest.DifficultyLevel) |  | How difficult this game should be |






<a name="solarium.NewGameResponse"></a>

### NewGameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GameID | [string](#string) |  | The Id with which to pass back for all queries relating to this game. |
| Description | [string](#string) |  | A short description of the game |
| Name | [string](#string) |  | A name for the game, this isnt stored server side so you can choose to use it in your client or make your own! |






<a name="solarium.Player"></a>

### Player



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  | The ID Of the player, in general this is also the secret key from a JoinGameResponse, so it should be kept a secret from other clients. |
| Name | [string](#string) |  | The Name of a player. |





 


<a name="solarium.NewGameRequest.DifficultyLevel"></a>

### NewGameRequest.DifficultyLevel
A List of valid difficulties.

| Name | Number | Description |
| ---- | ------ | ----------- |
| EASY | 0 | A nice lesuirely game, plenty of room for mistakes |
| NORMAL | 1 | A bit more of a challenge, mistakes have consequences |
| HARD | 2 | You will need to work together to survive |
| INSANE | 3 | You will not survive, but you can try to last as long as possible |



<a name="solarium.NewGameRequest.GameMode"></a>

### NewGameRequest.GameMode
A List of valid gamemodes.

| Name | Number | Description |
| ---- | ------ | ----------- |
| DESERTPLANET | 0 | A survival game, where the players must work together to gather supplies and escape. |
| THEWOLFGAME | 1 | A pvp game, where the objective is for the villagers to route out the werewolves before they manage to kill them all! |


 

 


<a name="solarium.Solarium"></a>

### Solarium
The primary GRPC interface

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| NewGame | [NewGameRequest](#solarium.NewGameRequest) | [NewGameResponse](#solarium.NewGameResponse) | Constructs a new game with the given parameters. |
| JoinGame | [JoinGameRequest](#solarium.JoinGameRequest) | [JoinGameResponse](#solarium.JoinGameResponse) | Join an existing game |
| GameUpdate | [GameUpdateRequest](#solarium.GameUpdateRequest) | [GameUpdateResponse](#solarium.GameUpdateResponse) stream | Subscribe to events from a given game. |
| GlobalUpdate | [GlobalUpdateRequest](#solarium.GlobalUpdateRequest) | [GlobalUpdateResponse](#solarium.GlobalUpdateResponse) stream | Subscribe to all global events from the server, including annoucements, new game notifications etc. |
| DoAction | [DoActionRequest](#solarium.DoActionRequest) | [DoActionResponse](#solarium.DoActionResponse) | Allows a player to perform an action in a game |
| GameStatus | [GameStatusRequest](#solarium.GameStatusRequest) | [GameStatusResponse](#solarium.GameStatusResponse) | Single call to request the current state of the game. |

 



<a name="pkg/gamemodes/desert-planet/proto/desert.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pkg/gamemodes/desert-planet/proto/desert.proto



<a name=".DesertPlanetAction"></a>

### DesertPlanetAction
Wrapper for desert planet actions
Only pass a single one of these, in the
even multiple are passed, only the first
will be used.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GatherWater | [DesertPlanetGatherWater](#DesertPlanetGatherWater) |  |  |
| GatherFood | [DesertPlanetGatherFood](#DesertPlanetGatherFood) |  |  |
| GatherComponents | [DesertPlanetGatherComponents](#DesertPlanetGatherComponents) |  |  |






<a name=".DesertPlanetEvent"></a>

### DesertPlanetEvent
Wrapper for a desert planet event, these
are all the possible events that a client can recieve
when subscribing to the gameevent stream for a 
desert planet type game.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| DesertPlanetGatheredWater | [DesertPlanetGatheredWater](#DesertPlanetGatheredWater) |  |  |
| DesertPlanetGatheredFood | [DesertPlanetGatheredFood](#DesertPlanetGatheredFood) |  |  |
| DesertPlanetGatheredComponents | [DesertPlanetGatheredComponents](#DesertPlanetGatheredComponents) |  |  |
| DesertPlanetFailed | [DesertPlanetFailed](#DesertPlanetFailed) |  |  |
| DesertPlanetSucceeded | [DesertPlanetSucceeded](#DesertPlanetSucceeded) |  |  |






<a name=".DesertPlanetFailed"></a>

### DesertPlanetFailed
A special message to indicate that the players have
failed the scenario.


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
A status message describing the state
of a single player.


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
A message describing the current state of the game
including all player statuses&#39;


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
A special message to indiciate that the players have
won the scenario


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Score | [int32](#int32) |  |  |





 

 

 

 



<a name="pkg/gamemodes/thewolfgame/proto/wolf.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pkg/gamemodes/thewolfgame/proto/wolf.proto



<a name=".TheWolfGameAction"></a>

### TheWolfGameAction
Wraps all the-wolf-game actions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PlayerID | [string](#string) |  | The ID of the player who&#39;s acting |
| Vote | [TheWolfGameAction.VoteMurder](#TheWolfGameAction.VoteMurder) |  | What they have voted to do |
| StartVote | [TheWolfGameAction.VoteStart](#TheWolfGameAction.VoteStart) |  | Special vote in order to start the game |
| VillageVictory | [TheWolfGameAction.VillagerVictory](#TheWolfGameAction.VillagerVictory) |  | Populated in the event of the village victory |
| WolfVictory | [TheWolfGameAction.WerewolfVictory](#TheWolfGameAction.WerewolfVictory) |  | Populated in the event of a wolf victory |






<a name=".TheWolfGameAction.VillagerVictory"></a>

### TheWolfGameAction.VillagerVictory







<a name=".TheWolfGameAction.VoteMurder"></a>

### TheWolfGameAction.VoteMurder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PlayerId | [string](#string) |  | The player to lynch/murder |






<a name=".TheWolfGameAction.VoteStart"></a>

### TheWolfGameAction.VoteStart







<a name=".TheWolfGameAction.WerewolfVictory"></a>

### TheWolfGameAction.WerewolfVictory







<a name=".TheWolfGameEvent"></a>

### TheWolfGameEvent







<a name=".TheWolfGameStatus"></a>

### TheWolfGameStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Players | [TheWolfGameStatusPlayer](#TheWolfGameStatusPlayer) | repeated |  |
| IsNight | [bool](#bool) |  |  |






<a name=".TheWolfGameStatusPlayer"></a>

### TheWolfGameStatusPlayer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| IsAlive | [bool](#bool) |  |  |





 

 

 

 



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

