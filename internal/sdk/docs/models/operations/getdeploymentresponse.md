# GetDeploymentResponse


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ContentType`                                               | *string*                                                    | :heavy_check_mark:                                          | HTTP response content type for this operation               |
| `StatusCode`                                                | *int*                                                       | :heavy_check_mark:                                          | HTTP response status code for this operation                |
| `RawResponse`                                               | [*http.Response](https://pkg.go.dev/net/http#Response)      | :heavy_check_mark:                                          | Raw HTTP response; suitable for custom response parsing     |
| `DeploymentV3`                                              | [*shared.DeploymentV3](../../models/shared/deploymentv3.md) | :heavy_minus_sign:                                          | Ok                                                          |