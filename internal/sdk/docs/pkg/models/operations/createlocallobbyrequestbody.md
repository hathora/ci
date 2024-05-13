# CreateLocalLobbyRequestBody


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `InitialConfig`                                                               | [shared.LobbyInitialConfig](../../../pkg/models/shared/lobbyinitialconfig.md) | :heavy_check_mark:                                                            | User input to initialize the game state. Object must be smaller than 64KB.    |
| `Region`                                                                      | [shared.Region](../../../pkg/models/shared/region.md)                         | :heavy_check_mark:                                                            | N/A                                                                           |