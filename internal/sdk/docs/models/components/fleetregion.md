# FleetRegion

A fleet region is a region in which a fleet can be deployed.
You can update cloudMinVcpus once every five minutes. It must be a multiple of
scaleIncrementVcpus


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `CloudMinVcpusUpdatedAt`                  | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `CloudMinVcpus`                           | *int*                                     | :heavy_check_mark:                        | N/A                                       |
| `ScaleIncrementVcpus`                     | *int*                                     | :heavy_check_mark:                        | N/A                                       |