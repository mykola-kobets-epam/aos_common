# UM Protocol

To trigger update procedure Update manager should implements client's part of update protocol.
**System manager** implements server part of UM protocol.
Communication between SM and UM done using [gRPC framework](https://grpc.io/).

## Upgrade sequence

On startup each Update manager should subscribe on notifications about update requests. Returned gRPC stream should be used fro upcoming communications.

```mermaid
sequenceDiagram
    participant cloud
    participant SM
    participant UMctrl
     Note right of UMctrl: part of SM
    participant UMDomF
    participant UMDomD

    UMDomD->>UMctrl: RegisterUM(UmID, status)
    UMDomF->>UMctrl: RegisterUM(UmID, status)

    cloud->>SM: updateRequest

    activate SM
    SM->>UMctrl: process Update
    activate UMctrl
    UMctrl->>UMctrl: download and save package
    UMctrl->>+UMDomD: PrepareUpdateRequest(path)
    UMctrl->>+UMDomF: PrepareUpdateRequest(path)
    UMDomD->>-UMctrl: PrepareUpdateResponse(true)
    UMDomF->>-UMctrl: PrepareUpdateResponse(true)
    UMctrl->>UMctrl: wait for all responses
    UMctrl->>+UMDomD: StartUpdateNotification()
    UMDomD->>-UMctrl: updateStatus(status)
    Note over UMctrl,UMDomD: disconnect or reboot
    deactivate UMctrl

    activate UMctrl
    UMDomD->>UMctrl: RegisterUM(UmID, status)
    UMctrl->>+UMDomD: ApplyUpdateNotification()
    UMDomD->>-UMctrl: updateStatus(state)
    Note over UMctrl,UMDomD: disconnect or reboot
    deactivate UMctrl

    activate UMctrl
    UMDomD->>UMctrl: RegisterUM(UmID, status)
    UMctrl->>+UMDomF: StartUpdateNotification()
    UMDomF->>-UMctrl: updateStatus(status)
    Note over UMctrl,DomF: disconnect or reboot
    deactivate UMctrl

    activate UMctrl
    UMDomF->>UMctrl: RegisterUM(UmID, status)
    UMctrl->>+UMDomF: ApplyUpdateNotification()
    UMDomF->>-UMctrl: updateStatus(state)
    Note over UMctrl,UMDomF: disconnect or reboot
    deactivate UMctrl

    activate UMctrl
    UMDomD->>UMctrl: RegisterUM(UmID, status)
    UMctrl->>SM: updateStatus(OK)

    SM->>cloud: update finish

    deactivate SM
```
