# AuthConfiguration

Configure [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service) for your application. Use Hathora's built-in auth providers or use your own [custom authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service#custom-auth-provider).


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Google`                                                                      | [*components.Google](../../models/components/google.md)                       | :heavy_minus_sign:                                                            | Enable google auth for your application.                                      |
| `Nickname`                                                                    | [*components.RecordStringNever](../../models/components/recordstringnever.md) | :heavy_minus_sign:                                                            | Construct a type with a set of properties K of type T                         |
| `Anonymous`                                                                   | [*components.RecordStringNever](../../models/components/recordstringnever.md) | :heavy_minus_sign:                                                            | Construct a type with a set of properties K of type T                         |